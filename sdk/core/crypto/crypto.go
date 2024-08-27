/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package crypto

import (
	"fmt"

	"github.com/vmware-tanzu/secrets-manager/sdk/lib/crypto"
)

// Id generates a cryptographically-unique secure random string.
func Id() string {
	id, err := crypto.RandomString(8)
	if err != nil {
		id = fmt.Sprintf("CRYPTO-ERR: %s", err.Error())
	}
	return id
}
