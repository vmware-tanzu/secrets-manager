/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package internal

import (
	"fmt"

	"github.com/vmware-tanzu/secrets-manager/core/crypto"
)

// PrintDecryptedKeys retrieves and prints the decrypted keys along with their
// metadata.
//
// The `secrets` function should return a structure with an `Algorithm` field
// and a `Secrets` field.
// Each element in the `Secrets` slice should have a `Name`, `EncryptedValue`,
// `Created`, and `Updated` field.
// The `crypto.Decrypt` function is used to decrypt the encrypted values.
func PrintDecryptedKeys() {
	ss := secrets()

	algorithm := ss.Algorithm

	const ruler = "---"

	fmt.Println("Algorithm:", algorithm)
	fmt.Println(ruler)
	for _, secret := range ss.Secrets {
		fmt.Println("Name:", secret.Name)

		values := secret.EncryptedValue

		for i, v := range values {
			dv, err := crypto.Decrypt([]byte(v), algorithm)
			if err != nil {
				fmt.Println("Error decrypting value:", err.Error())
				continue
			}
			fmt.Printf("Value[%d]: %s\n", i, dv)
		}

		fmt.Println("Created:", secret.Created.String())
		fmt.Println("Updated:", secret.Updated.String())
		fmt.Println(ruler)
	}
}
