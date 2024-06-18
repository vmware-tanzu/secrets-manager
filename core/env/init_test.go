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
	"reflect"
	"testing"
	"time"
)

func TestInitContainerPollInterval(t *testing.T) {
	tests := []struct {
		name    string
		setup   func() error
		cleanup func() error
		want    time.Duration
	}{
		{
			name: "default_container_poll_interval",
			want: 5000 * time.Millisecond,
		},
		{
			name: "container_poll_interval_from_env",
			setup: func() error {
				return os.Setenv("VSECM_INIT_CONTAINER_POLL_INTERVAL", "2000")
			},
			cleanup: func() error {
				return os.Unsetenv("VSECM_INIT_CONTAINER_POLL_INTERVAL")
			},
			want: 2000 * time.Millisecond,
		},
		{
			name: "container_poll_interval_from_env_with_invalid_value",
			setup: func() error {
				return os.Setenv("VSECM_INIT_CONTAINER_POLL_INTERVAL", "2a")
			},
			cleanup: func() error {
				return os.Unsetenv("VSECM_INIT_CONTAINER_POLL_INTERVAL")
			},
			want: 5000 * time.Millisecond,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				if err := tt.setup(); err != nil {
					t.Errorf("PollIntervalForInitContainer() = failed to setup with error: %+v", err)
				}
			}
			defer func() {
				if tt.cleanup != nil {
					if err := tt.cleanup(); err != nil {
						t.Errorf("PollIntervalForInitContainer() = failed to cleanup with error: %+v", err)
					}
				}
			}()
			if got := PollIntervalForInitContainer(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PollIntervalForInitContainer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSystemNamespace(t *testing.T) {
	tests := []struct {
		name    string
		setup   func() error
		cleanup func() error
		want    string
	}{
		{
			name: "system_namespace_from_env",
			setup: func() error {
				return os.Setenv("VSECM_SYSTEM_NAMESPACE", "vsecm-system")
			},
			cleanup: func() error {
				return os.Unsetenv("VSECM_SYSTEM_NAMESPACE")
			},
			want: "vsecm-system",
		},
		//{
		//	name: "empty_system_namespace",
		//	want: "",
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				if err := tt.setup(); err != nil {
					t.Errorf("SystemNamespace() = failed to setup, with error: %v", err)
				}
			}
			defer func() {
				if tt.cleanup != nil {
					if err := tt.cleanup(); err != nil {
						t.Errorf("SystemNamespace() = failed to cleanup, with error: %v", err)
					}
				}
			}()
			if got := NamespaceForVSecMSystem(); got != tt.want {
				t.Errorf("SystemNamespace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWaitBeforeExitForInitContainer(t *testing.T) {
	tests := []struct {
		name     string
		envValue string
		expected time.Duration
	}{
		{"Empty environment variable", "", 0 * time.Millisecond},
		{"Valid millisecond value", "100", 100 * time.Millisecond},
		{"Invalid value", "invalid", 0 * time.Millisecond},
		{"Negative value", "-100", -100 * time.Millisecond},
		{"Large value", "1000000", 1000000 * time.Millisecond},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set the environment variable
			_ = os.Setenv("VSECM_INIT_CONTAINER_WAIT_BEFORE_EXIT", tt.envValue)
			defer func() {
				err := os.Unsetenv("VSECM_INIT_CONTAINER_WAIT_BEFORE_EXIT")
				if err != nil {
					fmt.Println(err.Error())
				}
			}()

			// Call the function
			result := WaitBeforeExitForInitContainer()

			// Check if the result is as expected
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
