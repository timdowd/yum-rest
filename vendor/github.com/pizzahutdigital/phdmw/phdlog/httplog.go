package phdlog

import (
	"net"
	"net/http"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type httpHeaders struct {
	headers http.Header
}

func (httpHeader httpHeaders) MarshalLogObject(enc zapcore.ObjectEncoder) error {

	for k, v := range httpHeader.headers {
		redacted := []string{"auth", "key"}

		testString := strings.ToLower(k)
		strValue := strings.Join(v, ",")
		for i := 0; i < len(redacted); i++ {
			if strings.Contains(testString, redacted[i]) {
				strValue = "PHDRedacted"
			}
		}
		enc.AddString(k, strValue)
	}
	return nil
}

func calculateURI(req *http.Request) string {
	uri := req.RequestURI

	// Requests using the CONNECT method over HTTP/2.0 must use
	// the authority field (aka r.Host) to identify the target.
	// Refer: https://httpwg.github.io/specs/rfc7540.html#CONNECT
	if req.ProtoMajor == 2 && req.Method == "CONNECT" {
		uri = req.Host
	}
	if uri == "" {
		uri = req.URL.RequestURI()
	}

	return uri
}

func calculateHost(req *http.Request) string {
	host, _, err := net.SplitHostPort(req.RemoteAddr)

	if err != nil {
		host = req.RemoteAddr
	}

	return host
}

// HTTPHandlerSummary - Handler summary log
func HTTPHandlerSummary(req *http.Request, size int, status int, duration int64, conversationID string, handlerName string) {

	WrappedInfo("httpHandlerSummary", conversationID,
		zap.String("uri", calculateURI(req)),
		zap.Int64("duration", duration),
		zap.String("method", req.Method),
		zap.String("host", calculateHost(req)),
		zap.String("proto", req.Proto),
		zap.Int("status", status),
		zap.Int("size", size),
		zap.String("handler", handlerName),
	)
}

// HTTPReqIn - Inbound HTTP Request Log
func HTTPReqIn(req *http.Request, conversationID string) {

	WrappedInfo("httpReq", conversationID,
		zap.String("uri", calculateURI(req)),
		zap.String("direction", "in"),
		zap.String("method", req.Method),
		zap.String("host", calculateHost(req)),
		zap.String("proto", req.Proto),
		zap.Object("headers", httpHeaders{headers: req.Header}),
	)
}

// HTTPReqOut - Outbound HTTP Request Log
func HTTPReqOut(req *http.Request, conversationID string) {

	WrappedInfo("httpReq", conversationID,
		zap.String("uri", calculateURI(req)),
		zap.String("direction", "out"),
		zap.String("method", req.Method),
		zap.String("proto", req.Proto),
		zap.Object("headers", httpHeaders{headers: req.Header}),
	)
}

// HTTPRspIn - Inbound HTTP Response Log
func HTTPRspIn(rsp *http.Response, conversationID string) {

	WrappedInfo("httpRsp", conversationID,
		zap.String("direction", "in"),
		zap.Int("status", rsp.StatusCode),
		zap.Int64("contentLen", rsp.ContentLength),
		zap.String("proto", rsp.Proto),
		zap.Object("headers", httpHeaders{headers: rsp.Header}),
	)
}

// HTTPRspOut - Outbound HTTP Response Log
func HTTPRspOut(rsp *http.Response, conversationID string) {

	WrappedInfo("httpRsp", conversationID,
		zap.String("direction", "out"),
		zap.Int("status", rsp.StatusCode),
		zap.Int64("contentLen", rsp.ContentLength),
		zap.String("proto", rsp.Proto),
		zap.Object("headers", httpHeaders{headers: rsp.Header}),
	)
}
