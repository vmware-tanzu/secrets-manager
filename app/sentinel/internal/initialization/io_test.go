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
	"bufio"
	"context"
	"errors"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestInitializer_commandFileScanner(t *testing.T) {
	tests := []struct {
		name          string
		setupMocks    func(*MockFileOpener, *MockEnvReader, *MockLogger)
		expectFile    bool
		expectScanner bool
	}{
		{
			name: "Successfully open file",
			setupMocks: func(mfo *MockFileOpener, mer *MockEnvReader, ml *MockLogger) {
				mer.On("InitCommandPathForSentinel").Return("/path/to/file")
				mfo.On("Open", "/path/to/file").Return(os.NewFile(0, "testfile"), nil)
				ml.On("TraceLn", mock.Anything, mock.Anything).Return()
			},
			expectFile:    true,
			expectScanner: true,
		},
		{
			name: "Fail to open file",
			setupMocks: func(mfo *MockFileOpener, mer *MockEnvReader, ml *MockLogger) {
				mer.On("InitCommandPathForSentinel").Return("/path/to/file")
				mfo.On("Open", "/path/to/file").Return((*os.File)(nil), errors.New("file not found"))
				ml.On("InfoLn", mock.Anything, mock.Anything).Return()
			},
			expectFile:    false,
			expectScanner: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockFileOpener := &MockFileOpener{}
			mockEnvReader := &MockEnvReader{}
			mockLogger := &MockLogger{}

			tt.setupMocks(mockFileOpener, mockEnvReader, mockLogger)

			initializer := &Initializer{
				FileOpener: mockFileOpener,
				EnvReader:  mockEnvReader,
				Logger:     mockLogger,
			}

			cid := "test-cid"
			file, scanner := initializer.commandFileScanner(&cid)

			if tt.expectFile {
				assert.NotNil(t, file)
			} else {
				assert.Nil(t, file)
			}

			if tt.expectScanner {
				assert.NotNil(t, scanner)
			} else {
				assert.Nil(t, scanner)
			}

			mockFileOpener.AssertExpectations(t)
			mockEnvReader.AssertExpectations(t)
			mockLogger.AssertExpectations(t)
		})
	}
}

