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
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
	"github.com/spiffe/go-spiffe/v2/workloadapi"

	"github.com/vmware-tanzu/secrets-manager/core/constants/key"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	"github.com/vmware-tanzu/secrets-manager/core/validation"
)

// Get fetches secrets from the VSecM Safe API, optionally including encrypted
// secrets in the response. This function constructs a secure client and makes
// a GET request to the API based on the input parameters. The response is then
// returned as a string, or an error is generated if the process fails at any
// step.
//
// Parameters:
//   - ctx: a context.Context that must contain a 'correlationId' used for
//     logging.
//   - r: the *http.Request containing the original HTTP request details.
//     Headers from this request may be propagated to the API request.
//   - showEncryptedSecrets: a boolean indicating whether to retrieve encrypted
//     secrets.
//
// Returns:
//   - A string containing the API response if the request is successful.
//   - An error detailing what went wrong during the operation if unsuccessful.
//
// Usage:
//
//	response, err := secrets.Get(ctx, req, true)
//	if err != nil {
//	    log.Println("Error fetching secrets:", err)
//	} else {
//	    log.Println("Fetched secrets:", response)
//	}
func Get(
	ctx context.Context, r *http.Request, showEncryptedSecrets bool,
) (string, error) {
	cid := ctx.Value(key.CorrelationId).(*string)
	log.Println(cid, "Get: start")

	source, proceed := acquireSource(ctx)
	defer func(s *workloadapi.X509Source) {
		if s == nil {
			return
		}
		err := s.Close()
		if err != nil {
			log.Println(cid,
				"Get: Problem closing the workload source.", err.Error())
		}
	}(source)
	if !proceed {
		return "", errors.New("could not proceed")
	}

	authorizer := tlsconfig.AdaptMatcher(func(id spiffeid.ID) error {
		if validation.IsSafe(id.String()) {
			return nil
		}

		return errors.New("I don't know you, and it's crazy: '" +
			id.String() + "'")
	})

	safeUrl := "/sentinel/v1/secrets"
	if showEncryptedSecrets {
		safeUrl = "/sentinel/v1/secrets?reveal=true"
	}

	p, err := url.JoinPath(env.EndpointUrlForSafe(), safeUrl)
	if err != nil {
		log.Println(cid, "Get: I am having problem generating "+
			"VSecM Safe secrets api endpoint URL.", err.Error())
		return "", fmt.Errorf("get: I am having problem "+
			"generating VSecM Safe secrets api endpoint URL: %v", err.Error())
	}

	tlsConfig := tlsconfig.MTLSClientConfig(source, source, authorizer)
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	req, err := http.NewRequest(http.MethodGet, p, nil)
	if err != nil {
		log.Println(cid, "Get: Problem creating request.", err.Error())
		return "", fmt.Errorf("get: Problem creating request: %v", err.Error())
	}

	selectedHeaders := []string{"Authorization",
		"Content-Type", "ClientId", "ClientSecret", "UserName"}
	for _, headerName := range selectedHeaders {
		if value := r.Header.Get(headerName); value != "" {
			req.Header.Set(headerName, value)
		}
	}

	response, err := client.Do(req)
	if err != nil {
		log.Println(cid,
			"Get: Problem connecting to VSecM Safe API endpoint URL.",
			err.Error())
		return "",
			fmt.Errorf("get: Problem connecting to VSecM"+
				" Safe API endpoint URL: %v",
				err.Error())
	}

	defer func(b io.ReadCloser) {
		if b == nil {
			return
		}
		err := b.Close()
		if err != nil {
			log.Println(cid, "Get: Problem closing request body.", err.Error())
		}
	}(response.Body)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println(cid,
			"Get: Unable to read the response body from VSecM Safe.",
			err.Error())
		return "",
			fmt.Errorf("get: Unable to read"+
				" the response body from VSecM Safe: %v",
				err.Error())
	}

	return string(body), nil
}
