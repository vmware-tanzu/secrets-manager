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
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
	"github.com/spiffe/go-spiffe/v2/workloadapi"

	"github.com/vmware-tanzu/secrets-manager/core/constants/key"
	u "github.com/vmware-tanzu/secrets-manager/core/constants/url"
	"github.com/vmware-tanzu/secrets-manager/core/entity/v1/data"
	entity "github.com/vmware-tanzu/secrets-manager/core/entity/v1/reqres/safe"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
	"github.com/vmware-tanzu/secrets-manager/core/validation"
)

// CheckInitialization verifies if VSecM Sentinel has executed its init commands
// stanza successfully. This function utilizes a SPIFFE-based mTLS
// authentication mechanism to securely connect to a specified API endpoint.
//
// Parameters:
//   - ctx context.Context: The context carrying the correlation ID used for
//     logging and tracing the operation across different system components.
//     The correlation ID is extracted from the context for error logging
//     purposes.
//   - source *workloadapi.X509Source: A pointer to an X509Source, which
//     provides the credentials necessary for mTLS configuration. The source
//     must not be nil, as it is essential for establishing the TLS connection.
//
// Returns:
//   - bool: Returns true if VSecM Sentinel is initialized; false otherwise .
//   - error: Returns an error if the workload source is nil, URL joining fails,
//     the API call fails, the response body cannot be read, or the JSON
//     response cannot be unmarshalled. The error will provide a detailed
//     message about the nature of the failure.
func CheckInitialization(
	ctx context.Context, source *workloadapi.X509Source,
) (bool, error) {
	cid := ctx.Value(key.CorrelationId).(*string)

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

	checkUrl := u.SentinelKeystone

	p, err := url.JoinPath(env.EndpointUrlForSafe(), checkUrl)
	if err != nil {
		return false, errors.Join(
			err,
			fmt.Errorf(
				"check: I am having problem"+
					" generating VSecM Safe secrets api endpoint URL: %s\n",
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
		return false, errors.Join(
			err,
			fmt.Errorf(
				"check: Problem connecting"+
					" to VSecM Safe API endpoint URL: %s\n",
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
		return false, errors.Join(
			err,
			errors.New("check: Unable to read the response body from VSecM Safe"),
		)
	}

	log.TraceLn(cid, "json result: ' ", string(res), " ' status: ",
		r.Status, "code", r.StatusCode)

	var result entity.KeystoneStatusResponse

	if err := json.Unmarshal(res, &result); err != nil {
		log.ErrorLn(cid, "error unmarshalling JSON: %v", err.Error())
		return false, err
	}

	return result.Status == data.Ready, nil
}
