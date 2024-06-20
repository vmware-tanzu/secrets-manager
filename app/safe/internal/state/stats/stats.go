/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package stats

import (
	"sync"

	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/state/secret/queue/insertion"
	"github.com/vmware-tanzu/secrets-manager/core/entity/v1/data"
)

// CurrentState is a global Status object that represents the current state of
// the Secrets Manager.
var CurrentState = data.Status{
	SecretQueueLen: 0,
	SecretQueueCap: 0,
	K8sQueueLen:    0,
	K8sQueueCap:    0,
	NumSecrets:     0,
	Lock:           sync.RWMutex{},
}

// Stats returns a copy of the CurrentState Status object, providing a snapshot
// of the current status of VSecM.
func Stats() data.Status {
	CurrentState.Lock.RLock()
	defer CurrentState.Lock.RUnlock()

	// Note that this is a side effect.
	// Another alternative would be to update these values in a
	// background process, but that will be an overkill, and it will
	// consume more resources.
	//
	// This approach is a more pragmatic alternative.
	CurrentState.K8sQueueCap = cap(insertion.K8sSecretUpsertQueue)
	CurrentState.K8sQueueLen = len(insertion.K8sSecretUpsertQueue)
	CurrentState.SecretQueueCap = cap(insertion.K8sSecretUpsertQueue)
	CurrentState.SecretQueueLen = len(insertion.K8sSecretUpsertQueue)

	// Note that this is a copy, and it is intended to be read, not
	// modified. Also note that, since it is a copy, it maintains a
	// separate .Lock instance.
	return CurrentState
}
