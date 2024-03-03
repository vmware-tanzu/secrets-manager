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

func TestTlsPort(t *testing.T) {
	tests := []struct {
		name    string
		setup   func() error
		cleanup func() error
		want    string
	}{
		{
			name: "default_safe_tls_port",
			want: ":8443",
		},
		{
			name: "safe_tls_port_from_env",
			setup: func() error {
				return os.Setenv("VSECM_SAFE_TLS_PORT", ":8976")
			},
			cleanup: func() error {
				return os.Unsetenv("VSECM_SAFE_TLS_PORT")
			},
			want: ":8976",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				if err := tt.setup(); err != nil {
					t.Errorf("TlsPort() = failed to setup, with error: %+v", err)
				}
			}
			defer func() {
				if tt.cleanup != nil {
					if err := tt.cleanup(); err != nil {
						t.Errorf("TlsPort() = failed to cleanup, with error: %+v", err)
					}
				}
			}()
			if got := TlsPort(); got != tt.want {
				t.Errorf("TlsPort() = %v, want %v", got, tt.want)
			}
		})
	}
}
