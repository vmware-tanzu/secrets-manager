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

import "filippo.io/age"

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
