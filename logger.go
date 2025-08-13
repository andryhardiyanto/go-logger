// Package logger provides a structured logging interface built on top of zap.
// It offers a simple, consistent API for logging with support for different output formats,
// log levels, and New Relic integration.
package logger

import (
	"context"
	"errors"

	"github.com/newrelic/go-agent/v3/integrations/logcontext-v2/nrzap"
	"go.uber.org/zap"
)

// Logger defines the interface for structured logging operations.
// All logging methods accept a message string and optional structured fields.
// The interface supports different log levels and allows creating child loggers with additional fields.
type Logger interface {
	// Info logs a message at InfoLevel with optional structured fields
	Info(ctx context.Context, msg string, fields ...Field)
	// Warn logs a message at WarnLevel with optional structured fields
	Warn(ctx context.Context, msg string, fields ...Field)
	// Error logs a message at ErrorLevel with optional structured fields
	Error(ctx context.Context, msg string, fields ...Field)
	// Debug logs a message at DebugLevel with optional structured fields
	Debug(ctx context.Context, msg string, fields ...Field)
	// Fatal logs a message at FatalLevel with optional structured fields, then calls os.Exit(1)
	Fatal(ctx context.Context, msg string, fields ...Field)
	// Panic logs a message at PanicLevel with optional structured fields, then panics
	Panic(ctx context.Context, msg string, fields ...Field)
	// With creates a child logger with additional structured fields
	With(fields ...Field) Logger
	// GetLogger returns the underlying zap.Logger instance for advanced usage
	GetLogger() *zap.Logger
}

// logger is a wrapper struct that implements the Logger interface.
// It provides a consistent API while delegating actual logging operations to the underlying Logger implementation.
type logger struct {
	logger Logger
}

// NewLogger creates a new logger.

// NewLogger creates a new logger instance with the provided configuration options.
// It applies default configuration first, then applies user-provided options.
// The logger is configured based on the application mode (development, staging, production).
//
// Parameters:
//   - opts: Variable number of Option functions to configure the logger
//
// Returns:
//   - Logger: A configured logger instance ready for use
//   - error: An error if logger creation fails
//
// Example:
//   logger, err := NewLogger(
//       WithLevel(LevelInfo),
//       WithEncoding(EncodingJSON),
//       WithAppMode(AppModeProduction),
//   )
//   if err != nil {
//       log.Fatal(err)
//   }
//   logger.Info(context.Background(), "Application started")
func NewLogger(opts ...Option) (Logger, error) {
	cfg := &config{}

	//  set default config
	WithDefaultConfig()(cfg)

	// apply options
	for _, opt := range opts {
		opt(cfg)
	}

	var (
		zapConfig zap.Config
	)

	switch cfg.AppMode {
	case AppModeDevelopment, AppModeStaging:
		zapConfig = zap.NewDevelopmentConfig()
	case AppModeProduction:
		zapConfig = zap.NewProductionConfig()
	default:
		return nil, errors.New("invalid app mode")
	}

	level, err := parseLevel(cfg.Level)
	if err != nil {
		return nil, err
	}

	zapConfig.Level = level
	zapConfig.Encoding = cfg.Encoding.String()
	zapConfig.EncoderConfig.TimeKey = cfg.TimeKey
	zapConfig.EncoderConfig.LevelKey = cfg.LevelKey
	zapConfig.EncoderConfig.NameKey = cfg.NameKey
	zapConfig.EncoderConfig.CallerKey = cfg.CallerKey
	zapConfig.EncoderConfig.MessageKey = cfg.MessageKey
	zapConfig.EncoderConfig.StacktraceKey = cfg.StacktraceKey
	zapConfig.DisableStacktrace = cfg.DisableStacktrace
	zapConfig.DisableCaller = cfg.DisableCaller
	zapConfig.OutputPaths = cfg.OutputPaths
	zapConfig.ErrorOutputPaths = cfg.ErrorOutputPaths
	zaplog, err := zapConfig.Build()
	if err != nil {
		return nil, err
	}

	if cfg.NewRelicApp != nil {
		backgroundCore, err := nrzap.WrapBackgroundCore(zaplog.Core(), cfg.NewRelicApp)
		if err != nil {
			return nil, err
		}

		zaplog = zap.New(backgroundCore)
	}

	return &logger{
		logger: &zapLogger{zapLogger: zaplog},
	}, nil
}

// GetLogger returns the underlying zap logger.
func (l *logger) GetLogger() *zap.Logger {
	return l.logger.GetLogger()
}

// Info logs a message at InfoLevel with optional structured fields.
// Use this for general informational messages about application flow.
func (l *logger) Info(ctx context.Context, msg string, fields ...Field) {
	l.logger.Info(ctx, msg, fields...)
}

// Warn logs a message at WarnLevel with optional structured fields.
// Use this for potentially harmful situations that don't stop the application.
func (l *logger) Warn(ctx context.Context, msg string, fields ...Field) {
	l.logger.Warn(ctx, msg, fields...)
}

// Error logs a message at ErrorLevel with optional structured fields.
// Use this for error events that might still allow the application to continue.
func (l *logger) Error(ctx context.Context, msg string, fields ...Field) {
	l.logger.Error(ctx, msg, fields...)
}

// Debug logs a message at DebugLevel with optional structured fields.
// Use this for detailed information useful during development and debugging.
func (l *logger) Debug(ctx context.Context, msg string, fields ...Field) {
	l.logger.Debug(ctx, msg, fields...)
}

// Panic logs a message at PanicLevel with optional structured fields, then panics.
// Use this for severe errors that should cause the application to panic.
func (l *logger) Panic(ctx context.Context, msg string, fields ...Field) {
	l.logger.Panic(ctx, msg, fields...)
}

// Fatal logs a message at FatalLevel with optional structured fields, then calls os.Exit(1).
// Use this for critical errors that should terminate the application.
func (l *logger) Fatal(ctx context.Context, msg string, fields ...Field) {
	l.logger.Fatal(ctx, msg, fields...)
}

// With creates a child logger with additional structured fields.
// The returned logger will include these fields in all subsequent log entries.
// Example: childLogger := logger.With(zap.String("component", "database"))
func (l *logger) With(fields ...Field) Logger {
	return l.logger.With(fields...)
}
