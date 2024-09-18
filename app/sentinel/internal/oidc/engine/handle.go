/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package engine

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/vmware-tanzu/secrets-manager/core/constants/key"
	"github.com/vmware-tanzu/secrets-manager/core/entity/v1/data"
	"github.com/vmware-tanzu/secrets-manager/core/entity/v1/reqres/sentinel"
)

const defaultNamespace = "default"

var (
	ErrUnsupportedMethod  = errors.New("unsupported method")
	ErrInvalidInput       = errors.New("invalid input for secret modification")
	ErrInvalidRequestBody = errors.New("invalid request body")
	ErrUnauthorized       = errors.New("unauthorized: please provide correct credentials")
)

// HandleSecrets processes incoming HTTP requests related to secrets management.
// This function is specifically designed to handle POST requests.
// It handles both listing and modifying secrets based on the provided SecretRequest.
// It ensures that the request is authorized using JWT and processes
// the request accordingly.
//
// Parameters:
//   - w: the http.ResponseWriter used to write the HTTP response.
//   - r: the *http.Request containing all the request data including headers,
//     query parameters, and the body.
//
// The function performs the following steps:
//  1. Checks if the request method is POST.
//  2. Decodes the request body into a SecretRequest struct.
//  3. Generates a unique correlation ID for the request.
//  4. Creates a context with the correlation ID for logging and tracing.
//  5. Checks if the request is authorized using JWT.
//  6. If the request is authorized, it processes the request based on the
//     `req.List` flag.
//     - If `req.List` is true, it retrieves and sends the list of secrets.
//     - If `req.List` is false, it validates the input and performs the
//     requested secret management action (e.g., create, update, delete).
//
// The function returns appropriate HTTP status codes and messages in case
// of errors.
//
// Example request body:
//
//	{
//	    "list": true,
//	    "encrypt": false,
//	    "workloads": ["workload1"],
//	    "secret": "my-secret",
//	    "namespaces": ["namespace1"],
//	}
//
// Usage:
//
//	http.HandleFunc("/api/secrets", engine.HandleSecrets)
//
// Errors:
//   - Returns HTTP 400 if the request body cannot be decoded into
//     SecretRequest.
//   - Returns HTTP 400 if the request body is invalid for secret modification.
//   - Returns HTTP 401 if the request is not authorized.
//   - Returns HTTP 405 if the method is not POST.
//   - Returns HTTP 500 if there is an error during secret retrieval or modification.
func (e *Engine) HandleSecrets(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, ErrUnsupportedMethod.Error(), http.StatusMethodNotAllowed)
		return
	}

	var req sentinel.SecretRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, ErrInvalidRequestBody.Error(), http.StatusBadRequest)
		return
	}

	ctx, cancel := e.createContext()
	defer cancel()

	if !e.authorizer.IsAuthorized(ctx.Value(key.CorrelationId).(string), r) {
		http.Error(w, ErrUnauthorized.Error(), http.StatusUnauthorized)
		return
	}

	if req.List {
		e.handleListSecrets(ctx, w, r, &req)
	} else {
		e.handleModifySecrets(ctx, w, r, &req)
	}
}

// handleListSecrets handles the listing of secrets.
// It retrieves the secrets from the safe and sends them to the client.
func (e *Engine) handleListSecrets(ctx context.Context,
	w http.ResponseWriter, r *http.Request, req *sentinel.SecretRequest) {
	secrets, err := e.safeOperations.GetSecrets(ctx, r, req.Encrypt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	e.sendJSONResponse(w, secrets)
}

// handleModifySecrets handles the modification of secrets.
// It validates the input and performs the requested secret management action.
func (e *Engine) handleModifySecrets(ctx context.Context, w http.ResponseWriter, r *http.Request, req *sentinel.SecretRequest) {
	if err := e.validateModifyRequest(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cmd := e.createSentinelCommand(req)
	result, err := e.safeOperations.UpdateSecrets(ctx, r, cmd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	e.sendJSONResponse(w, result)
}

// validateModifyRequest validates the input for secret modification.
func (e *Engine) validateModifyRequest(req *sentinel.SecretRequest) error {
	if len(req.Namespaces) == 0 {
		req.Namespaces = []string{defaultNamespace}
	}

	if !isValidSecretModification(req.Workloads, req.Encrypt,
		req.SerializedRootKeys, req.Secret, req.Delete) {
		return ErrInvalidInput
	}
	return nil
}

// createSentinelCommand creates a SentinelCommand from a SecretRequest.
func (e *Engine) createSentinelCommand(
	req *sentinel.SecretRequest) data.SentinelCommand {
	return data.SentinelCommand{
		WorkloadIds:        req.Workloads,
		Secret:             req.Secret,
		Namespaces:         req.Namespaces,
		Template:           req.Template,
		Format:             req.Format,
		Encrypt:            req.Encrypt,
		DeleteSecret:       req.Delete,
		AppendSecret:       req.Append,
		SerializedRootKeys: req.SerializedRootKeys,
		NotBefore:          req.NotBefore,
		Expires:            req.Expires,
	}
}

// sendJSONResponse sends the response to the client.
// It sets the content type to application/json and writes the response body.
// If there is an error in writing the response, it logs the error.
func (e *Engine) sendJSONResponse(w http.ResponseWriter, data string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err := fmt.Fprint(w, data); err != nil {
		e.logger.ErrorLn(nil, "Error in writing response", err.Error())
	}
}
