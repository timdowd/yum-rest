package server

import (
	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/pizzahutdigital/phdmw/grpcmw"
	"github.com/pizzahutdigital/phdmw/phdlog"
	"github.com/pizzahutdigital/yum-rest/handlers"
	pb "github.com/pizzahutdigital/yum-rest/protobufs"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.opencensus.io/exporter/stackdriver"
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/trace"
	"go.uber.org/zap"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
)

// RunRPC starts rpc server for siskel
func RunRPC() error {

	// Get flags
	traceserviceaccountfile := viper.GetString("traceserviceaccountfile")
	if traceserviceaccountfile == "" {
		return errors.New("You must supply a valid service account for tracing using the `traceserviceaccountfile` flag")
	}

	projectid := viper.GetString("projectid")
	if projectid == "" {
		return errors.New("You must provide a valid project id using the `projectid` argument")
	}

	// Create a stackdriver exporter for traces.
	stackExporter, err := stackdriver.NewExporter(stackdriver.Options{
		ProjectID: projectid,
		TraceClientOptions: []option.ClientOption{
			option.WithCredentialsFile(traceserviceaccountfile),
		},
	})
	if err != nil {
		werr := errors.Wrap(err, "stackdriver.NewExporter")
		phdlog.Info(logMessage,
			"",
			zap.String("processStatus", "unable to create stackdriver exporter"),
			zap.String("error", werr.Error()))
		return werr
	}
	// Register the stackdriver exporter.
	trace.RegisterExporter(stackExporter)

	rpcPort := ":" + viper.GetString("rpc-port")
	if rpcPort == ":" {
		return errors.New("You must supply a valid port using the 'rpc-port' argument")
	}
	lis, err := net.Listen("tcp", rpcPort)
	if err != nil {
		return errors.Wrap(err, "failed to initialize TCP listen: %v")
	}

	defer func() {
		if ferr := lis.Close(); err != nil {
			phdlog.Error(logMessage, "", zap.String("error", ferr.Error()))
		}
	}()

	rpcServer := grpc.NewServer(
		grpc.StatsHandler(&ocgrpc.ServerHandler{
			StartOptions: trace.StartOptions{
				Sampler: trace.AlwaysSample(),
			},
		}),
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpcmw.ConversationIDMiddleware(),
				grpcmw.LoggerMiddleware(),
			),
		),
	)
	var service *handlers.RestServiceServer
	service, err = handlers.NewRest()
	if err != nil {
		return err
	}

	pb.RegisterRestServiceServer(rpcServer, service)

	phdlog.Info(logMessage, "", zap.String("RPC Listening on", lis.Addr().String()))
	return rpcServer.Serve(lis)
}
