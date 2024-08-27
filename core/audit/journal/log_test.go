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
	"encoding/json"
	"github.com/vmware-tanzu/secrets-manager/core/constants/audit"
	"github.com/vmware-tanzu/secrets-manager/core/entity/v1/data"
	"net/http"
	"regexp"
	"testing"

	reqres "github.com/vmware-tanzu/secrets-manager/core/entity/v1/reqres/safe"
)

func TestLog(t *testing.T) {
	// Compile the regular expression
	re, err := regexp.Compile(regexPattern)
	if err != nil {
		t.Fatal(err)
	}

	toJsonStr := func(s any) string {
		bytes, err := json.Marshal(s)
		if err != nil {
			return ""
		}
		return string(bytes)
	}

	type args struct {
		e data.JournalEntry
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "nil_JournalEntry",
			args: args{
				e: data.JournalEntry{},
			},
		},
		{
			name: "Entity_default",
			args: args{
				e: data.JournalEntry{
					CorrelationId: "1234",
					Payload:       "",
					Method:        "test_method",
					Url:           "test_url",
					SpiffeId:      "test_spiffeid",
					Event:         "test_event",
				},
			},
		},
		{
			name: "Entity_type_SecretDeleteRequest",
			args: args{
				e: data.JournalEntry{
					CorrelationId: "1234",
					Payload:       toJsonStr(reqres.SecretDeleteRequest{WorkloadIds: []string{"test_workloadid"}, Err: "test_err"}),
					Method:        "test_method",
					Url:           "test_url",
					SpiffeId:      "test_spiffeid",
					Event:         "test_event",
				},
			},
		},
		{
			name: "Entity_type_SecretDeleteResponse",
			args: args{
				e: data.JournalEntry{
					CorrelationId: "1234",
					Payload:       toJsonStr(reqres.SecretDeleteResponse{Err: "test_err"}),
					Method:        "test_method",
					Url:           "test_url",
					SpiffeId:      "test_spiffeid",
					Event:         "test_event",
				},
			},
		},
		{
			name: "Entity_type_SecretFetchRequest",
			args: args{
				e: data.JournalEntry{
					CorrelationId: "1234",
					Payload:       toJsonStr(reqres.SecretFetchRequest{Err: "test_err"}),
					Method:        "test_method",
					Url:           "test_url",
					SpiffeId:      "test_spiffeid",
					Event:         "test_event",
				},
			},
		},
		{
			name: "Entity_type_SecretFetchResponse",
			args: args{
				e: data.JournalEntry{
					CorrelationId: "1234",
					Payload:       toJsonStr(reqres.SecretFetchResponse{Data: "test_data", Created: "test_created", Updated: "test_updated", Err: "test_err"}),
					Method:        "test_method",
					Url:           "test_url",
					SpiffeId:      "test_spiffeid",
					Event:         "test_event",
				},
			},
		},
		{
			name: "Entity_type_SecretUpsertRequest",
			args: args{
				e: data.JournalEntry{
					CorrelationId: "1234",
					Payload: toJsonStr(reqres.SecretUpsertRequest{
						WorkloadIds: []string{"test_workloadid"},
						Namespaces:  []string{"test_namespace"},
						Value:       "test_value",
						Template:    "test_template",
						Format:      data.SecretFormat("test_format"),
					}),
				},
			},
		},
		{
			name: "Entity_type_SecretUpsertResponse",
			args: args{
				e: data.JournalEntry{
					CorrelationId: "1234",
					Payload:       toJsonStr(reqres.SecretUpsertResponse{Err: "test_err"}),
					Method:        "test_method",
					Url:           "test_url",
					SpiffeId:      "test_spiffeid",
					Event:         "test_event",
				},
			},
		},
		{
			name: "Entity_type_SecretListRequest",
			args: args{
				e: data.JournalEntry{
					CorrelationId: "1234",
					Payload:       toJsonStr(reqres.SecretListRequest{Err: "test_err"}),
					Method:        "test_method",
					Url:           "test_url",
					SpiffeId:      "test_spiffeid",
					Event:         "test_event",
				},
			},
		},
		{
			name: "Entity_type_SecretListResponse",
			args: args{
				e: data.JournalEntry{
					CorrelationId: "1234",
					Payload:       toJsonStr(reqres.SecretListResponse{Err: "test_err"}),
					Method:        "test_method",
					Url:           "test_url",
					SpiffeId:      "test_spiffeid",
					Event:         "test_event",
				},
			},
		},
		{
			name: "Entity_type_SecretEncryptedListResponse",
			args: args{
				e: data.JournalEntry{
					CorrelationId: "1234",
					Payload:       toJsonStr(reqres.SecretEncryptedListResponse{Err: "test_err"}),
					Method:        "test_method",
					Url:           "test_url",
					SpiffeId:      "test_spiffeid",
					Event:         "test_event",
				},
			},
		},
		{
			name: "Entity_type_KeyInputRequest",
			args: args{
				e: data.JournalEntry{
					CorrelationId: "1234",
					Payload:       toJsonStr(reqres.KeyInputRequest{}),
					Method:        "test_method",
					Url:           "test_url",
					SpiffeId:      "test_spiffeid",
					Event:         "test_event",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logLine := captureOutput(func() {
				Log(tt.args.e)
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
			event := matches[4]
			method := matches[5]
			url := matches[6]
			spiffeid := matches[7]
			payload := matches[8]

			// Check if extracted values match expected arguments
			if !(correlationID == tt.args.e.CorrelationId &&
				event == string(tt.args.e.Event) &&
				method == tt.args.e.Method &&
				url == tt.args.e.Url &&
				spiffeid == tt.args.e.SpiffeId &&
				payload == tt.args.e.Payload) {
				t.Errorf("printAudit() = %v; expected %v", logLine, regexPattern)
			}
		})
	}
}

func TestCreateDefaultEntry(t *testing.T) {
	type args struct {
		cid      string
		spiffeid string
		r        *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		//{ // This test case is not applicable as the function is not expected to handle nil requests
		//	name: "nil_http_request",
		//	args: args{
		//		cid:      "1234",
		//		spiffeid: "abcd",
		//		r:        nil,
		//	},
		//},
		{
			name: "empty_values",
			args: args{
				cid:      "",
				spiffeid: "",
				r:        &http.Request{},
			},
		},
		{
			name: "valid_values",
			args: args{
				cid:      "1234",
				spiffeid: "abcd",
				r: &http.Request{
					Method:     "GET",
					RequestURI: "http://localhost:5000/",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := CreateDefaultEntry(tt.args.cid, tt.args.spiffeid, tt.args.r)
			if !(actual.CorrelationId == tt.args.cid &&
				actual.Method == tt.args.r.Method &&
				actual.Url == tt.args.r.RequestURI &&
				actual.SpiffeId == tt.args.spiffeid &&
				actual.Event == audit.Enter) {
				t.Errorf("CreateDefaultEntry() = %v; expected %v", actual, tt.args)
			}
		})
	}
}
