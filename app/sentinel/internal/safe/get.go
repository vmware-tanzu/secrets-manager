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
	"context"
	"errors"
	"fmt"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	"github.com/vmware-tanzu/secrets-manager/core/validation"
	"io"
	"log"
	"net/http"
	"net/url"
)

func acquireSource(ctx context.Context) (*workloadapi.X509Source, bool) {
	source, err := workloadapi.NewX509Source(
		ctx, workloadapi.WithClientOptions(workloadapi.WithAddr(env.SpiffeSocketUrl())),
	)

	if err != nil {
		fmt.Println("Post: I cannot execute command because I cannot talk to SPIRE.")
		fmt.Println("")
		return nil, false
	}

	svid, err := source.GetX509SVID()
	if err != nil {
		fmt.Println("Post: I am having trouble fetching my identity from SPIRE.")
		fmt.Println("Post: I won’t proceed until you put me in a secured container.")
		fmt.Println("")
		return source, false
	}

	// Make sure that the binary is enclosed in a Pod that we trust.
	if !validation.IsSentinel(svid.ID.String()) {
		fmt.Println("I don’t know you, and it’s crazy: '" + svid.ID.String() + "'")
		fmt.Println("`safe` can only run from within the Sentinel container.")
		fmt.Println("")
		return source, false
	}

	return source, true
}

func Get() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	source, proceed := acquireSource(ctx)
	defer func() {
		if source == nil {
			return
		}
		err := source.Close()
		if err != nil {
			log.Println("Problem closing the workload source.")
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

	p, err := url.JoinPath(env.SafeEndpointUrl(), "/sentinel/v1/secrets")
	if err != nil {
		fmt.Println("I am having problem generating VSecM Safe secrets api endpoint URL.")
		fmt.Println("")
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
		fmt.Println("Get: Problem connecting to VSecM Safe API endpoint URL.", err.Error())
		fmt.Println("")
		return
	}

	defer func(b io.ReadCloser) {
		if b == nil {
			return
		}
		err := b.Close()
		if err != nil {
			log.Println("Get: Problem closing request body.")
		}
	}(r.Body)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Get: Unable to read the response body from VSecM Safe.")
		fmt.Println("")
		return
	}

	fmt.Println("")
	fmt.Println(string(body))
	fmt.Println("")
}
