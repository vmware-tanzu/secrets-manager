/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

// Package log provides a simple and flexible logging library with various
// log levels.
package std

import (
	"testing"
)

func TestSetLevel(t *testing.T) {
	type args struct {
		l Level
	}
	tests := []struct {
		name   string
		args   args
		setup  func()
		verify func()
	}{
		{
			name: "set_level_info",
			args: args{
				l: Info,
			},
			verify: func() {
				if GetLevel() != Info {
					t.Errorf("currentLevel = %v, want %v", GetLevel(), Info)
				}
			},
		},
		{
			name: "set_level_more_than_trace",
			args: args{
				l: Level(8),
			},
			setup: func() {
				SetLevel(Trace)
			},
			verify: func() {
				if GetLevel() != Trace {
					t.Errorf("currentLevel = %v, want %v", GetLevel(), Trace)
				}
			},
		},
		{
			name: "set_level_less_than_off",
			args: args{
				l: -1,
			},
			setup: func() {
				SetLevel(Audit)
			},
			verify: func() {
				if GetLevel() != Audit {
					t.Errorf("currentLevel = %v, want %v", GetLevel(), Audit)
				}
			},
		},
	}
	for _, tt := range tests {
		if tt.setup != nil {
			tt.setup()
		}
		t.Run(tt.name, func(t *testing.T) {
			SetLevel(tt.args.l)
		})
		if tt.verify != nil {
			tt.verify()
		}
	}
}

func TestGetLevel(t *testing.T) {
	tests := []struct {
		name  string
		setup func()
		want  Level
	}{
		{
			name: "success",
			setup: func() {
				SetLevel(Info)
			},
			want: Info,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup()
			}
			if got := GetLevel(); got != tt.want {
				t.Errorf("GetLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}
