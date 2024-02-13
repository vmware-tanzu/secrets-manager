/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

// Package log provides a simple and flexible logging library with various
// log levels.
package log

import (
	"github.com/vmware-tanzu/secrets-manager/core/env"
	"log"
	"sync"
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

func init() {
	// Initialize currentLevel with the value from the environment.
	currentLevel = Level(env.LogLevel())
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

// logMessage logs a message with the specified level, correlation ID, and message arguments.
// It checks the current log level to decide if the message should be logged.
func logMessage(level Level, prefix string, correlationID *string, v ...any) {
	if level != Audit && GetLevel() < level {
		return
	}

	args := make([]any, 0, len(v)+2)
	args = append(args, prefix)
	if correlationID != nil {
		args = append(args, *correlationID)
	}
	args = append(args, v...)

	if level == Fatal {
		log.Fatalln(args...)
	} else {
		log.Println(args...)
	}
}

// FatalLn logs a fatal level message and exits.
func FatalLn(correlationID *string, v ...any) {
	logMessage(Fatal, "[FATAL]", correlationID, v...)
}

// ErrorLn logs an error level message.
func ErrorLn(correlationID *string, v ...any) {
	logMessage(Error, "[ERROR]", correlationID, v...)
}

// WarnLn logs a warning level message.
func WarnLn(correlationID *string, v ...any) {
	logMessage(Warn, "[WARN]", correlationID, v...)
}

// InfoLn logs an info level message.
func InfoLn(correlationID *string, v ...any) {
	logMessage(Info, "[INFO]", correlationID, v...)
}

// AuditLn logs an audit level message.
func AuditLn(correlationID *string, v ...any) {
	logMessage(Audit, "[AUDIT]", correlationID, v...)
}

// DebugLn logs a debug level message.
func DebugLn(correlationID *string, v ...any) {
	logMessage(Debug, "[DEBUG]", correlationID, v...)
}

// TraceLn logs a trace level message.
func TraceLn(correlationID *string, v ...any) {
	logMessage(Trace, "[TRACE]", correlationID, v...)
}
