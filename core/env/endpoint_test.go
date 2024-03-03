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

func TestSafeEndpointUrl(t *testing.T) {
	tests := []struct {
		name    string
		setup   func() error
		cleanup func() error
		want    string
	}{
		{
			name: "endpoint_url_from_env_variable",
			setup: func() error {
				return os.Setenv("VSECM_SAFE_ENDPOINT_URL", "https://vsecm-safe.vsecm-system.svc.cluster.local:5000/")
			},
			cleanup: func() error {
				return os.Unsetenv("VSECM_SAFE_ENDPOINT_URL")
			},
			want: "https://vsecm-safe.vsecm-system.svc.cluster.local:5000/",
		},
		{
			name: "default_endpoint_url",
			want: "https://vsecm-safe.vsecm-system.svc.cluster.local:8443/",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				if err := tt.setup(); err != nil {
					t.Errorf("EndpointUrlForSafe() = failed to setup, with error: %+v", err)
				}
			}
			defer func() {
				if tt.cleanup != nil {
					if err := tt.cleanup(); err != nil {
						t.Errorf("EndpointUrlForSafe() = failed to cleanup, with error: %+v", err)
					}
				}
			}()
			if got := EndpointUrlForSafe(); got != tt.want {
				t.Errorf("EndpointUrlForSafe() = %v, want %v", got, tt.want)
			}
		})
	}
}
