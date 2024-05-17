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
	"github.com/vmware-tanzu/secrets-manager/core/crypto"
	entity "github.com/vmware-tanzu/secrets-manager/core/entity/v1/data"
)

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

func HandleCommandSecrets(
	w http.ResponseWriter, r *http.Request, req *SecretRequest,
) {
	id := crypto.Id()

	ctx, cancel := context.WithCancel(
		context.WithValue(context.Background(), "correlationId", &id),
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
