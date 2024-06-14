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
	"github.com/vmware-tanzu/secrets-manager/core/constants"
	"math"
	"strconv"
	"time"
)

// infiniteDuration is used to indicate that no synchronization should occur.
const infiniteDuration = time.Duration(math.MaxInt64)

// RootKeySyncIntervalForSafe retrieves the synchronization interval for root keys from an environment variable.
// If the variable is unset or set to "never", it returns an infinite duration, effectively disabling the synchronization.
//
// Returns:
//   - A time.Duration representing the interval at which root keys should be synchronized.
//   - Returns an infinite duration if the interval is set to "never" or if there is an error in parsing the interval.
func RootKeySyncIntervalForSafe() time.Duration {
	p := constants.GetEnv(constants.VSecMSafeSyncRootKeyInterval)
	if p == "" || constants.Never(p) {
		return infiniteDuration
	}

	i, err := strconv.ParseInt(p, 10, 32)
	if err != nil {
		return infiniteDuration
	}

	return time.Duration(i) * time.Millisecond
}

// SecretsSyncIntervalForSafe retrieves the synchronization interval for secrets from an environment variable.
// Similar to RootKeySyncIntervalForSafe, it returns an infinite duration if the interval is set to "never" or on error.
//
// Returns:
//   - A time.Duration representing the interval at which secrets should be synchronized.
func SecretsSyncIntervalForSafe() time.Duration {
	p := constants.GetEnv(constants.VSecMSafeSyncSecretsInterval)
	if p == "" || constants.Never(p) {
		return infiniteDuration
	}

	i, err := strconv.ParseInt(p, 10, 32)
	if err != nil {
		return infiniteDuration
	}

	return time.Duration(i) * time.Millisecond
}

// SyncDeletedSecretsForSafe checks if deleted secrets should be synchronized.
// It reads from an environment variable and returns true if synchronization
// is enabled.
//
// Returns:
//   - A bool indicating whether deleted secrets should be synchronized.
func SyncDeletedSecretsForSafe() bool {
	p := constants.GetEnv(constants.VSecMSafeSyncDeletedSecrets)
	if p == "" {
		return false
	}
	return constants.True(p)
}

// SyncInterpolatedKubernetesSecretsForSafe checks if interpolated Kubernetes
// secrets should be synchronized. It returns true if the respective environment
// variable is set to "true".
//
// Returns:
//   - A bool indicating whether interpolated Kubernetes secrets should be
//     synchronized.
func SyncInterpolatedKubernetesSecretsForSafe() bool {
	p := constants.GetEnv(constants.VSecMSafeSyncInterpolatedK8sSecrets)
	if p == "" {
		return false
	}
	return constants.True(p)
}

// SyncExpiredSecretsSecretsForSafe checks if expired secrets should be
// synchronized. It returns true if the respective environment variable is
// set to "true".
//
// Returns:
//   - A bool indicating whether expired secrets should be synchronized.
func SyncExpiredSecretsSecretsForSafe() bool {
	p := constants.GetEnv(constants.VSecMSafeSyncExpiredSecrets)
	if p == "" {
		return false
	}
	return constants.True(p)
}
