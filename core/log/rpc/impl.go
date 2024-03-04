/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package rpc

import (
	core "github.com/vmware-tanzu/secrets-manager/core/log"
)

// FatalLn logs a fatal message with the provided correlationId and message
// arguments. The application will exit after the message is logged.
func FatalLn(correlationId *string, v ...any) {
	message := build("[FATAL]", correlationId, v)
	log(message)
}

// ErrorLn logs an error message with the provided correlationId and message
// arguments if the current log level is Error or lower.
func ErrorLn(correlationId *string, v ...any) {
	l := core.GetLevel()
	if l < core.Error {
		return
	}

	message := build("[ERROR]", correlationId, v)
	log(message)
}

// WarnLn logs a warning message with the provided correlationId and message
// arguments if the current log level is Warn or lower.
func WarnLn(correlationId *string, v ...any) {
	l := core.GetLevel()
	if l < core.Warn {
		return
	}

	message := build("[WARN]", correlationId, v)
	log(message)
}

// InfoLn logs an informational message with the provided correlationId and
// message arguments if the current log level is Info or lower.
func InfoLn(correlationId *string, v ...any) {
	l := core.GetLevel()
	if l < core.Info {
		return
	}

	message := build("[INFO]", correlationId, v)
	log(message)
}

// AuditLn logs an audit message with the provided correlationId and message
// arguments. Audit messages are always logged, regardless of the current log
// level.
func AuditLn(correlationId *string, v ...any) {
	message := build("[AUDIT]", correlationId, v)
	log(message)
}

// DebugLn logs a debug message with the provided correlationId and message
// arguments if the current log level is Debug or lower.
func DebugLn(correlationId *string, v ...any) {
	l := core.GetLevel()
	if l < core.Debug {
		return
	}

	message := build("[DEBUG]", correlationId, v)
	log(message)
}

// TraceLn logs a trace message with the provided correlationId and message
// arguments if the current log level is Trace or lower.
func TraceLn(correlationId *string, v ...any) {
	l := core.GetLevel()
	if l < core.Trace {
		return
	}

	message := build("[TRACE]", correlationId, v)
	log(message)
}
