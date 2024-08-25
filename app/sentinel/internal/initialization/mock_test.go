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
	"github.com/agiledragon/gomonkey/v2"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	"os"
	"testing"
	"time"

	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"github.com/stretchr/testify/mock"
	entity "github.com/vmware-tanzu/secrets-manager/core/entity/v1/data"
)

// MockFileOpener mocks the FileOpener interface
type MockFileOpener struct {
	mock.Mock
}

func (m *MockFileOpener) Open(name string) (*os.File, error) {
	args := m.Called(name)
	return args.Get(0).(*os.File), args.Error(1)
}

// MockEnvReader mocks the EnvReader interface
type MockEnvReader struct {
	mock.Mock
}

func (m *MockEnvReader) InitCommandPathForSentinel() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockEnvReader) InitCommandRunnerWaitBeforeExecIntervalForSentinel() time.Duration {
	args := m.Called()
	return args.Get(0).(time.Duration)
}

func (m *MockEnvReader) InitCommandRunnerWaitIntervalBeforeInitComplete() time.Duration {
	args := m.Called()
	return args.Get(0).(time.Duration)
}

func (m *MockEnvReader) NamespaceForVSecMSystem() string {
	args := m.Called()
	return args.String(0)
}

// MockLogger mocks the Logger interface
type MockLogger struct {
	mock.Mock
}

func (m *MockLogger) InfoLn(correlationID *string, v ...any) {
	m.Called(correlationID, v)
}

func (m *MockLogger) ErrorLn(correlationID *string, v ...any) {
	m.Called(correlationID, v)
}

func (m *MockLogger) TraceLn(correlationID *string, v ...any) {
	m.Called(correlationID, v)
}

func (m *MockLogger) WarnLn(correlationID *string, v ...any) {
	m.Called(correlationID, v)
}

func (m *MockLogger) FatalLn(correlationID *string, v ...any) {
	m.Called(correlationID, v)
}

// MockSafeOps mocks the SafeOps interface
type MockSafeOps struct {
	mock.Mock
}

func (m *MockSafeOps) Check(ctx context.Context, src *workloadapi.X509Source) error {
	args := m.Called(ctx, src)
	return args.Error(0)
}

func (m *MockSafeOps) CheckInitialization(ctx context.Context, src *workloadapi.X509Source) (bool, error) {
	args := m.Called(ctx, src)
	return args.Bool(0), args.Error(1)
}

func (m *MockSafeOps) Post(ctx context.Context, sc entity.SentinelCommand) error {
	args := m.Called(ctx, sc)
	return args.Error(0)
}

// MockSpiffeOps mocks the SpiffeOps interface
type MockSpiffeOps struct {
	mock.Mock
}

func (m *MockSpiffeOps) AcquireSourceForSentinel(ctx context.Context) (*workloadapi.X509Source, bool) {
	args := m.Called(ctx)
	return args.Get(0).(*workloadapi.X509Source), args.Bool(1)
}

// MonkeyPatch is a struct that contains the gomonkey.Patches struct and the testing.T struct.
type MonkeyPatch struct {
	t       *testing.T
	Patches *gomonkey.Patches
}

// NewMonkeyPatch is a constructor function that returns a pointer to the MonkeyPatch struct.
// It is not thread-safe and should be used only in the test functions.
func NewMonkeyPatch(t *testing.T) *MonkeyPatch {
	return &MonkeyPatch{
		t:       t,
		Patches: gomonkey.NewPatches(),
	}
}

// ApplyBackoffPatches is a function that applies the monkey patches for the backoff functions.
func (m *MonkeyPatch) ApplyBackoffPatches() {
	m.Patches.ApplyFuncReturn(env.BackoffMaxRetries, int64(1))
	m.Patches.ApplyFuncReturn(env.BackoffDelay, time.Millisecond) // 1ms for testing
	m.Patches.ApplyFuncReturn(env.BackoffMode, string(env.Exponential))
	m.Patches.ApplyFuncReturn(env.BackoffMaxWait, time.Millisecond) // 1ms for testing
}

// Reset is a function that resets the monkey patches after the test is completed.
func (m *MonkeyPatch) Reset() {
	m.Patches.Reset()
}

// MonkeyPatch is a function that replaces the backoff.RetryExponential function with a mock function for testing.
// This function returns a pointer to the gomonkey.Patches struct which contains the monkey patches applied.
// The t.Cleanup function is used to reset the monkey patches after the test is completed.
//func MonkeyPatch(t *testing.T) {
//	t.Helper()
//	// monkey patching for testing backoff functions in test cases
//	patches := gomonkey.NewPatches()
//	patches.ApplyFuncReturn(env.BackoffMaxRetries, int64(1))
//	patches.ApplyFuncReturn(env.BackoffDelay, time.Millisecond) // 1ms for testing
//	patches.ApplyFuncReturn(env.BackoffMode, string(env.Exponential))
//	patches.ApplyFuncReturn(env.BackoffMaxWait, time.Millisecond) // 1ms for testing
//
//	t.Cleanup(patches.Reset)
//}
