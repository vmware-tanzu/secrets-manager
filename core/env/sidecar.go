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

// SidecarSecretsPath returns the path to the secrets file used by the sidecar.
// The path is determined by the VSECM_SIDECAR_SECRETS_PATH environment variable,
// with a default value of "/opt/vsecm/secrets.json" if the variable is not set.
func SidecarSecretsPath() string {
	p := os.Getenv("VSECM_SIDECAR_SECRETS_PATH")
	if p == "" {
		p = "/opt/vsecm/secrets.json"
	}
	return p
}
