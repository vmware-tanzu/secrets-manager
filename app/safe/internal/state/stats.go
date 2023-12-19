/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware, Inc.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package state

import "sync"

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

var currentState = Status{
	SecretQueueLen: 0,
	SecretQueueCap: 0,
	K8sQueueLen:    0,
	K8sQueueCap:    0,
	NumSecrets:     0,
	lock:           &sync.RWMutex{},
}

// Increment is a method for the Status struct that increments the NumSecrets
// field by 1 if the provided secret name is not found in the in-memory store.
func (s *Status) Increment(name string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	_, ok := secrets.Load(name)
	if !ok {
		s.NumSecrets++
	}
}

// Decrement is a method for the Status struct that decrements the NumSecrets
// field by 1 if the provided secret name is found in the in-memory store.
func (s *Status) Decrement(name string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	_, ok := secrets.Load(name)
	if ok {
		s.NumSecrets--
	}
}

// Stats returns a copy of the currentState Status object, providing a snapshot
// of the current status of the secret manager.
func Stats() Status {
	currentState.lock.RLock()
	defer currentState.lock.RUnlock()

	currentState.K8sQueueCap = cap(k8sSecretQueue)
	currentState.K8sQueueLen = len(k8sSecretQueue)
	currentState.SecretQueueCap = cap(secretQueue)
	currentState.SecretQueueLen = len(secretQueue)

	return currentState
}
