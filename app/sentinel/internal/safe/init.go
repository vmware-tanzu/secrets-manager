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
	"encoding/json"
	"fmt"
	data "github.com/vmware-tanzu/secrets-manager/core/entity/data/v1"
	entity "github.com/vmware-tanzu/secrets-manager/core/entity/reqres/safe/v1"
	"io"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
	"github.com/spiffe/go-spiffe/v2/workloadapi"

	"github.com/vmware-tanzu/secrets-manager/core/env"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
	"github.com/vmware-tanzu/secrets-manager/core/validation"
)

func CheckInitialization(ctx context.Context, source *workloadapi.X509Source) (bool, error) {
	cid := ctx.Value("correlationId").(*string)

	if source == nil {
		return false, errors.New("check: workload source is nil")
	}

	authorizer := tlsconfig.AdaptMatcher(func(id spiffeid.ID) error {
		if validation.IsSafe(id.String()) {
			return nil
		}

		return errors.New(
			"I don't know you, and it's crazy: '" + id.String() + "'",
		)
	})

	checkUrl := "/sentinel/entity/keystone"

	p, err := url.JoinPath(env.EndpointUrlForSafe(), checkUrl)
	if err != nil {
		return false, errors.Wrap(
			err,
			fmt.Sprintf(
				"check: I am having problem generating VSecM Safe secrets api endpoint URL: %s\n",
				checkUrl,
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
		return false, errors.Wrap(
			err,
			fmt.Sprintf(
				"check: Problem connecting to VSecM Safe API endpoint URL: %s\n",
				checkUrl,
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

	res, err := io.ReadAll(r.Body)
	if err != nil {
		return false, errors.Wrap(
			err, "check: Unable to read the response body from VSecM Safe",
		)
	}

	var result entity.KeystoneStatusResponse

	if err := json.Unmarshal(res, &result); err != nil {
		log.ErrorLn(cid, "error unmarshaling JSON: %v", err.Error())
		return false, err
	}

	return result.Status == data.Ready, nil
}
