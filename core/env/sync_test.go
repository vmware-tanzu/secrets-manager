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
	"testing"
	"time"
)

func TestRootKeySyncIntervalForSafe(t *testing.T) {
	tests := []struct {
		envValue string
		expected time.Duration
	}{
		{"", infiniteDuration},
		{"never", infiniteDuration},
		{"1000", 1000 * time.Millisecond},
		{"invalid", infiniteDuration},
	}

	for _, test := range tests {
		_ = os.Setenv("VSECM_SAFE_SYNC_ROOT_KEY_INTERVAL", test.envValue)
		actual := RootKeySyncIntervalForSafe()
		if actual != test.expected {
			t.Errorf("RootKeySyncIntervalForSafe() with env value %q = %v; expected %v", test.envValue, actual, test.expected)
		}
	}
}

func TestSecretSyncIntervalForSafe(t *testing.T) {
	tests := []struct {
		envValue string
		expected time.Duration
	}{
		{"", infiniteDuration},
		{"never", infiniteDuration},
		{"2000", 2000 * time.Millisecond},
		{"invalid", infiniteDuration},
	}

	for _, test := range tests {
		_ = os.Setenv("VSECM_SAFE_SYNC_SECRETS_INTERVAL", test.envValue)
		actual := SecretsSyncIntervalForSafe()
		if actual != test.expected {
			t.Errorf("SecretsSyncIntervalForSafe() with env value %q = %v; expected %v", test.envValue, actual, test.expected)
		}
	}
}

func TestSyncDeletedSecretsForSafe(t *testing.T) {
	tests := []struct {
		envValue string
		expected bool
	}{
		{"", false},
		{"true", true},
		{"false", false},
		{"TRUE", true},
		{"FALSE", false},
	}

	for _, test := range tests {
		_ = os.Setenv("VSECM_SAFE_SYNC_DELETED_SECRETS", test.envValue)
		actual := SyncDeletedSecretsForSafe()
		if actual != test.expected {
			t.Errorf("SyncDeletedSecretsForSafe() with env value %q = %v; expected %v", test.envValue, actual, test.expected)
		}
	}
}

func TestSyncInterpolatedKubernetesSecretsForSafe(t *testing.T) {
	tests := []struct {
		envValue string
		expected bool
	}{
		{"", false},
		{"true", true},
		{"false", false},
		{"TRUE", true},
		{"FALSE", false},
	}

	for _, test := range tests {
		_ = os.Setenv("VSECM_SAFE_SYNC_INTERPOLATED_K8S_SECRETS", test.envValue)
		actual := SyncInterpolatedKubernetesSecretsForSafe()
		if actual != test.expected {
			t.Errorf("SyncInterpolatedKubernetesSecretsForSafe() with env value %q = %v; expected %v", test.envValue, actual, test.expected)
		}
	}
}

func TestSyncExpiredSecretsSecretsForSafe(t *testing.T) {
	tests := []struct {
		envValue string
		expected bool
	}{
		{"", false},
		{"true", true},
		{"false", false},
		{"TRUE", true},
		{"FALSE", false},
	}

	for _, test := range tests {
		_ = os.Setenv("VSECM_SAFE_SYNC_EXPIRED_SECRETS", test.envValue)
		actual := SyncExpiredSecretsSecretsForSafe()
		if actual != test.expected {
			t.Errorf("SyncExpiredSecretsSecretsForSafe() with env value %q = %v; expected %v", test.envValue, actual, test.expected)
		}
	}
}
