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
	"github.com/vmware-tanzu/secrets-manager/core/constants/env"
	"strconv"
	"strings"
	"time"
)

// Redefine some constants to avoid import cycle.

// Mode is the type for backoff mode.
type Mode string

var Exponential Mode = "exponential"
var Linear Mode = "linear"

var backoff = struct {
	Exponential Mode
	Linear      Mode
}{
	Exponential: Exponential,
	Linear:      Linear,
}

// BackoffMaxRetries reads the "VSECM_BACKOFF_MAX_RETRIES" environment variable,
// parses its value as an int64, and returns the parsed number. If the
// environment variable is not set or cannot be parsed, a default value of
// 10 is returned. This function is useful for configuring the maximum number
// of retries in backoff algorithms, particularly in scenarios where operations
// might fail transiently and require repeated attempts to succeed.
//
// Returns:
// int64 - the maximum number of retries.
func BackoffMaxRetries() int64 {
	p := env.Value(env.VSecMBackoffMaxRetries)
	if p == "" {
		p = string(env.VSecMBackoffMaxRetriesDefault)
	}
	i, err := strconv.ParseInt(p, 10, 32)
	if err != nil {
		i, _ := strconv.Atoi(string(env.VSecMBackoffMaxRetriesDefault))
		return int64(i)
	}

	return i
}

// BackoffDelay reads the "VSECM_BACKOFF_DELAY" environment variable, parses its
// value as an int64, and returns the parsed number as a time.Duration in
// milliseconds. If the environment variable is not set or cannot be parsed,
// a default delay of  1000 milliseconds is returned. This function facilitates
// configuring the initial delay for backoff algorithms, which is essential for
// handling operations that might need a waiting period before retrying after
// a failure.
//
// Returns:
// time.Duration - the initial backoff delay duration.
func BackoffDelay() time.Duration {
	p := env.Value(env.VSecMBackoffDelay)
	if p == "" {
		p = string(env.VSecMBackoffDelayDefault)
	}

	i, err := strconv.ParseInt(p, 10, 32)
	if err != nil {
		return 1000 * time.Millisecond
	}

	return time.Duration(i) * time.Millisecond
}

// BackoffMode reads the "VSECM_BACKOFF_MODE" environment variable and determines
// the backoff strategy to be used. If the environment variable is not set, or if
// its value is "exponential", "exponential" is returned. For any other non-empty
// value, "linear" is returned. This allows for dynamic adjustment of the backoff
// strategy based on external configuration, supporting both linear and
// exponential backoff modes depending on the requirements of the operation or
// the system.
//
// Returns:
// string - the backoff mode, either "exponential" or "linear".
func BackoffMode() string {
	p := env.Value(env.VSecMBackoffMode)
	p = strings.TrimSpace(p)

	if p == "" {
		return string(backoff.Exponential)
	}

	if p != string(backoff.Exponential) {
		return string(backoff.Linear)
	}

	return string(backoff.Exponential)
}

// BackoffMaxWait reads the "VSECM_BACKOFF_MAX_WAIT" environment variable,
// parses its value as an int64, and returns the parsed number as a time.Duration
// in milliseconds. If the environment variable is not set or cannot be parsed,
// a default maximum duration of 30000 milliseconds is returned. This function is
// crucial for defining the upper limit on the duration to which backoff delay can
// grow, ensuring that retry mechanisms do not result in excessively long wait
// times.
//
// Returns:
// time.Duration - the maximum backoff duration.
func BackoffMaxWait() time.Duration {
	p := env.Value(env.VSecMBackoffMaxWait)
	if p == "" {
		p = string(env.VSecMBackoffMaxWaitDefault)
	}

	i, err := strconv.ParseInt(p, 10, 32)
	if err != nil {
		return 30000 * time.Millisecond
	}

	return time.Duration(i) * time.Millisecond
}
