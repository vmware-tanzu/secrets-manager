/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware, Inc.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package safe

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
	data "github.com/vmware-tanzu/secrets-manager/core/entity/data/v1"
	reqres "github.com/vmware-tanzu/secrets-manager/core/entity/reqres/safe/v1"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	"github.com/vmware-tanzu/secrets-manager/core/validation"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
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
		return env.SafeBackingStore()
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

func newSecretUpsertRequest(workloadId, secret, namespace, backingStore string,
	useKubernetes bool, template string, format string, encrypt, appendSecret bool,
) reqres.SecretUpsertRequest {
	bs := decideBackingStore(backingStore)
	f := decideSecretFormat(format)

	return reqres.SecretUpsertRequest{
		WorkloadId:    workloadId,
		BackingStore:  bs,
		Namespace:     namespace,
		UseKubernetes: useKubernetes,
		Template:      template,
		Format:        f,
		Encrypt:       encrypt,
		AppendValue:   appendSecret,
		Value:         secret,
	}
}

func respond(r *http.Response) {
	if r == nil {
		return
	}

	defer func(b io.ReadCloser) {
		if b == nil {
			return
		}
		err := b.Close()
		if err != nil {
			log.Println("Post: Problem closing request body.", err.Error())
		}
	}(r.Body)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Post: Unable to read the response body from VSecM Safe.", err.Error())
		fmt.Println("")
		return
	}

	fmt.Println("")
	fmt.Println(string(body))
	fmt.Println("")
}

func printEndpointError(err error) {
	fmt.Println("Post: I am having problem generating VSecM Safe "+
		"secrets api endpoint URL.", err.Error())
	fmt.Println("")
}

func printPayloadError(err error) {
	fmt.Println("Post: I am having problem generating the payload.", err.Error())
	fmt.Println("")
}

func doDelete(client *http.Client, p string, md []byte) {
	req, err := http.NewRequest(http.MethodDelete, p, bytes.NewBuffer(md))
	if err != nil {
		fmt.Println("Post:Delete: Problem connecting to VSecM Safe API endpoint URL.", err.Error())
		fmt.Println("")
		return
	}
	req.Header.Set("Content-Type", "application/json")
	r, err := client.Do(req)
	if err != nil {
		fmt.Println("Post:Delete: Problem connecting to VSecM Safe API endpoint URL.", err.Error())
		fmt.Println("")
		return
	}
	respond(r)
}

func doPost(client *http.Client, p string, md []byte) {
	r, err := client.Post(p, "application/json", bytes.NewBuffer(md))
	if err != nil {
		fmt.Println("Post: Problem connecting to VSecM Safe API endpoint URL.", err.Error())
		fmt.Println("")
		return
	}
	respond(r)
}

func Post(parentContext context.Context, workloadId, secret, namespace, backingStore string,
	useKubernetes bool, template string, format string, encrypt, deleteSecret, appendSecret bool, inputKeys string,
) {
	// TODO: add that env var to the configuration documentation.
	ctxWithTimeout, cancel := context.WithTimeout(
		parentContext,
		env.SafeSourceAcquisitionTimeout(),
	)
	defer cancel()

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
			fmt.Println("Post: I cannot execute command because I cannot talk to SPIRE.")
			fmt.Println("")
			return
		}

		fmt.Println("Post: Operation was cancelled due to an unknown reason.")
	case source := <-sourceChan:
		defer func() {
			if source == nil {
				return
			}
			err := source.Close()
			if err != nil {
				log.Println("Problem closing the workload source.")
			}
		}()

		proceed := <-proceedChan

		if !proceed {
			return
		}

		authorizer := createAuthorizer()

		if inputKeys != "" {
			p, err := url.JoinPath(env.SafeEndpointUrl(), "/sentinel/v1/keys")
			if err != nil {
				printEndpointError(err)
				return
			}

			tlsConfig := tlsconfig.MTLSClientConfig(source, source, authorizer)
			client := &http.Client{
				Transport: &http.Transport{
					TLSClientConfig: tlsConfig,
				},
			}

			parts := strings.Split(inputKeys, "\n")
			if len(parts) != 3 {
				printPayloadError(errors.New("post: Bad data! Very bad data"))
				return
			}

			sr := newInputKeysRequest(parts[0], parts[1], parts[2])
			md, err := json.Marshal(sr)
			if err != nil {
				printPayloadError(err)
				return
			}

			doPost(client, p, md)
			return
		}

		p, err := url.JoinPath(env.SafeEndpointUrl(), "/sentinel/v1/secrets")
		if err != nil {
			printEndpointError(err)
			return
		}

		tlsConfig := tlsconfig.MTLSClientConfig(source, source, authorizer)
		client := &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: tlsConfig,
			},
		}

		sr := newSecretUpsertRequest(workloadId, secret, namespace, backingStore,
			useKubernetes, template, format, encrypt, appendSecret)
		md, err := json.Marshal(sr)
		if err != nil {
			printPayloadError(err)
			return
		}

		if deleteSecret {
			doDelete(client, p, md)
			return
		}

		doPost(client, p, md)
	}
}
