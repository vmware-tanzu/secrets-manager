/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware, Inc.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package env

import "os"

// SafeEndpointUrl returns the URL for the VSecM Safe endpoint
// used in the VMware Secres Manager system.
// The URL is obtained from the environment variable VSECM_SAFE_ENDPOINT_URL.
// If the variable is not set, the default URL is used.
func SafeEndpointUrl() string {
	u := os.Getenv("VSECM_SAFE_ENDPOINT_URL")
	if u == "" {
		u = "https://vsecm-safe.vsecm-system.svc.cluster.local:8443/"
	}
	return u
}
