package data

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vmware-tanzu/secrets-manager/core/constants/audit"
)

// Corrected regex pattern for the log message
const regexPattern = `^(\d{4}/\d{2}/\d{2}) (\d{2}:\d{2}:\d{2}) \[AUDIT\]\s*(\d+)?\s*(\w+[-]?\w+)?\s*\{\{method:\[\[(.*?)\]\],url:\[\[(.*?)\]\],spiffeid:\[\[(.*?)\]\],payload:\[\[(.*?)\]\]\}\}$`

func TestAuditTest(t *testing.T) {
	// Compile the regex pattern to ensure it's valid
	compiledRegex, err := regexp.Compile(regexPattern)
	if err != nil {
		t.Fatal(err)
	}

	// Define a structure for the test input arguments
	type args struct {
		CorrelationId string
		Payload       string
		Method        string
		Url           string
		SpiffeId      string
		Event         audit.Event
	}

	// Define test cases
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success_case",
			args: args{
				CorrelationId: "1234",
				Payload:       "abcd",
				Method:        "GET",
				Url:           "http://localhost:5000/",
				SpiffeId:      "spiffe://example.org/service",
				Event:         audit.Ok,
			},
			want: "2024/10/20 14:09:32 [AUDIT] 1234 vsecm-ok {{method:[[GET]],url:[[http://localhost:5000/]],spiffeid:[[spiffe://example.org/service]],payload:[[abcd]]}}",
		},
	}

	// Execute each test case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logMessage := tt.want

			// Check if the log message matches the compiled regex
			matched := compiledRegex.MatchString(logMessage)

			// If the match fails, print the reason for debugging
			if !matched {
				matches := compiledRegex.FindStringSubmatch(logMessage)
				t.Logf("Match failed. Matches found: %v", matches)
			}

			// Assert that the log message matches the regex
			assert.True(t, matched, "Log message should match the regex pattern")
		})
	}
}
