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
	"fmt"
	"log"
	"net/http"

	"github.com/vmware-tanzu/secrets-manager/app/sentinel/internal/oidc/safe"
	"github.com/vmware-tanzu/secrets-manager/core/constants/key"
	"github.com/vmware-tanzu/secrets-manager/core/crypto"
	entity "github.com/vmware-tanzu/secrets-manager/core/entity/v1/data"
)

// SecretRequest encapsulates a VSecM Safe REST command payload.
type SecretRequest struct {
	Workloads          []string `json:"workload"`
	Secret             string   `json:"secret"`
	Namespaces         []string `json:"namespaces,omitempty"`
	Encrypt            bool     `json:"encrypt,omitempty"`
	Delete             bool     `json:"delete,omitempty"`
	Append             bool     `json:"append,omitempty"`
	List               bool     `json:"list,omitempty"`
	Template           string   `json:"template,omitempty"`
	Format             string   `json:"format,omitempty"`
	SerializedRootKeys string   `json:"root-keys,omitempty"`
	NotBefore          string   `json:"nbf,omitempty"`
	Expires            string   `json:"exp,omitempty"`
}

// HandleCommandSecrets processes HTTP requests related to secret management.
//
// This function handles both listing and modifying secrets based on the
// provided SecretRequest.
// It ensures that the request is authorized using JWT and processes
// the request accordingly.
//
// Parameters:
//   - w: The HTTP response writer to send the response.
//   - r: The HTTP request being processed.
//   - req: A pointer to a SecretRequest struct containing the details of the
//     secret management request.
//
// The function performs the following steps:
//  1. Generates a unique correlation ID for the request.
//  2. Creates a context with the correlation ID for logging and tracing.
//  3. Checks if the request is authorized using JWT.
//  4. If the request is authorized, it processes the request based on the
//     `req.List` flag.
//     - If `req.List` is true, it retrieves and sends the list of secrets.
//     - If `req.List` is false, it validates the input and performs the
//     requested secret management action (e.g., create, update, delete).
//
// The function returns appropriate HTTP status codes and messages in case
// of errors.
//
// Example usage:
//
//	req := &SecretRequest{
//	    List:       true,
//	    Encrypt:    false,
//	    Workloads:  []string{"workload1"},
//	    Secret:     "my-secret",
//	    Namespaces: []string{"namespace1"},
//	}
//	HandleCommandSecrets(w, r, req)
//
// Error Handling:
//   - If the request is not authorized, it returns a 400 Bad Request status
//     with an appropriate message.
//   - If there is an error during secret retrieval or modification, it returns
//     a 500 Internal Server Error status with the error message.
func HandleCommandSecrets(
	w http.ResponseWriter, r *http.Request, req *SecretRequest,
) {
	id := crypto.Id()

	ctx, cancel := context.WithCancel(
		context.WithValue(context.Background(),
			key.CorrelationId, &id),
	)
	defer cancel()

	ok := AuthorizedJWT(id, r)
	if !ok {
		http.Error(w,
			"isAuthorizedJWT : Please provide correct credentials",
			http.StatusBadRequest)
		return
	}

	if req.List {
		encrypt := req.Encrypt
		responseBody, err := safe.Get(ctx, r, encrypt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		_, err = fmt.Fprintf(w, responseBody)
		if err != nil {
			log.Println("Error in writing response", err.Error())
			return
		}
		return
	}

	if req.Namespaces == nil || len(req.Namespaces) == 0 {
		req.Namespaces = []string{"default"}
	}

	if invalidInput(req.Workloads, req.Encrypt,
		req.SerializedRootKeys, req.Secret, req.Delete,
	) {
		http.Error(w, "Input Validation Failure",
			http.StatusInternalServerError)
		return
	}

	responseBody, err := safe.Post(ctx, r,
		entity.SentinelCommand{
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
		})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	log.Println(&id, responseBody)
	_, err = fmt.Fprintf(w, responseBody)
	if err != nil {
		log.Println("Error in writing response", err.Error())
		return
	}
}
