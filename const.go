package logger

import (
	"go.uber.org/zap"
)

type (
	AppMode    string
	Encoding   string
	Level      string
	ContextKey string
	Field      = zap.Field

	LoggingField struct {
		Key   ContextKey
		Value string
	}
)

const (
	AppModeDevelopment AppMode = "development"
	AppModeStaging     AppMode = "staging"
	AppModeProduction  AppMode = "production"
	AppModeEmpty       AppMode = "empty"

	EncodingJson    Encoding = "json"
	EncodingConsole Encoding = "console"

	LevelDebug   Level = "debug"
	LevelInfo    Level = "info"
	LevelWarning Level = "warning"
	LevelError   Level = "error"
	LevelPanic   Level = "panic"
	LevelFatal   Level = "fatal"

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

// ContextKeys contains all predefined context keys for automatic field extraction.
// This slice is used by GetLoggingFields to iterate through all possible context values.
// Add new context keys to this slice to enable automatic extraction.
var contextKeys = []ContextKey{
	ContextKeyUserID,
	ContextKeyTraceID,
	ContextKeySpanID,
	ContextKeyEntityGuid,
	ContextKeyHostname,
	ContextKeyApplicationName,
	ContextKeyApplicationEnvironment,
	ContextKeyRequestID,
	ContextKeyAcceptLanguage,
	ContextKeyUserContext,
	ContextKeyIpAddress,
}

var (
	validAppMode = map[AppMode]bool{
		AppModeDevelopment: true,
		AppModeProduction:  true,
		AppModeStaging:     true,
	}
	validEncoding = map[Encoding]bool{
		EncodingConsole: true,
		EncodingJson:    true,
	}
)

// String returns the string representation of AppMode.
// Returns AppModeEmpty if the mode is not valid.
// Implements the Stringer interface for better debugging and logging.
func (am AppMode) String() string {
	if validAppMode[am] {
		return string(am)
	}
	return string(AppModeEmpty)
}

// String returns the string representation of Encoding.
// Returns EncodingConsole if the encoding is not valid.
// Implements the Stringer interface for better debugging and logging.
func (e Encoding) String() string {
	if validEncoding[e] {
		return string(e)
	}
	return string(EncodingConsole)
}

// String returns the string representation of ContextKey.
// Implements the Stringer interface for better debugging and logging.
func (ck ContextKey) String() string {
	return string(ck)
}
