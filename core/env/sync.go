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
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

const infiniteDuration = time.Duration(math.MaxInt64)

func RootKeySyncIntervalForSafe() time.Duration {
	p := os.Getenv("VSECM_SAFE_SYNC_ROOT_KEY_INTERVAL")
	if p == "" || p == "never" {
		return infiniteDuration
	}

	i, err := strconv.ParseInt(p, 10, 32)
	if err != nil {
		return infiniteDuration
	}

	return time.Duration(i) * time.Millisecond
}

func SecretSyncIntervalForSafe() time.Duration {
	p := os.Getenv("VSECM_SAFE_SYNC_SECRETS_INTERVAL")
	if p == "" || p == "never" {
		return infiniteDuration
	}

	i, err := strconv.ParseInt(p, 10, 32)
	if err != nil {
		return infiniteDuration
	}

	return time.Duration(i) * time.Millisecond
}

func SyncDeletedSecretsForSafe() bool {
	p := os.Getenv("VSECM_SAFE_SYNC_DELETED_SECRETS")
	if p == "" {
		return false
	}
	return strings.ToLower(strings.TrimSpace(p)) == "true"
}

func SyncInterpolatedKubernetesSecretsForSafe() bool {
	p := os.Getenv("VSECM_SAFE_SYNC_INTERPOLATED_K8S_SECRETS")
	if p == "" {
		return false
	}
	return strings.ToLower(strings.TrimSpace(p)) == "true"
}

func SyncExpiredSecretsSecretsForSafe() bool {
	p := os.Getenv("VSECM_SAFE_SYNC_EXPIRED_SECRETS")
	if p == "" {
		return false
	}
	return strings.ToLower(strings.TrimSpace(p)) == "true"
}
