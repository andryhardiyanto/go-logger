// Package logger provides configuration options for the logging system.
// This file defines the configuration structure and option functions
// that allow flexible logger setup and customization.
package logger

import "github.com/newrelic/go-agent/v3/newrelic"

type (
	// Option represents a configuration option function.
	// Options are applied during logger creation to customize behavior.
	// This pattern allows for flexible and extensible configuration.
	Option func(*config)

	// config holds the internal configuration for the logger.
	// This struct contains all the settings that control logger behavior,
	// including output format, log levels, and integration settings.
	config struct {
		// Level sets the minimum log level that will be output.
		// Messages below this level will be filtered out.
		Level Level
		// Encoding determines the output format (JSON or console).
		// JSON is recommended for production, console for development.
		Encoding Encoding
		// AppMode sets the application environment mode.
		// Different modes have different default configurations.
		AppMode AppMode
		// NewRelicApp enables New Relic integration when provided.
		// Allows automatic forwarding of logs to New Relic for monitoring.
		NewRelicApp *newrelic.Application
		// TimeKey specifies the key name for timestamp in log output.
		TimeKey string
		// LevelKey specifies the key name for log level in log output.
		LevelKey string
		// NameKey specifies the key name for logger name in log output.
		NameKey string
		// CallerKey specifies the key name for caller info in log output.
		CallerKey string
		// MessageKey specifies the key name for log message in log output.
		MessageKey string
		// StacktraceKey specifies the key name for stack trace in log output.
		StacktraceKey string
		// DisableStacktrace controls whether stack traces are included.
		DisableStacktrace bool
		// DisableCaller controls whether caller information is included.
		DisableCaller bool
		// OutputPaths specifies where normal log messages are written.
		OutputPaths []string
		// ErrorOutputPaths specifies where error-level messages are written.
		ErrorOutputPaths []string
	}
)

// WithLevel sets the minimum log level for the logger.
// Messages below this level will be filtered out to reduce noise.
//
// Parameters:
//   - level: The minimum log level (debug, info, warning, error, panic, fatal)
//
// Example:
//
//	logger := NewLogger(WithLevel(LevelInfo)) // Only info and above
func WithLevel(level Level) Option {
	return func(c *config) {
		c.Level = level
	}
}

// WithEncoding sets the output format for log messages.
// Choose between human-readable console format or machine-readable JSON.
//
// Parameters:
//   - encoding: The output format (EncodingConsole or EncodingJson)
//
// Example:
//
//	logger := NewLogger(WithEncoding(EncodingJson)) // JSON output
func WithEncoding(encoding Encoding) Option {
	return func(c *config) {
		c.Encoding = encoding
	}
}

// WithAppMode sets the application environment mode.
// Different modes have optimized defaults for their use cases.
//
// Parameters:
//   - appMode: The application mode (development, staging, production)
//
// Example:
//
//	logger := NewLogger(WithAppMode(AppModeProduction)) // Production optimized
func WithAppMode(appMode AppMode) Option {
	return func(c *config) {
		c.AppMode = appMode
	}
}

// WithNewRelicApp enables New Relic integration for log forwarding.
// When provided, logs will be automatically sent to New Relic for monitoring.
//
// Parameters:
//   - newRelicApp: The New Relic application instance
//
// Example:
//
//	app, _ := newrelic.NewApplication(config)
//	logger := NewLogger(WithNewRelicApp(app))
func WithNewRelicApp(newRelicApp *newrelic.Application) Option {
	return func(c *config) {
		c.NewRelicApp = newRelicApp
	}
}

// WithDefaultConfig applies sensible default configuration values.
// This provides a good starting point for most applications.
//
// Default values:
//   - Level: Info (filters out debug messages)
//   - Encoding: JSON (machine-readable format)
//   - AppMode: Development (includes caller info)
//   - Output: stdout for logs, stderr for errors
//   - Stacktrace and caller info: disabled for performance
//
// Example:
//
//	logger := NewLogger(WithDefaultConfig()) // Use recommended defaults
func WithDefaultConfig() Option {
	return func(c *config) {
		c.Level = LevelInfo
		c.Encoding = EncodingJson
		c.AppMode = AppModeDevelopment
		c.TimeKey = "time"
		c.LevelKey = "level"
		c.NameKey = "logger"
		c.CallerKey = "caller"
		c.MessageKey = "message"
		c.StacktraceKey = "stack_trace"
		c.DisableStacktrace = true
		c.DisableCaller = true
		c.OutputPaths = []string{"stdout"}
		c.ErrorOutputPaths = []string{"stderr"}
	}
}

