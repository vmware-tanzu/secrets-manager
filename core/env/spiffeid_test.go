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
	"github.com/vmware-tanzu/secrets-manager/core/constants/env"
	"os"
	"testing"
)

func TestSentinelSpiffeIdPrefix(t *testing.T) {
	tests := []struct {
		name    string
		setup   func() error
		cleanup func() error
		want    string
	}{
		{
			name: "default_sentinel_spiffeid_prefix",
			want: string(env.VSecMSpiffeIdPrefixSentinelDefault),
		},
		{
			name: "sentinel_spiffeid_prefix_from_env",
			setup: func() error {
				return os.Setenv("VSECM_SPIFFEID_PREFIX_SENTINEL", "spiffe://vsecm.com/workload/vsecm-sentinel/test")
			},
			cleanup: func() error {
				return os.Unsetenv("VSECM_SPIFFEID_PREFIX_SENTINEL")
			},
			want: "spiffe://vsecm.com/workload/vsecm-sentinel/test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				if err := tt.setup(); err != nil {
					t.Errorf("SpiffeIdPrefixForSentinel() = failed to setup, with error: %+v", err)
				}
			}
			defer func() {
				if tt.cleanup != nil {
					if err := tt.cleanup(); err != nil {
						t.Errorf("SpiffeIdPrefixForSentinel() = failed to cleanup, with error: %+v", err)
					}
				}
			}()
			if got := SpiffeIdPrefixForSentinel(); got != tt.want {
				t.Errorf("SpiffeIdPrefixForSentinel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSafeSpiffeIdPrefix(t *testing.T) {
	tests := []struct {
		name    string
		setup   func() error
		cleanup func() error
		want    string
	}{
		{
			name: "default_safe_spiffeid_prefix",
			want: string(env.VSecMSpiffeIdPrefixSafeDefault),
		},
		{
			name: "safe_spiffeid_prefix_from_env",
			setup: func() error {
				return os.Setenv("VSECM_SPIFFEID_PREFIX_SAFE", "spiffe://vsecm.com/workload/vsecm-safe/test")
			},
			cleanup: func() error {
				return os.Unsetenv("VSECM_SPIFFEID_PREFIX_SAFE")
			},
			want: "spiffe://vsecm.com/workload/vsecm-safe/test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				if err := tt.setup(); err != nil {
					t.Errorf("SpiffeIdPrefixForSafe() = failed to setup, with error: %+v", err)
				}
			}
			defer func() {
				if tt.cleanup != nil {
					if err := tt.cleanup(); err != nil {
						t.Errorf("SpiffeIdPrefixForSafe() = failed to cleanup, with error: %+v", err)
					}
				}
			}()
			if got := SpiffeIdPrefixForSafe(); got != tt.want {
				t.Errorf("SpiffeIdPrefixForSafe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWorkloadSpiffeIdPrefix(t *testing.T) {
	tests := []struct {
		name    string
		setup   func() error
		cleanup func() error
		want    string
	}{
		{
			name: "default_safe_spiffeid_prefix",
			want: string(env.VSecMSpiffeIdPrefixWorkloadDefault),
		},
		{
			name: "safe_spiffeid_prefix_from_env",
			setup: func() error {
				return os.Setenv("VSECM_SPIFFEID_PREFIX_WORKLOAD", "spiffe://vsecm.com/workload/test/ns/test/sa/test/n/test")
			},
			cleanup: func() error {
				return os.Unsetenv("VSECM_SPIFFEID_PREFIX_WORKLOAD")
			},
			want: "spiffe://vsecm.com/workload/test/ns/test/sa/test/n/test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				if err := tt.setup(); err != nil {
					t.Errorf("SpiffeIdPrefixForWorkload() = failed to setup, with error: %+v", err)
				}
			}
			defer func() {
				if tt.cleanup != nil {
					if err := tt.cleanup(); err != nil {
						t.Errorf("SpiffeIdPrefixForWorkload() = failed to cleanup, with error: %+v", err)
					}
				}
			}()
			if got := SpiffeIdPrefixForWorkload(); got != tt.want {
				t.Errorf("SpiffeIdPrefixForWorkload() = %v, want %v", got, tt.want)
			}
		})
	}
}
