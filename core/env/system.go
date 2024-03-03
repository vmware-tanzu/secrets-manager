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

// SystemNamespace returns namespace from metadata,
// metadata.namespace should be passed as environment variable
// as VSECM_SYSTEM_NAMESPACE to the container.
func SystemNamespace() string {
	return os.Getenv("VSECM_SYSTEM_NAMESPACE")
}
