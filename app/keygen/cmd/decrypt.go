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
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/vmware-tanzu/secrets-manager/core/crypto"
	entity "github.com/vmware-tanzu/secrets-manager/core/entity/reqres/safe/v1"
	"github.com/vmware-tanzu/secrets-manager/core/env"
)

func rootKeyTriplet(content string) (string, string, string) {
	if content == "" {
		return "", "", ""
	}

	parts := strings.Split(content, "\n")

	if len(parts) != 3 {
		return "", "", ""
	}

	return parts[0], parts[1], parts[2]
}

func keys() (string, string, string) {
	p := env.RootKeyPathForKeyGen()

	content, err := os.ReadFile(p)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	trimmed := strings.TrimSpace(string(content))

	return rootKeyTriplet(trimmed)
}

func decrypt(value []byte, algorithm crypto.Algorithm) (string, error) {
	privateKey, _, aesKey := keys()

	decodedValue, err := base64.StdEncoding.DecodeString(string(value))
	if err != nil {
		return "", err
	}

	if algorithm == crypto.Age {
		res, err := crypto.DecryptBytesAge(decodedValue, privateKey)

		if err != nil {
			return "", err
		}

		return string(res), nil
	}

	res, err := crypto.DecryptBytesAes(decodedValue, aesKey)

	if err != nil {
		return "", err
	}

	return string(res), nil
}

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
			dv, err := decrypt([]byte(v), algorithm)
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
