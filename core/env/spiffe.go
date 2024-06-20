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
	"github.com/vmware-tanzu/secrets-manager/core/constants/env"
)

// SpiffeSocketUrl returns the URL for the SPIFFE endpoint socket used in the
// VMware Secrets Manager system. The URL is obtained from the environment variable
// SPIFFE_ENDPOINT_SOCKET. If the variable is not set, the default URL is used.
func SpiffeSocketUrl() string {
	p := env.Value(env.SpiffeEndpointSocket)
	if p == "" {
		p = string(env.SpiffeEndpointSocketDefault)
	}
	return p
}

// SpiffeTrustDomain retrieves the SPIFFE trust domain from environment
// variables.
//
// This function looks for the trust domain using the environment variable
// defined by `constants.SpiffeTrustDomain`. If the environment variable is not
// set or is an empty string, it defaults to the value specified by
// `constants.SpiffeTrustDomainDefault`.
//
// Returns:
//   - A string representing the SPIFFE trust domain.
func SpiffeTrustDomain() string {
	p := env.Value(env.SpiffeTrustDomain)
	if p == "" {
		p = string(env.SpiffeTrustDomainDefault)
	}
	return p
}
