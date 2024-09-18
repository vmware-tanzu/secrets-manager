/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package safe

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"

	"github.com/vmware-tanzu/secrets-manager/core/entity/v1/data"
	reqres "github.com/vmware-tanzu/secrets-manager/core/entity/v1/reqres/safe"
	"github.com/vmware-tanzu/secrets-manager/core/validation"
)

func createAuthorizer() tlsconfig.Authorizer {
	return tlsconfig.AdaptMatcher(func(id spiffeid.ID) error {
		if validation.IsSafe(id.String()) {
			return nil
		}

		return errors.New("Post: I don't know you, and it's crazy: '" +
			id.String() + "'")
	})
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

func newSecretUpsertRequest(
	workloadIds []string, secret string, namespaces []string,
	template string, format string,
	encrypt, appendSecret bool, notBefore string, expires string,
) reqres.SecretUpsertRequest {
	f := decideSecretFormat(format)

	if notBefore == "" {
		notBefore = "now"
	}

	if expires == "" {
		expires = "never"
	}

	return reqres.SecretUpsertRequest{
		WorkloadIds: workloadIds,
		Namespaces:  namespaces,
		Template:    template,
		Format:      f,
		Encrypt:     encrypt,
		AppendValue: appendSecret,
		Value:       secret,
		NotBefore:   notBefore,
		Expires:     expires,
	}
}

func respond(cid *string, r *http.Response) (string, error) {
	if r == nil {
		return "", errors.New("post: Response is null")
	}

	defer func(b io.ReadCloser) {
		if b == nil {
			return
		}
		err := b.Close()
		if err != nil {
			log.Println(cid, "Post: Problem closing request body : %v",
				err.Error())
		}
	}(r.Body)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(cid,
			"Post: Unable to read the response body from VSecM Safe : %v",
			err.Error())
		return "",
			fmt.Errorf("post: Unable to read the"+
				" response body from VSecM Safe : %v",
				err.Error())
	}

	return string(body), nil
}

func printEndpointError(cid *string, err error) error {
	log.Println(cid,
		"Post: I am having problem generating"+
			" VSecM Safe secrets api endpoint URL : %v", err.Error())
	return fmt.Errorf(
		"post: I am having problem generating"+
			" VSecM Safe secrets api endpoint URL: %v", err.Error())
}

func printPayloadError(cid *string, err error) error {
	log.Println(cid,
		"Post: I am having problem generating the payload : %v", err.Error())
	return fmt.Errorf(
		"post: I am having problem generating the payload: %v", err.Error())
}

func doDelete(
	cid *string, client *http.Client, r *http.Request, p string, md []byte,
) (string, error) {
	req, err := http.NewRequest(http.MethodDelete, p, bytes.NewBuffer(md))
	if err != nil {
		log.Println(cid, "Post:Delete: Problem creating request : %v",
			err.Error())
		return "", fmt.Errorf("post:Delete: Problem creating request : %v",
			err.Error())
	}

	selectedHeaders := []string{
		"Authorization", "Content-Type", "ClientId", "ClientSecret", "UserName",
	}
	for _, headerName := range selectedHeaders {
		if value := r.Header.Get(headerName); value != "" {
			req.Header.Set(headerName, value)
		}
	}

	req.Header.Set("Content-Type", "application/json")
	response, err := client.Do(req)
	if err != nil {
		log.Println(cid,
			"Post:Delete: Problem connecting"+
				" to VSecM Safe API endpoint URL : %v", err.Error())
		return "", fmt.Errorf("post:Delete: Problem connecting"+
			" to VSecM Safe API endpoint URL : %v", err.Error())
	}
	return respond(cid, response)
}

func doPost(
	cid *string, client *http.Client, r *http.Request, p string, md []byte,
) (string, error) {
	req, err := http.NewRequest(http.MethodPost, p, bytes.NewBuffer(md))
	if err != nil {
		log.Println(cid, "Post: Problem creating request : %v", err.Error())
		return "", fmt.Errorf("post: Problem creating request : %v",
			err.Error())
	}

	selectedHeaders := []string{
		"Authorization", "Content-Type", "ClientId", "ClientSecret", "UserName",
	}
	for _, headerName := range selectedHeaders {
		if value := r.Header.Get(headerName); value != "" {
			req.Header.Set(headerName, value)
		}
	}

	response, err := client.Do(req)
	if err != nil {
		log.Println(cid,
			"Post: Problem connecting to"+
				" VSecM Safe API endpoint URL : %v", err.Error())
		return "", fmt.Errorf(
			"post:doPost: Problem connecting to"+
				" VSecM Safe API endpoint URL : %v", err.Error())
	}

	return respond(cid, response)
}
