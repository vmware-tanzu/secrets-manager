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
	"github.com/vmware-tanzu/secrets-manager/core/constants/symbol"
	"strings"
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
	return rkt.PrivateKey +
		symbol.RootKeySeparator + rkt.PublicKey +
		symbol.RootKeySeparator + rkt.AesSeed
}

func (rkt *RootKeyCollection) Empty() bool {
	return rkt.PrivateKey == "" && rkt.PublicKey == "" && rkt.AesSeed == ""
}

func (rkt *RootKeyCollection) UpdateFromSerialized(serialized string) {
	parts := strings.Split(serialized, symbol.RootKeySeparator)

	if len(parts) < 3 {
		return
	}

	rkt.PrivateKey = parts[0]
	rkt.PublicKey = parts[1]
	rkt.AesSeed = parts[2]
}
