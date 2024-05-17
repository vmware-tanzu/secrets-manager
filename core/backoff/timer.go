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
	"time"

	"github.com/vmware-tanzu/secrets-manager/core/env"
)

var maxInterval = env.MaxPollIntervalForSidecar()
var factor = env.ExponentialBackoffMultiplierForSidecar()
var successThreshold = env.SuccessThresholdForSidecar()
var errorThreshold = env.ErrorThresholdForSidecar()

var InitialInterval = env.PollIntervalForSidecar()

// ExponentialBackoff calculates the next interval for retrying an operation,
// based on the outcome of the current attempt (success or failure), the current
// interval, and the counts of successive successes and errors. It adjusts the
// interval according to an exponential backoff strategy, which expands the
// interval after failures and shrinks it following a series of successes.
//
// Parameters:
// - success (bool): Indicates whether the current attempt was successful.
// - interval (time.Duration): The current interval between attempts.
// - successCount (int64): The number of successive successful attempts so far.
// - errorCount (int64): The number of successive failed attempts so far.
//
// Returns:
//   - time.Duration: The next interval to use for the following attempt. This
//     value is adjusted based on the outcome of the current attempt, subject to
//     the defined minimum and maximum boundaries.
//   - int64: The updated count of successive successes, reset to 0 after an
//     interval adjustment or continued incrementation upon success.
//   - int64: The updated count of successive errors, reset to 0 after an
//     interval adjustment or continued incrementation upon failure.
func ExponentialBackoff(
	success bool, interval time.Duration, successCount, errorCount int64,
) (time.Duration, int64, int64) {
	// #region Boundary Corrections
	if factor < 1 {
		factor = 1
	}
	if InitialInterval > maxInterval {
		InitialInterval = maxInterval
	}
	// #endregion

	// Decide whether to shrink, expand, or keep the interval the same
	// based on the success and error count so far.
	if success {
		nextSuccessCount := successCount + 1

		// We have a success, so the interval "may" shrink.
		shrinkInterval := nextSuccessCount >= successThreshold
		if shrinkInterval {
			nextInterval := time.Duration(int64(interval) / factor)

			// boundary check:
			if nextInterval < InitialInterval {
				nextInterval = InitialInterval
			}

			// Interval shrank.
			return nextInterval, 0, 0
		}

		// Success count increased, interval is intact.
		return interval, nextSuccessCount, 0
	}

	nextErrorCount := errorCount + 1

	// We have an error, so the interval "may" expand.
	expandInterval := nextErrorCount >= errorThreshold
	if expandInterval {
		nextInterval := time.Duration(int64(interval) * factor)

		// boundary check:
		if nextInterval > maxInterval {
			nextInterval = maxInterval
		}

		// Interval expanded.
		return nextInterval, 0, 0
	}

	// Error count increased, interval is intact.
	return interval, 0, nextErrorCount
}
