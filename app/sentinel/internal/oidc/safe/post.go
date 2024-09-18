/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package safe

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
	"github.com/spiffe/go-spiffe/v2/workloadapi"

	"github.com/vmware-tanzu/secrets-manager/core/constants/key"
	"github.com/vmware-tanzu/secrets-manager/core/entity/v1/data"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	"github.com/vmware-tanzu/secrets-manager/lib/template"
)

// Post handles the HTTP POST request for secret management using the provided
// SentinelCommand.
//
// This function performs the following steps:
//  1. Creates a context with a timeout based on the parent context and
//     environment settings.
//  2. Acquires a workload source and proceeds only if the source acquisition
//     is successful.
//  3. Depending on the SentinelCommand, it either posts new secrets or deletes
//     existing ones.
//
// Parameters:
//   - parentContext: The parent context for the request, used for tracing and
//     cancellation.
//   - r: The HTTP request being processed.
//   - sc: The SentinelCommand containing details for the secret management
//     operation.
//
// Returns:
//   - A string representing the response body or an error if the operation
//     fails.
//
// Example usage:
//
//	parentContext := context.Background()
//	r, _ := http.NewRequest("POST", "http://example.com", nil)
//	sc := data.SentinelCommand{
//	    WorkloadIds:        []string{"workload1"},
//	    Secret:             "my-secret",
//	    Namespaces:         []string{"namespace1"},
//	    SerializedRootKeys: "key1\nkey2\nkey3",
//	}
//	response, err := Post(parentContext, r, sc)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(response)
//
// Error Handling:
//   - If the context times out or is canceled, it logs the error and returns
//     an appropriate message.
//   - If there is an error during source acquisition, secret generation, or
//     payload processing, it returns an error with details.
func Post(
	parentContext context.Context, r *http.Request, sc data.SentinelCommand,
) (string, error) {
	ctxWithTimeout, cancel := context.WithTimeout(
		parentContext,
		env.SourceAcquisitionTimeoutForSafe(),
	)
	defer cancel()

	cid := ctxWithTimeout.Value(key.CorrelationId).(*string)

	sourceChan := make(chan *workloadapi.X509Source)
	proceedChan := make(chan bool)

	go func() {
		source, proceed := acquireSource(ctxWithTimeout)
		sourceChan <- source
		proceedChan <- proceed
	}()
	select {
	case <-ctxWithTimeout.Done():
		if errors.Is(ctxWithTimeout.Err(), context.DeadlineExceeded) {
			log.Println(cid,
				"Post: I cannot execute command because I cannot talk to SPIRE.")
			return "",
				errors.New(
					"post: I cannot execute command because I cannot talk to SPIRE")
		}

		log.Println(cid, "Post: Operation was cancelled due to an unknown reason.")
	case source := <-sourceChan:
		defer func(s *workloadapi.X509Source) {
			if s == nil {
				return
			}
			err := s.Close()
			if err != nil {
				log.Println(cid, "Post: Problem closing the workload source.")
			}
		}(source)

		proceed := <-proceedChan

		if !proceed {
			return "", printPayloadError(cid,
				errors.New("post: Could not proceed"))
		}

		authorizer := createAuthorizer()

		if sc.SerializedRootKeys != "" {
			p, err := url.JoinPath(env.EndpointUrlForSafe(), "/sentinel/v1/keys")
			if err != nil {
				return "", printEndpointError(cid, err)
			}

			tlsConfig := tlsconfig.MTLSClientConfig(source, source, authorizer)
			client := &http.Client{
				Transport: &http.Transport{
					TLSClientConfig: tlsConfig,
				},
			}

			parts := sc.SplitRootKeys()

			if len(parts) != 3 {
				return "", printPayloadError(
					cid, errors.New("post: Bad data! Very bad data"))
			}

			sr := newInputKeysRequest(parts[0], parts[1], parts[2])
			md, err := json.Marshal(sr)
			if err != nil {
				return "", printPayloadError(cid, err)
			}

			return doPost(cid, client, r, p, md)
		}

		// Generate pattern-based random secrets if the secret has the prefix.
		if strings.HasPrefix(sc.Secret, env.SecretGenerationPrefix()) {
			sc.Secret = strings.Replace(
				sc.Secret, env.SecretGenerationPrefix(), "", 1,
			)
			newSecret, err := template.Value(sc.Secret)
			if err != nil {
				sc.Secret = "ParseError:" + sc.Secret
			} else {
				sc.Secret = newSecret
			}
		}

		p, err := url.JoinPath(env.EndpointUrlForSafe(), "/sentinel/v1/secrets")
		if err != nil {
			return "", printEndpointError(cid, err)
		}

		tlsConfig := tlsconfig.MTLSClientConfig(source, source, authorizer)
		client := &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: tlsConfig,
			},
		}

		sr := newSecretUpsertRequest(sc.WorkloadIds, sc.Secret, sc.Namespaces,
			sc.Template, sc.Format,
			sc.Encrypt, sc.AppendSecret, sc.NotBefore, sc.Expires)

		md, err := json.Marshal(sr)
		if err != nil {
			return "", printPayloadError(cid, err)

		}

		if sc.DeleteSecret {
			return doDelete(cid, client, r, p, md)
		}

		return doPost(cid, client, r, p, md)
	}
	return "", fmt.Errorf("post: An error occured")
}
