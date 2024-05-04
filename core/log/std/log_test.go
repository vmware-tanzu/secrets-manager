/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

// Package log provides a simple and flexible logging library with various
// log levels.
package std

//
//func TestSetLevel(t *testing.T) {
//	type args struct {
//		l level.Level
//	}
//	tests := []struct {
//		name   string
//		args   args
//		setup  func()
//		verify func()
//	}{
//		{
//			name: "set_level_info",
//			args: args{
//				l: level.Info,
//			},
//			verify: func() {
//				if level.Get() != level.Info {
//					t.Errorf("currentLevel = %v, want %v", level.Get(), level.Info)
//				}
//			},
//		},
//		{
//			name: "set_level_more_than_trace",
//			args: args{
//				l: log.Level(8),
//			},
//			setup: func() {
//				log.SetLevel(log.Trace)
//			},
//			verify: func() {
//				if log.GetLevel() != log.Trace {
//					t.Errorf("currentLevel = %v, want %v", log.GetLevel(), log.Trace)
//				}
//			},
//		},
//		{
//			name: "set_level_less_than_off",
//			args: args{
//				l: -1,
//			},
//			setup: func() {
//				log.SetLevel(log.Audit)
//			},
//			verify: func() {
//				if log.GetLevel() != log.Audit {
//					t.Errorf("currentLevel = %v, want %v", log.GetLevel(), log.Audit)
//				}
//			},
//		},
//	}
//	for _, tt := range tests {
//		if tt.setup != nil {
//			tt.setup()
//		}
//		t.Run(tt.name, func(t *testing.T) {
//			log.SetLevel(tt.args.l)
//		})
//		if tt.verify != nil {
//			tt.verify()
//		}
//	}
//}
//
//func TestGetLevel(t *testing.T) {
//	tests := []struct {
//		name  string
//		setup func()
//		want  log.Level
//	}{
//		{
//			name: "success",
//			setup: func() {
//				log.SetLevel(log.Info)
//			},
//			want: log.Info,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if tt.setup != nil {
//				tt.setup()
//			}
//			if got := log.GetLevel(); got != tt.want {
//				t.Errorf("GetLevel() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
