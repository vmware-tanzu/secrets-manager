/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package level

import (
	"sync"

	"github.com/vmware-tanzu/secrets-manager/sdk/core/env"
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

// Get retrieves the current global log level.
func Get() Level {
	mux.RLock()
	defer mux.RUnlock()

	return currentLevel
}
