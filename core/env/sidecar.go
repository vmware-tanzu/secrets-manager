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

// SecretsPathForSidecar returns the path to the secrets file used by the sidecar.
// The path is determined by the VSECM_SIDECAR_SECRETS_PATH environment variable,
// with a default value of "/opt/vsecm/secrets.json" if the variable is not set.
func SecretsPathForSidecar() string {
	p := env.Value(env.VSecMSidecarSecretsPath)
	if p == "" {
		p = string(env.VSecMSidecarSecretsPathDefault)
	}
	return p
}
