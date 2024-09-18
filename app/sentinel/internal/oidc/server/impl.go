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
	GetSecrets(ctx context.Context,
		r *http.Request, encrypt bool) (string, error)
	UpdateSecrets(ctx context.Context,
		r *http.Request, cmd data.SentinelCommand) (string, error)
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

// SafeClient implements the SafeOperations interface for interacting
// with the VSecM Safe.
type SafeClient struct{}

var _ SafeOperations = (*SafeClient)(nil)

// GetSecrets retrieves secrets from the safe.
// It takes a context, an HTTP request, and a boolean indicating whether to
// encrypt the secrets.
// It returns the secrets as a string and any error encountered.
func (SafeClient) GetSecrets(ctx context.Context,
	r *http.Request, encrypt bool) (string, error) {
	return safe.Get(ctx, r, encrypt)
}

// UpdateSecrets updates secrets in the safe.
// It takes a context, an HTTP request, and a SentinelCommand containing
// the update details.
// It returns the result as a string and any error encountered.
func (SafeClient) UpdateSecrets(ctx context.Context,
	r *http.Request, cmd data.SentinelCommand) (string, error) {
	return safe.Post(ctx, r, cmd)
}

// RpcLogger implements the Logger interface for RPC logging.
type RpcLogger struct{}

var _ Logger = (*RpcLogger)(nil)

// InfoLn logs information messages.
// It takes a correlation ID and a variable number of arguments to log.
func (RpcLogger) InfoLn(correlationID *string, v ...any) {
	rpc.InfoLn(correlationID, v...)
}

// ErrorLn logs error messages.
// It takes a correlation ID and a variable number of arguments to log.
func (RpcLogger) ErrorLn(correlationID *string, v ...any) {
	rpc.ErrorLn(correlationID, v...)
}