// WithTimeKey sets the key name for timestamp in log output.
// Useful for customizing log format to match existing systems.
//
// Parameters:
//   - timeKey: The key name for timestamp field
//
// Example:
//
//	logger := NewLogger(WithTimeKey("timestamp")) // Use "timestamp" instead of "time"
func WithTimeKey(timeKey string) Option {
	return func(c *config) {
		c.TimeKey = timeKey
	}
}

// WithLevelKey sets the key name for log level in log output.
// Allows customization of the field name for log levels.
//
// Parameters:
//   - levelKey: The key name for log level field
//
// Example:
//
//	logger := NewLogger(WithLevelKey("severity")) // Use "severity" instead of "level"
func WithLevelKey(levelKey string) Option {
	return func(c *config) {
		c.LevelKey = levelKey
	}
}

// WithNameKey sets the key name for logger name in log output.
// Useful when you want to identify different loggers in the same application.
//
// Parameters:
//   - nameKey: The key name for logger name field
//
// Example:
//
//	logger := NewLogger(WithNameKey("component")) // Use "component" instead of "logger"
func WithNameKey(nameKey string) Option {
	return func(c *config) {
		c.NameKey = nameKey
	}
}

// WithCallerKey sets the key name for caller information in log output.
// Controls the field name for file and line number information.
//
// Parameters:
//   - callerKey: The key name for caller info field
//
// Example:
//
//	logger := NewLogger(WithCallerKey("source")) // Use "source" instead of "caller"
func WithCallerKey(callerKey string) Option {
	return func(c *config) {
		c.CallerKey = callerKey
	}
}

// WithMessageKey sets the key name for log message in log output.
// Allows customization of the main message field name.
//
// Parameters:
//   - messageKey: The key name for message field
//
// Example:
//
//	logger := NewLogger(WithMessageKey("msg")) // Use "msg" instead of "message"
func WithMessageKey(messageKey string) Option {
	return func(c *config) {
		c.MessageKey = messageKey
	}
}

// WithStacktraceKey sets the key name for stack trace in log output.
// Controls the field name for stack trace information in error logs.
//
// Parameters:
//   - stacktraceKey: The key name for stack trace field
//
// Example:
//
//	logger := NewLogger(WithStacktraceKey("stack")) // Use "stack" instead of "stack_trace"
func WithStacktraceKey(stacktraceKey string) Option {
	return func(c *config) {
		c.StacktraceKey = stacktraceKey
	}
}

// WithDisableStacktrace controls whether stack traces are included in logs.
// Stack traces are useful for debugging but add overhead and verbosity.
//
// Parameters:
//   - disableStacktrace: true to disable stack traces, false to enable
//
// Example:
//
//	logger := NewLogger(WithDisableStacktrace(false)) // Enable stack traces for debugging
func WithDisableStacktrace(disableStacktrace bool) Option {
	return func(c *config) {
		c.DisableStacktrace = disableStacktrace
	}
}

// WithDisableCaller controls whether caller information is included in logs.
// Caller info shows file name and line number but adds performance overhead.
//
// Parameters:
//   - disableCaller: true to disable caller info, false to enable
//
// Example:
//
//	logger := NewLogger(WithDisableCaller(false)) // Enable caller info for debugging
func WithDisableCaller(disableCaller bool) Option {
	return func(c *config) {
		c.DisableCaller = disableCaller
	}
}

// WithOutputPaths sets the output destinations for normal log messages.
// Supports multiple outputs including files, stdout, stderr, and URLs.
//
// Parameters:
//   - outputPaths: slice of output destinations
//
// Example:
//
//	logger := NewLogger(WithOutputPaths([]string{"stdout", "/var/log/app.log"}))
func WithOutputPaths(outputPaths []string) Option {
	return func(c *config) {
		c.OutputPaths = outputPaths
	}
}

// WithErrorOutputPaths sets the output destinations for error-level log messages.
// Typically used to separate error logs from normal logs.
//
// Parameters:
//   - errorOutputPaths: slice of error output destinations
//
// Example:
//
//	logger := NewLogger(WithErrorOutputPaths([]string{"stderr", "/var/log/error.log"}))
func WithErrorOutputPaths(errorOutputPaths []string) Option {
	return func(c *config) {
		c.ErrorOutputPaths = errorOutputPaths
	}
}
