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
	"errors"
)

var reader = rand.Read

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
		return "", errors.Join(
			err,
			errors.New("generateAesSeed: failed to generate random key"),
		)
	}

	return hex.EncodeToString(key), nil
}
