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
	"os"

	"github.com/vmware-tanzu/secrets-manager/core/entity/v1/data"
)

// RootKeyStoreTypeForSafe determines the root key store type for
// VSecM Safe.
//
// The function retrieves this configuration from an environment variable.
// If the variable is not set or is explicitly set to "k8s", which is the
// default.
//
// Returns:
//   - The configured root key store as a data.BackingStore.
//   - Panics if the environment variable is set to anything other than "k8s",
//     as it is currently the only supported root key store type.
//
// Usage:
//
//	storeType := config.RootKeyStoreTypeForSafe()
func RootKeyStoreTypeForSafe() data.BackingStore {
	s := env.Value(env.VSecMSafeRootKeyStore)
	if s == "" {
		return data.Kubernetes
	}

	if s != string(data.Kubernetes) {
		panic("Only Kubernetes is supported as a root key store")
	}
	return data.Kubernetes
}

// BackingStoreForSafe determines the backing store type for the VSecM Safe.
// This configuration is retrieved from an environment variable. If the variable
// is not set or explicitly set to "file", and "file" is used as the default and
// only supported backing store.
//
// Returns:
//   - The configured backing store as a data.BackingStore.
//   - Panics if the environment variable is set to anything other than "file",
//     as it is currently the only supported backing store type.
//
// Usage:
//
//	backingStore := config.BackingStoreForSafe()
func BackingStoreForSafe() data.BackingStore {
	s := os.Getenv(string(env.VSecMSafeBackingStore))
	if s == "" {
		return data.File
	}

	if s != string(data.File) {
		panic("Only File is supported as a backing store")
	}

	return data.File
}