func TestInitializer_parseCommandsFile(t *testing.T) {
	mp := NewMonkeyPatch(t)
	t.Cleanup(mp.Reset)

	// monkey patching for testing backoff functions in test cases
	mp.ApplyBackoffPatches()

	tests := []struct {
		name        string
		input       string
		setupMocks  func(*MockLogger, *MockSafeOps)
		expectPanic bool
	}{
		{
			name:  "Successfully parse commands - workload, secret, and exit",
			input: "w:workload1\ns:secret1\n--\nexit\n",
			setupMocks: func(ml *MockLogger, ms *MockSafeOps) {
				ml.On("TraceLn", mock.Anything, mock.Anything).Return()
				ml.On("InfoLn", mock.Anything, mock.Anything).Return()
				ms.On("Post", mock.Anything, mock.Anything).Return(nil)
			},
			expectPanic: false,
		},
		{
			name:  "Successfully parse commands - workload, secret, transformation, and exit",
			input: "w:workload1\ns:secret1\nt:1000\n--\nexit\n",
			setupMocks: func(ml *MockLogger, ms *MockSafeOps) {
				ml.On("TraceLn", mock.Anything, mock.Anything).Return()
				ml.On("InfoLn", mock.Anything, mock.Anything).Return()
				ms.On("Post", mock.Anything, mock.Anything).Return(nil)
			},
			expectPanic: false,
		},
		{
			name:  "Successfully parse commands - workload, secret, transformation, sleep, and exit",
			input: "w:workload1\ns:secret1\nt:1000\nsleep:1000\n--\nexit\n",
			setupMocks: func(ml *MockLogger, ms *MockSafeOps) {
				ml.On("TraceLn", mock.Anything, mock.Anything).Return()
				ml.On("InfoLn", mock.Anything, mock.Anything).Return()
			},
			expectPanic: false,
		},
		{
			name:  "Successfully parse commands - workload, secret, encryption, and exit",
			input: "w:workload1\ns:secret1\ne:encryption1\n--\nexit\n",
			setupMocks: func(ml *MockLogger, ms *MockSafeOps) {
				ml.On("TraceLn", mock.Anything, mock.Anything).Return()
				ml.On("InfoLn", mock.Anything, mock.Anything).Return()
				ms.On("Post", mock.Anything, mock.Anything).Return(nil)
			},
			expectPanic: false,
		},
		{
			name:  "Successfully parse commands - workload, secret, remove, and exit",
			input: "w:workload1\ns:secret1\nd:remove1\n--\nexit\n",
			setupMocks: func(ml *MockLogger, ms *MockSafeOps) {
				ml.On("TraceLn", mock.Anything, mock.Anything).Return()
				ml.On("InfoLn", mock.Anything, mock.Anything).Return()
				ms.On("Post", mock.Anything, mock.Anything).Return(nil)
			},
			expectPanic: false,
		},
		{
			name:  "Successfully parse commands - workload, secret, join, and exit",
			input: "w:workload1\ns:secret1\na:join1\n--\nexit\n",
			setupMocks: func(ml *MockLogger, ms *MockSafeOps) {
				ml.On("TraceLn", mock.Anything, mock.Anything).Return()
				ml.On("InfoLn", mock.Anything, mock.Anything).Return()
				ms.On("Post", mock.Anything, mock.Anything).Return(nil)
			},
			expectPanic: false,
		},
		{
			name:  "Successfully parse commands - workload, secret, format, and exit",
			input: "w:workload1\ns:secret1\nf:format1\n--\nexit\n",
			setupMocks: func(ml *MockLogger, ms *MockSafeOps) {
				ml.On("TraceLn", mock.Anything, mock.Anything).Return()
				ml.On("InfoLn", mock.Anything, mock.Anything).Return()
				ms.On("Post", mock.Anything, mock.Anything).Return(nil)
			},
			expectPanic: false,
		},
		{
			name:  "Successfully parse commands - workload, secret, keys, and exit",
			input: "w:workload1\ns:secret1\ni:keys1\n--\nexit\n",
			setupMocks: func(ml *MockLogger, ms *MockSafeOps) {
				ml.On("TraceLn", mock.Anything, mock.Anything).Return()
				ml.On("InfoLn", mock.Anything, mock.Anything).Return()
				ms.On("Post", mock.Anything, mock.Anything).Return(nil)
			},
			expectPanic: false,
		},
		{
			name:  "Fail to process command block",
			input: "w:workload1\ns:secret1\n--\n",
			setupMocks: func(ml *MockLogger, ms *MockSafeOps) {
				ml.On("TraceLn", mock.Anything, mock.Anything).Return()
				ml.On("ErrorLn", mock.Anything, mock.Anything).Return()
				ms.On("Post", mock.Anything, mock.Anything).Return(errors.New("post error"))
			},
			expectPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockLogger := &MockLogger{}
			mockSafe := &MockSafeOps{}

			tt.setupMocks(mockLogger, mockSafe)

			initializer := &Initializer{
				Logger: mockLogger,
				Safe:   mockSafe,
			}

			ctx := context.Background()
			cid := "test-cid"
			scanner := bufio.NewScanner(strings.NewReader(tt.input))

			if tt.expectPanic {
				assert.Panics(t, func() {
					initializer.parseCommandsFile(ctx, &cid, scanner)
				})
			} else {
				assert.NotPanics(t, func() {
					initializer.parseCommandsFile(ctx, &cid, scanner)
				})
			}

			mockLogger.AssertExpectations(t)
			mockSafe.AssertExpectations(t)
		})
	}
}
