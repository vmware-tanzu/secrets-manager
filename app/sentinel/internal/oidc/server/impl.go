package server

import (
	"context"
	"net/http"

	"github.com/vmware-tanzu/secrets-manager/app/sentinel/internal/oidc/safe"
	"github.com/vmware-tanzu/secrets-manager/core/entity/v1/data"
	"github.com/vmware-tanzu/secrets-manager/core/log/rpc"
)

// SafeOperations is an interface that defines the methods for safe operations.
type SafeOperations interface {
	GetSecrets(ctx context.Context, r *http.Request, encrypt bool) (string, error)
	UpdateSecrets(ctx context.Context, r *http.Request, cmd data.SentinelCommand) (string, error)
}

// Authorizer is an interface that defines the methods for authorizing requests.
type Authorizer interface {
	IsAuthorized(id string, r *http.Request) bool
}

// Logger is an interface that defines the methods for logging.
type Logger interface {
	InfoLn(correlationID *string, v ...any)
	ErrorLn(correlationID *string, v ...any)
}

type SafeClient struct{}

var _ SafeOperations = (*SafeClient)(nil)

func (SafeClient) GetSecrets(ctx context.Context, r *http.Request, encrypt bool) (string, error) {
	return safe.Get(ctx, r, encrypt)
}

func (SafeClient) UpdateSecrets(ctx context.Context, r *http.Request, cmd data.SentinelCommand) (string, error) {
	return safe.Post(ctx, r, cmd)
}

type RpcLogger struct{}

var _ Logger = (*RpcLogger)(nil)

func (RpcLogger) InfoLn(correlationID *string, v ...any) {
	rpc.InfoLn(correlationID, v...)
}

func (RpcLogger) ErrorLn(correlationID *string, v ...any) {
	rpc.ErrorLn(correlationID, v...)
}
