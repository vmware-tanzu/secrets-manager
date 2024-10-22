/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package state_test

/*
import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/server/route/base/state"
	entity "github.com/vmware-tanzu/secrets-manager/core/entity/v1/data"
)
// Mock the UpsertSecret function from the collection package
type mockCollection struct {
	mock.Mock
}
func (m *mockCollection) UpsertSecret(secretToStore entity.SecretStored, appendValue bool) {
	m.Called(secretToStore, appendValue)
}
// Mock the Log function from the journal package
type mockJournal struct {
	mock.Mock
}
func (m *mockJournal) Log(j entity.JournalEntry) {
	m.Called(j)
}
// CustomResponseWriter is a custom HTTP response writer that simulates an error when writing a response.
type CustomResponseWriter struct {
	httptest.ResponseRecorder
	writeError bool
}
func (w *CustomResponseWriter) Write(b []byte) (int, error) {
	if w.writeError {
		return 0, errors.New("simulated write error")
	}
	return w.ResponseRecorder.Write(b)
}
// Test the Upsert function in the state package
func TestUpsert_Success(t *testing.T) {
	// Initialize mocks
	mockCol := new(mockCollection)
	mockJour := new(mockJournal)
	// Prepare test data
	secret := entity.SecretStored{Name: "test-secret"}
	journalEntry := entity.JournalEntry{
		CorrelationId: "1234",
		Payload:       "test-payload",
		Method:        "POST",
		Url:           "https://vsecm-safe.vsecm-system.svc.cluster.local:8443/workload/v1/secrets",
		SpiffeId:      "spiffe://example.org/service",
	}
	workloadId := "workload-1"
	cid := "1234"
	// Setup mock expectations for UpsertSecret
	mockCol.On("UpsertSecret", mock.MatchedBy(func(secret entity.SecretStored) bool {
		// Match relevant fields like Name and ignore timestamps
		return secret.Name == "test-secret"
	}), false).Return()
	// Setup mock expectations for Log
	mockJour.On("Log", mock.MatchedBy(func(j entity.JournalEntry) bool {
		// Ensure journal entry contains the correct data
		return j.CorrelationId == "1234" && j.Payload == "test-payload"
	})).Return()
	// Create a ResponseRecorder to record the response.
	w := httptest.NewRecorder()
	// Call the Upsert function
	state.Upsert(secret, false, workloadId, cid, journalEntry, w)
	// Assert the expectations
	mockCol.AssertCalled(t, "UpsertSecret", mock.AnythingOfType("Up"), false)
	mockJour.AssertCalled(t, "Log", mock.AnythingOfType("entity.JournalEntry"))
	// Assert HTTP response
	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	assert.Equal(t, "OK", string(body))
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
func TestUpsert_ResponseWriteError(t *testing.T) {
	// Initialize mocks
	mockCol := new(mockCollection)
	mockJour := new(mockJournal)
	// Prepare test data
	secret := entity.SecretStored{Name: "test-secret"}
	journalEntry := entity.JournalEntry{
		CorrelationId: "1234",
		Payload:       "test-payload",
		Method:        "POST",
		Url:           "https://vsecm-safe.vsecm-system.svc.cluster.local:8443/workload/v1/secrets",
		SpiffeId:      "spiffe://example.org/service",
	}
	workloadId := "workload-1"
	cid := "1234"
	// Setup mock expectations
	mockCol.On("UpsertSecret", secret, true).Return()
	mockJour.On("Log", mock.AnythingOfType("entity.JournalEntry")).Return()
	// Create a CustomResponseWriter that simulates a write error
	w := &CustomResponseWriter{
		ResponseRecorder: *httptest.NewRecorder(),
		writeError:       true,
	}
	// Call the function
	state.Upsert(secret, false, workloadId, cid, journalEntry, w)
	// Assert the expectations
	mockCol.AssertCalled(t, "UpsertSecret", secret, false)
	mockJour.AssertCalled(t, "Log", mock.AnythingOfType("entity.JournalEntry"))
	// Assert that the error was logged correctly (if applicable)
	assert.Contains(t, w.Body.String(), "") // Body will be empty due to write error
}
*/
