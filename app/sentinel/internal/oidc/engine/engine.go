package engine

import (
	"context"
	"net/http"

	"github.com/vmware-tanzu/secrets-manager/core/constants/key"
	"github.com/vmware-tanzu/secrets-manager/core/crypto"
	"github.com/vmware-tanzu/secrets-manager/core/entity/v1/data"
)

// safeOperations is an interface that defines the methods for safe operations.
type safeOperations interface {
	GetSecrets(ctx context.Context, r *http.Request, encrypt bool) (string, error)
	UpdateSecrets(ctx context.Context,
		r *http.Request, cmd data.SentinelCommand) (string, error)
}

// Authorizer is an interface that defines the methods for authorizing requests.
type authorizer interface {
	IsAuthorized(id string, r *http.Request) bool
}

// Logger is an interface that defines the methods for logging.
type logger interface {
	InfoLn(correlationID *string, v ...any)
	ErrorLn(correlationID *string, v ...any)
}

type Engine struct {
	safeOperations safeOperations
	authorizer     authorizer
	logger         logger
}

// New creates a new Engine instance with the provided safe operations and logger.
// It internally uses the newEngine function to create the Engine instance.
func New(safeOps safeOperations, log logger) *Engine {
	return newEngine(safeOps, newAuth(withLogger(log)), log)
}

// newEngine creates a new Engine instance with the provided safe operations,
// authorizer, and logger.
// It is used internally by the NewEngine function to create the Engine instance.
// Better to use NewEngine function to create the Engine instance for unit testing.
func newEngine(safeOps safeOperations, auth authorizer, log logger) *Engine {
	return &Engine{
		safeOperations: safeOps,
		authorizer:     auth,
		logger:         log,
	}
}

func (e *Engine) createContext() (context.Context, context.CancelFunc) {
	id := crypto.Id()
	return context.WithCancel(
		context.WithValue(context.Background(), key.CorrelationId, id),
	)
}
