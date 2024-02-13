package logger

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

// Level represents log levels.
type Level int

// Define log levels as constants.
const (
	Off Level = iota
	Fatal
	Error
	Warn
	Info
	Audit
	Debug
	Trace
)

var (
	currentLevel Level        // Holds the current log level.
	mux          sync.RWMutex // Protects access to currentLevel.
)

var currentTime = func() string {
	return time.Now().Local().Format(time.DateTime)
}

func init() {
	// Initialize currentLevel with the value from the environment.
	currentLevel = Level(LogLevel())
}

// SetLevel updates the global log level to the provided level if it is valid.
func SetLevel(level Level) {
	mux.Lock()
	defer mux.Unlock()

	if level >= Off && level <= Trace {
		currentLevel = level
	}
}

// GetLevel retrieves the current global log level.
func GetLevel() Level {
	mux.RLock()
	defer mux.RUnlock()
	return currentLevel
}

// FatalLn logs a fatal message with the provided correlationId and message
// arguments. The application will exit after the message is logged.
func FatalLn(v ...any) {
	message := LogTextBuilder("[SENTINEL_FATAL]", v)
	SendLogMessage(message)
}

// ErrorLn logs an error message with the provided correlationId and message
// arguments if the current log level is Error or lower.
func ErrorLn(v ...any) {
	l := GetLevel()
	if l < Error {
		return
	}

	message := LogTextBuilder("[SENTINEL_ERROR]", v)
	SendLogMessage(message)
}

// WarnLn logs a warning message with the provided correlationId and message
// arguments if the current log level is Warn or lower.
func WarnLn(v ...any) {
	l := GetLevel()
	if l < Warn {
		return
	}

	message := LogTextBuilder("[SENTINEL_WARN]", v)
	SendLogMessage(message)
}

// InfoLn logs an informational message with the provided correlationId and
// message arguments if the current log level is Info or lower.
func InfoLn(correlationId *string, v ...any) {
	l := GetLevel()
	if l < Info {
		return
	}

	message := LogTextBuilder("[SENTINEL_INFO]", v)
	SendLogMessage(message)
}

// AuditLn logs an audit message with the provided correlationId and message
// arguments. Audit messages are always logged, regardless of the current log
// level.
func AuditLn(v ...any) {
	message := LogTextBuilder("[SENTINEL_AUDIT]", v)
	SendLogMessage(message)
}

// DebugLn logs a debug message with the provided correlationId and message
// arguments if the current log level is Debug or lower.
func DebugLn(v ...any) {
	l := GetLevel()
	if l < Debug {
		return
	}

	message := LogTextBuilder("[SENTINEL_DEBUG]", v)
	SendLogMessage(message)
}

// TraceLn logs a trace message with the provided correlationId and message
// arguments if the current log level is Trace or lower.
func TraceLn(v ...any) {
	l := GetLevel()
	if l < Trace {
		return
	}

	message := LogTextBuilder("[SENTINEL_TRACE]", v)
	SendLogMessage(message)
}

func LogTextBuilder(logHeader string, a ...any) string {
	logPrefix := fmt.Sprintf("%s[%s]", logHeader, currentTime())
	var messageParts []string

	for _, element := range a {
		messageParts = append(messageParts, fmt.Sprintf("%v", element))
	}

	message := strings.Join(messageParts, " ")
	finalLog := fmt.Sprintf("%s %s\n", logPrefix, message)

	return finalLog
}
