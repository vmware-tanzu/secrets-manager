/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package journal

import (
	"github.com/vmware-tanzu/secrets-manager/core/constants/audit"
	"log"
	"os"
	"regexp"
	"strings"
	"testing"
)

// Regular expression to match the log line
const regexPattern = `^(\d{4}/\d{2}/\d{2}) (\d{2}:\d{2}:\d{2}) \[AUDIT\]\s*(\d+)?\s*(\w+)?\s*\{\{method:\[\[(.*?)\]\],url:\[\[(.*?)\]\],spiffeid:\[\[(.*?)\]\],payload:\[\[(.*?)\]\]\}\}$` // Updated regex to allow for optional correlation ID and entity name

func Test_printAudit(t *testing.T) {
	// Compile the regular expression
	re, err := regexp.Compile(regexPattern)
	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		correlationId string
		entityName    string
		method        string
		url           string
		spiffeid      string
		message       string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success_case",
			args: args{
				correlationId: "1234",
				entityName:    "abcd",
				method:        "GET",
				url:           "http://localhost:5000/",
				spiffeid:      "abcd1234",
				message:       "testing audit func",
			},
		},
		{
			name: "empty_values",
			args: args{
				correlationId: "",
				entityName:    "",
				method:        "",
				url:           "",
				spiffeid:      "",
				message:       "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logLine := captureOutput(func() {
				printAudit(tt.args.correlationId, audit.Event(tt.args.entityName), tt.args.method, tt.args.url, tt.args.spiffeid, tt.args.message)
			})
			// Match the log line against the regular expression
			matches := re.FindStringSubmatch(clean(logLine))
			if matches == nil {
				t.Fatalf("printAudit() = %v; expected %v", logLine, regexPattern)
			}

			// Extract components from the matched groups
			_ = matches[1] // date
			_ = matches[2] // time
			correlationID := matches[3]
			entityName := matches[4]
			method := matches[5]
			url := matches[6]
			spiffeid := matches[7]
			payload := matches[8]

			// Check if extracted values match expected arguments
			if !(correlationID == tt.args.correlationId &&
				entityName == tt.args.entityName &&
				method == tt.args.method &&
				url == tt.args.url &&
				spiffeid == tt.args.spiffeid &&
				payload == tt.args.message) {
				t.Errorf("printAudit() = %v; expected %v", logLine, regexPattern)
			}
		})
	}
}

// clean prepares log output for comparison.
func clean(s string) string {
	if len(s) > 0 && s[len(s)-1] == '\n' { // Remove trailing newline
		s = s[:len(s)-1]
	}
	return strings.ReplaceAll(s, "\n", "~") // Replace newline with tilde for comparison purposes
}

// captureOutput captures log output. It sets the log output to a buffer, runs the function, and returns the buffer contents.
func captureOutput(f func()) string {
	var buf strings.Builder
	log.SetOutput(&buf)
	defer log.SetOutput(os.Stderr)
	f()
	return buf.String()
}
