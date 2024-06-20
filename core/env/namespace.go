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

// NamespaceForVSecMSystem returns the namespace for the VSecM apps.
// The namespace is determined by the environment variable
// "VSECM_NAMESPACE_SYSTEM". If the variable is not set or is empty,
// it defaults to "vsecm-system".
//
// Returns:
//
//	string: The namespace to be used for the VSecM system.
func NamespaceForVSecMSystem() string {
	u := env.Value(env.VSecMNamespaceSystem)
	if u == "" {
		u = string(env.VSecMSystem)
	}
	return u
}
