/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package initialization

import (
	"context"
	"errors"
	"testing"

	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/vmware-tanzu/secrets-manager/core/constants/key"
)

func TestInitializer_initCommandsExecutedAlready(t *testing.T) {
	mp := NewMonkeyPatch(t)
	t.Cleanup(mp.Reset)
	// monkey patching for testing backoff functions in test cases
	mp.ApplyBackoffPatches()

	tests := []struct {
		name           string
		setupMocks     func(*MockLogger, *MockSafeOps)
		expectedResult bool
		expectPanic    bool
	}{
		{
			name: "Initialization check succeeds and returns true",
			setupMocks: func(ml *MockLogger, ms *MockSafeOps) {
				ml.On("TraceLn", mock.Anything, mock.Anything).Return()
				ms.On("CheckInitialization", mock.Anything, mock.Anything).Return(true, nil)
			},
			expectedResult: true,
			expectPanic:    false,
		},
		{
			name: "Initialization check succeeds and returns false",
			setupMocks: func(ml *MockLogger, ms *MockSafeOps) {
				ml.On("TraceLn", mock.Anything, mock.Anything).Return()
				ms.On("CheckInitialization", mock.Anything, mock.Anything).Return(false, nil)
			},
			expectedResult: false,
			expectPanic:    false,
		},
		{
			name: "Initialization check fails with error",
			setupMocks: func(ml *MockLogger, ms *MockSafeOps) {
				ml.On("TraceLn", mock.Anything, mock.Anything).Return()
				ms.On("CheckInitialization", mock.Anything, mock.Anything).Return(false, errors.New("check failed"))
			},
			expectedResult: false,
			expectPanic:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockLogger := &MockLogger{}
			mockSafeOps := &MockSafeOps{}

			tt.setupMocks(mockLogger, mockSafeOps)

			initializer := &Initializer{
				Logger: mockLogger,
				Safe:   mockSafeOps,
			}

			cid := "test-cid"
			ctx := context.WithValue(context.Background(), key.CorrelationId, &cid)
			src := &workloadapi.X509Source{}

			if tt.expectPanic {
				assert.Panics(t, func() {
					initializer.initCommandsExecutedAlready(ctx, src)
				})
			} else {
				result := initializer.initCommandsExecutedAlready(ctx, src)
				assert.Equal(t, tt.expectedResult, result)
			}

			mockLogger.AssertExpectations(t)
			mockSafeOps.AssertExpectations(t)
		})
	}
}
