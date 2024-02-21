/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package log

import (
	"github.com/vmware-tanzu/secrets-manager/core/env"
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

var mux sync.RWMutex // Protects access to currentLevel.

// Initialize currentLevel with the value from the environment.
var currentLevel = Level(env.LogLevel())

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
