package mw

import (
	"os"

	"cloud.google.com/go/trace"
	"go.uber.org/zap"
)

var (
	log    *zap.Logger
	tracer *trace.Client
)

func init() {
	// configure the zap logger
	if os.Getenv("ENV") == "dev" {
		log, _ = zap.NewDevelopment()
	} else {
		log, _ = zap.NewProduction()
	}
}

// SetGlobalTracer sets the tracer used by all parts of the application
func SetGlobalTracer(t *trace.Client) {
	tracer = t
}

// GlobalTracer returns the global tracer
func GlobalTracer() *trace.Client {
	return tracer
}
