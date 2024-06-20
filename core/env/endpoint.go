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

// EndpointUrlForSafe returns the URL for the VSecM Safe endpoint
// used in the VMware Secrets Manager system.
// The URL is obtained from the environment variable VSECM_SAFE_ENDPOINT_URL.
// If the variable is not set, the default URL is used.
func EndpointUrlForSafe() string {
	u := env.Value(env.VSecMSafeEndpointUrl)
	if u == "" {
		u = string(env.VSecMSafeEndpointUrlDefault)
	}
	return u
}
