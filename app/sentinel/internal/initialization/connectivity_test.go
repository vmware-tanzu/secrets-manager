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

func TestInitializer_ensureApiConnectivity(t *testing.T) {
	// helper function to create default mocks for testing purpose
	defaultMocks := func() (*MockLogger, *MockSpiffeOps, *MockSafeOps) {
		return &MockLogger{}, &MockSpiffeOps{}, &MockSafeOps{}
	}

	mp := NewMonkeyPatch(t)
	t.Cleanup(mp.Reset)

	// monkey patching for testing backoff functions in test cases
	mp.ApplyBackoffPatches()

	tests := []struct {
		name        string
		setupMocks  func(*MockLogger, *MockSpiffeOps, *MockSafeOps)
		expectPanic bool
	}{
		{
			name: "Successful API connectivity",
			setupMocks: func(ml *MockLogger, ms *MockSpiffeOps, mso *MockSafeOps) {
				ml.On("TraceLn", mock.Anything, mock.Anything).Return()
				ms.On("AcquireSourceForSentinel", mock.Anything).Return(&workloadapi.X509Source{}, true)
				mso.On("Check", mock.Anything, mock.Anything).Return(nil)
			},
			expectPanic: false,
		},
		{
			name: "Fail to acquire source",
			setupMocks: func(ml *MockLogger, ms *MockSpiffeOps, mso *MockSafeOps) {
				ml.On("TraceLn", mock.Anything, mock.Anything).Return()
				ms.On("AcquireSourceForSentinel", mock.Anything).Return((*workloadapi.X509Source)(nil), false)
			},
			expectPanic: true,
		},
		{
			name: "Fail safe check",
			setupMocks: func(ml *MockLogger, ms *MockSpiffeOps, mso *MockSafeOps) {
				ml.On("TraceLn", mock.Anything, mock.Anything).Return()
				ms.On("AcquireSourceForSentinel", mock.Anything).Return(&workloadapi.X509Source{}, true)
				mso.On("Check", mock.Anything, mock.Anything).Return(errors.New("check failed"))
			},
			expectPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockLogger, mockSpiffe, mockSafe := defaultMocks()

			tt.setupMocks(mockLogger, mockSpiffe, mockSafe)

			initializer := &Initializer{
				Logger: mockLogger,
				Spiffe: mockSpiffe,
				Safe:   mockSafe,
			}

			ctx := context.Background()
			cid := "test-cid"

			if tt.expectPanic {
				assert.Panics(t, func() {
					initializer.ensureApiConnectivity(ctx, &cid)
				})
			} else {
				assert.NotPanics(t, func() {
					initializer.ensureApiConnectivity(ctx, &cid)
				})
			}

			mockLogger.AssertExpectations(t)
			mockSpiffe.AssertExpectations(t)
			mockSafe.AssertExpectations(t)
		})
	}
}

func TestInitializer_ensureSourceAcquisition(t *testing.T) {
	// helper function to create default mocks for testing purpose
	defaultMocks := func() (*MockLogger, *MockSpiffeOps) {
		return &MockLogger{}, &MockSpiffeOps{}
	}

	mp := NewMonkeyPatch(t)
	t.Cleanup(mp.Reset)

	// monkey patching for testing backoff functions in test cases
	mp.ApplyBackoffPatches()

	tests := []struct {
		name            string
		setupMocks      func(*MockLogger, *MockSpiffeOps)
		expectPanic     bool
		expectNilSource bool
	}{
		{
			name: "Successfully acquire source",
			setupMocks: func(ml *MockLogger, ms *MockSpiffeOps) {
				ml.On("TraceLn", mock.Anything, mock.Anything).Return()
				ms.On("AcquireSourceForSentinel", mock.Anything).Return(&workloadapi.X509Source{}, true)
			},
			expectPanic:     false,
			expectNilSource: false,
		},
		{
			name: "Fail to acquire source",
			setupMocks: func(ml *MockLogger, ms *MockSpiffeOps) {
				ml.On("TraceLn", mock.Anything, mock.Anything).Return()
				ms.On("AcquireSourceForSentinel", mock.Anything).Return((*workloadapi.X509Source)(nil), false)
			},
			expectPanic:     true,
			expectNilSource: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockLogger, mockSpiffe := defaultMocks()

			tt.setupMocks(mockLogger, mockSpiffe)

			initializer := &Initializer{
				Logger: mockLogger,
				Spiffe: mockSpiffe,
			}

			cid := "test-cid"
			ctx := context.WithValue(context.Background(), key.CorrelationId, &cid)

			if tt.expectPanic {
				assert.Panics(t, func() {
					initializer.ensureSourceAcquisition(ctx)
				})
			} else {
				result := initializer.ensureSourceAcquisition(ctx)
				if tt.expectNilSource {
					assert.Nil(t, result)
				} else {
					assert.NotNil(t, result)
				}
			}

			mockLogger.AssertExpectations(t)
			mockSpiffe.AssertExpectations(t)
		})
	}
}
