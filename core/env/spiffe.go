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

import "os"

// SpiffeSocketUrl returns the URL for the SPIFFE endpoint socket used in the
// VMware Secrets Manager system. The URL is obtained from the environment variable
// SPIFFE_ENDPOINT_SOCKET. If the variable is not set, the default URL is used.
func SpiffeSocketUrl() string {
	p := os.Getenv("SPIFFE_ENDPOINT_SOCKET")
	if p == "" {
		p = "unix:///spire-agent-socket/agent.sock"
	}
	return p
}
