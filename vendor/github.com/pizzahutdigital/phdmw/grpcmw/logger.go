package grpcmw

import (
	"context"
	"time"

	"github.com/pizzahutdigital/phdmw/phdcid"
	"github.com/pizzahutdigital/phdmw/phdlog"
	"google.golang.org/grpc"
)

type key int

const (
	// HandlerNameKey is the key for context values sent from handler to the logger.  This is to prevent package conflicts.
	HandlerNameKey key = iota
)

// HandlerStart - Used at begining of handler to add handler name in the context so it can be logged in the HTTPHandlerSummary Event log
func HandlerStart(r context.Context, handlerName string) string {
	// Prevent panics during tests.
	if r == nil {
		return ""
	}

	iPtr := r.Value(HandlerNameKey)
	if iPtr != nil {
		ptr := iPtr.(*string)
		*ptr = handlerName
	}
	return phdcid.GetCIDFromContext(r)
}

// LoggerMW returns a new unary server interceptors that adds zap.Logger to the context.
func LoggerMiddleware() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		startTime := time.Now()

		handlerName := "unknown"

		ctx = context.WithValue(ctx, HandlerNameKey, &handlerName)
		resp, err := handler(ctx, req)

		errorString := ""
		if err != nil {
			errorString = err.Error()
		}
		duration := ((time.Since(startTime).Nanoseconds()) / 1000000)
		phdlog.GRPCHandlerSummary(errorString,
			int64(duration),
			phdcid.GetCIDFromContext(ctx),
			handlerName)

		// re-extract logger from newCtx, as it may have extra fields that changed in the holder.
		// ctx_zap.Extract(newCtx).Check(level, "finished unary call with code "+code.String()).Write(
		// 	zap.Error(err),
		// 	zap.String("grpc.code", code.String()),
		// 	o.durationFunc(time.Since(startTime)),
		// )

		return resp, err
	}
}
