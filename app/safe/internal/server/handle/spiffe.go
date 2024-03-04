/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package handle

import (
	"net/http"

	"github.com/pkg/errors"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/go-spiffe/v2/svid/x509svid"
)

func spiffeIdFromRequest(r *http.Request) (*spiffeid.ID, error) {
	tlsConnectionState := r.TLS
	if len(tlsConnectionState.PeerCertificates) == 0 {
		return nil, errors.New("no peer certs")
	}

	id, err := x509svid.IDFromCert(tlsConnectionState.PeerCertificates[0])
	if err != nil {
		return nil, errors.Wrap(err, "problem extracting svid")
	}

	return &id, nil
}
