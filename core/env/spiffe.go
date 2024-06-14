/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package env

import (
	"github.com/vmware-tanzu/secrets-manager/core/constants"
)

// SpiffeSocketUrl returns the URL for the SPIFFE endpoint socket used in the
// VMware Secrets Manager system. The URL is obtained from the environment variable
// SPIFFE_ENDPOINT_SOCKET. If the variable is not set, the default URL is used.
func SpiffeSocketUrl() string {
	p := constants.GetEnv(constants.SpiffeEndpointSocket)
	if p == "" {
		p = string(constants.SpiffeEndpointSocketDefault)
	}
	return p
}

func SpiffeTrustDomain() string {
	p := constants.GetEnv(constants.SpiffeTrustDomain)
	if p == "" {
		p = string(constants.SpiffeTrustDomainDefault)
	}
	return p
}
