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
	"strings"
	"sync"

	"filippo.io/age"
)

const BlankRootKeyValue = "{}"

// RootKey is the key used for encryption, decryption, backup, and restore.
var RootKey = ""
var RootKeyLock sync.RWMutex

// SetRootKeyInMemory sets the age key to be used for encryption and decryption.
// This function is called in several instances:
//
// 1. During bootstrapping of VSecM Safe to set the initial root key from
// a mounted backing store.
// 2. When an operator sets a new root key through VSecM Sentinel or other
// similar means.
func SetRootKeyInMemory(k string) {
	RootKeyLock.Lock()
	defer RootKeyLock.Unlock()

	RootKey = k
}

// RootKeySetInMemory returns true if the root key has been set.
func RootKeySetInMemory() bool {
	RootKeyLock.RLock()
	defer RootKeyLock.RUnlock()

	return RootKey != ""
}

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
func (rkt RootKeyCollection) Combine() string {
	return rkt.PrivateKey + "\n" + rkt.PublicKey + "\n" + rkt.AesSeed
}

// RktFromMemory creates a new Rkt struct from the RootKey stored in memory.
func RootKeyCollectionFromMemory() RootKeyCollection {
	RootKeyLock.RLock()
	defer RootKeyLock.RUnlock()

	if RootKey == "" {
		return RootKeyCollection{}
	}

	parts := strings.Split(RootKey, "\n")

	if len(parts) != 3 {
		return RootKeyCollection{}
	}

	return RootKeyCollection{
		PrivateKey: parts[0],
		PublicKey:  parts[1],
		AesSeed:    parts[2],
	}
}

// NewRootKeyCollection creates a new cryptographic key pair and an AES seed.
// It utilizes the X25519 algorithm for key generation and includes both the
// private and public keys in the returned Rkt structure. The function also
// generates an AES seed that can be used for symmetric encryption.
//
// Returns:
//   - Rkt: A struct containing the private key, public key, and AES seed. The
//     Rkt struct should be defined elsewhere in your codebase with the
//     respective fields: PrivateKey, PublicKey, and AesSeed.
//   - error: An error object that reports issues in the key generation process,
//     such as failures in generating the X25519 identity or the AES seed. If
//     the function executes without encountering any issues, the error will be
//     nil.
//
// Example usage:
//
//	keys, err := NewRootKeyCollection()
//	if err != nil {
//	    log.Fatalf("Key generation failed: %v", err)
//	}
//	fmt.Printf("Private Key: %s\n", keys.PrivateKey)
//	fmt.Printf("Public Key: %s\n", keys.PublicKey)
//	fmt.Printf("AES Seed: %s\n", keys.AesSeed)
//
// Note:
//
//	The NewRkt function depends on the 'age' package for generating the
//	X25519 identity and an implementation of generateAesSeed, which must be
//	provided in your codebase or through an external library.
func NewRootKeyCollection() (RootKeyCollection, error) {
	identity, err := age.GenerateX25519Identity()

	if err != nil {
		return RootKeyCollection{}, err
	}

	privateKey := identity.String()
	publicKey := identity.Recipient().String()
	aesSeed, err := generateAesSeed()

	if err != nil {
		return RootKeyCollection{}, err
	}

	return RootKeyCollection{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
		AesSeed:    aesSeed,
	}, nil
}
