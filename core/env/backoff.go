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
	"time"
)

// BackoffMaxRetries reads the "VSECM_BACKOFF_MAX_RETRIES" environment variable,
// parses its value as an int64, and returns the parsed number. If the environment
// variable is not set or cannot be parsed, a default value of 20 is returned.
// This function is useful for configuring the maximum number of retries in backoff
// algorithms, particularly in scenarios where operations might fail transiently and
// require repeated attempts to succeed.
//
// Returns:
// int64 - the maximum number of retries.
func BackoffMaxRetries() int64 {
	p := os.Getenv("VSECM_BACKOFF_MAX_RETRIES")
	if p == "" {
		p = "20"
	}
	i, err := strconv.ParseInt(p, 10, 32)
	if err != nil {
		return 20
	}

	return i
}

// BackoffDelay reads the "VSECM_BACKOFF_DELAY" environment variable, parses its
// value as an int64, and returns the parsed number as a time.Duration in milliseconds.
// If the environment variable is not set or cannot be parsed, a default delay of
// 1000 milliseconds is returned. This function facilitates configuring the initial
// delay for backoff algorithms, which is essential for handling operations that
// might need a waiting period before retrying after a failure.
//
// Returns:
// time.Duration - the initial backoff delay duration.
func BackoffDelay() time.Duration {
	p := os.Getenv("VSECM_BACKOFF_DELAY")
	if p == "" {
		p = "1000"
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
// strategy based on external configuration, supporting both linear and exponential
// backoff modes depending on the requirements of the operation or the system.
//
// Returns:
// string - the backoff mode, either "exponential" or "linear".
func BackoffMode() string {
	p := os.Getenv("VSECM_BACKOFF_MODE")
	p = strings.TrimSpace(p)

	if p == "" {
		return "exponential"
	}

	if p != "exponential" {
		return "linear"
	}

	return "exponential"
}

// BackoffMaxDuration reads the "VSECM_BACKOFF_MAX_DURATION" environment variable,
// parses its value as an int64, and returns the parsed number as a time.Duration
// in milliseconds. If the environment variable is not set or cannot be parsed,
// a default maximum duration of 30000 milliseconds is returned. This function is
// crucial for defining the upper limit on the duration to which backoff delay can
// grow, ensuring that retry mechanisms do not result in excessively long wait times.
//
// Returns:
// time.Duration - the maximum backoff duration.
func BackoffMaxDuration() time.Duration {
	p := os.Getenv("VSECM_BACKOFF_MAX_DURATION")
	if p == "" {
		p = "30000"
	}
	i, err := strconv.ParseInt(p, 10, 32)
	if err != nil {
		return 30000 * time.Millisecond
	}

	return time.Duration(i) * time.Millisecond
}
