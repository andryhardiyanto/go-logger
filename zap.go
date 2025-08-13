package logger

import (
	"context"

	"go.uber.org/zap"
)

// zapLogger is the concrete implementation of the Logger interface using zap.
// It wraps a zap.Logger instance and provides the standardized logging methods.
// This implementation is optimized for performance and provides structured logging capabilities.
type zapLogger struct {
	zapLogger *zap.Logger
}

// Info logs a message at InfoLevel using the underlying zap logger.
// This method is optimized for performance and supports structured logging.
// Fields are added as key-value pairs to the log entry for better searchability.
func (z *zapLogger) Info(ctx context.Context, msg string, fields ...Field) {
	z.zapLogger.With(z.extractTrace(ctx)...).Info(msg, fields...)
}

// Warn logs a message at WarnLevel using the underlying zap logger.
// Use this for potentially harmful situations that are not errors.
// The message and fields are structured for easy parsing and analysis.
func (z *zapLogger) Warn(ctx context.Context, msg string, fields ...Field) {
	z.zapLogger.With(z.extractTrace(ctx)...).Warn(msg, fields...)
}

// Error logs a message at ErrorLevel using the underlying zap logger.
// This method should be used for error conditions that don't require immediate termination.
// Structured fields help with error tracking and debugging.
func (z *zapLogger) Error(ctx context.Context, msg string, fields ...Field) {
	z.zapLogger.With(z.extractTrace(ctx)...).Error(msg, fields...)
}

// Debug logs a message at DebugLevel using the underlying zap logger.
// Debug messages are typically disabled in production for performance.
// Use this for detailed diagnostic information during development.
func (z *zapLogger) Debug(ctx context.Context, msg string, fields ...Field) {
	z.zapLogger.With(z.extractTrace(ctx)...).Debug(msg, fields...)
}

// Panic logs a message at PanicLevel using the underlying zap logger, then panics.
// This method should only be used for severe errors that require immediate attention.
// The application will terminate after logging the message.
func (z *zapLogger) Panic(ctx context.Context, msg string, fields ...Field) {
	z.zapLogger.With(z.extractTrace(ctx)...).Panic(msg, fields...)
}

// Fatal logs a message at FatalLevel using the underlying zap logger, then calls os.Exit(1).
// This method should be used for critical errors that require application termination.
// The application will exit immediately after logging the message.
func (z *zapLogger) Fatal(ctx context.Context, msg string, fields ...Field) {
	z.zapLogger.With(z.extractTrace(ctx)...).Fatal(msg, fields...)
}

// With creates a new logger instance with additional structured fields.
// The returned logger will include the provided fields in all subsequent log entries.
// This is useful for adding context-specific information like request IDs or user IDs.
//
// Example:
//   requestLogger := logger.With(zap.String("requestID", "12345"))
//   requestLogger.Info("Processing request") // Will include requestID field
func (z *zapLogger) With(fields ...Field) Logger {
	return &zapLogger{zapLogger: z.zapLogger.With(fields...)}
}

// GetLogger returns the underlying zap.Logger instance.
// This method provides access to advanced zap features not exposed by the Logger interface.
// Use with caution as it bypasses the standardized logging interface.
//
// Example:
//   zapLogger := logger.GetLogger()
//   zapLogger.Sugar().Infof("Formatted message: %s", value)
func (z *zapLogger) GetLogger() *zap.Logger {
	return z.zapLogger
}
func (z *zapLogger) extractTrace(ctx context.Context) []Field {
	var fields []Field
	ctxFields := GetLoggingFields(ctx)
	for _, field := range ctxFields {
		fields = append(fields, zap.String(field.Key.String(), field.Value))
	}

	return fields
}
