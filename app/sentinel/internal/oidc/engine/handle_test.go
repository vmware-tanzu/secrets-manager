/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package engine

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/vmware-tanzu/secrets-manager/core/entity/v1/reqres/sentinel"
)

func TestEngine_HandleSecrets(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		request        *sentinel.SecretRequest
		setupMocks     func(safeOps *MockSafeOperations, authorizer *MockAuthorizer, logger *MockLogger)
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Invalid method",
			method:         http.MethodGet,
			request:        nil,
			setupMocks:     func(safeOps *MockSafeOperations, authorizer *MockAuthorizer, logger *MockLogger) {},
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   "unsupported method\n",
		},
		{
			name:           "Invalid JSON",
			method:         http.MethodPost,
			request:        nil,
			setupMocks:     func(safeOps *MockSafeOperations, authorizer *MockAuthorizer, logger *MockLogger) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "invalid request body\n",
		},
		{
			name:   "List secrets - authorized",
			method: http.MethodPost,
			request: &sentinel.SecretRequest{
				List: true,
			},
			setupMocks: func(safeOps *MockSafeOperations, authorizer *MockAuthorizer, logger *MockLogger) {
				authorizer.On("IsAuthorized", mock.Anything, mock.Anything).Return(true)
				safeOps.On("GetSecrets", mock.Anything, mock.Anything, false).Return("listed secrets", nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   "listed secrets",
		},
		{
			name:   "Modify secrets - authorized",
			method: http.MethodPost,
			request: &sentinel.SecretRequest{
				List:      false,
				Workloads: []string{"workload1"},
				Secret:    "mysecret",
			},
			setupMocks: func(safeOps *MockSafeOperations, authorizer *MockAuthorizer, logger *MockLogger) {
				authorizer.On("IsAuthorized", mock.Anything, mock.Anything).Return(true)
				safeOps.On("UpdateSecrets", mock.Anything, mock.Anything, mock.Anything).Return("secret modified", nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   "secret modified",
		},
		{
			name:   "Unauthorized request",
			method: http.MethodPost,
			request: &sentinel.SecretRequest{
				List: true,
			},
			setupMocks: func(safeOps *MockSafeOperations, authorizer *MockAuthorizer, logger *MockLogger) {
				authorizer.On("IsAuthorized", mock.Anything, mock.Anything).Return(false)
			},
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   "unauthorized: please provide correct credentials\n",
		},
		{
			name:   "Invalid input",
			method: http.MethodPost,
			request: &sentinel.SecretRequest{
				List:      false,
				Workloads: []string{},
				Secret:    "",
			},
			setupMocks: func(safeOps *MockSafeOperations, authorizer *MockAuthorizer, logger *MockLogger) {
				authorizer.On("IsAuthorized", mock.Anything, mock.Anything).Return(true)
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "invalid input for secret modification\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSafeOps := &MockSafeOperations{}
			mockAuthorizer := &MockAuthorizer{}
			mockLogger := &MockLogger{}
			tt.setupMocks(mockSafeOps, mockAuthorizer, mockLogger)
			engine := newEngine(mockSafeOps, mockAuthorizer, mockLogger)

			var reqBody []byte
			var err error
			if tt.request != nil {
				reqBody, err = json.Marshal(tt.request)
				assert.NoError(t, err)
			}

			req, _ := http.NewRequest(tt.method, "/secrets", bytes.NewBuffer(reqBody))
			rr := httptest.NewRecorder()

			engine.HandleSecrets(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)
			assert.Equal(t, tt.expectedBody, rr.Body.String())

			mockSafeOps.AssertExpectations(t)
			mockAuthorizer.AssertExpectations(t)
			mockLogger.AssertExpectations(t)
		})
	}
}
