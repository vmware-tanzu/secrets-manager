/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package data

import "sync"

// InitStatus is the initialization status of VSecM Sentinel
// and other VSecM components.
type InitStatus string

var (
	Pending InitStatus = "pending"
	Ready   InitStatus = "ready"
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
}

var statusLock sync.RWMutex // Protects access to the Status struct.

// Increment is a method for the Status struct that increments the NumSecrets
// field by 1 if the provided secret name is not found in the in-memory store.
func (s *Status) Increment(name string, loader func(name any) (any, bool)) {
	statusLock.Lock()
	defer statusLock.Unlock()
	_, ok := loader(name)
	if !ok {
		s.NumSecrets++
	}
}

// Decrement is a method for the Status struct that decrements the NumSecrets
// field by 1 if the provided secret name is found in the in-memory store.
func (s *Status) Decrement(name string, loader func(name any) (any, bool)) {
	statusLock.Lock()
	defer statusLock.Unlock()
	_, ok := loader(name)
	if ok {
		s.NumSecrets--
	}
}
