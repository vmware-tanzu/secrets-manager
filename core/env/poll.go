/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware, Inc.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package env

import (
	"os"
	"strconv"
	"time"
)

// SidecarMaxPollInterval returns the maximum interval for polling by the
// sidecar process. The value is read from the environment variable
// `VSECM_SIDECAR_MAX_POLL_INTERVAL` or returns 300000 milliseconds as default.
func SidecarMaxPollInterval() time.Duration {
	p := os.Getenv("VSECM_SIDECAR_MAX_POLL_INTERVAL")
	if p == "" {
		p = "300000"
	}
	i, err := strconv.ParseInt(p, 10, 32)
	if err != nil {
		return 300000 * time.Millisecond
	}
	return time.Duration(i) * time.Millisecond
}

// SidecarExponentialBackoffMultiplier returns the multiplier for exponential
// backoff by the sidecar process.
// The value is read from the environment variable
// `VSECM_SIDECAR_EXPONENTIAL_BACKOFF_MULTIPLIER` or returns 2 as default.
func SidecarExponentialBackoffMultiplier() int64 {
	p := os.Getenv("VSECM_SIDECAR_EXPONENTIAL_BACKOFF_MULTIPLIER")
	if p == "" {
		p = "2"
	}
	i, err := strconv.ParseInt(p, 10, 32)
	if err != nil {
		return 2
	}
	return i
}

// SidecarSuccessThreshold returns the number of consecutive successful
// polls before reducing the interval. The value is read from the environment
// variable `VSECM_SIDECAR_SUCCESS_THRESHOLD` or returns 3 as default.
func SidecarSuccessThreshold() int64 {
	p := os.Getenv("VSECM_SIDECAR_SUCCESS_THRESHOLD")
	if p == "" {
		p = "3"
	}
	i, err := strconv.ParseInt(p, 10, 32)
	if err != nil {
		return 3
	}
	return i
}

// SidecarErrorThreshold returns the number of consecutive failed polls
// before increasing the interval. The value is read from the environment
// variable `VSECM_SIDECAR_ERROR_THRESHOLD` or returns 2 as default.
func SidecarErrorThreshold() int64 {
	p := os.Getenv("VSECM_SIDECAR_ERROR_THRESHOLD")
	if p == "" {
		p = "2"
	}
	i, err := strconv.ParseInt(p, 10, 32)
	if err != nil {
		return 2
	}
	return i
}

// SidecarPollInterval returns the polling interval for sentry in time.Duration
// The interval is determined by the VSECM_SIDECAR_POLL_INTERVAL environment
// variable, with a default value of 20000 milliseconds if the variable is not
// set or if there is an error in parsing the value.
func SidecarPollInterval() time.Duration {
	p := os.Getenv("VSECM_SIDECAR_POLL_INTERVAL")
	if p == "" {
		p = "20000"
	}
	i, err := strconv.ParseInt(p, 10, 32)
	if err != nil {
		return 20000 * time.Millisecond
	}
	return time.Duration(i) * time.Millisecond
}
