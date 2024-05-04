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

import "github.com/vmware-tanzu/secrets-manager/core/log/level"

// FatalLn logs a fatal level message and exits.
func FatalLn(correlationID *string, v ...any) {
	logMessage(level.Fatal, "[FATAL]", correlationID, v...)
}

// ErrorLn logs an error level message.
func ErrorLn(correlationID *string, v ...any) {
	logMessage(level.Error, "[ERROR]", correlationID, v...)
}

// WarnLn logs a warning level message.
func WarnLn(correlationID *string, v ...any) {
	logMessage(level.Warn, "[WARN]", correlationID, v...)
}

// InfoLn logs an info level message.
func InfoLn(correlationID *string, v ...any) {
	logMessage(level.Info, "[INFO]", correlationID, v...)
}

// AuditLn logs an audit level message.
func AuditLn(correlationID *string, v ...any) {
	logMessage(level.Audit, "[AUDIT]", correlationID, v...)
}

// DebugLn logs a debug level message.
func DebugLn(correlationID *string, v ...any) {
	logMessage(level.Debug, "[DEBUG]", correlationID, v...)
}

// TraceLn logs a trace level message.
func TraceLn(correlationID *string, v ...any) {
	logMessage(level.Trace, "[TRACE]", correlationID, v...)
}
