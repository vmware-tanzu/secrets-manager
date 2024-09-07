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

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestInitializer_markKeystone(t *testing.T) {
	mp := NewMonkeyPatch(t)
	t.Cleanup(mp.Reset)

	// monkey patching for testing backoff functions in test cases
	mp.ApplyBackoffPatches()

	tests := []struct {
		name        string
		setupMocks  func(*MockLogger, *MockSafeOps, *MockEnvReader)
		expectPanic bool
		expectTrue  bool
	}{
		{
			name: "Successfully mark keystone",
			setupMocks: func(ml *MockLogger, ms *MockSafeOps, me *MockEnvReader) {
				ml.On("TraceLn", mock.Anything, mock.Anything).Return()
				ms.On("Post", mock.Anything, mock.Anything).Return(nil)
				me.On("NamespaceForVSecMSystem").Return("vsecm-system")
			},
			expectPanic: false,
			expectTrue:  true,
		},
		{
			name: "Fail to mark keystone",
			setupMocks: func(ml *MockLogger, ms *MockSafeOps, me *MockEnvReader) {
				ml.On("TraceLn", mock.Anything, mock.Anything).Return()
				ml.On("ErrorLn", mock.Anything, mock.Anything).Return()
				ms.On("Post", mock.Anything, mock.Anything).Return(errors.New("post error"))
				me.On("NamespaceForVSecMSystem").Return("vsecm-system")
			},
			expectPanic: true,
			expectTrue:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockLogger := &MockLogger{}
			mockSafe := &MockSafeOps{}
			mockEnvReader := &MockEnvReader{}

			tt.setupMocks(mockLogger, mockSafe, mockEnvReader)

			initializer := &Initializer{
				Logger:    mockLogger,
				Safe:      mockSafe,
				EnvReader: mockEnvReader,
			}

			ctx := context.Background()
			cid := "test-cid"

			if tt.expectPanic {
				assert.Panics(t, func() {
					initializer.markKeystone(ctx, &cid)
				})
			} else {
				result := initializer.markKeystone(ctx, &cid)
				assert.Equal(t, tt.expectTrue, result)
			}

			mockLogger.AssertExpectations(t)
			mockSafe.AssertExpectations(t)
			mockEnvReader.AssertExpectations(t)
		})
	}
}
