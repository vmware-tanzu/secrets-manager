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

func TestRootKeyPathForKeyGen(t *testing.T) {
	tests := []struct {
		name     string
		envValue string
		expected string
	}{
		{
			name:     "Environment variable not set",
			envValue: "",
			expected: "/opt/vsecm/keys.txt",
		},
		{
			name:     "Environment variable set to non-empty value",
			envValue: "/custom/key/path",
			expected: "/custom/key/path",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.envValue != "" {
				err := os.Setenv("VSECM_KEYGEN_ROOT_KEY_PATH", tt.envValue)
				if err != nil {
					t.Errorf("Error setting environment variable: %v", err)
					return
				}
			} else {
				err := os.Unsetenv("VSECM_KEYGEN_ROOT_KEY_PATH")
				if err != nil {
					t.Errorf("Error unsetting environment variable: %v", err)
					return
				}
			}

			result := RootKeyPathForKeyGen()

			if result != tt.expected {
				t.Errorf("Expected %s, but got %s", tt.expected, result)
			}

			_ = os.Unsetenv("VSECM_KEYGEN_ROOT_KEY_PATH")
		})
	}
}

func TestExportedSecretPathForKeyGen(t *testing.T) {
	tests := []struct {
		name     string
		envValue string
		expected string
	}{
		{
			name:     "Environment variable not set",
			envValue: "",
			expected: "/opt/vsecm/secrets.json",
		},
		{
			name:     "Environment variable set to non-empty value",
			envValue: "/custom/secrets/path",
			expected: "/custom/secrets/path",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.envValue != "" {
				err := os.Setenv("VSECM_KEYGEN_EXPORTED_SECRET_PATH", tt.envValue)
				if err != nil {
					t.Errorf("Error setting environment variable: %v", err)
					return
				}
			} else {
				err := os.Unsetenv("VSECM_KEYGEN_EXPORTED_SECRET_PATH")
				if err != nil {
					t.Errorf("Error unsetting environment variable: %v", err)
					return
				}
			}

			result := ExportedSecretPathForKeyGen()

			if result != tt.expected {
				t.Errorf("Expected %s, but got %s", tt.expected, result)
			}

			_ = os.Unsetenv("VSECM_KEYGEN_EXPORTED_SECRET_PATH")
		})
	}
}

func TestKeyGenDecrypt(t *testing.T) {
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
				err := os.Setenv("VSECM_KEYGEN_DECRYPT", tt.envValue)
				if err != nil {
					t.Errorf("Error setting environment variable: %v", err)
					return
				}
			} else {
				err := os.Unsetenv("VSECM_KEYGEN_DECRYPT")
				if err != nil {
					t.Errorf("Error unsetting environment variable: %v", err)
					return
				}
			}

			result := KeyGenDecrypt()

			if result != tt.expected {
				t.Errorf("Expected %v, but got %v", tt.expected, result)
			}

			// Clean up
			_ = os.Unsetenv("VSECM_KEYGEN_DECRYPT")
		})
	}
}
