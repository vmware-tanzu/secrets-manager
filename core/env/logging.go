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
	"os"
	"strconv"
)

// LogLevel returns the value set by VSECM_LOG_LEVEL environment
// variable, or a default level.
//
// VSECM_LOG_LEVEL determines the verbosity of the logs.
// 0: logs are off, 7: highest verbosity (TRACE).
func LogLevel() int {
	p := os.Getenv("VSECM_LOG_LEVEL")
	if p == "" {
		return 3 // WARN
	}
	l, _ := strconv.Atoi(p)
	if l == 0 {
		return 3 // WARN
	}
	if l < 0 || l > 7 {
		return 3 // WARN
	}
	return l
}
