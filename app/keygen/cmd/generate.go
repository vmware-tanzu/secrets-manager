/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package main

import (
	"fmt"

	"github.com/vmware-tanzu/secrets-manager/core/crypto"
)

func printGeneratedKeys() {
	rkt, err := crypto.NewRootKeyCollection()

	if err != nil {
		fmt.Println()
		fmt.Println("Failed to generate keys:")
		fmt.Println(err.Error())
		fmt.Println()
		return
	}

	fmt.Println()
	fmt.Println(rkt.Combine())
	fmt.Println()
}
