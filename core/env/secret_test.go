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
)

func TestSecretGenerationPrefix(t *testing.T) {
	tests := []struct {
		name     string
		envValue string
		expected string
	}{
		{
			name:     "ENV not set",
			envValue: "",
			expected: "gen:",
		},
		{
			name:     "ENV set to non-empty value",
			envValue: "custom:",
			expected: "custom:",
		},
		{
			name:     "ENV set to empty string",
			envValue: "",
			expected: "gen:",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set environment variables
			if tt.envValue != "" {
				err := os.Setenv("VSECM_SENTINEL_SECRET_GENERATION_PREFIX", tt.envValue)
				if err != nil {
					t.Errorf("Error setting environment variable: %v", err)
					return
				}
			} else {
				err := os.Unsetenv("VSECM_SENTINEL_SECRET_GENERATION_PREFIX")
				if err != nil {
					t.Errorf("Error unsetting environment variable: %v", err)
					return
				}
			}

			result := SecretGenerationPrefix()

			if result != tt.expected {
				t.Errorf("Expected %s, but got %s", tt.expected, result)
			}

			err := os.Unsetenv("VSECM_SENTINEL_SECRET_GENERATION_PREFIX")
			if err != nil {
				t.Errorf("Error unsetting environment variable: %v", err)
				return
			}
		})
	}
}

func TestStoreWorkloadAsK8sSecretPrefix(t *testing.T) {
	tests := []struct {
		name     string
		envValue string
		expected string
	}{
		{
			name:     "ENV not set",
			envValue: "",
			expected: "k8s:",
		},
		{
			name:     "ENV set to non-empty value",
			envValue: "custom:",
			expected: "custom:",
		},
		{
			name:     "ENV set to empty string",
			envValue: "",
			expected: "k8s:",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set environment variables
			if tt.envValue != "" {
				err := os.Setenv("VSECM_SAFE_STORE_WORKLOAD_SECRET_AS_K8S_SECRET_PREFIX", tt.envValue)
				if err != nil {
					t.Errorf("Error setting environment variable: %v", err)
					return
				}
			} else {
				err := os.Unsetenv("VSECM_SAFE_STORE_WORKLOAD_SECRET_AS_K8S_SECRET_PREFIX")
				if err != nil {
					t.Errorf("Error unsetting environment variable: %v", err)
					return
				}
			}

			result := StoreWorkloadAsK8sSecretPrefix()

			if result != tt.expected {
				t.Errorf("Expected %s, but got %s", tt.expected, result)
			}

			err := os.Unsetenv("VSECM_SAFE_STORE_WORKLOAD_SECRET_AS_K8S_SECRET_PREFIX")
			if err != nil {
				t.Errorf("Error unsetting environment variable: %v", err)
				return
			}
		})
	}
}
