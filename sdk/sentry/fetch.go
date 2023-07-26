/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware, Inc.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package sentry

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"github.com/vmware-tanzu/secrets-manager/core/crypto"
	reqres "github.com/vmware-tanzu/secrets-manager/core/entity/reqres/safe/v1"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	"github.com/vmware-tanzu/secrets-manager/core/log"
	"github.com/vmware-tanzu/secrets-manager/core/validation"
	"io"
	"net/http"
	"net/url"
)

var ErrSecretNotFound = errors.New("Secret does not exist")

// Fetch fetches the up-to-date secret that has been registered to the workload.
//
//	secret, err := sentry.Fetch()
//
// In case of a problem, Fetch will return an empty string and an error explaining
// what went wrong.
//
// Fetch can ONLY be called from a registered workload; and it ONLY delivers
// the secret that the workload is associated with.
func Fetch() (reqres.SecretFetchResponse, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cid, _ := crypto.RandomString(8)
	if cid == "" {
		cid = "VSECMSDK"
	}

	var source *workloadapi.X509Source

	source, err := workloadapi.NewX509Source(
		ctx, workloadapi.WithClientOptions(
			workloadapi.WithAddr(env.SpiffeSocketUrl()),
		),
	)
	if err != nil {
		return reqres.SecretFetchResponse{}, errors.Wrap(
			err, "Fetch: failed getting SVID Bundle from the SPIRE Workload API",
		)
	}

	defer func() {
		err := source.Close()
		if err != nil {
			log.InfoLn(&cid, "Fetch: problem closing source: ", err.Error())
		}
	}()

	svid, err := source.GetX509SVID()
	if err != nil {
		return reqres.SecretFetchResponse{}, errors.Wrap(err, "Fetch: error getting SVID from source")
	}

	// Make sure that we are calling Safe from a workload that VSecM knows about.
	if !validation.IsWorkload(svid.ID.String()) {
		return reqres.SecretFetchResponse{}, errors.New("Fetch: untrusted workload")
	}

	authorizer := tlsconfig.AdaptMatcher(func(id spiffeid.ID) error {
		if validation.IsSafe(id.String()) {
			return nil
		}

		return errors.New("Fetch: I don’t know you, and it’s crazy: '" + id.String() + "'")
	})

	p, err := url.JoinPath(env.SafeEndpointUrl(), "/workload/v1/secrets")
	if err != nil {
		return reqres.SecretFetchResponse{},
			errors.New("Fetch: problem generating server url")
	}

	client := &http.Client{
		Transport: &http.Transport{
			// Use the connection to serve a single http request only.
			// This is not a web server; there is no need to keep the
			// connection open for multiple requests. This will also
			// save a good chunk of memory, especially when polling
			// interval is shorter. [1]
			DisableKeepAlives: true,
			TLSClientConfig:   tlsconfig.MTLSClientConfig(source, source, authorizer),
		},
	}

	log.TraceLn(&cid, "Sentry:Fetch", p)

	r, err := client.Get(p)
	if err != nil {
		return reqres.SecretFetchResponse{}, errors.Wrap(
			err, "Fetch: problem connecting to VMware Secrets Manager Safe API endpoint",
		)
	}

	defer func() {
		err := r.Body.Close()
		if err != nil {
			if err != nil {
				log.InfoLn(&cid, "Fetch: problem closing response body: ", err.Error())
			}
		}
	}()

	// Related to [1]. Hint the server that we wish to close the connection
	// as soon as we are done with it.
	r.Close = true

	if r.StatusCode == http.StatusNotFound {
		return reqres.SecretFetchResponse{}, ErrSecretNotFound
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return reqres.SecretFetchResponse{}, errors.Wrap(
			err, "unable to read the response body from VMware Secrets Manager Safe API endpoint",
		)
	}

	var sfr reqres.SecretFetchResponse
	err = json.Unmarshal(body, &sfr)
	if err != nil {
		return reqres.SecretFetchResponse{}, errors.Wrap(
			err, "unable to deserialize response",
		)
	}

	return sfr, nil
}
