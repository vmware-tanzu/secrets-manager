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

func TestSpiffeSocketUrl(t *testing.T) {
	tests := []struct {
		name    string
		setup   func() error
		cleanup func() error
		want    string
	}{
		{
			name: "default_spiffe_endpoint_socket",
			want: "unix:///spire-agent-socket/spire-agent.sock",
		},
		{
			name: "spiffe_endpoint_socket_from_env",
			setup: func() error {
				return os.Setenv("SPIFFE_ENDPOINT_SOCKET", "unix:///spire-agent-custom-socket/spire-agent.sock")
			},
			cleanup: func() error {
				return os.Unsetenv("SPIFFE_ENDPOINT_SOCKET")
			},
			want: "unix:///spire-agent-custom-socket/spire-agent.sock",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				if err := tt.setup(); err != nil {
					t.Errorf("SpiffeSocketUrl() = failed to setup, with error: %+v", err)
				}
			}
			defer func() {
				if tt.cleanup != nil {
					if err := tt.cleanup(); err != nil {
						t.Errorf("SpiffeSocketUrl() = failed to cleanup, with error: %+v", err)
					}
				}
			}()
			if got := SpiffeSocketUrl(); got != tt.want {
				t.Errorf("SpiffeSocketUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
