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
	"os"
	"testing"
	"time"

	"github.com/vmware-tanzu/secrets-manager/core/constants/key"

	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestInitializer_RunInitCommands(t *testing.T) {
	mp := NewMonkeyPatch(t)
	t.Cleanup(mp.Reset)
	// monkey patching for testing backoff functions in test cases
	mp.ApplyBackoffPatches()

	tests := []struct {
		name        string
		setupMocks  func(*MockFileOpener, *MockEnvReader, *MockLogger, *MockSafeOps, *MockSpiffeOps)
		expectPanic bool
	}{
		{
			name: "Successful initialization",
			setupMocks: func(mfo *MockFileOpener, mer *MockEnvReader, ml *MockLogger, ms *MockSafeOps, msp *MockSpiffeOps) {
				mer.On("InitCommandRunnerWaitBeforeExecIntervalForSentinel").Return(time.Millisecond)
				mer.On("InitCommandRunnerWaitIntervalBeforeInitComplete").Return(time.Millisecond)
				mer.On("InitCommandPathForSentinel").Return("/path/to/file")
				mer.On("NamespaceForVSecMSystem").Return("vsecm-system")
				msp.On("AcquireSourceForSentinel", mock.Anything).Return(&workloadapi.X509Source{}, true)
				ms.On("Check", mock.Anything, mock.Anything).Return(nil)
				ms.On("CheckInitialization", mock.Anything, mock.Anything).Return(false, nil)
				ms.On("Post", mock.Anything, mock.Anything).Return(nil)
				mfo.On("Open", "/path/to/file").Return(os.NewFile(0, "testfile"), nil)
				ml.On("TraceLn", mock.Anything, mock.Anything).Return()
				ml.On("InfoLn", mock.Anything, mock.Anything).Return()
			},
			expectPanic: false,
		},
		{
			name: "Fail to acquire source",
			setupMocks: func(mfo *MockFileOpener, mer *MockEnvReader, ml *MockLogger, ms *MockSafeOps, msp *MockSpiffeOps) {
				mer.On("InitCommandRunnerWaitBeforeExecIntervalForSentinel").Return(time.Millisecond)
				msp.On("AcquireSourceForSentinel", mock.Anything).Return((*workloadapi.X509Source)(nil), false)
				ml.On("TraceLn", mock.Anything, mock.Anything).Return()
			},
			expectPanic: true,
		},
		{
			name: "Fail API connectivity",
			setupMocks: func(mfo *MockFileOpener, mer *MockEnvReader, ml *MockLogger, ms *MockSafeOps, msp *MockSpiffeOps) {
				mer.On("InitCommandRunnerWaitBeforeExecIntervalForSentinel").Return(time.Millisecond)
				msp.On("AcquireSourceForSentinel", mock.Anything).Return(&workloadapi.X509Source{}, true)
				ms.On("Check", mock.Anything, mock.Anything).Return(errors.New("check error"))
				ml.On("TraceLn", mock.Anything, mock.Anything).Return()
			},
			expectPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockFileOpener := &MockFileOpener{}
			mockEnvReader := &MockEnvReader{}
			mockLogger := &MockLogger{}
			mockSafe := &MockSafeOps{}
			mockSpiffe := &MockSpiffeOps{}

			tt.setupMocks(mockFileOpener, mockEnvReader, mockLogger, mockSafe, mockSpiffe)

			initializer := &Initializer{
				FileOpener: mockFileOpener,
				EnvReader:  mockEnvReader,
				Logger:     mockLogger,
				Safe:       mockSafe,
				Spiffe:     mockSpiffe,
			}

			cid := "test-cid"
			ctx := context.WithValue(context.Background(), key.CorrelationId, &cid)

			if tt.expectPanic {
				assert.Panics(t, func() {
					initializer.RunInitCommands(ctx)
				})
			} else {
				assert.NotPanics(t, func() {
					initializer.RunInitCommands(ctx)
				})
			}

			mockFileOpener.AssertExpectations(t)
			mockEnvReader.AssertExpectations(t)
			mockLogger.AssertExpectations(t)
			mockSafe.AssertExpectations(t)
			mockSpiffe.AssertExpectations(t)
		})
	}
}
