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
	"github.com/vmware-tanzu/secrets-manager/core/spiffe"
	"io"
	"net/http"
	"net/url"

	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	log "github.com/vmware-tanzu/secrets-manager/core/log/rpc"
	"github.com/vmware-tanzu/secrets-manager/core/validation"
)

func Get(ctx context.Context, showEncryptedSecrets bool) {
	cid := ctx.Value("correlationId").(*string)

	source, proceed := spiffe.AcquireSourceForSentinel(ctx)
	defer func() {
		if source == nil {
			return
		}
		err := source.Close()
		if err != nil {
			log.ErrorLn(cid, "Get: Problem closing the workload source.")
		}
	}()
	if !proceed {
		return
	}

	authorizer := tlsconfig.AdaptMatcher(func(id spiffeid.ID) error {
		if validation.IsSafe(id.String()) {
			return nil
		}

		return errors.New("I don’t know you, and it’s crazy: '" + id.String() + "'")
	})

	safeUrl := "/sentinel/v1/secrets"
	if showEncryptedSecrets {
		safeUrl = "/sentinel/v1/secrets?reveal=true"
	}

	p, err := url.JoinPath(env.EndpointUrlForSafe(), safeUrl)
	if err != nil {
		log.ErrorLn(
			cid,
			"Get: I am having problem generating VSecM Safe secrets api endpoint URL.",
		)
		return
	}

	tlsConfig := tlsconfig.MTLSClientConfig(source, source, authorizer)
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	r, err := client.Get(p)
	if err != nil {
		log.ErrorLn(cid,
			"Get: Problem connecting to VSecM Safe API endpoint URL.", err.Error(),
		)
		return
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
		log.ErrorLn(cid, "Get: Unable to read the response body from VSecM Safe.")
		return
	}

	fmt.Println("")
	fmt.Println(string(body))
	fmt.Println("")
}
