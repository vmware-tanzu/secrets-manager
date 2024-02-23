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
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
	"github.com/spiffe/go-spiffe/v2/workloadapi"

	"github.com/vmware-tanzu/secrets-manager/core/crypto"
	data "github.com/vmware-tanzu/secrets-manager/core/entity/data/v1"
	entity "github.com/vmware-tanzu/secrets-manager/core/entity/data/v1"
	reqres "github.com/vmware-tanzu/secrets-manager/core/entity/reqres/safe/v1"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	log "github.com/vmware-tanzu/secrets-manager/core/log/rpc"
	"github.com/vmware-tanzu/secrets-manager/core/validation"
)

func createAuthorizer() tlsconfig.Authorizer {
	return tlsconfig.AdaptMatcher(func(id spiffeid.ID) error {
		if validation.IsSafe(id.String()) {
			return nil
		}

		return errors.New("Post: I don’t know you, and it’s crazy: '" +
			id.String() + "'",
		)
	})
}

func decideBackingStore(backingStore string) data.BackingStore {
	switch data.BackingStore(backingStore) {
	case data.File:
		return data.File
	case data.Memory:
		return data.Memory
	default:
		return env.BackingStoreForSafe()
	}
}

func decideSecretFormat(format string) data.SecretFormat {
	switch data.SecretFormat(format) {
	case data.Json:
		return data.Json
	case data.Yaml:
		return data.Yaml
	default:
		return data.Json
	}
}

func newInputKeysRequest(ageSecretKey, agePublicKey, aesCipherKey string,
) reqres.KeyInputRequest {
	return reqres.KeyInputRequest{
		AgeSecretKey: ageSecretKey,
		AgePublicKey: agePublicKey,
		AesCipherKey: aesCipherKey,
	}
}

func newInitCompletedRequest() reqres.SentinelInitCompleteRequest {
	return reqres.SentinelInitCompleteRequest{}
}

func newSecretUpsertRequest(workloadId, secret string, namespaces []string,
	backingStore string, useKubernetes bool, template string, format string,
	encrypt, appendSecret bool, notBefore string, expires string,
) reqres.SecretUpsertRequest {
	bs := decideBackingStore(backingStore)
	f := decideSecretFormat(format)

	if notBefore == "" {
		notBefore = "now"
	}

	if expires == "" {
		expires = "never"
	}

	return reqres.SecretUpsertRequest{
		WorkloadId:    workloadId,
		BackingStore:  bs,
		Namespaces:    namespaces,
		UseKubernetes: useKubernetes,
		Template:      template,
		Format:        f,
		Encrypt:       encrypt,
		AppendValue:   appendSecret,
		Value:         secret,
		NotBefore:     notBefore,
		Expires:       expires,
	}
}

func respond(cid *string, r *http.Response) {
	if r == nil {
		return
	}

	defer func(b io.ReadCloser) {
		if b == nil {
			return
		}
		err := b.Close()
		if err != nil {
			log.ErrorLn(cid, "Post: Problem closing request body.", err.Error())
		}
	}(r.Body)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.ErrorLn(cid, "Post: Unable to read the response body from VSecM Safe.", err.Error())
		return
	}

	fmt.Println("")
	fmt.Println(string(body))
	fmt.Println("")
}

func printEndpointError(cid *string, err error) {
	log.ErrorLn(cid, "Post: I am having problem generating VSecM Safe "+
		"secrets api endpoint URL.", err.Error())
}

func printPayloadError(cid *string, err error) {
	log.ErrorLn(cid, "Post: I am having problem generating the payload.", err.Error())
}

func doDelete(cid *string, client *http.Client, p string, md []byte) {
	req, err := http.NewRequest(http.MethodDelete, p, bytes.NewBuffer(md))
	if err != nil {
		log.ErrorLn(cid, "Post:Delete: Problem connecting to VSecM Safe API endpoint URL.", err.Error())
		return
	}
	req.Header.Set("Content-Type", "application/json")
	r, err := client.Do(req)
	if err != nil {
		log.ErrorLn(cid, "Post:Delete: Problem connecting to VSecM Safe API endpoint URL.", err.Error())
		return
	}
	respond(cid, r)
}

func doPost(cid *string, client *http.Client, p string, md []byte) {
	r, err := client.Post(p, "application/json", bytes.NewBuffer(md))
	if err != nil {
		log.ErrorLn(cid, "Post: Problem connecting to VSecM Safe API endpoint URL.", err.Error())
		return
	}
	respond(cid, r)
}

