/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware, Inc.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package server

import (
	"github.com/pkg/errors"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/server/handle"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	"github.com/vmware-tanzu/secrets-manager/core/validation"
	"net/http"
)

func Serve(source *workloadapi.X509Source, serverStarted chan<- bool) error {
	if source == nil {
		return errors.New("serve: got nil source while trying to serve")
	}

	handle.InitializeRoutes()

	authorizer := tlsconfig.AdaptMatcher(func(id spiffeid.ID) error {
		if validation.IsWorkload(id.String()) {
			return nil
		}

		return errors.New(
			"TLS Config: I don’t know you, and it’s crazy '" + id.String() + "'",
		)
	})

	tlsConfig := tlsconfig.MTLSServerConfig(source, source, authorizer)
	server := &http.Server{
		Addr:      env.TlsPort(),
		TLSConfig: tlsConfig,
	}

	serverStarted <- true

	if err := server.ListenAndServeTLS("", ""); err != nil {
		return errors.Wrap(err, "serve: failed to listen and serve")
	}

	return nil
}
