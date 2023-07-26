/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware, Inc.
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

var currentLevel = Level(env.LogLevel())
var mux sync.RWMutex

// SetLevel sets the global log level to the provided level.
//
// The log level is only updated if the provided level is valid (Off, Fatal,
// Error, Warn, Info, Audit, Debug, or Trace).
func SetLevel(l Level) {
	mux.Lock()
	defer mux.Unlock()
	if l < Off || l > Trace {
		return
	}
	currentLevel = l
}

// GetLevel returns the current global log level.
func GetLevel() Level {
	mux.RLock()
	defer mux.RUnlock()
	return currentLevel
}

// FatalLn logs a fatal message with the provided correlationId and message
// arguments. The application will exit after the message is logged.
func FatalLn(correlationId *string, v ...any) {
	var args []any
	args = append(args, "[FATAL]")
	args = append(args, *correlationId)
	args = append(args, v...)
	log.Fatalln(args...)
}

// ErrorLn logs an error message with the provided correlationId and message
// arguments if the current log level is Error or lower.
func ErrorLn(correlationId *string, v ...any) {
	l := GetLevel()
	if l < Error {
		return
	}

	var args []any
	args = append(args, "[ERROR]")
	args = append(args, *correlationId)
	args = append(args, v...)
	log.Println(args...)
}

// WarnLn logs a warning message with the provided correlationId and message
// arguments if the current log level is Warn or lower.
func WarnLn(correlationId *string, v ...any) {
	l := GetLevel()
	if l < Warn {
		return
	}

	var args []any
	args = append(args, "[WARN ]")
	args = append(args, *correlationId)
	args = append(args, v...)
	log.Println(args...)
}

// InfoLn logs an informational message with the provided correlationId and
// message arguments if the current log level is Info or lower.
func InfoLn(correlationId *string, v ...any) {
	l := GetLevel()
	if l < Info {
		return
	}

	var args []any
	args = append(args, "[INFO ]")
	args = append(args, *correlationId)
	args = append(args, v...)
	log.Println(args...)
}

// AuditLn logs an audit message with the provided correlationId and message
// arguments. Audit messages are always logged, regardless of the current log
// level.
func AuditLn(correlationId *string, v ...any) {
	// Audit is always logged.
	var args []any
	args = append(args, "[AUDIT]")
	args = append(args, *correlationId)
	args = append(args, v...)
	log.Println(args...)
}

// DebugLn logs a debug message with the provided correlationId and message
// arguments if the current log level is Debug or lower.
func DebugLn(correlationId *string, v ...any) {
	l := GetLevel()
	if l < Debug {
		return
	}

	var args []any
	args = append(args, "[DEBUG]")
	args = append(args, *correlationId)
	args = append(args, v...)
	log.Println(args...)
}

// TraceLn logs a trace message with the provided correlationId and message
// arguments if the current log level is Trace or lower.
func TraceLn(correlationId *string, v ...any) {
	l := GetLevel()
	if l < Trace {
		return
	}

	var args []any
	args = append(args, "[TRACE]")
	args = append(args, *correlationId)
	args = append(args, v...)
	log.Println(args...)
}
