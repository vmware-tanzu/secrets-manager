/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware, Inc.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package env

import (
	"os"
	"testing"
)

func TestSentinelSvidPrefix(t *testing.T) {
	tests := []struct {
		name    string
		setup   func() error
		cleanup func() error
		want    string
	}{
		{
			name: "default_sentinel_svid_prefix",
			want: "spiffe://vsecm.com/workload/vsecm-sentinel/ns/vsecm-system/sa/vsecm-sentinel/n/",
		},
		{
			name: "sentinel_svid_prefix_from_env",
			setup: func() error {
				return os.Setenv("VSECM_SENTINEL_SVID_PREFIX", "spiffe://vsecm.com/workload/vsecm-sentinel/test")
			},
			cleanup: func() error {
				return os.Unsetenv("VSECM_SENTINEL_SVID_PREFIX")
			},
			want: "spiffe://vsecm.com/workload/vsecm-sentinel/test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				if err := tt.setup(); err != nil {
					t.Errorf("SentinelSvidPrefix() = failed to setup, with error: %+v", err)
				}
			}
			defer func() {
				if tt.cleanup != nil {
					if err := tt.cleanup(); err != nil {
						t.Errorf("SentinelSvidPrefix() = failed to cleanup, with error: %+v", err)
					}
				}
			}()
			if got := SentinelSvidPrefix(); got != tt.want {
				t.Errorf("SentinelSvidPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSafeSvidPrefix(t *testing.T) {
	tests := []struct {
		name    string
		setup   func() error
		cleanup func() error
		want    string
	}{
		{
			name: "default_safe_svid_prefix",
			want: "spiffe://vsecm.com/workload/vsecm-safe/ns/vsecm-system/sa/vsecm-safe/n/",
		},
		{
			name: "safe_svid_prefix_from_env",
			setup: func() error {
				return os.Setenv("VSECM_SAFE_SVID_PREFIX", "spiffe://vsecm.com/workload/vsecm-safe/test")
			},
			cleanup: func() error {
				return os.Unsetenv("VSECM_SAFE_SVID_PREFIX")
			},
			want: "spiffe://vsecm.com/workload/vsecm-safe/test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				if err := tt.setup(); err != nil {
					t.Errorf("SafeSvidPrefix() = failed to setup, with error: %+v", err)
				}
			}
			defer func() {
				if tt.cleanup != nil {
					if err := tt.cleanup(); err != nil {
						t.Errorf("SafeSvidPrefix() = failed to cleanup, with error: %+v", err)
					}
				}
			}()
			if got := SafeSvidPrefix(); got != tt.want {
				t.Errorf("SafeSvidPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotarySvidPrefix(t *testing.T) {
	tests := []struct {
		name    string
		setup   func() error
		cleanup func() error
		want    string
	}{
		{
			name: "default_notary_svid_prefix",
			want: "spiffe://vsecm.com/workload/vsecm-notary/ns/vsecm-system/sa/vsecm-notary/n/",
		},
		{
			name: "notary_svid_prefix_from_env",
			setup: func() error {
				return os.Setenv("VSECM_NOTARY_SVID_PREFIX", "spiffe://vsecm.com/workload/vsecm-notary/test")
			},
			cleanup: func() error {
				return os.Unsetenv("VSECM_NOTARY_SVID_PREFIX")
			},
			want: "spiffe://vsecm.com/workload/vsecm-notary/test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				if err := tt.setup(); err != nil {
					t.Errorf("NotarySvidPrefix() = failed to setup, with error: %+v", err)
				}
			}
			defer func() {
				if tt.cleanup != nil {
					if err := tt.cleanup(); err != nil {
						t.Errorf("NotarySvidPrefix() = failed to cleanup, with error: %+v", err)
					}
				}
			}()
			if got := NotarySvidPrefix(); got != tt.want {
				t.Errorf("NotarySvidPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWorkloadSvidPrefix(t *testing.T) {
	tests := []struct {
		name    string
		setup   func() error
		cleanup func() error
		want    string
	}{
		{
			name: "default_safe_svid_prefix",
			want: "spiffe://vsecm.com/workload/",
		},
		{
			name: "safe_svid_prefix_from_env",
			setup: func() error {
				return os.Setenv("VSECM_WORKLOAD_SVID_PREFIX", "spiffe://vsecm.com/workload/test/")
			},
			cleanup: func() error {
				return os.Unsetenv("VSECM_WORKLOAD_SVID_PREFIX")
			},
			want: "spiffe://vsecm.com/workload/test/",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				if err := tt.setup(); err != nil {
					t.Errorf("WorkloadSvidPrefix() = failed to setup, with error: %+v", err)
				}
			}
			defer func() {
				if tt.cleanup != nil {
					if err := tt.cleanup(); err != nil {
						t.Errorf("WorkloadSvidPrefix() = failed to cleanup, with error: %+v", err)
					}
				}
			}()
			if got := WorkloadSvidPrefix(); got != tt.want {
				t.Errorf("WorkloadSvidPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
