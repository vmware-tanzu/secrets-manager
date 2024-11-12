package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEngine(t *testing.T) {
	safeOps := &MockSafeOperations{}
	authorizer := &MockAuthorizer{}
	logger := &MockLogger{}

	engine := newEngine(safeOps, authorizer, logger)

	assert.NotNil(t, engine)
	assert.IsType(t, &Engine{}, engine)
}

func TestNew(t *testing.T) {
	safeOps := &MockSafeOperations{}
	logger := &MockLogger{}

	engine := New(safeOps, logger)

	assert.NotNil(t, engine)
	assert.IsType(t, &Engine{}, engine)
}
