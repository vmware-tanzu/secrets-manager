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
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/vmware-tanzu/secrets-manager/core/crypto"
	entity "github.com/vmware-tanzu/secrets-manager/core/entity/v1/reqres/safe"
	"github.com/vmware-tanzu/secrets-manager/core/env"
)

func secrets() entity.SecretStringTimeListResponse {
	p := env.ExportedSecretPathForKeyGen()

	content, err := os.ReadFile(p)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	var secrets entity.SecretStringTimeListResponse

	err = json.Unmarshal(content, &secrets)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	return secrets
}

func printDecryptedKeys() {
	ss := secrets()

	algorithm := ss.Algorithm

	println("Algorithm:", algorithm)
	println("---")
	for _, secret := range ss.Secrets {
		println("Name:", secret.Name)

		values := secret.EncryptedValue

		for i, v := range values {
			dv, err := crypto.Decrypt([]byte(v), algorithm)
			if err != nil {
				println("Error decrypting value:", err.Error())
				continue
			}
			fmt.Printf("Value[%d]: %s\n", i, dv)
		}

		println("Created:", secret.Created)
		println("Updated:", secret.Updated)
		println("---")
	}
}
