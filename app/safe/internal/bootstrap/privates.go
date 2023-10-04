/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware, Inc.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package bootstrap

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/pkg/errors"
)

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

	_, err := rand.Read(key)
	if err != nil {
		return "", errors.Wrap(err, "generateAesSeed: failed to generate random key")
	}

	return hex.EncodeToString(key), nil
}
