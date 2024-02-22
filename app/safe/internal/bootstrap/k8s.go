/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package bootstrap

import (
	"github.com/vmware-tanzu/secrets-manager/core/env"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
)

func ValidateEnvironment() {
	id := "VSECMSAFE"

	// getting metadata.namespace, passed through environment variable VSECM_SYSTEM_NAMESPACE
	if len(env.SystemNamespace()) == 0 {
		log.FatalLn(&id, "Failed to get pod namespace",
			"Pod namespace should be exported into environment as VSECM_SYSTEM_NAMESPACE")
	}
}
