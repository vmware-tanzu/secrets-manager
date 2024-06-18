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
	"fmt"
	"os"
	"testing"
)

func TestLogLevel(t *testing.T) {
	tests := []struct {
		name    string
		setup   func() error
		cleanup func() error
		want    int
	}{
		{
			name: "default_log_level_warn",
			want: 3,
		},
		{
			name: "log_level_with_invalid_env_value",
			setup: func() error {
				return os.Setenv("VSECM_LOG_LEVEL", "a")
			},
			want: 3,
		},
		{
			name: "env_log_level_0",
			setup: func() error {
				return os.Setenv("VSECM_LOG_LEVEL", "0")
			},
			want: 3,
		},
		{
			name: "env_log_level_11",
			setup: func() error {
				return os.Setenv("VSECM_LOG_LEVEL", "11")
			},
			want: 3,
		},
		{
			name: "env_log_level_2",
			setup: func() error {
				return os.Setenv("VSECM_LOG_LEVEL", "2")
			},
			want: 2,
		},
	}
	fmt.Println("###############################################")
	for _, tt := range tests {
		fmt.Println("!!!!!!running", tt.name)
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println("#####running", tt.name)

			if tt.setup != nil {
				if err := tt.setup(); err != nil {
					t.Errorf("LogLevel() = failed to setup, with error: %+v", err)
				}
			}
			defer func() {
				if tt.cleanup != nil {
					if err := tt.cleanup(); err != nil {
						t.Errorf("LogLevel() = failed to cleanup, with error: %+v", err)
					}
				}
			}()
			if got := LogLevel(); got != tt.want {
				t.Errorf("!!! LogLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogSecretFingerprints(t *testing.T) {
	tests := []struct {
		name     string
		envValue string
		expected bool
	}{
		{"Empty environment variable", "", false},
		{"Exact 'true' value", "true", true},
		{"Exact 'false' value", "false", false},
		{"Whitespace around 'true'", "  true  ", true},
		{"Uppercase 'TRUE'", "TRUE", true},
		{"Mixed case 'TrUe'", "TrUe", true},
		{"Random string", "random", false},
		{"Whitespace only", "   ", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set the environment variable
			_ = os.Setenv("VSECM_LOG_SECRET_FINGERPRINTS", tt.envValue)
			defer func() {
				err := os.Unsetenv("VSECM_LOG_SECRET_FINGERPRINTS")
				if err != nil {
					fmt.Println(err.Error())
				}
			}()

			// Call the function
			result := LogSecretFingerprints()

			// Check if the result is as expected
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
