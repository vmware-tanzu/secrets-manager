/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package rpc

import (
	"github.com/vmware-tanzu/secrets-manager/core/log"
	"regexp"
	"testing"
)

func TestSetAndGetLevel(t *testing.T) {
	tests := []struct {
		name     string
		setLevel log.Level
		want     log.Level
	}{
		{"Set to Off", log.Off, log.Off},
		{"Set to Error", log.Error, log.Error},
		{"Set to Debug", log.Debug, log.Debug},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			log.SetLevel(tc.setLevel)
			if got := log.GetLevel(); got != tc.want {
				t.Errorf("After SetLevel(%v), GetLevel() = %v, want %v", tc.setLevel, got, tc.want)
			}
		})
	}
}

func TestLogTextBuilder(t *testing.T) {
	// Define a regex pattern for the timestamp
	// This example assumes an ISO 8601 format, adjust as needed
	timeRegex := `\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}`

	tests := []struct {
		name      string
		logHeader string
		messages  []any
		wantRegex string
	}{
		{
			name:      "Info level log",
			logHeader: "[SENTINEL_INFO]",
			messages:  []any{"Test", "message"},
			wantRegex: `^\[SENTINEL_INFO\]\[` + timeRegex + `\] Test message\n$`,
		},
		{
			name:      "Debug level log",
			logHeader: "[SENTINEL_DEBUG]",
			messages:  []any{"Another", "test", 123},
			wantRegex: `^\[SENTINEL_DEBUG\]\[` + timeRegex + `\] Another test 123\n$`,
		},
	}

	cid := "VSECMSENTINEL"

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := build(tc.logHeader, &cid, tc.messages...)
			matched, err := regexp.MatchString(tc.wantRegex, got)
			if err != nil {
				t.Fatalf("Regex match failed: %v", err)
			}
			if !matched {
				t.Errorf("build() = %q, want match regex %q", got, tc.wantRegex)
			}
		})
	}
}
