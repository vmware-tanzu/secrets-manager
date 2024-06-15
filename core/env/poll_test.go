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
	"reflect"
	"testing"
	"time"
)

func TestSidecarMaxPollInterval(t *testing.T) {
	tests := []struct {
		name    string
		setup   func() error
		cleanup func() error
		want    time.Duration
	}{
		{
			name: "default_vsecm_sidecar_max_poll_interval",
			want: 300000 * time.Millisecond,
		},
		{
			name: "vsecm_sidecar_max_poll_interval_from_env",
			setup: func() error {
				return os.Setenv("VSECM_SIDECAR_MAX_POLL_INTERVAL", "7900")
			},
			cleanup: func() error {
				return os.Unsetenv("VSECM_SIDECAR_MAX_POLL_INTERVAL")
			},
			want: 7900 * time.Millisecond,
		},
		{
			name: "invalid_vsecm_sidecar_max_poll_interval_from_env",
			setup: func() error {
				return os.Setenv("VSECM_SIDECAR_MAX_POLL_INTERVAL", "test")
			},
			cleanup: func() error {
				return os.Unsetenv("VSECM_SIDECAR_MAX_POLL_INTERVAL")
			},
			want: 300000 * time.Millisecond,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				if err := tt.setup(); err != nil {
					t.Errorf("MaxPollIntervalForSidecar() = failed to setup, with error: %+v", err)
				}
			}
			defer func() {
				if tt.cleanup != nil {
					if err := tt.cleanup(); err != nil {
						t.Errorf("MaxPollIntervalForSidecar() = failed to cleanup, with error: %+v", err)
					}
				}
			}()
			if got := MaxPollIntervalForSidecar(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaxPollIntervalForSidecar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSidecarExponentialBackoffMultiplier(t *testing.T) {
	tests := []struct {
		name    string
		setup   func() error
		cleanup func() error
		want    int64
	}{
		{
			name: "default_vsecm_sidecar_exponential_backoff_multiplier",
			want: 2,
		},
		{
			name: "vsecm_sidecar_exponential_backoff_multiplier_from_env",
			setup: func() error {
				return os.Setenv("VSECM_SIDECAR_EXPONENTIAL_BACKOFF_MULTIPLIER", "2000")
			},
			cleanup: func() error {
				return os.Unsetenv("VSECM_SIDECAR_EXPONENTIAL_BACKOFF_MULTIPLIER")
			},
			want: 2000,
		},
		{
			name: "invalid_vsecm_sidecar_exponential_backoff_multiplier_from_env",
			setup: func() error {
				return os.Setenv("VSECM_SIDECAR_EXPONENTIAL_BACKOFF_MULTIPLIER", "test")
			},
			cleanup: func() error {
				return os.Unsetenv("VSECM_SIDECAR_EXPONENTIAL_BACKOFF_MULTIPLIER")
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				if err := tt.setup(); err != nil {
					t.Errorf("ExponentialBackoffMultiplierForSidecar() = failed to setup, with error: %+v", err)
				}
			}
			defer func() {
				if tt.cleanup != nil {
					if err := tt.cleanup(); err != nil {
						t.Errorf("ExponentialBackoffMultiplierForSidecar() = failed to cleanup, with error: %+v", err)
					}
				}
			}()
			if got := ExponentialBackoffMultiplierForSidecar(); got != tt.want {
				t.Errorf("ExponentialBackoffMultiplierForSidecar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSidecarSuccessThreshold(t *testing.T) {
	tests := []struct {
		name    string
		setup   func() error
		cleanup func() error
		want    int64
	}{
		{
			name: "default_vsecm_sidecar_success_threshold",
			want: 3,
		},
		{
			name: "vsecm_sidecar_success_threshold_from_env",
			setup: func() error {
				return os.Setenv("VSECM_SIDECAR_SUCCESS_THRESHOLD", "599")
			},
			cleanup: func() error {
				return os.Unsetenv("VSECM_SIDECAR_SUCCESS_THRESHOLD")
			},
			want: 599,
		},
		{
			name: "invalid_vsecm_sidecar_success_threshold_from_env",
			setup: func() error {
				return os.Setenv("VSECM_SIDECAR_SUCCESS_THRESHOLD", "test")
			},
			cleanup: func() error {
				return os.Unsetenv("VSECM_SIDECAR_SUCCESS_THRESHOLD")
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				if err := tt.setup(); err != nil {
					t.Errorf("SuccessThresholdForSidecar() = failed to setup, with error: %+v", err)
				}
			}
			defer func() {
				if tt.cleanup != nil {
					if err := tt.cleanup(); err != nil {
						t.Errorf("SuccessThresholdForSidecar() = failed to cleanup, with error: %+v", err)
					}
				}
			}()
			if got := SuccessThresholdForSidecar(); got != tt.want {
				t.Errorf("SuccessThresholdForSidecar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSidecarErrorThreshold(t *testing.T) {
	tests := []struct {
		name    string
		setup   func() error
		cleanup func() error
		want    int64
	}{
		{
			name: "default_vsecm_sidecar_error_threshold",
			want: 3,
		},
		{
			name: "vsecm_sidecar_error_threshold_from_env",
			setup: func() error {
				return os.Setenv("VSECM_SIDECAR_ERROR_THRESHOLD", "595")
			},
			cleanup: func() error {
				return os.Unsetenv("VSECM_SIDECAR_ERROR_THRESHOLD")
			},
			want: 595,
		},
		{
			name: "invalid_vsecm_sidecar_error_threshold_from_env",
			setup: func() error {
				return os.Setenv("VSECM_SIDECAR_ERROR_THRESHOLD", "test")
			},
			cleanup: func() error {
				return os.Unsetenv("VSECM_SIDECAR_ERROR_THRESHOLD")
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				if err := tt.setup(); err != nil {
					t.Errorf("ErrorThresholdForSidecar() = failed to setup, with error: %+v", err)
				}
			}
			defer func() {
				if tt.cleanup != nil {
					if err := tt.cleanup(); err != nil {
						t.Errorf("ErrorThresholdForSidecar() = failed to cleanup, with error: %+v", err)
					}
				}
			}()
			if got := ErrorThresholdForSidecar(); got != tt.want {
				t.Errorf("ErrorThresholdForSidecar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSidecarPollInterval(t *testing.T) {
	tests := []struct {
		name    string
		setup   func() error
		cleanup func() error
		want    time.Duration
	}{
		{
			name: "default_vsecm_sidecar_poll_interval",
			want: 20000 * time.Millisecond,
		},
		{
			name: "vsecm_sidecar_poll_interval_from_env",
			setup: func() error {
				return os.Setenv("VSECM_SIDECAR_POLL_INTERVAL", "400")
			},
			cleanup: func() error {
				return os.Unsetenv("VSECM_SIDECAR_POLL_INTERVAL")
			},
			want: 400 * time.Millisecond,
		},
		{
			name: "invalid_vsecm_sidecar_poll_interval_from_env",
			setup: func() error {
				return os.Setenv("VSECM_SIDECAR_POLL_INTERVAL", "test")
			},
			cleanup: func() error {
				return os.Unsetenv("VSECM_SIDECAR_POLL_INTERVAL")
			},
			want: 20000 * time.Millisecond,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				if err := tt.setup(); err != nil {
					t.Errorf("PollIntervalForSidecar() = failed to setup, with error: %+v", err)
				}
			}
			defer func() {
				if tt.cleanup != nil {
					if err := tt.cleanup(); err != nil {
						t.Errorf("PollIntervalForSidecar() = failed to cleanup, with error: %+v", err)
					}
				}
			}()
			if got := PollIntervalForSidecar(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PollIntervalForSidecar() = %v, want %v", got, tt.want)
			}
		})
	}
}
