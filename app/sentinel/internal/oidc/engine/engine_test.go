package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEngine(t *testing.T) {
	safeOps := new(MockSafeOperations)
	authorizer := new(MockAuthorizer)
	logger := new(MockLogger)

	engine := newEngine(safeOps, authorizer, logger)

	assert.NotNil(t, engine)
	assert.IsType(t, &Engine{}, engine)
}
