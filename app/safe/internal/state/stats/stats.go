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
)

// Status is a struct representing the current state of the secret manager,
// including the lengths and capacities of the secret queues and the total
// number of secrets stored.
type Status struct {
	SecretQueueLen int
	SecretQueueCap int
	K8sQueueLen    int
	K8sQueueCap    int
	NumSecrets     int
	lock           *sync.RWMutex
}

// CurrentState is a global Status object that represents the current state of
// the Secrets Manager.
var CurrentState = Status{
	SecretQueueLen: 0,
	SecretQueueCap: 0,
	K8sQueueLen:    0,
	K8sQueueCap:    0,
	NumSecrets:     0,
	lock:           &sync.RWMutex{},
}

// Increment is a method for the Status struct that increments the NumSecrets
// field by 1 if the provided secret name is not found in the in-memory store.
func (s *Status) Increment(name string, loader func(name any) (any, bool)) {
	s.lock.Lock()
	defer s.lock.Unlock()
	_, ok := loader(name)
	if !ok {
		s.NumSecrets++
	}
}

// Decrement is a method for the Status struct that decrements the NumSecrets
// field by 1 if the provided secret name is found in the in-memory store.
func (s *Status) Decrement(name string, loader func(name any) (any, bool)) {
	s.lock.Lock()
	defer s.lock.Unlock()
	_, ok := loader(name)
	if ok {
		s.NumSecrets--
	}
}

// Stats returns a copy of the CurrentState Status object, providing a snapshot
// of the current status of the secret manager.
func Stats() Status {
	CurrentState.lock.RLock()
	defer CurrentState.lock.RUnlock()

	CurrentState.K8sQueueCap = cap(insertion.K8sSecretUpsertQueue)
	CurrentState.K8sQueueLen = len(insertion.K8sSecretUpsertQueue)
	CurrentState.SecretQueueCap = cap(insertion.K8sSecretUpsertQueue)
	CurrentState.SecretQueueLen = len(insertion.K8sSecretUpsertQueue)

	return CurrentState
}
