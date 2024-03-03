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
	"math"
	"math/rand"
	"time"
)

// Strategy is a configuration for the backoff strategy to use when retrying operations.
type Strategy struct {
	// Maximum number of retries before giving up (inclusive)
	// Default is 5
	MaxRetries int // Maximum number of retries before giving up (inclusive)
	// Maximum delay to use between retries (in milliseconds)
	// If Exponential is true, this is the initial delay
	// Default is 1000
	Delay int

	// Whether to use exponential backoff or not (if false, constant delay is used)
	// Default is false
	Exponential bool
	// Maximum duration to wait between retries (in milliseconds)
	// Default is 10 seconds
	MaxDuration time.Duration
}

// Retry is a helper function to retry an operation with the given strategy.
// It returns an error if the operation fails after all retries.
func Retry(ns string, f func() error, s Strategy) error {
	s = withDefaults(s)
	var err error

	for i := 0; i <= s.MaxRetries; i++ {
		err = f()
		if err == nil {
			return nil
		}

		var retryDelay time.Duration
		// if exponential backoff is enabled then delay increases exponentially
		if s.Exponential {
			// Calculate the delay for the current attempt. The delay is 2^i seconds.
			delay := time.Duration(math.Pow(2, float64(i))) * time.Second
			// Some randomness to avoid the thundering herd problem.
			delay += time.Duration(rand.Intn(s.Delay)) * time.Millisecond
			if delay > s.MaxDuration {
				delay = s.MaxDuration
			}

			retryDelay = delay
		} else { // otherwise delay is constant for all retries
			retryDelay = time.Duration(s.Delay) * time.Millisecond
		}

		time.Sleep(retryDelay)
		println("Retrying after", retryDelay, "ms", "for", ns, "namespace --",
			"attempt", i+1, "of", s.MaxRetries+1)
	}

	return err
}

// RetryExponential is a helper function to retry an operation with exponential backoff.
func RetryExponential(ns string, f func() error) error {
	return Retry(ns, f, Strategy{
		MaxRetries:  5,
		Delay:       1000,
		Exponential: true,
		MaxDuration: 10 * time.Second,
	})
}

// RetryFixed is a helper function to retry an operation with fixed backoff.
func RetryFixed(ns string, f func() error) error {
	return Retry(ns, f, Strategy{
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
