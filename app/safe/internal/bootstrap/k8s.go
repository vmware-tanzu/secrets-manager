/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware, Inc.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package bootstrap

import (
	"github.com/vmware-tanzu/secrets-manager/core/env"
	"github.com/vmware-tanzu/secrets-manager/core/log"
)

func ValidateEnvironment() {
	id := "VSECMSAFE"

	// getting metadata.namespace, passed through environment variable VSECM_SYSTEM_NAMESPACE
	if len(env.SystemNamespace()) == 0 {
		log.FatalLn(&id, "Failed to get pod namespace",
			"Pod namespace should be exported into environment as VSECM_SYSTEM_NAMESPACE")
	}
}
