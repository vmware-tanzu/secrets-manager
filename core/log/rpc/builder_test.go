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

import (
	"regexp"
	"testing"
)

func TestBuild(t *testing.T) {

	testCorrelationID := "23B3BB0E-DB5D-4F90-B980-F0215232A017"
	dateTimeRegex := `\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}`

	tests := []struct {
		name          string
		logHeader     string
		correlationId *string
		wantRegex     string
		messageParts  []any
	}{
		{
			name:          "Valid values",
			logHeader:     "LOG",
			correlationId: &testCorrelationID,
			messageParts:  []any{"Test", "message"},
			wantRegex:     `^LOG\[` + dateTimeRegex + `\] 23B3BB0E-DB5D-4F90-B980-F0215232A017 Test message\n$`,
		},
		{
			name:          "Empty log header / mixed type message parts",
			logHeader:     "",
			correlationId: &testCorrelationID,
			messageParts:  []any{"Test", nil, "message", 3},
			wantRegex:     `^\[` + dateTimeRegex + `\] 23B3BB0E-DB5D-4F90-B980-F0215232A017 Test <nil> message 3\n$`,
		},
		{
			name:          "Nil correlationId",
			logHeader:     "LOG",
			correlationId: nil,
			messageParts:  []any{"Test", "message", 4},
			wantRegex:     `^LOG\[` + dateTimeRegex + `\] Test message 4\n$`,
		},
		{
			name:          "Nil message",
			logHeader:     "LOG",
			correlationId: &testCorrelationID,
			messageParts:  nil,
			wantRegex:     `^LOG\[` + dateTimeRegex + `\] 23B3BB0E-DB5D-4F90-B980-F0215232A017\n$`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := build(tc.logHeader, tc.correlationId, tc.messageParts[:]...)
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
