package logger

import (
	"context"
	"errors"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// parseLevel converts a custom Level type to zap.AtomicLevel.
// It validates the input level and returns the corresponding zap atomic level.
//
// Parameters:
//   - level: The logging level to convert (debug, info, warning, error, panic, fatal)
//
// Returns:
//   - zap.AtomicLevel: The converted zap atomic level
//   - error: An error if the level is invalid or unsupported
//
// Example:
//   zapLevel, err := parseLevel(LevelInfo)
//   if err != nil {
//       log.Fatal("Failed to parse level:", err)
//   }
func parseLevel(level Level) (zap.AtomicLevel, error) {
	var zapLevel zap.AtomicLevel

	if level == "" {
		return zapLevel, errors.New("logging level cannot be empty. Please specify one of: debug, info, warning, error, panic, fatal")
	}

	switch level {
	case LevelDebug:
		zapLevel = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	case LevelInfo:
		zapLevel = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	case LevelWarning:
		zapLevel = zap.NewAtomicLevelAt(zapcore.WarnLevel)
	case LevelError:
		zapLevel = zap.NewAtomicLevelAt(zapcore.ErrorLevel)
	case LevelPanic:
		zapLevel = zap.NewAtomicLevelAt(zapcore.PanicLevel)
	case LevelFatal:
		zapLevel = zap.NewAtomicLevelAt(zapcore.FatalLevel)
	default:
		return zapLevel, errors.New("invalid logging level: '" + string(level) + "'. Supported levels are: debug, info, warning, error, panic, fatal. Please check your configuration and ensure you're using one of these valid levels")
	}

	return zapLevel, nil
}

// getStringFromContext safely extracts a string value from context using the provided key.
// This function performs safe type assertion to prevent runtime panics when extracting context values.
//
// Parameters:
//   - ctx: The context to extract the value from
//   - key: The context key to look for
//
// Returns:
//   - string: The extracted string value (empty string if not found or not a string)
//   - bool: True if the key was found and the value is a string, false otherwise
//
// Example:
//   userID, found := getStringFromContext(ctx, ContextKeyUserID)
//   if found {
//       log.Info("User ID found:", userID)
//   }
func getStringFromContext(ctx context.Context, key ContextKey) (string, bool) {
	if ctx == nil {
		return "", false
	}

	if key == "" {
		return "", false
	}

	val := ctx.Value(key)
	if val == nil {
		return "", false
	}

	if str, ok := val.(string); ok {
		return str, true
	}
	return "", false
}

// GetLoggingFields extracts all logging-related fields from the provided context.
// This function iterates through all predefined context keys and safely extracts their string values.
// It's designed to be used internally by the logger to automatically include context information in log entries.
//
// Parameters:
//   - ctx: The context to extract logging fields from
//
// Returns:
//   - []LoggingField: A slice containing all found key-value pairs from the context.
//                     Returns empty slice if context is nil or no fields found.
//
// Example:
//   fields := GetLoggingFields(ctx)
//   for _, field := range fields {
//       fmt.Printf("Key: %s, Value: %s\n", field.Key, field.Value)
//   }
//
// Note: This function is safe to call with nil context and will not panic.
func GetLoggingFields(ctx context.Context) []LoggingField {
	if ctx == nil {
		return make([]LoggingField, 0)
	}

	// Pre-allocate slice with estimated capacity to reduce memory allocations
	fields := make([]LoggingField, 0, len(contextKeys))

	for _, key := range contextKeys {
		if key == "" {
			continue // Skip empty keys to avoid potential issues
		}

		if value, ok := getStringFromContext(ctx, key); ok {
			fields = append(fields, LoggingField{
				Key:   key,
				Value: value,
			})
		}
	}

	return fields
}

func AppendContextKeys(keys ...ContextKey) {
	contextKeys = append(contextKeys, keys...)
}
