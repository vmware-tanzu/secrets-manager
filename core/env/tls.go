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

// TlsPort returns the secure port for VSecM Safe to listen on.
// It checks the VSECM_SAFE_TLS_PORT environment variable. If the variable
// is not set, it defaults to ":8443".
func TlsPort() string {
	p := env.Value(env.VSecMSafeTlsPort)
	if p == "" {
		p = string(env.VSecMSafeTlsPortDefault)
	}
	return p
}
