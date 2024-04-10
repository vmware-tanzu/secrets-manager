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
	"strings"
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

// LogSecretFingerprints checks the "VSECM_LOG_SECRET_FINGERPRINTS" environment variable,
// normalizes its value by trimming whitespace and converting it to lowercase, and
// evaluates whether logging of secret fingerprints is enabled or not. The function
// returns true if the environment variable is explicitly set to "true", otherwise,
// it defaults to false.
//
// When `true`, VSecM logs will include partial hashes for the secrets. This
// approach will be useful to verify changes to a secret without revealing it
// in the logs. The partial hash is a cryptographically secure string, and there
// is no way to retrieve the original secret from it.
//
// If not provided in the environment variables, this flag will be set to `false`
// by default.
//
// Returns:
// bool - true if logging of secret fingerprints is enabled, false otherwise.
func LogSecretFingerprints() bool {
	p := os.Getenv("VSECM_LOG_SECRET_FINGERPRINTS")
	p = strings.ToLower(strings.TrimSpace(p))
	if p == "" {
		return false
	}
	return p == "true"
}
