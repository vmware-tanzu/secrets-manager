/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package core

import (
	"context"
	"fmt"
	"github.com/vmware-tanzu/secrets-manager/app/sentinel/rest/safe"
	"github.com/vmware-tanzu/secrets-manager/core/crypto"
	entity "github.com/vmware-tanzu/secrets-manager/core/entity/data/v1"
	"log"
	"net/http"
)

type SecretRequest struct {
	Workload     string   `json:"workload"`
	Secret       string   `json:"secret"`
	Namespaces   []string `json:"namespaces,omitempty"`
	UseK8s       bool     `json:"use-k8s,omitempty"`
	Encrypt      bool     `json:"encrypt,omitempty"`
	Delete       bool     `json:"delete,omitempty"`
	Append       bool     `json:"append,omitempty"`
	List         bool     `json:"list,omitempty"`
	Template     string   `json:"template,omitempty"`
	Format       string   `json:"format,omitempty"`
	BackingStore string   `json:"store,omitempty"`
	InputKeys    string   `json:"input-keys,omitempty"`
	NotBefore    string   `json:"nbf,omitempty"`
	Expires      string   `json:"exp,omitempty"`
}

func HandleCommandSecrets(w http.ResponseWriter, r *http.Request, req *SecretRequest) {
	id, err := crypto.RandomString(8)
	if err != nil {
		id = "VSECSENTINELREST"
	}

	ctx, cancel := context.WithCancel(
		context.WithValue(context.Background(), "correlationId", &id),
	)
	defer cancel()

	ok := IsAuthorizedJWT(id, r)
	if !ok {
		http.Error(w, "isAuthorizedJWT : Please provide correct credentials", http.StatusBadRequest)
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

		fmt.Fprintf(w, responseBody)
		return
	}

	if req.Namespaces == nil || len(req.Namespaces) == 0 {
		req.Namespaces = []string{"default"}
	}

	if InputValidationFailure(req.Workload, req.Encrypt, req.InputKeys, req.Secret, req.Delete) {
		http.Error(w, "Input Validation Failure", http.StatusInternalServerError)
		return
	}

	responseBody, err := safe.Post(ctx, r,
		entity.SentinelCommand{
			WorkloadId:    req.Workload,
			Secret:        req.Secret,
			Namespaces:    req.Namespaces,
			BackingStore:  req.BackingStore,
			UseKubernetes: req.UseK8s,
			Template:      req.Template,
			Format:        req.Format,
			Encrypt:       req.Encrypt,
			DeleteSecret:  req.Delete,
			AppendSecret:  req.Append,
			InputKeys:     req.InputKeys,
			NotBefore:     req.NotBefore,
			Expires:       req.Expires,
		})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	log.Println(&id, responseBody)
	fmt.Fprintf(w, responseBody)
}
