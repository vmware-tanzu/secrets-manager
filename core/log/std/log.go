/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

// Package std provides a simple and flexible logging library with various
// log levels.
package std

import (
	"github.com/vmware-tanzu/secrets-manager/core/log"
	stdlib "log"
)

// logMessage logs a message with the specified level, correlation ID, and
// message arguments. It checks the current log level to decide if the message
// should be logged.
func logMessage(level log.Level, prefix string, correlationID *string, v ...any) {
	if level != log.Audit && log.GetLevel() < level {
		return
	}

	args := make([]any, 0, len(v)+2)
	args = append(args, prefix)
	if correlationID != nil {
		args = append(args, *correlationID)
	}
	args = append(args, v...)

	if level == log.Fatal {
		stdlib.Fatalln(args...)
	} else {
		stdlib.Println(args...)
	}
}

// FatalLn logs a fatal level message and exits.
func FatalLn(correlationID *string, v ...any) {
	logMessage(log.Fatal, "[FATAL]", correlationID, v...)
}

// ErrorLn logs an error level message.
func ErrorLn(correlationID *string, v ...any) {
	logMessage(log.Error, "[ERROR]", correlationID, v...)
}

// WarnLn logs a warning level message.
func WarnLn(correlationID *string, v ...any) {
	logMessage(log.Warn, "[WARN]", correlationID, v...)
}

// InfoLn logs an info level message.
func InfoLn(correlationID *string, v ...any) {
	logMessage(log.Info, "[INFO]", correlationID, v...)
}

// AuditLn logs an audit level message.
func AuditLn(correlationID *string, v ...any) {
	logMessage(log.Audit, "[AUDIT]", correlationID, v...)
}

// DebugLn logs a debug level message.
func DebugLn(correlationID *string, v ...any) {
	logMessage(log.Debug, "[DEBUG]", correlationID, v...)
}

// TraceLn logs a trace level message.
func TraceLn(correlationID *string, v ...any) {
	logMessage(log.Trace, "[TRACE]", correlationID, v...)
}
