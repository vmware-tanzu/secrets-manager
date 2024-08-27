/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package env

import (
	"github.com/vmware-tanzu/secrets-manager/sdk/core/constants/env"
	"strconv"
)

type Level int

// Redefine log levels to avoid import cycle.
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

var level = struct {
	Off   Level
	Fatal Level
	Error Level
	Warn  Level
	Info  Level
	Audit Level
	Debug Level
	Trace Level
}{
	Off:   Off,
	Fatal: Fatal,
	Error: Error,
	Warn:  Warn,
	Info:  Info,
	Audit: Audit,
	Debug: Debug,
	Trace: Trace,
}

// LogLevel returns the value set by VSECM_LOG_LEVEL environment
// variable, or a default level.
//
// VSECM_LOG_LEVEL determines the verbosity of the logs.
// 0: logs are off, 7: highest verbosity (TRACE).
func LogLevel() int {
	p := env.Value(env.VSecMLogLevel)
	if p == "" {
		return int(level.Warn)
	}

	l, _ := strconv.Atoi(p)
	if l == int(level.Off) {
		return int(level.Warn)
	}

	if l < int(level.Off) || l > int(level.Trace) {
		return int(level.Warn)
	}

	return l
}
