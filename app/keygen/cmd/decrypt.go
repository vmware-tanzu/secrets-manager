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
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/vmware-tanzu/secrets-manager/core/crypto"
	entity "github.com/vmware-tanzu/secrets-manager/core/entity/reqres/safe/v1"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	"log"
	"os"
	"strings"
)

func ageKeyTriplet(content string) (string, string, string) {
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
	p := env.KeyGenRootKeyPath()

	content, err := os.ReadFile(p)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	trimmed := strings.TrimSpace(string(content))

	return ageKeyTriplet(trimmed)
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
	p := env.KeyGenExportedSecretPath()

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

	fmt.Println("Algorithm:", algorithm)
	fmt.Println("---")
	for _, secret := range ss.Secrets {
		fmt.Println("Name:", secret.Name)

		values := secret.EncryptedValue

		for i, v := range values {
			dv, err := decrypt([]byte(v), algorithm)
			if err != nil {
				fmt.Println("Error decrypting value:", err.Error())
				continue
			}
			fmt.Printf("Value[%d]: %s\n", i, dv)
		}

		fmt.Println("Created:", secret.Created)
		fmt.Println("Updated:", secret.Updated)
		fmt.Println("---")
	}
}
