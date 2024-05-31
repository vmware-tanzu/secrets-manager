/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package std

import (
	"github.com/vmware-tanzu/secrets-manager/core/log/level"
	"log"
	"os"
	"strings"
	"testing"
)

func TestMaxLen(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{[]string{"SHORT", "LONGER", "LONGEST"}, 7},
		{[]string{"one", "two", "three"}, 5},
		{[]string{}, 0},
	}

	for _, test := range tests {
		actual := maxLen(test.input)
		if actual != test.expected {
			t.Errorf("maxLen(%v) = %v; expected %v", test.input, actual, test.expected)
		}
	}
}

func captureOutput(f func()) string {
	var buf strings.Builder
	log.SetOutput(&buf)
	defer log.SetOutput(os.Stderr)
	f()
	return buf.String()
}

func TestLogMessage(t *testing.T) {
	// We will use log.Println for simplicity, assuming the level package has
	// necessary constants and methods.

	correlationID := "corr-id"
	expectedOutput := "PREFIX corr-id Test message\n"

	out := captureOutput(func() {
		logMessage(level.Audit, "PREFIX", &correlationID, "Test message")
	})

	// Remove timestamp
	output := strings.Join(strings.Split(out, " ")[2:], " ")
	if output != expectedOutput {
		t.Errorf("logMessage() = %v; expected %v", output, expectedOutput)
	}
}
