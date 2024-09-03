package engine

import (
	"context"
	"net/http"

	"github.com/stretchr/testify/mock"
	"github.com/vmware-tanzu/secrets-manager/core/entity/v1/data"
)

type MockSafeOperations struct {
	mock.Mock
}

var _ safeOperations = (*MockSafeOperations)(nil)

func (m *MockSafeOperations) GetSecrets(ctx context.Context, r *http.Request, encrypt bool) (string, error) {
	args := m.Called(ctx, r, encrypt)
	return args.String(0), args.Error(1)
}

func (m *MockSafeOperations) UpdateSecrets(ctx context.Context, r *http.Request, cmd data.SentinelCommand) (string, error) {
	args := m.Called(ctx, r, cmd)
	return args.String(0), args.Error(1)
}

type MockAuthorizer struct {
	mock.Mock
}

var _ authorizer = (*MockAuthorizer)(nil)

func (m *MockAuthorizer) IsAuthorized(id string, r *http.Request) bool {
	args := m.Called(id, r)
	return args.Bool(0)
}

type MockLogger struct {
	mock.Mock
}

var _ logger = (*MockLogger)(nil)

func (m *MockLogger) InfoLn(correlationID *string, v ...any) {
	m.Called(correlationID, v)
}

func (m *MockLogger) ErrorLn(correlationID *string, v ...any) {
	m.Called(correlationID, v)
}
