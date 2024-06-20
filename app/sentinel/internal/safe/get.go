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
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"github.com/vmware-tanzu/secrets-manager/core/constants/key"
	"github.com/vmware-tanzu/secrets-manager/core/spiffe"

	u "github.com/vmware-tanzu/secrets-manager/core/constants/url"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	log "github.com/vmware-tanzu/secrets-manager/core/log/rpc"
	"github.com/vmware-tanzu/secrets-manager/core/validation"
)

// Check validates the connectivity to VSecM Safe by calling the "list secrets"
// API and expecting a successful response. The successful return (`nil`) from
// this method means that VSecM Safe is up, and VSecM Sentinel is able to
// establish an authorized request and get a meaningful response body.
//
// Parameters:
//   - ctx: Context used for operation cancellation and passing metadata such as
//     "correlationId" for logging purposes.
//   - source: A pointer to a workloadapi.X509Source that provides the necessary
//     credentials for mTLS communication.
//
// Returns:
//   - An error if the validation fails, the workload source is nil, there's an
//     issue with constructing the API endpoint URL, problems occur during the
//     HTTP request to the VSecM Safe API endpoint, or the response body cannot
//     be read. The error includes a descriptive message indicating the nature
//     of the failure.
func Check(ctx context.Context, source *workloadapi.X509Source) error {
	cid := ctx.Value(key.CorrelationId).(*string)

	if source == nil {
		return errors.New("check: workload source is nil")
	}

	authorizer := tlsconfig.AdaptMatcher(func(id spiffeid.ID) error {
		if validation.IsSafe(id.String()) {
			return nil
		}

		return errors.New(
			"I don't know you, and it's crazy: '" + id.String() + "'",
		)
	})

	safeUrl := u.SentinelSecrets

	p, err := url.JoinPath(env.EndpointUrlForSafe(), safeUrl)
	if err != nil {
		return errors.Join(
			err,
			fmt.Errorf(
				"check: I am having problem generating"+
					" VSecM Safe secrets api endpoint URL: %s\n",
				safeUrl,
			),
		)
	}

	tlsConfig := tlsconfig.MTLSClientConfig(source, source, authorizer)
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	r, err := client.Get(p)
	if err != nil {
		return errors.Join(
			err,
			fmt.Errorf(
				"check: Problem connecting to"+
					" VSecM Safe API endpoint URL: %s\n",
				safeUrl,
			),
		)
	}

	defer func(b io.ReadCloser) {
		if b == nil {
			return
		}
		err := b.Close()
		if err != nil {
			log.ErrorLn(cid, "Get: Problem closing request body.")
		}
	}(r.Body)

	_, err = io.ReadAll(r.Body)
	if err != nil {
		return errors.Join(
			err,
			errors.New("check: Unable to read the response body from VSecM Safe"),
		)
	}

	return nil
}

// Get retrieves secrets from a VSecM Safe API endpoint based on the context and
// whether encrypted secrets should be shown.
// The function uses SPIFFE for secure communication, establishing mTLS with
// the server.
//
// Parameters:
//   - ctx: Context used for operation cancellation and passing metadata across
//     API boundaries. It must contain a "correlationId" value.
//   - showEncryptedSecrets: A boolean flag indicating whether to retrieve
//     encrypted secrets. If true, secrets are shown in encrypted form.
func Get(ctx context.Context, showEncryptedSecrets bool) error {
	cid := ctx.Value(key.CorrelationId).(*string)

	log.AuditLn(cid, "Sentinel:Get")

	source, proceed := spiffe.AcquireSourceForSentinel(ctx)
	defer func(s *workloadapi.X509Source) {
		if s == nil {
			return
		}
		err := s.Close()
		if err != nil {
			log.ErrorLn(cid, "Get: Problem closing the workload source.")
		}
	}(source)
	if !proceed {
		return errors.New("get: Problem acquiring source")
	}

	authorizer := tlsconfig.AdaptMatcher(func(id spiffeid.ID) error {
		if validation.IsSafe(id.String()) {
			return nil
		}

		return errors.New("I don't know you, and it's crazy: '" +
			id.String() + "'")
	})

	safeUrl := u.SentinelSecrets
	if showEncryptedSecrets {
		safeUrl = u.SentinelSecretsWithReveal
	}

	p, err := url.JoinPath(env.EndpointUrlForSafe(), safeUrl)
	if err != nil {
		return errors.Join(
			err,
			errors.New("problem generating VSecM Safe secrets api endpoint URL"),
		)
	}

	tlsConfig := tlsconfig.MTLSClientConfig(source, source, authorizer)
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	r, err := client.Get(p)
	if err != nil {
		return errors.Join(
			err,
			errors.New("problem connecting to VSecM Safe API endpoint URL"),
		)
	}

	defer func(b io.ReadCloser) {
		if b == nil {
			return
		}
		err := b.Close()
		if err != nil {
			log.ErrorLn(cid, "Get: Problem closing request body.")
		}
	}(r.Body)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return errors.Join(
			err,
			errors.New("unable to read the response body from VSecM Safe"),
		)
	}

	fmt.Println("")
	fmt.Println(string(body))
	fmt.Println("")

	return nil
}
