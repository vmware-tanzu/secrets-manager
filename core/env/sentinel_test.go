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

func TestInitCommandPathForSentinel(t *testing.T) {
	tests := []struct {
		name     string
		envValue string
		expected string
	}{
		{
			name:     "env not set",
			envValue: "",
			expected: "/opt/vsecm-sentinel/init/data",
		},
		{
			name:     "env set to non-empty value",
			envValue: "/custom/path",
			expected: "/custom/path",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.envValue != "" {
				err := os.Setenv("VSECM_SENTINEL_INIT_COMMAND_PATH", tt.envValue)
				if err != nil {
					t.Errorf("Error setting environment variable: %v", err)
					return
				}
			} else {
				err := os.Unsetenv("VSECM_SENTINEL_INIT_COMMAND_PATH")
				if err != nil {
					t.Errorf("Error unsetting environment variable: %v", err)
					return
				}
			}

			result := InitCommandPathForSentinel()

			if result != tt.expected {
				t.Errorf("Expected %s, but got %s", tt.expected, result)
			}

			_ = os.Unsetenv("VSECM_SENTINEL_INIT_COMMAND_PATH")
		})
	}
}

func TestInitCommandRunnerWaitBeforeExecIntervalForSentinel(t *testing.T) {
	tests := []struct {
		name     string
		envValue string
		expected time.Duration
	}{
		{
			name:     "Environment variable not set",
			envValue: "",
			expected: 0 * time.Millisecond,
		},
		{
			name:     "Environment variable set to valid integer",
			envValue: "100",
			expected: 100 * time.Millisecond,
		},
		{
			name:     "Environment variable set to invalid integer",
			envValue: "invalid",
			expected: 0 * time.Millisecond,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.envValue != "" {
				err := os.Setenv("VSECM_SENTINEL_INIT_COMMAND_WAIT_BEFORE_EXEC", tt.envValue)
				if err != nil {
					t.Errorf("Error setting environment variable: %v", err)
					return
				}
			} else {
				err := os.Unsetenv("VSECM_SENTINEL_INIT_COMMAND_WAIT_BEFORE_EXEC")
				if err != nil {
					t.Errorf("Error unsetting environment variable: %v", err)
					return
				}
			}

			result := InitCommandRunnerWaitBeforeExecIntervalForSentinel()

			if result != tt.expected {
				t.Errorf("Expected %v, but got %v", tt.expected, result)
			}

			_ = os.Unsetenv("VSECM_SENTINEL_INIT_COMMAND_WAIT_BEFORE_EXEC")
		})
	}
}

func TestInitCommandRunnerWaitIntervalBeforeInitComplete(t *testing.T) {
	tests := []struct {
		name     string
		envValue string
		expected time.Duration
	}{
		{
			name:     "Environment variable not set",
			envValue: "",
			expected: 0 * time.Millisecond,
		},
		{
			name:     "Environment variable set to valid integer",
			envValue: "100",
			expected: 100 * time.Millisecond,
		},
		{
			name:     "Environment variable set to invalid integer",
			envValue: "invalid",
			expected: 0 * time.Millisecond,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.envValue != "" {
				err := os.Setenv("VSECM_SENTINEL_INIT_COMMAND_WAIT_AFTER_INIT_COMPLETE", tt.envValue)
				if err != nil {
					t.Errorf("Error setting environment variable: %v", err)
					return
				}
			} else {
				err := os.Unsetenv("VSECM_SENTINEL_INIT_COMMAND_WAIT_AFTER_INIT_COMPLETE")
				if err != nil {
					t.Errorf("Error unsetting environment variable: %v", err)
					return
				}
			}

			result := InitCommandRunnerWaitIntervalBeforeInitComplete()

			if result != tt.expected {
				t.Errorf("Expected %v, but got %v", tt.expected, result)
			}

			_ = os.Unsetenv("VSECM_SENTINEL_INIT_COMMAND_WAIT_AFTER_INIT_COMPLETE")
		})
	}
}

func TestOIDCProviderBaseUrlForSentinel(t *testing.T) {
	tests := []struct {
		name     string
		envValue string
		expected string
	}{
		{
			name:     "Environment variable not set",
			envValue: "",
			expected: "",
		},
		{
			name:     "Environment variable set to non-empty value",
			envValue: "https://oidc.example.com",
			expected: "https://oidc.example.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.envValue != "" {
				err := os.Setenv("VSECM_SENTINEL_OIDC_PROVIDER_BASE_URL", tt.envValue)
				if err != nil {
					t.Errorf("Error setting environment variable: %v", err)
					return
				}
			} else {
				err := os.Unsetenv("VSECM_SENTINEL_OIDC_PROVIDER_BASE_URL")
				if err != nil {
					t.Errorf("Error unsetting environment variable: %v", err)
					return
				}
			}

			result := OIDCProviderBaseUrlForSentinel()

			if result != tt.expected {
				t.Errorf("Expected %s, but got %s", tt.expected, result)
			}

			_ = os.Unsetenv("VSECM_SENTINEL_OIDC_PROVIDER_BASE_URL")
		})
	}
}

func TestSentinelEnableOIDCResourceServer(t *testing.T) {
	tests := []struct {
		name     string
		envValue string
		expected bool
	}{
		{
			name:     "Environment variable not set",
			envValue: "",
			expected: false,
		},
		{
			name:     "Environment variable set to true",
			envValue: "true",
			expected: true,
		},
		{
			name:     "Environment variable set to false",
			envValue: "false",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.envValue != "" {
				err := os.Setenv("VSECM_SENTINEL_OIDC_ENABLE_RESOURCE_SERVER", tt.envValue)
				if err != nil {
					t.Errorf("Error setting environment variable: %v", err)
					return
				}
			} else {
				err := os.Unsetenv("VSECM_SENTINEL_OIDC_ENABLE_RESOURCE_SERVER")
				if err != nil {
					t.Errorf("Error unsetting environment variable: %v", err)
					return
				}
			}

			result := SentinelEnableOIDCResourceServer()

			if result != tt.expected {
				t.Errorf("Expected %v, but got %v", tt.expected, result)
			}

			_ = os.Unsetenv("VSECM_SENTINEL_OIDC_ENABLE_RESOURCE_SERVER")
		})
	}
}
