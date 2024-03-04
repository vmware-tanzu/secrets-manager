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

func TestSidecarSecretsPath(t *testing.T) {
	tests := []struct {
		name    string
		setup   func() error
		cleanup func() error
		want    string
	}{
		{
			name: "default_sidecar_secrets_path",
			want: "/opt/vsecm/secrets.json",
		},
		{
			name: "sidecar_secrets_path_from_env",
			setup: func() error {
				return os.Setenv("VSECM_SIDECAR_SECRETS_PATH", "/opt/data/secrets.json")
			},
			cleanup: func() error {
				return os.Unsetenv("VSECM_SIDECAR_SECRETS_PATH")
			},
			want: "/opt/data/secrets.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				if err := tt.setup(); err != nil {
					t.Errorf("SecretsPathForSidecar() = failed to setup, with error: %+v", err)
				}
			}
			defer func() {
				if tt.cleanup != nil {
					if err := tt.cleanup(); err != nil {
						t.Errorf("SecretsPathForSidecar() = failed to cleanup, with error: %+v", err)
					}
				}
			}()
			if got := SecretsPathForSidecar(); got != tt.want {
				t.Errorf("SecretsPathForSidecar() = %v, want %v", got, tt.want)
			}
		})
	}
}
