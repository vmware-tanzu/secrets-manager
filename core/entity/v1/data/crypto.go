/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package data

import (
	"strings"

	"github.com/vmware-tanzu/secrets-manager/core/constants/symbol"
)

type RootKeyCollection struct {
	PrivateKey string
	PublicKey  string
	AesSeed    string
}

// Combine takes the private key, a public key, and an AES seed,
// and combines them into a single string, separating each with a newline.
//
// Returns:
//   - A single string containing the private key, public key, and AES seed,
//     each separated by a newline.
func (rkt *RootKeyCollection) Combine() string {
	if rkt.PrivateKey == "" || rkt.PublicKey == "" || rkt.AesSeed == "" {
		return ""
	}

	return rkt.PrivateKey +
		symbol.RootKeySeparator + rkt.PublicKey +
		symbol.RootKeySeparator + rkt.AesSeed
}

// Empty checks if the RootKeyCollection contains any key data.
// It returns true if the PrivateKey, PublicKey, and AesSeed fields are all
// empty.
func (rkt *RootKeyCollection) Empty() bool {
	return rkt.PrivateKey == "" && rkt.PublicKey == "" && rkt.AesSeed == ""
}

// UpdateFromSerialized updates the RootKeyCollection from a serialized string.
//
// The serialized string is expected to be a concatenation of the private key,
// public key, and AES seed separated by the RootKeySeparator symbol. If the
// serialized string is empty, the PrivateKey, PublicKey, and AesSeed fields
// will be set to empty strings. If the serialized string is improperly
// formatted, the function will not update the fields.
//
// Parameters:
//   - serialized: A string containing the serialized key data.
func (rkt *RootKeyCollection) UpdateFromSerialized(serialized string) {
	serialized = strings.TrimSpace(serialized)

	if serialized == "" {
		rkt.PrivateKey = ""
		rkt.PublicKey = ""
		rkt.AesSeed = ""
	}

	parts := strings.Split(serialized, symbol.RootKeySeparator)

	if len(parts) < 3 {
		return
	}

	rkt.PrivateKey = parts[0]
	rkt.PublicKey = parts[1]
	rkt.AesSeed = parts[2]
}
