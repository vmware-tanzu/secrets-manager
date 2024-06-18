/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package engine

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
	"github.com/spiffe/go-spiffe/v2/workloadapi"

	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/server/handle"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	"github.com/vmware-tanzu/secrets-manager/core/validation"
)

// Serve initializes and starts an mTLS-secured HTTP server using the given
// X509Source and TLS configuration. It also signals that the server has started
// by sending a message on the provided channel.
//
// Parameters:
//   - source: A pointer to workloadapi.X509Source, which provides X.509 SVIDs
//     for mTLS.
//   - serverStarted: A channel that will receive a boolean value to signal
//     server startup.
//
// Returns:
//   - error: An error object if the server fails to start or run; otherwise,
//     returns nil.
//
// The function performs the following operations:
//  1. Validates the source and initializes routes.
//  2. Sets up a custom authorizer using the TLS configuration.
//  3. Configures and starts the HTTP server with mTLS enabled.
//  4. Signals server startup by sending a boolean value on the `serverStarted`
//     channel.
//  5. Listens and serves incoming HTTP requests.
//
// The function will return an error if any of the following conditions occur:
//   - The source is nil.
//   - Server fails to listen and serve.
//
// Note: Serve should be called only once during the application lifecycle to
// initialize the HTTP server.
func Serve(source *workloadapi.X509Source, serverStarted chan<- bool) error {
	if source == nil {
		return errors.New("serve: got nil source while trying to serve")
	}

	handle.InitializeRoutes(source)

	authorizer := tlsconfig.AdaptMatcher(func(id spiffeid.ID) error {
		if validation.IsWorkload(id.String()) {
			return nil
		}

		return fmt.Errorf(
			"TLS Config: I don't know you, and it's crazy '%s'", id.String(),
		)
	})

	tlsConfig := tlsconfig.MTLSServerConfig(source, source, authorizer)
	server := &http.Server{
		Addr:      env.TlsPort(),
		TLSConfig: tlsConfig,
	}

	serverStarted <- true

	if err := server.ListenAndServeTLS("", ""); err != nil {
		return errors.Join(
			err,
			errors.New("serve: failed to listen and serve"),
		)
	}

	return nil
}
