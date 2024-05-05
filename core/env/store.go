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
	"github.com/vmware-tanzu/secrets-manager/core/entity/v1/data"
	"os"
)

func RootKeyStoreTypeForSafe() data.BackingStore {
	s := os.Getenv("VSECM_SAFE_ROOT_KEY_STORE")
	if s == "" {
		return data.Kubernetes
	}

	// TODO: implement other store options too.
	return data.Kubernetes
}

func BackingStoreForSafe() data.BackingStore {
	s := os.Getenv("VSECM_SAFE_BACKING_STORE")
	if s == "" {
		return data.File
	}

	// TODO: implement other store options too.
	return data.File
}
