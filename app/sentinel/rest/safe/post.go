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
	"github.com/vmware-tanzu/secrets-manager/core/validation"
	"io"
	"log"
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
)

func createAuthorizer() tlsconfig.Authorizer {
	return tlsconfig.AdaptMatcher(func(id spiffeid.ID) error {
		if validation.IsSafe(id.String()) {
			return nil
		}

		return errors.New("Post: I don’t know you, and it’s crazy: '" + id.String() + "'")
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

func newSecretUpsertRequest(workloadId, secret string, namespaces []string,
	backingStore string, template string, format string,
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
		WorkloadIds:  workloadId,
		BackingStore: bs,
		Namespaces:   namespaces,
		Template:     template,
		Format:       f,
		Encrypt:      encrypt,
		AppendValue:  appendSecret,
		Value:        secret,
		NotBefore:    notBefore,
		Expires:      expires,
	}
}

func respond(cid *string, r *http.Response) (string, error) {
	if r == nil {
		return "", fmt.Errorf("post: Response is null")
	}

	defer func(b io.ReadCloser) {
		if b == nil {
			return
		}
		err := b.Close()
		if err != nil {
			log.Println(cid, "Post: Problem closing request body : %v", err.Error())
		}
	}(r.Body)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(cid, "Post: Unable to read the response body from VSecM Safe : %v", err.Error())
		return "", fmt.Errorf("post: Unable to read the response body from VSecM Safe : %v", err.Error())
	}

	return string(body), nil
}

func printEndpointError(cid *string, err error) error {
	log.Println(cid, "Post: I am having problem generating VSecM Safe secrets api endpoint URL : %v", err.Error())
	return fmt.Errorf("post: I am having problem generating VSecM Safe secrets api endpoint URL: %v", err.Error())
}

func printPayloadError(cid *string, err error) error {
	log.Println(cid, "Post: I am having problem generating the payload : %v", err.Error())
	return fmt.Errorf("post: I am having problem generating the payload: %v", err.Error())
}

func doDelete(cid *string, client *http.Client, r *http.Request, p string, md []byte) (string, error) {
	req, err := http.NewRequest(http.MethodDelete, p, bytes.NewBuffer(md))
	selectedHeaders := []string{"Authorization", "Content-Type", "ClientId", "ClientSecret", "UserName"}
	for _, headerName := range selectedHeaders {
		if value := r.Header.Get(headerName); value != "" {
			req.Header.Set(headerName, value)
		}
	}
	if err != nil {
		log.Println(cid, "Post:Delete: Problem connecting to VSecM Safe API endpoint URL : %v", err.Error())
		return "", fmt.Errorf("post:Delete: Problem connecting to VSecM Safe API endpoint URL : %v", err.Error())
	}
	req.Header.Set("Content-Type", "application/json")
	response, err := client.Do(req)
	if err != nil {
		log.Println(cid, "Post:Delete: Problem connecting to VSecM Safe API endpoint URL : %v", err.Error())
		return "", fmt.Errorf("post:Delete: Problem connecting to VSecM Safe API endpoint URL : %v", err.Error())
	}
	return respond(cid, response)
}

func doPost(cid *string, client *http.Client, r *http.Request, p string, md []byte) (string, error) {
	req, err := http.NewRequest(http.MethodPost, p, bytes.NewBuffer(md))
	selectedHeaders := []string{"Authorization", "Content-Type", "ClientId", "ClientSecret", "UserName"}
	for _, headerName := range selectedHeaders {
		if value := r.Header.Get(headerName); value != "" {
			req.Header.Set(headerName, value)
		}
	}
	response, err := client.Do(req)
	if err != nil {
		log.Println(cid, "Post: Problem connecting to VSecM Safe API endpoint URL : %v", err.Error())
		return "", fmt.Errorf("post:doPost: Problem connecting to VSecM Safe API endpoint URL : %v", err.Error())
	}
	return respond(cid, response)
}

func Post(parentContext context.Context, r *http.Request, sc entity.SentinelCommand) (string, error) {
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
			log.Println(cid, "Post: I cannot execute command because I cannot talk to SPIRE.")
			return "", fmt.Errorf("post: I cannot execute command because I cannot talk to SPIRE")
		}

		log.Println(cid, "Post: Operation was cancelled due to an unknown reason.")
	case source := <-sourceChan:
		defer func() {
			if source == nil {
				return
			}
			err := source.Close()
			if err != nil {
				log.Println(cid, "Post: Problem closing the workload source.")
			}
		}()

		proceed := <-proceedChan

		if !proceed {
			return "", printPayloadError(cid, errors.New("post: Could not proceed"))
		}

		authorizer := createAuthorizer()

		if sc.InputKeys != "" {
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

			parts := strings.Split(sc.InputKeys, "\n")
			if len(parts) != 3 {
				return "", printPayloadError(cid, errors.New("post: Bad data! Very bad data"))
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
			newSecret, err := crypto.GenerateValue(sc.Secret)
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
			sc.BackingStore, sc.Template, sc.Format,
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
