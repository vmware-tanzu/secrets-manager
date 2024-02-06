/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package timer

import (
	"github.com/vmware-tanzu/secrets-manager/core/env"
	"time"
)

var maxInterval = env.SidecarMaxPollInterval()
var factor = env.SidecarExponentialBackoffMultiplier()
var successThreshold = env.SidecarSuccessThreshold()
var errorThreshold = env.SidecarErrorThreshold()

var InitialInterval = env.SidecarPollInterval()

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

		// We have a success, so the interval “may” shrink.
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

	// We have an error, so the interval “may” expand.
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
