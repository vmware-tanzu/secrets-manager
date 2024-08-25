package initialization

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewInitializer(t *testing.T) {
	mockFileOpener := &MockFileOpener{}
	mockEnvReader := &MockEnvReader{}
	mockLogger := &MockLogger{}
	mockSafeOps := &MockSafeOps{}
	mockSpiffeOps := &MockSpiffeOps{}

	initializer := NewInitializer(
		mockFileOpener,
		mockEnvReader,
		mockLogger,
		mockSafeOps,
		mockSpiffeOps,
	)

	assert.NotNil(t, initializer)
	assert.Equal(t, mockFileOpener, initializer.FileOpener)
	assert.Equal(t, mockEnvReader, initializer.EnvReader)
	assert.Equal(t, mockLogger, initializer.Logger)
	assert.Equal(t, mockSafeOps, initializer.Safe)
	assert.Equal(t, mockSpiffeOps, initializer.Spiffe)
}
