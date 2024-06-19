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

func TestBackoffMaxRetries(t *testing.T) {
	tests := []struct {
		name     string
		envValue string
		expected int64
	}{
		{
			name:     "Environment variable not set",
			envValue: "",
			expected: 10,
		},
		{
			name:     "Environment variable set to valid value",
			envValue: "5",
			expected: 5,
		},
		{
			name:     "Environment variable set to invalid value",
			envValue: "invalid",
			expected: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.envValue != "" {
				err := os.Setenv("VSECM_BACKOFF_MAX_RETRIES", tt.envValue)
				if err != nil {
					t.Errorf("Error setting environment variable: %v", err)
					return
				}
			} else {
				err := os.Unsetenv("VSECM_BACKOFF_MAX_RETRIES")
				if err != nil {
					t.Errorf("Error unsetting environment variable: %v", err)
					return
				}
			}

			result := BackoffMaxRetries()

			if result != tt.expected {
				t.Errorf("Expected %d, but got %d", tt.expected, result)
			}

			_ = os.Unsetenv("VSECM_BACKOFF_MAX_RETRIES")
		})
	}
}

func TestBackoffDelay(t *testing.T) {
	tests := []struct {
		name     string
		envValue string
		expected time.Duration
	}{
		{
			name:     "Environment variable not set",
			envValue: "",
			expected: 1000 * time.Millisecond,
		},
		{
			name:     "Environment variable set to valid value",
			envValue: "500",
			expected: 500 * time.Millisecond,
		},
		{
			name:     "Environment variable set to invalid value",
			envValue: "invalid",
			expected: 1000 * time.Millisecond,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.envValue != "" {
				err := os.Setenv("VSECM_BACKOFF_DELAY", tt.envValue)
				if err != nil {
					return
				}
			} else {
				err := os.Unsetenv("VSECM_BACKOFF_DELAY")
				if err != nil {
					t.Errorf("Error unsetting environment variable: %v", err)
					return
				}
			}

			result := BackoffDelay()

			if result != tt.expected {
				t.Errorf("Expected %v, but got %v", tt.expected, result)
			}

			_ = os.Unsetenv("VSECM_BACKOFF_DELAY")
		})
	}
}

func TestBackoffMode(t *testing.T) {
	tests := []struct {
		name     string
		envValue string
		expected string
	}{
		{
			name:     "Environment variable not set",
			envValue: "",
			expected: "exponential",
		},
		{
			name:     "Environment variable set to exponential",
			envValue: "exponential",
			expected: "exponential",
		},
		{
			name:     "Environment variable set to linear",
			envValue: "linear",
			expected: "linear",
		},
		{
			name:     "Environment variable set to invalid value",
			envValue: "invalid",
			expected: "linear",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.envValue != "" {
				err := os.Setenv("VSECM_BACKOFF_MODE", tt.envValue)
				if err != nil {
					t.Errorf("Error setting environment variable: %v", err)
					return
				}
			} else {
				err := os.Unsetenv("VSECM_BACKOFF_MODE")
				if err != nil {
					t.Errorf("Error unsetting environment variable: %v", err)
					return
				}
			}

			result := BackoffMode()

			if result != tt.expected {
				t.Errorf("Expected %s, but got %s", tt.expected, result)
			}

			_ = os.Unsetenv("VSECM_BACKOFF_MODE")
		})
	}
}

func TestBackoffMaxDuration(t *testing.T) {
	tests := []struct {
		name     string
		envValue string
		expected time.Duration
	}{
		{
			name:     "Environment variable not set",
			envValue: "",
			expected: 30000 * time.Millisecond,
		},
		{
			name:     "Environment variable set to valid value",
			envValue: "15000",
			expected: 15000 * time.Millisecond,
		},
		{
			name:     "Environment variable set to invalid value",
			envValue: "invalid",
			expected: 30000 * time.Millisecond,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.envValue != "" {
				err := os.Setenv("VSECM_BACKOFF_MAX_WAIT", tt.envValue)
				if err != nil {
					t.Errorf("Error setting environment variable: %v", err)
					return
				}
			} else {
				err := os.Unsetenv("VSECM_BACKOFF_MAX_WAIT")
				if err != nil {
					t.Errorf("Error unsetting environment variable: %v", err)
					return
				}
			}

			result := BackoffMaxWait()

			if result != tt.expected {
				t.Errorf("Expected %v, but got %v", tt.expected, result)
			}

			_ = os.Unsetenv("VSECM_BACKOFF_MAX_WAIT")
		})
	}
}
