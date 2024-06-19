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
	"crypto/rand"
	"fmt"
)

type Algorithm string

const Age = Algorithm("age")
const Aes = Algorithm("aes")

const letters = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var reader = rand.Read

// RandomString generates a cryptographically-unique secure random string.
func RandomString(n int) (string, error) {
	bytes := make([]byte, n)

	if _, err := reader(bytes); err != nil {
		return "", err
	}

	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}

	return string(bytes), nil
}

// Id generates a cryptographically-unique secure random string.
func Id() string {
	id, err := RandomString(8)
	if err != nil {
		id = fmt.Sprintf("CRYPTO-ERR: %s", err.Error())
	}
	return id
}