// PostInitializationComplete is a function that signals the completion of a
// post-initialization process.
// It takes a parent context as an argument and performs several steps involving
// timeout management, source acquisition, error handling, and sending a
// notification about the initialization completion.
//
// In a separate goroutine, it tries to acquire a source and sends the source and
// a proceed signal back to the main function through channels. The main function
// then waits for either a timeout or a source to be returned.
//
// If a timeout occurs, it logs an error depending on whether it’s due to deadline
// exceeded or an unknown reason. If a source is received, it checks whether to
// proceed. If not, it returns early.
//
// If proceeding, the function then creates an authorizer and builds a client with
// mutual TLS configuration. It creates a new request payload, marshals it to
// JSON, and sends a POST request to a specified endpoint.
//
// Parameters:
//   - parentContext (context.Context): The parent context from which the function
//     will derive its context.
func PostInitializationComplete(parentContext context.Context) {
	ctxWithTimeout, cancel := context.WithTimeout(
		parentContext,
		env.SourceAcquisitionTimeoutForSafe(),
	)
	defer cancel()

	cid := ctxWithTimeout.Value("correlationId").(*string)

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
			log.ErrorLn(cid, "PostInit: I cannot execute command because I cannot talk to SPIRE.")
			return
		}

		log.ErrorLn(cid, "PostInit: Operation was cancelled due to an unknown reason.")
	case source := <-sourceChan:
		defer func() {
			if source == nil {
				return
			}
			err := source.Close()
			if err != nil {
				log.ErrorLn(cid, "Post: Problem closing the workload source.")
			}
		}()

		proceed := <-proceedChan

		if !proceed {
			return
		}

		authorizer := createAuthorizer()

		p, err := url.JoinPath(env.EndpointUrlForSafe(), "/sentinel/v1/init-completed")
		if err != nil {
			printEndpointError(cid, err)
			return
		}

		tlsConfig := tlsconfig.MTLSClientConfig(source, source, authorizer)
		client := &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: tlsConfig,
			},
		}

		sr := newInitCompletedRequest()

		md, err := json.Marshal(sr)
		if err != nil {
			printPayloadError(cid, err)
			return
		}

		doPost(cid, client, p, md)
	}
}

func Post(parentContext context.Context,
	sc entity.SentinelCommand,
) {
	ctxWithTimeout, cancel := context.WithTimeout(
		parentContext,
		env.SourceAcquisitionTimeoutForSafe(),
	)
	defer cancel()

	cid := ctxWithTimeout.Value("correlationId").(*string)

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
			log.ErrorLn(
				cid,
				"Post: I cannot execute command because I cannot talk to SPIRE.",
			)
			return
		}

		log.ErrorLn(
			cid,
			"Post: Operation was cancelled due to an unknown reason.",
		)
	case source := <-sourceChan:
		defer func() {
			if source == nil {
				return
			}
			err := source.Close()
			if err != nil {
				log.ErrorLn(cid, "Post: Problem closing the workload source.")
			}
		}()

		proceed := <-proceedChan

		if !proceed {
			return
		}

		authorizer := createAuthorizer()

		if sc.InputKeys != "" {
			p, err := url.JoinPath(env.EndpointUrlForSafe(), "/sentinel/v1/keys")
			if err != nil {
				printEndpointError(cid, err)
				return
			}

			tlsConfig := tlsconfig.MTLSClientConfig(source, source, authorizer)
			client := &http.Client{
				Transport: &http.Transport{
					TLSClientConfig: tlsConfig,
				},
			}

			parts := strings.Split(sc.InputKeys, "\n")
			if len(parts) != 3 {
				printPayloadError(cid, errors.New("post: Bad data! Very bad data"))
				return
			}

			sr := newInputKeysRequest(parts[0], parts[1], parts[2])
			md, err := json.Marshal(sr)
			if err != nil {
				printPayloadError(cid, err)
				return
			}

			doPost(cid, client, p, md)
			return
		}

		// Generate pattern-based random secrets if the secret has the prefix.
		if strings.HasPrefix(sc.Secret, env.SecretGenerationPrefix()) {
			sc.Secret = strings.Replace(
				sc.Secret, env.SecretGenerationPrefix(), "", 1,
			)
			newSecret, err := crypto.GenerateValue(sc.Secret)
			if err != nil {
				sc.Secret = "ParseError:" + sc.Secret
			} else {
				sc.Secret = newSecret
			}
		}

		p, err := url.JoinPath(env.EndpointUrlForSafe(), "/sentinel/v1/secrets")
		if err != nil {
			printEndpointError(cid, err)
			return
		}

		tlsConfig := tlsconfig.MTLSClientConfig(source, source, authorizer)
		client := &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: tlsConfig,
			},
		}

		sr := newSecretUpsertRequest(sc.WorkloadId, sc.Secret, sc.Namespaces,
			sc.BackingStore, sc.UseKubernetes, sc.Template, sc.Format,
			sc.Encrypt, sc.AppendSecret, sc.NotBefore, sc.Expires)

		md, err := json.Marshal(sr)
		if err != nil {
			printPayloadError(cid, err)
			return
		}

		if sc.DeleteSecret {
			doDelete(cid, client, p, md)
			return
		}

		doPost(cid, client, p, md)
	}
}
