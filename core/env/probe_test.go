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

func TestProbeLivenessPort(t *testing.T) {
	tests := []struct {
		name    string
		setup   func() error
		cleanup func() error
		want    string
	}{
		{
			name: "default_liveness_port_value",
			want: ":8081",
		},
		{
			name: "liveness_port_value_from_env",
			setup: func() error {
				return os.Setenv("VSECM_PROBE_LIVENESS_PORT", ":5050")
			},
			cleanup: func() error {
				return os.Unsetenv("VSECM_PROBE_LIVENESS_PORT")
			},
			want: ":5050",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				if err := tt.setup(); err != nil {
					t.Errorf("ProbeLivenessPort() = failed to setup, with error: %+v", err)
				}
			}
			defer func() {
				if tt.cleanup != nil {
					if err := tt.cleanup(); err != nil {
						t.Errorf("ProbeLivenessPort() = failed to cleanup, with error: %+v", err)
					}
				}
			}()
			if got := ProbeLivenessPort(); got != tt.want {
				t.Errorf("ProbeLivenessPort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProbeReadinessPort(t *testing.T) {
	tests := []struct {
		name    string
		setup   func() error
		cleanup func() error
		want    string
	}{
		{
			name: "default_readiness_port_value",
			want: ":8082",
		},
		{
			name: "readiness_port_value_from_env",
			setup: func() error {
				return os.Setenv("VSECM_PROBE_READINESS_PORT", ":5052")
			},
			cleanup: func() error {
				return os.Unsetenv("VSECM_PROBE_READINESS_PORT")
			},
			want: ":5052",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				if err := tt.setup(); err != nil {
					t.Errorf("ProbeReadinessPort() = failed to setup, with error: %+v", err)
				}
			}
			defer func() {
				if tt.cleanup != nil {
					if err := tt.cleanup(); err != nil {
						t.Errorf("ProbeReadinessPort() = failed to cleanup, with error: %+v", err)
					}
				}
			}()
			if got := ProbeReadinessPort(); got != tt.want {
				t.Errorf("ProbeReadinessPort() = %v, want %v", got, tt.want)
			}
		})
	}
}
