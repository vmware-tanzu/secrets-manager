/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package rpc

//
//import (
//	core "github.com/vmware-tanzu/secrets-manager/core/log"
//	"regexp"
//	"testing"
//)
//
//func TestSetAndGetLevel(t *testing.T) {
//	tests := []struct {
//		name     string
//		setLevel core.Level
//		want     core.Level
//	}{
//		{"Set to Off", core.Off, core.Off},
//		{"Set to Error", core.Error, core.Error},
//		{"Set to Debug", core.Debug, core.Debug},
//	}
//
//	for _, tc := range tests {
//		t.Run(tc.name, func(t *testing.T) {
//			core.SetLevel(tc.setLevel)
//			if got := core.GetLevel(); got != tc.want {
//				t.Errorf("After SetLevel(%v), GetLevel() = %v, want %v", tc.setLevel, got, tc.want)
//			}
//		})
//	}
//}
//
//func TestLogTextBuilder(t *testing.T) {
//	// Define a regexp pattern for the timestamp
//	// This example assumes an ISO 8601 format, adjust as needed
//	timeRegex := `\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}`
//
//	tests := []struct {
//		name      string
//		logHeader string
//		messages  []any
//		wantRegex string
//	}{
//		{
//			name:      "Info level log",
//			logHeader: "[INFO]",
//			messages:  []any{"Test", "message"},
//			wantRegex: `^\[INFO\]\[` + timeRegex + `\] \w+ Test message\n$`,
//		},
//		{
//			name:      "Debug level log",
//			logHeader: "[DEBUG]",
//			messages:  []any{"Another", "test", 123},
//			wantRegex: `^\[DEBUG\]\[` + timeRegex + `\] \w+ Another test 123\n$`,
//		},
//	}
//
//	cid := "VSECMSENTINEL"
//
//	for _, tc := range tests {
//		t.Run(tc.name, func(t *testing.T) {
//			got := build(tc.logHeader, &cid, tc.messages...)
//			matched, err := regexp.MatchString(tc.wantRegex, got)
//			if err != nil {
//				t.Fatalf("Regex match failed: %v", err)
//			}
//			if !matched {
//				t.Errorf("build() = %q, want match regex %q", got, tc.wantRegex)
//			}
//		})
//	}
//}
