package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/pizzahutdigital/phdmw/httpmw"
	"github.com/pizzahutdigital/phdmw/phdlog"
	"github.com/pizzahutdigital/yum-rest/mw"
	pb "github.com/pizzahutdigital/yum-rest/protobufs"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

const (
	logMessage = "yum-rest"
)

type errorBody struct {
	Err string `json:"error,omitempty"`
}

// CustomHTTPError custom error
func CustomHTTPError(ctx context.Context, _ *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, _ *http.Request, err error) {
	const fallback = `{"error": "failed to marshal error message"}`

	w.Header().Set("Content-type", marshaler.ContentType())
	w.WriteHeader(runtime.HTTPStatusFromCode(grpc.Code(err)))
	jErr := json.NewEncoder(w).Encode(errorBody{
		Err: grpc.ErrorDesc(err),
	})

	if jErr != nil {
		if _, err := w.Write([]byte(fallback)); err != nil {
			phdlog.Error(logMessage, "", zap.String("error", err.Error()))
		}
	}
}

// RunREST Start rest server for yum-rest
func RunREST() error {
	servicename := viper.GetString("servicename")
	if servicename == "" {
		return errors.New("You must supply a valid servicename for logging using the `servicename` flag")
	}

	phdlog.InitLogger(servicename, zap.InfoLevel)

	serverIP := viper.GetString("server-ip")
	if serverIP == "" {
		return errors.New("You must supply a valid server-ip using the 'server-id' argument")
	}

	runtime.HTTPError = CustomHTTPError

	restPort := ":" + viper.GetString("rest-port")
	if restPort == ":" {
		return errors.New("You must supply a valid port using the 'rest-port' argument")
	}

	rpcPort := viper.GetString("rpc-port")
	if rpcPort == "" {
		return errors.New("You must supply a valid port using the 'rpc-port' argument")
	}

	rpcIP := viper.GetString("rpc-address")
	if rpcPort == "" {
		return errors.New("You must supply a valid address using the 'rpc-address' argument")
	}

	opts := []grpc.DialOption{grpc.WithInsecure()}

	headerToMatch := map[string]bool{
		"location": true,
	}

	mux := runtime.NewServeMux(runtime.WithOutgoingHeaderMatcher(func(header string) (string, bool) {
		return header, headerToMatch[header]
	}))

	err := pb.RegisterRestServiceHandlerFromEndpoint(context.Background(), mux, fmt.Sprintf("%s:%s", rpcIP, rpcPort), opts)
	if err != nil {
		return errors.Wrap(err, "failed to start HTTP server: %v")
	}

	phdlog.Info(logMessage, "", zap.String("HTTP Listening on", restPort))

	mwHandlers := []httpmw.Constructor{
		httpmw.ConversationIDMiddleware,
		httpmw.LoggingMiddleware(),
		mw.GlobalTracer().HTTPHandler,
		mw.CORSMiddleware,
	}

	if os.Getenv("ENV") == "dev" {
		mwHandlers = mwHandlers[:len(mwHandlers)-1]
	}

	return http.ListenAndServe(restPort, httpmw.New(mwHandlers...).Then(mux))
}
