# Go Logger

A high-performance, structured logging library for Go applications built on top of [Uber's Zap](https://github.com/uber-go/zap). This logger provides a simple, consistent API with support for different output formats, log levels, context-aware logging, and New Relic integration.

## Features

- 🚀 **High Performance**: Built on Uber's Zap for maximum performance
- 📊 **Structured Logging**: JSON and console output formats
- 🎯 **Context-Aware**: Automatic extraction of context fields (trace ID, user ID, etc.)
- 🔧 **Configurable**: Flexible configuration with sensible defaults
- 📈 **New Relic Integration**: Built-in support for New Relic log forwarding
- 🛡️ **Type Safe**: Strongly typed configuration and field handling
- 🎨 **Multiple Environments**: Development, staging, and production presets

## Installation

```bash
go get github.com/andryhardiyanto/go-logger
```

## Quick Start

```go
package main

import (
    "context"
    "github.com/andryhardiyanto/go-logger"
    "go.uber.org/zap"
)

func main() {
    // Create a logger with default configuration
    log, err := logger.NewLogger()
    if err != nil {
        panic(err)
    }

    ctx := context.Background()
    
    // Basic logging
    log.Info(ctx, "Application started")
    log.Warn(ctx, "This is a warning")
    log.Error(ctx, "An error occurred", zap.String("error", "something went wrong"))
    
    // Create a child logger with additional fields
    requestLogger := log.With(zap.String("requestID", "12345"))
    requestLogger.Info(ctx, "Processing request")
}
```

## Configuration

### Basic Configuration

```go
log, err := logger.NewLogger(
    logger.WithLevel(logger.LevelInfo),
    logger.WithEncoding(logger.EncodingJson),
    logger.WithAppMode(logger.AppModeProduction),
)
```

### Advanced Configuration Examples

```go
// Using default configuration (recommended for quick setup)
log, err := logger.NewLogger(logger.WithDefaultConfig())

// Customizing field names
log, err := logger.NewLogger(
    logger.WithTimeKey("timestamp"),
    logger.WithLevelKey("severity"),
    logger.WithMessageKey("msg"),
    logger.WithCallerKey("source"),
    logger.WithStacktraceKey("stack"),
)

// Performance optimization (disable caller and stacktrace)
log, err := logger.NewLogger(
    logger.WithLevel(logger.LevelInfo),
    logger.WithDisableCaller(true),
    logger.WithDisableStacktrace(true),
)

// Custom output destinations
log, err := logger.NewLogger(
    logger.WithOutputPaths([]string{"stdout", "/var/log/app.log"}),
    logger.WithErrorOutputPaths([]string{"stderr", "/var/log/error.log"}),
)
```

### Available Options

| Option | Description | Values |
|--------|-------------|--------|
| `WithLevel` | Set minimum log level | `LevelDebug`, `LevelInfo`, `LevelWarning`, `LevelError`, `LevelPanic`, `LevelFatal` |
| `WithEncoding` | Set output format | `EncodingJson`, `EncodingConsole` |
| `WithAppMode` | Set application mode | `AppModeDevelopment`, `AppModeStaging`, `AppModeProduction` |
| `WithNewRelicApp` | Enable New Relic integration | `*newrelic.Application` |
| `WithDefaultConfig` | Apply sensible default configuration | No parameters |
| `WithTimeKey` | Customize timestamp field name | `string` (default: "time") |
| `WithLevelKey` | Customize log level field name | `string` (default: "level") |
| `WithNameKey` | Customize logger name field name | `string` (default: "logger") |
| `WithCallerKey` | Customize caller info field name | `string` (default: "caller") |
| `WithMessageKey` | Customize message field name | `string` (default: "message") |
| `WithStacktraceKey` | Customize stack trace field name | `string` (default: "stack_trace") |
| `WithDisableCaller` | Disable caller information | `true` or `false` |
| `WithDisableStacktrace` | Disable stack traces | `true` or `false` |
| `WithOutputPaths` | Set output destinations for normal logs | `[]string` (e.g., `["stdout", "/var/log/app.log"]`) |
| `WithErrorOutputPaths` | Set output destinations for error logs | `[]string` (e.g., `["stderr", "/var/log/error.log"]`) |

### Application Modes

- **Development**: Console encoding, debug level, caller info enabled
- **Staging**: JSON encoding, info level, balanced configuration
- **Production**: JSON encoding, warn level, optimized for performance

## Context-Aware Logging

The logger automatically extracts and includes context fields in log entries:

```go
ctx := context.WithValue(context.Background(), logger.ContextKeyUserID, "user123")
ctx = context.WithValue(ctx, logger.ContextKeyTraceID, "trace456")

log.Info(ctx, "User action performed")
// Output: {"level":"info","msg":"User action performed","user_id":"user123","trace_id":"trace456"}
```

### Supported Context Keys

- `user_id`: User identifier
- `trace_id`: Distributed tracing ID
- `span_id`: Span identifier
- `request_id`: HTTP request ID
- `entity_guid`: New Relic entity GUID
- `hostname`: Server hostname
- `application_name`: Application name
- `application_environment`: Environment name
- `accept_language`: Client language preference
- `user_context`: Additional user context
- `ip_address`: Client IP address

## Log Levels

```go
log.Debug(ctx, "Detailed debug information")
log.Info(ctx, "General information")
log.Warn(ctx, "Warning message")
log.Error(ctx, "Error occurred", zap.Error(err))
log.Panic(ctx, "Critical error - will panic")
log.Fatal(ctx, "Fatal error - will exit")
```

## Structured Fields

Add structured data to your logs using Zap fields:

```go
log.Info(ctx, "User login",
    zap.String("username", "john_doe"),
    zap.Int("attempt", 3),
    zap.Duration("duration", time.Since(start)),
    zap.Bool("success", true),
)
```

## Child Loggers

Create child loggers with additional context:

```go
// Create a component-specific logger
dbLogger := log.With(
    zap.String("component", "database"),
    zap.String("version", "1.2.3"),
)

// All logs from dbLogger will include these fields
dbLogger.Info(ctx, "Connection established")
dbLogger.Error(ctx, "Query failed", zap.String("query", "SELECT * FROM users"))
```

## New Relic Integration

```go
import "github.com/newrelic/go-agent/v3/newrelic"

app, err := newrelic.NewApplication(
    newrelic.ConfigAppName("My App"),
    newrelic.ConfigLicense("your-license-key"),
)
log, err := logger.NewLogger(
    logger.WithNewRelicApp(app),
    logger.WithAppMode(logger.AppModeProduction),
)
```

## Advanced Usage

### Custom Output Destinations

```go
log, err := logger.NewLogger(
    logger.WithOutputPaths([]string{
        "stdout",
        "/var/log/app.log",
    }),
    logger.WithErrorOutputPaths([]string{
        "stderr",
        "/var/log/app-error.log",
    }),
)
```

### Accessing Underlying Zap Logger

```go
zapLogger := log.GetLogger()
sugar := zapLogger.Sugar()
sugar.Infof("Formatted message: %s", value)
```

## Error Handling

The logger provides descriptive error messages for configuration issues:

```go
log, err := logger.NewLogger(
    logger.WithLevel("invalid-level"),
)
if err != nil {
    // Error: invalid logging level: 'invalid-level'. Supported levels are: debug, info, warning, error, panic, fatal. Please check your configuration and ensure you're using one of these valid levels
    fmt.Println(err)
}
```

## Performance Considerations

- Use appropriate log levels for different environments
- Prefer structured fields over string formatting
- Consider disabling debug logs in production
- Use child loggers to avoid repeating common fields

## Best Practices

1. **Use appropriate log levels**:
   - `Debug`: Detailed diagnostic information
   - `Info`: General application flow
   - `Warn`: Potentially harmful situations
   - `Error`: Error conditions that don't require termination
   - `Panic/Fatal`: Critical errors requiring immediate attention

2. **Include relevant context**:
   ```go
   log.Error(ctx, "Database connection failed",
       zap.String("database", "users"),
       zap.String("host", "db.example.com"),
       zap.Error(err),
   )
   ```

3. **Use structured fields instead of string formatting**:
   ```go
   // Good
   log.Info(ctx, "User created", zap.String("userID", userID))
   
   // Avoid
   log.Info(ctx, fmt.Sprintf("User %s created", userID))
   ```

4. **Create component-specific loggers**:
   ```go
   dbLogger := log.With(zap.String("component", "database"))
   apiLogger := log.With(zap.String("component", "api"))
   ```

## Examples

### Web Server with Request Logging

```go
import (
    "context"
    "net/http"
    "time"
    "github.com/google/uuid"
    "github.com/andryhardiyanto/go-logger"
    "go.uber.org/zap"
)

func loggingMiddleware(log logger.Logger) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            requestID := uuid.New().String()
            ctx := context.WithValue(r.Context(), logger.ContextKeyRequestID, requestID)
            
            requestLogger := log.With(
                zap.String("method", r.Method),
                zap.String("path", r.URL.Path),
                zap.String("remote_addr", r.RemoteAddr),
            )
            
            start := time.Now()
            requestLogger.Info(ctx, "Request started")
            
            next.ServeHTTP(w, r.WithContext(ctx))
            
            requestLogger.Info(ctx, "Request completed",
                zap.Duration("duration", time.Since(start)),
            )
        })
    }
}
```

### Database Operations

```go
import (
    "context"
    "database/sql"
    "github.com/andryhardiyanto/go-logger"
    "go.uber.org/zap"
)

type User struct {
    Email string
    // other fields...
}

type UserService struct {
    log logger.Logger
    db  *sql.DB
}

func NewUserService(log logger.Logger, db *sql.DB) *UserService {
    return &UserService{
        log: log.With(zap.String("component", "user_service")),
        db:  db,
    }
}

func (s *UserService) CreateUser(ctx context.Context, user *User) error {
    s.log.Info(ctx, "Creating user", zap.String("email", user.Email))
    
    if err := s.db.QueryRowContext(ctx, "INSERT INTO users (email) VALUES ($1)", user.Email).Err(); err != nil {
        s.log.Error(ctx, "Failed to create user",
            zap.String("email", user.Email),
            zap.Error(err),
        )
        return err
    }
    
    s.log.Info(ctx, "User created successfully", zap.String("email", user.Email))
    return nil
}
```

## API Reference

### Logger Interface

```go
type Logger interface {
    Info(ctx context.Context, msg string, fields ...Field)
    Warn(ctx context.Context, msg string, fields ...Field)
    Error(ctx context.Context, msg string, fields ...Field)
    Debug(ctx context.Context, msg string, fields ...Field)
    Fatal(ctx context.Context, msg string, fields ...Field)
    Panic(ctx context.Context, msg string, fields ...Field)
    With(fields ...Field) Logger
    GetLogger() *zap.Logger
}
```

### Configuration Options

```go
// Level options
const (
    LevelDebug   Level = "debug"
    LevelInfo    Level = "info"
    LevelWarning Level = "warning"
    LevelError   Level = "error"
    LevelPanic   Level = "panic"
    LevelFatal   Level = "fatal"
)

// Encoding options
const (
    EncodingJson    Encoding = "json"
    EncodingConsole Encoding = "console"
)

// Application modes
const (
    AppModeDevelopment AppMode = "development"
    AppModeStaging     AppMode = "staging"
    AppModeProduction  AppMode = "production"
)

// Context keys for automatic field extraction
const (
    ContextKeyUserID                 ContextKey = "user_id"
    ContextKeyTraceID                ContextKey = "trace_id"
    ContextKeySpanID                 ContextKey = "span_id"
    ContextKeyEntityGuid             ContextKey = "entity_guid"
    ContextKeyHostname               ContextKey = "hostname"
    ContextKeyApplicationName        ContextKey = "application_name"
    ContextKeyApplicationEnvironment ContextKey = "application_environment"
    ContextKeyRequestID              ContextKey = "request_id"
    ContextKeyAcceptLanguage         ContextKey = "accept_language"
    ContextKeyUserContext            ContextKey = "user_context"
    ContextKeyIpAddress              ContextKey = "ip_address"
)
```

### Helper Functions

```go
// GetLoggingFields extracts all logging fields from context
func GetLoggingFields(ctx context.Context) []LoggingField

// AppendContextKeys adds new context keys for automatic extraction
func AppendContextKeys(keys ...ContextKey)
```

## Dependencies

- [go.uber.org/zap](https://github.com/uber-go/zap) - High-performance logging
- [github.com/newrelic/go-agent/v3](https://github.com/newrelic/go-agent) - New Relic integration


## Contributing

We welcome contributions! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

## Support

If you encounter any issues or have questions, please open an issue on GitHub.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.