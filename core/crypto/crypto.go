/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware, Inc.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package crypto

import (
	"crypto/rand"
	"encoding/hex"

	"filippo.io/age"
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
	// Generate a 256 bit key
	key := make([]byte, 32)

	_, err := reader(key)
	if err != nil {
		return "", errors.Wrap(err, "generateAesSeed: failed to generate random key")
	}

	return hex.EncodeToString(key), nil
}

// GenerateKeys generates a pair of X25519 keys for public key encryption
// using the age library, as well as an AES seed for symmetric encryption.
//
// Returns:
// - publicKey: The X25519 public key as a string.
// - privateKey: The X25519 private key as a string.
// - aesSeed: A generated AES seed for symmetric encryption.
// - error: An error object if any step in the process fails.
func GenerateKeys() (string, string, string, error) {
	identity, err := age.GenerateX25519Identity()

	if err != nil {
		return "", "", "", err
	}

	privateKey := identity.String()
	publicKey := identity.Recipient().String()
	aesSeed, err := generateAesSeed()

	if err != nil {
		return "", "", "", err
	}

	return privateKey, publicKey, aesSeed, err
}

// CombineKeys takes a private key, a public key, and an AES seed,
// and combines them into a single string, separating each with a newline.
//
// Parameters:
// - privateKey: The X25519 private key as a string.
// - publicKey: The X25519 public key as a string.
// - aesSeed: The AES seed for symmetric encryption as a string.
//
// Returns:
//   - A single string containing the private key, public key, and AES seed,
//     each separated by a newline.
func CombineKeys(privateKey, publicKey, aesSeed string) string {
	return privateKey + "\n" + publicKey + "\n" + aesSeed
}
