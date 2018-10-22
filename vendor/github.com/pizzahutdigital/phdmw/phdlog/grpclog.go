package phdlog

import (
	"go.uber.org/zap"
)

// GRPCHandlerSummary - Handler summary log
func GRPCHandlerSummary(err string, duration int64, conversationID string, handlerName string) {

	WrappedInfo("grpcHandlerSummary", conversationID,
		zap.Int64("duration", duration),
		zap.String("error", err),
		zap.String("handler", handlerName),
	)
}

// // HTTPReqIn - Inbound HTTP Request Log
// func HTTPReqIn(req *http.Request, conversationID string) {

// 	WrappedInfo("httpReq", conversationID,
// 		zap.String("uri", calculateURI(req)),
// 		zap.String("direction", "in"),
// 		zap.String("method", req.Method),
// 		zap.String("host", calculateHost(req)),
// 		zap.String("proto", req.Proto),
// 		zap.Object("headers", httpHeaders{headers: req.Header}),
// 	)
// }

// // HTTPReqOut - Outbound HTTP Request Log
// func HTTPReqOut(req *http.Request, conversationID string) {

// 	WrappedInfo("httpReq", conversationID,
// 		zap.String("uri", calculateURI(req)),
// 		zap.String("direction", "out"),
// 		zap.String("method", req.Method),
// 		zap.String("proto", req.Proto),
// 		zap.Object("headers", httpHeaders{headers: req.Header}),
// 	)
// }

// // HTTPRspIn - Inbound HTTP Response Log
// func HTTPRspIn(rsp *http.Response, conversationID string) {

// 	WrappedInfo("httpRsp", conversationID,
// 		zap.String("direction", "in"),
// 		zap.Int("status", rsp.StatusCode),
// 		zap.Int64("contentLen", rsp.ContentLength),
// 		zap.String("proto", rsp.Proto),
// 		zap.Object("headers", httpHeaders{headers: rsp.Header}),
// 	)
// }

// // HTTPRspOut - Outbound HTTP Response Log
// func HTTPRspOut(rsp *http.Response, conversationID string) {

// 	WrappedInfo("httpRsp", conversationID,
// 		zap.String("direction", "out"),
// 		zap.Int("status", rsp.StatusCode),
// 		zap.Int64("contentLen", rsp.ContentLength),
// 		zap.String("proto", rsp.Proto),
// 		zap.Object("headers", httpHeaders{headers: rsp.Header}),
// 	)
// }
