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

import "os"

// ProbeLivenessPort returns the port for liveness probe.
// It first checks the environment variable VSECM_PROBE_LIVENESS_PORT.
// If the variable is not set, it returns the default value ":8081".
func ProbeLivenessPort() string {
	u := os.Getenv("VSECM_PROBE_LIVENESS_PORT")
	if u == "" {
		u = ":8081"
	}
	return u
}

// ProbeReadinessPort returns the port for readiness probe.
// It first checks the environment variable VSECM_PROBE_READINESS_PORT.
// If the variable is not set, it returns the default value ":8082".
func ProbeReadinessPort() string {
	u := os.Getenv("VSECM_PROBE_READINESS_PORT")
	if u == "" {
		u = ":8082"
	}
	return u
}
