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
	"encoding/hex"

	"github.com/pkg/errors"
)

type Algorithm string

const Age = Algorithm("age")
const Aes = Algorithm("aes")

var reader = rand.Read

// RandomString generates a cryptographically-unique secure random string.
func RandomString(n int) (string, error) {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	bytes := make([]byte, n)

	if _, err := reader(bytes); err != nil {
		return "", err
	}

	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}

	return string(bytes), nil
}

// generateAesSeed generates a random 256-bit AES key, and returns it as a
// hexadecimal encoded string.
//
// Returns:
//  1. A hexadecimal string representation of the generated 256-bit AES key.
//  2. An error value that indicates if any error occurs during key generation
//     or encoding. A non-nil error indicates failure.
//
// Usage:
//
//	hexKey, err := generateAesSeed()
//	if err != nil {
//	    log.Fatal("Failed to generate AES key:", err)
//	}
//
// Example output:
//
//	hexKey: "5baa61e4c9b93f3f0682250b6cf8331b7ee68fd8"
//
// Note:
//
//		The function uses the crypto/rand package for secure random number
//	 generation, suitable for cryptographic purposes.
func generateAesSeed() (string, error) {
	// Generate a 256-bit key
	key := make([]byte, 32)

	_, err := reader(key)
	if err != nil {
		return "", errors.Wrap(err, "generateAesSeed: failed to generate random key")
	}

	return hex.EncodeToString(key), nil
}
