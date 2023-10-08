/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware, Inc.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package main

import (
	"fmt"
	"github.com/vmware-tanzu/secrets-manager/core/crypto"
)

func main() {
	privateKey, publicKey, aesSeed, err := crypto.GenerateKeys()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println()
	fmt.Println(crypto.CombineKeys(privateKey, publicKey, aesSeed))
	fmt.Println()
}
