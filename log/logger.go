package log

import (
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.SugaredLogger = newLogger()

// Info prints info level logs
func Info(args ...interface{}) {
	logger.Info(args...)
}

// Infof prints formatted info level logs
func Infof(template string, args ...interface{}) {
	logger.Infof(template, args...)
}

// Error prints error level logs
func Error(args ...interface{}) {
	logger.Error(args...)
}

// Errorf prints formatted error level logs
func Errorf(template string, args ...interface{}) {
	logger.Errorf(template, args...)
}

// Debug prints debug level logs
func Debug(args ...interface{}) {
	logger.Debug(args...)
}

// Debugf prints formatted debug level logs
func Debugf(template string, args ...interface{}) {
	logger.Debugf(template, args...)
}

// Fatal prints logs and calls os.Exit()
func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

// Fatalf prints logs and calls os.Exit()
func Fatalf(template string, args ...interface{}) {
	logger.Fatalf(template, args...)
}

// Warn prints warn level logs
func Warn(args ...interface{}) {
	logger.Warn(args...)
}

// Warnf prints formatted warn level logs
func Warnf(template string, args ...interface{}) {
	logger.Warnf(template, args...)
}

// WithLogOption returns a logger with given additional option
func WithLogOption(option zap.Option) *zap.SugaredLogger {
	core := zapcore.NewCore(getEncoder(), getLogWriter(), zapcore.DebugLevel)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(0))
	return logger.WithOptions(option).Sugar()
}

func newLogger() *zap.SugaredLogger {
	core := zapcore.NewCore(getEncoder(), getLogWriter(), zapcore.DebugLevel)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	// Use SugaredLogger
	return logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	// By default, logs are printed to console
	return zapcore.AddSync(os.Stdout)
}

func formatEncodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(fmt.Sprintf("%d%02d%02d_%02d%02d%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second()))
}
