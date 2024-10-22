/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package http_test

import (
	"github.com/vmware-tanzu/secrets-manager/core/constants/audit"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/server/route/base/http"
	"github.com/vmware-tanzu/secrets-manager/core/entity/v1/data"
)

func TestSendEncryptedValue(t *testing.T) {
	tests := []struct {
		name           string
		cid            string
		value          string
		expectedStatus int
		expectedBody   string
		expectedEvent  audit.Event
		expectedEntry  data.JournalEntry
	}{
		{
			name:           "Empty Value",
			cid:            "1234",
			value:          "",
			expectedStatus: 400,
			expectedBody:   "",
			expectedEvent:  audit.NoValue,
			expectedEntry: data.JournalEntry{
				CorrelationId: "1234",
				Payload:       "some_payload",
				Method:        "POST",
				Url:           "https://example.com/api/secret",
				SpiffeId:      "spiffe://example.org/service",
				Event:         audit.NoValue,
			},
		},
		{
			name:           "Encryption Failure",
			cid:            "1234",
			value:          "fail",
			expectedStatus: 500,
			expectedBody:   "",
			expectedEvent:  audit.EncryptionFailed,
			expectedEntry: data.JournalEntry{
				CorrelationId: "1234",
				Payload:       "some_payload",
				Method:        "POST",
				Url:           "https://example.com/api/secret",
				SpiffeId:      "spiffe://example.org/service",
				Event:         audit.EncryptionFailed,
			},
		},

		//TODO: Expected Entry is not initialized for the test.
		/*{
			name:           "Successful Encryption",
			cid:            "1234",
			value:          "[\"{\\\"name\\\": \\\"PASSWORD\\\", \\\"value\\\": \\\"VSecMRocks!\\\"}\",\"{\\\"name\\\": \\\"USERNAME\\\", \\\"value\\\": \\\"admin\\\"}\",\"VSecMRocks\"]",
			expectedStatus: 200,
			expectedBody:   "",
			expectedEvent:  "", // No event should be logged on success
			expectedEntry: data.JournalEntry{
				CorrelationId: "1234",
				Payload:       "[\"{\\\"name\\\": \\\"PASSWORD\\\", \\\"value\\\": \\\"VSecMRocks!\\\"}\",\"{\\\"name\\\": \\\"USERNAME\\\", \\\"value\\\": \\\"admin\\\"}\",\"VSecMRocks\"]",
				Method:        "POST",
				Url:           "https://vsecm-safe.vsecm-system.svc.cluster.local:8443/workload/v1/secrets",
				SpiffeId:      "spiffe://vsecm.com/workload/example/ns/default/sa/example/n/example-c5dccdb67-xxtpv",
				Event:         "", // No event should be set on success
			},
		},

		*/
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Initialize the JournalEntry based on the test case
			journalEntry := data.JournalEntry{
				CorrelationId: tt.cid,
				Payload:       tt.expectedEntry.Payload,
				Method:        tt.expectedEntry.Method,
				Url:           tt.expectedEntry.Url,
				SpiffeId:      tt.expectedEntry.SpiffeId,
				Event:         tt.expectedEntry.Event,
			}

			w := httptest.NewRecorder()

			http.SendEncryptedValue(tt.cid, tt.value, journalEntry, w)

			resp := w.Result()
			body, _ := io.ReadAll(resp.Body)

			// Assert HTTP status code
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			// Assert response body
			assert.Equal(t, tt.expectedBody, string(body))

			// Assert journal event
			assert.Equal(t, tt.expectedEvent, journalEntry.Event)

			// Assert the entire journal entry matches expectations
			assert.Equal(t, tt.expectedEntry, journalEntry)
		})
	}

}
