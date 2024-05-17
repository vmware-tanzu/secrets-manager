/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package backoff

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// Strategy is a configuration for the backoff strategy to use when retrying
// operations.
type Strategy struct {
	// Maximum number of retries before giving up (inclusive)
	// Default is 5
	MaxRetries int64 // Maximum number of retries before giving up (inclusive)
	// Maximum delay to use between retries (in milliseconds)
	// If Exponential is true, this is the initial delay
	// Default is 1000
	Delay time.Duration

	// Whether to use exponential backoff or not (if false, constant delay is
	// used)
	// Default is false
	Exponential bool
	// Maximum duration to wait between retries (in milliseconds)
	// Default is 10 seconds
	MaxDuration time.Duration
}

// Retry implements a retry mechanism for a function that can fail
// (return an error).
// It accepts a scope for logging or identification purposes, a function that
// it will attempt to execute, and a strategy defining the behavior of the retry
// logic.
//
// The retry strategy allows for setting maximum retries, initial delay, whether
// to use exponential backoff, and a maximum duration for the delay. If
// exponential backoff is enabled, the delay between retries increases
// exponentially with each attempt, combined with a small randomization to
// prevent synchronization issues (thundering herd problem).
// If the function succeeds (returns nil), Retry will terminate early.
// If all retries are exhausted, the last error is returned.
//
// Params:
//
//	scope string - A descriptive name or identifier for the context of the retry
//	operation.
//	f func() error - The function to execute and retry if it fails.
//	s Strategy - Struct defining the retry parameters including maximum retries,
//	delay strategy, and max delay.
//
// Returns:
//
//	error - The last error returned by the function after all retries, or nil
//	if the function eventually succeeds.
//
// Example of usage:
//
//	err := Retry("database_connection", connectToDatabase, Strategy{
//	    MaxRetries: 5,
//	    Delay: 100,
//	    Exponential: true,
//	    MaxDuration: 10 * time.Second,
//	})
//	if err != nil {
//	    fmt.Println("Failed to connect to database after retries:", err)
//	}
func Retry(scope string, f func() error, s Strategy) error {
	s = withDefaults(s)
	var err error

	for i := 0; i <= int(s.MaxRetries); i++ {
		err = f()
		if err == nil {
			return nil
		}

		var retryDelay time.Duration
		// if exponential backoff is enabled then delay increases exponentially
		if s.Exponential {
			// Calculate the delay for the current attempt. The delay is
			// 2^i seconds.
			delay := time.Duration(math.Pow(2, float64(i))) * time.Second
			// Some randomness to avoid the thundering herd problem.
			d := int(s.Delay)
			delay += time.Duration(rand.Intn(d)) * time.Millisecond
			if delay > s.MaxDuration {
				delay = s.MaxDuration
			}

			retryDelay = delay
		} else { // otherwise delay is constant for all retries
			retryDelay = s.Delay * time.Millisecond
		}

		time.Sleep(retryDelay)
		_, _ = fmt.Printf(
			"Retrying after %d ms for the scope '%s' -- attempt %d of %d",
			retryDelay, scope, i+1, s.MaxRetries+1,
		)
	}

	return err
}

// RetryExponential is a helper function to retry an operation with exponential
// backoff.
func RetryExponential(scope string, f func() error) error {
	return Retry(scope, f, Strategy{
		MaxRetries:  5,
		Delay:       1000,
		Exponential: true,
		MaxDuration: 10 * time.Second,
	})
}

// RetryFixed is a helper function to retry an operation with fixed backoff.
func RetryFixed(scope string, f func() error) error {
	return Retry(scope, f, Strategy{
		MaxRetries: 5,
		Delay:      1000,
	})
}

// withDefaults sets default values for the strategy if they are not set.
func withDefaults(s Strategy) Strategy {
	if s.MaxRetries == 0 {
		s.MaxRetries = 5
	}
	if s.Delay == 0 {
		s.Delay = 1000
	}
	if s.Exponential && s.MaxDuration == 0 {
		s.MaxDuration = 10 * time.Second
	}

	return s
}
