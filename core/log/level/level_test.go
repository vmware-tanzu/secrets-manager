/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package level

import (
	"fmt"
	"os"
	"testing"

	"github.com/vmware-tanzu/secrets-manager/core/env"
)

func TestLogLevel(t *testing.T) {
	tests := []struct {
		name    string
		setup   func() error
		cleanup func() error
		want    int
	}{
		{
			name: "default_log_level_warn",
			want: 3,
		},
		{
			name: "log_level_with_invalid_env_value",
			setup: func() error {
				err := os.Setenv("VSECM_LOG_LEVEL", "a")
				level := env.LogLevel()
				Set(Level(level))
				return err
			},
			want: 3,
		},
		{
			name: "env_log_level_0",
			setup: func() error {
				err := os.Setenv("VSECM_LOG_LEVEL", "0")
				level := env.LogLevel()
				Set(Level(level))
				return err
			},
			want: 3,
		},
		{
			name: "env_log_level_11",
			setup: func() error {
				err := os.Setenv("VSECM_LOG_LEVEL", "11")
				level := env.LogLevel()
				Set(Level(level))
				return err
			},
			want: 3,
		},
		{
			name: "env_log_level_2",
			setup: func() error {
				err := os.Setenv("VSECM_LOG_LEVEL", "2")
				level := env.LogLevel()
				Set(Level(level))
				return err
			},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				if err := tt.setup(); err != nil {
					t.Errorf("LogLevel() = failed to setup, with error: %+v", err)
				}
			}

			defer func() {
				if tt.cleanup != nil {
					if err := tt.cleanup(); err != nil {
						t.Errorf("LogLevel() = failed to cleanup, with error: %+v", err)
					}
				}
			}()

			if got := Get(); int(got) != tt.want {
				fmt.Println("Get env", os.Getenv("VSECM_LOG_LEVEL"))
			}
		})
	}
}
