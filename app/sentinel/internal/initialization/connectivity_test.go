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
	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/vmware-tanzu/secrets-manager/core/constants/key"
	"testing"
)

func TestInitializer_ensureApiConnectivity(t *testing.T) {
	// helper function to create default mocks for testing purpose
	defaultMocks := func() (*MockLogger, *MockSpiffeOps, *MockSafeOps) {
		return new(MockLogger), new(MockSpiffeOps), new(MockSafeOps)
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
		return new(MockLogger), new(MockSpiffeOps)
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

			ctx := context.WithValue(context.Background(), key.CorrelationId, new(string))

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

//func TestInitializer_ensureApiConnectivity(t *testing.T) {
//	mockLogger := new(MockLogger)
//	mockSpiffe := new(MockSpiffeOps)
//	mockSafe := new(MockSafeOps)
//
//	initializer := &Initializer{
//		Logger: mockLogger,
//		Spiffe: mockSpiffe,
//		Safe:   mockSafe,
//	}
//
//	ctx := context.Background()
//	cid := "test-cid"
//
//	mockLogger.On("TraceLn", mock.Anything, mock.Anything).Return()
//	mockSpiffe.On("AcquireSourceForSentinel", ctx).Return(&workloadapi.X509Source{}, true)
//	mockSafe.On("Check", ctx, mock.Anything).Return(nil)
//
//	assert.NotPanics(t, func() {
//		initializer.ensureApiConnectivity(ctx, &cid)
//	})
//
//	mockLogger.AssertExpectations(t)
//	mockSpiffe.AssertExpectations(t)
//	mockSafe.AssertExpectations(t)
//}
//
//func TestInitializer_ensureSourceAcquisition(t *testing.T) {
//	mockLogger := new(MockLogger)
//	mockSpiffe := new(MockSpiffeOps)
//
//	initializer := &Initializer{
//		Logger: mockLogger,
//		Spiffe: mockSpiffe,
//	}
//
//	cid := "test-cid"
//	ctx := context.WithValue(context.Background(), key.CorrelationId, &cid)
//
//	mockLogger.On("TraceLn", mock.Anything, mock.Anything).Return()
//	mockSpiffe.On("AcquireSourceForSentinel", ctx).Return(&workloadapi.X509Source{}, true)
//
//	result := initializer.ensureSourceAcquisition(ctx)
//
//	assert.NotNil(t, result)
//
//	mockLogger.AssertExpectations(t)
//	mockSpiffe.AssertExpectations(t)
//}
//
//func TestInitializer_ensureSourceAcquisition_Failure(t *testing.T) {
//	mockLogger := new(MockLogger)
//	mockSpiffe := new(MockSpiffeOps)
//
//	initializer := &Initializer{
//		Logger: mockLogger,
//		Spiffe: mockSpiffe,
//	}
//
//	cid := "test-cid"
//	ctx := context.WithValue(context.Background(), key.CorrelationId, &cid)
//
//	mockLogger.On("TraceLn", mock.Anything, mock.Anything).Return()
//	mockSpiffe.On("AcquireSourceForSentinel", ctx).Return((*workloadapi.X509Source)(nil), false)
//
//	assert.Panics(t, func() {
//		initializer.ensureSourceAcquisition(ctx)
//	})
//
//	mockLogger.AssertExpectations(t)
//	mockSpiffe.AssertExpectations(t)
//}

//
//import (
//	"context"
//	"errors"
//	"github.com/agiledragon/gomonkey/v2"
//	"github.com/vmware-tanzu/secrets-manager/app/sentinel/internal/safe"
//	"github.com/vmware-tanzu/secrets-manager/core/constants/key"
//	"github.com/vmware-tanzu/secrets-manager/core/env"
//	"testing"
//	"time"
//
//	"github.com/spiffe/go-spiffe/v2/workloadapi"
//	"github.com/stretchr/testify/assert"
//	"github.com/vmware-tanzu/secrets-manager/core/spiffe"
//)
//
//func TestEnsureApiConnectivity(t *testing.T) {
//	ctx := context.Background()
//	cid := "test-correlation-id"
//
//	defaultPatches := func() *gomonkey.Patches {
//		patches := gomonkey.NewPatches()
//		patches.ApplyFuncReturn(env.BackoffMaxRetries, int64(1))
//		patches.ApplyFuncReturn(env.BackoffDelay, time.Millisecond) // 1ms for testing
//		patches.ApplyFuncReturn(env.BackoffMode, string(env.Exponential))
//		patches.ApplyFuncReturn(env.BackoffMaxWait, time.Millisecond) // 1ms for testing
//
//		// add more patches here if needed
//		// ...
//
//		return patches
//	}
//
//	tests := []struct {
//		name        string
//		setupMock   func() *gomonkey.Patches
//		expectPanic bool
//	}{
//		{
//			name: "Successful API connectivity",
//			setupMock: func() *gomonkey.Patches {
//				p := defaultPatches()
//				p.ApplyFuncReturn(spiffe.AcquireSourceForSentinel, &workloadapi.X509Source{}, true)
//				p.ApplyFuncReturn(safe.Check, nil)
//				return p
//			},
//			expectPanic: false,
//		},
//		{
//			name: "API connectivity failure - AcquireSourceForSentinel fails",
//			setupMock: func() *gomonkey.Patches {
//				p := defaultPatches()
//				p.ApplyFuncReturn(spiffe.AcquireSourceForSentinel, nil, false)
//				return p
//			},
//			expectPanic: true,
//		},
//		{
//			name: "API connectivity failure - safe.Check fails",
//			setupMock: func() *gomonkey.Patches {
//				p := defaultPatches()
//				p.ApplyFuncReturn(spiffe.AcquireSourceForSentinel, &workloadapi.X509Source{}, true)
//				p.ApplyFuncReturn(safe.Check, errors.New("error"))
//				return p
//			},
//			expectPanic: true,
//		},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			m := tt.setupMock()
//			t.Cleanup(m.Reset)
//
//			if tt.expectPanic {
//				assert.Panics(t, func() { ensureApiConnectivity(ctx, &cid) })
//			} else {
//				assert.NotPanics(t, func() { ensureApiConnectivity(ctx, &cid) })
//			}
//		})
//	}
//}
//
//func TestEnsureSourceAcquisition(t *testing.T) {
//	ctx := context.Background()
//	cid := "test-correlation-id"
//	ctx = context.WithValue(ctx, key.CorrelationId, &cid)
//
//	defaultPatches := func() *gomonkey.Patches {
//		patches := gomonkey.NewPatches()
//		patches.ApplyFuncReturn(env.BackoffMaxRetries, int64(1))
//		patches.ApplyFuncReturn(env.BackoffDelay, time.Millisecond) // 1ms for testing
//		patches.ApplyFuncReturn(env.BackoffMode, string(env.Exponential))
//		patches.ApplyFuncReturn(env.BackoffMaxWait, time.Millisecond) // 1ms for testing
//		return patches
//	}
//
//	tests := []struct {
//		name        string
//		setupMock   func() *gomonkey.Patches
//		expectPanic bool
//	}{
//		{
//			name: "Successful source acquisition",
//			setupMock: func() *gomonkey.Patches {
//				p := defaultPatches()
//				p.ApplyFuncReturn(spiffe.AcquireSourceForSentinel, &workloadapi.X509Source{}, true)
//				return p
//			},
//			expectPanic: false,
//		},
//		{
//			name: "Source acquisition failure",
//			setupMock: func() *gomonkey.Patches {
//				p := defaultPatches()
//				p.ApplyFuncReturn(spiffe.AcquireSourceForSentinel, nil, false)
//				return p
//			},
//			expectPanic: true,
//		},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			m := tt.setupMock()
//			t.Cleanup(m.Reset)
//
//			if tt.expectPanic {
//				assert.Panics(t, func() { ensureSourceAcquisition(ctx) })
//			} else {
//				assert.NotPanics(t, func() { ensureSourceAcquisition(ctx) })
//			}
//		})
//	}
//}
