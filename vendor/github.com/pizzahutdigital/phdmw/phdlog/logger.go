package phdlog

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	// Dev sets environment to development; includes all debug level logging
	Dev = zap.DebugLevel
	// Prod sets environment to production; excludes all debug level logging
	Prod = zap.InfoLevel
)

var (
	logger        *zap.Logger
	wrappedLogger *zap.Logger
	cfg           zap.Config
)

func init() {
	cfg = zap.NewProductionConfig()
	InitLogger("Please Call InitLogger in main.", Prod)
}

// Debug - log debug messsage
func Debug(msg string, conversationID string, fields ...zapcore.Field) {
	logger.Debug(msg, append(fields, zap.String("conversationID", conversationID))...)
}

// Warn - log warning message
func Warn(msg string, conversationID string, fields ...zapcore.Field) {
	logger.Warn(msg, append(fields, zap.String("conversationID", conversationID))...)
}

// Error - log error message
func Error(msg string, conversationID string, fields ...zapcore.Field) {
	logger.Error(msg, append(fields, zap.String("conversationID", conversationID))...)
}

// Fatal - log fatal message
func Fatal(msg string, conversationID string, fields ...zapcore.Field) {
	logger.Fatal(msg, append(fields, zap.String("conversationID", conversationID))...)
}

// Info - log Info message
func Info(msg string, conversationID string, fields ...zapcore.Field) {
	logger.Info(msg, append(fields, zap.String("conversationID", conversationID))...)
}

/////////////////////
// Wrapped logger is used in cases that the log is wrapped in another function.
// The wrapped logged function will allow the caller parameter to behave correctely
////////////////////

// WrappedDebug - log debug messsage. Used in case the logger is wrapped by a helper function
func WrappedDebug(msg string, conversationID string, fields ...zapcore.Field) {
	wrappedLogger.Debug(msg, append(fields, zap.String("conversationID", conversationID))...)
}

// WrappedWarn - log warning message. Used in case the logger is wrapped by a helper function
func WrappedWarn(msg string, conversationID string, fields ...zapcore.Field) {
	wrappedLogger.Warn(msg, append(fields, zap.String("conversationID", conversationID))...)
}

// WrappedError - log error message. Used in case the logger is wrapped by a helper function
func WrappedError(msg string, conversationID string, fields ...zapcore.Field) {
	wrappedLogger.Error(msg, append(fields, zap.String("conversationID", conversationID))...)
}

// WrappedFatal - log fatal message. Used in case the logger is wrapped by a helper function
func WrappedFatal(msg string, conversationID string, fields ...zapcore.Field) {
	wrappedLogger.Fatal(msg, append(fields, zap.String("conversationID", conversationID))...)
}

// WrappedInfo - log Info message. Used in case the logger is wrapped by a helper function
func WrappedInfo(msg string, conversationID string, fields ...zapcore.Field) {
	wrappedLogger.Info(msg, append(fields, zap.String("conversationID", conversationID))...)
}

// InitLogger - init logger
func InitLogger(serviceName string, environment zapcore.Level) {
	var err error

	// adjust the production default encoder
	cfg.EncoderConfig.LevelKey = "severity"
	cfg.EncoderConfig.MessageKey = "message"
	cfg.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.UTC().Format("2006-01-02T15:04:05.000Z0700"))
	}

	cfg.OutputPaths = []string{"stdout"}
	cfg.ErrorOutputPaths = []string{"stderr"}
	cfg.Level = zap.NewAtomicLevelAt(environment)
	cfg.DisableStacktrace = true
	cfg.DisableCaller = false
	cfg.InitialFields = map[string]interface{}{"serviceName": serviceName}

	// Because zap is wrapped we need to tell zap to skip a frame.
	logger, err = cfg.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}

	// wrapped logger is the same as the logger except the call stack is adjusted
	wrappedLogger, err = cfg.Build(zap.AddCallerSkip(2))
	if err != nil {
		panic(err)
	}
}

///////////////////////////////////// New logging interface....

// PHDLogger - holds the logger configuration
type PHDLogger struct {
	cfg *zap.Logger
}

// NewPHDLogger - create a new PHDLogger.
func NewPHDLogger(callerSkip int) *PHDLogger {
	cfg, err := cfg.Build(zap.AddCallerSkip(1 + callerSkip))
	if err != nil {
		panic(err)
	}
	return &PHDLogger{cfg: cfg}
}

// Debug - log debug messsage. Used in case the logger is wrapped by a helper function
func (logger *PHDLogger) Debug(msg string, conversationID string, fields ...zapcore.Field) {
	logger.cfg.Debug(msg, append(fields, zap.String("conversationID", conversationID))...)
}

// Warn - log warning message. Used in case the logger is wrapped by a helper function
func (logger *PHDLogger) Warn(msg string, conversationID string, fields ...zapcore.Field) {
	logger.cfg.Warn(msg, append(fields, zap.String("conversationID", conversationID))...)
}

// Error - log error message. Used in case the logger is wrapped by a helper function
func (logger *PHDLogger) Error(msg string, conversationID string, fields ...zapcore.Field) {
	logger.cfg.Error(msg, append(fields, zap.String("conversationID", conversationID))...)
}

// Fatal - log fatal message. Used in case the logger is wrapped by a helper function
func (logger *PHDLogger) Fatal(msg string, conversationID string, fields ...zapcore.Field) {
	logger.cfg.Fatal(msg, append(fields, zap.String("conversationID", conversationID))...)
}

// Info - log Info message. Used in case the logger is wrapped by a helper function
func (logger *PHDLogger) Info(msg string, conversationID string, fields ...zapcore.Field) {
	logger.cfg.Info(msg, append(fields, zap.String("conversationID", conversationID))...)
}
