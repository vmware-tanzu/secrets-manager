/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package state

import (
	entity "github.com/vmware-tanzu/secrets-manager/core/entity/data/v1"
	"github.com/vmware-tanzu/secrets-manager/core/env"
)

// The secrets put here are synced with their Kubernetes Secret counterparts.
var k8sSecretDeleteQueue = make(chan entity.SecretStored, env.SafeK8sSecretDeleteBufferSize())

func processK8sSecretDeleteQueue() {
	// id := "AEGIHK8D"

	// No need to implement this; but we’ll keep the placeholder here, in case
	// we find a need for it in the future.
	//
	// @see https://github.com/vmware-tanzu/secrets-manager/issues/268
}
