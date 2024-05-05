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

func NamespaceForVSecMSystem() string {
	u := os.Getenv("VSECM_NAMESPACE_SYSTEM")
	if u == "" {
		u = "vsecm-system"
	}
	return u
}

func NamespaceForSpire() string {
	u := os.Getenv("VSECM_NAMESPACE_SPIRE")
	if u == "" {
		u = "spire-system"
	}
	return u
}
