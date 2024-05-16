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
func SetRootKeyInMemory(k string) {
	// id := Id()

	RootKeyLock.Lock()
	defer RootKeyLock.Unlock()

	// TODO: ensure that root key secret is updated too
	// TODO: root key secret shall have backing store option: file, database, various secret stores, in-memory
	// TODO: Periodically sync root key secret from the backing store into memory

	RootKey = k
}

// RootKeySet returns true if the root key has been set.
func RootKeySet() bool {
	RootKeyLock.RLock()
	defer RootKeyLock.RUnlock()

	return RootKey != ""
}

type RootKeyCollection struct {
	PrivateKey string
	PublicKey  string
	AesSeed    string
}

// RootKeyCollectionFromMemory splits the RootKey into three components, if it is properly
// formatted.
//
// The function returns a triplet of strings representing the parts of the RootKey,
// separated by newlines. If the RootKey is empty or does not contain exactly
// three parts, the function returns three empty strings.
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

//func Keys() (string, string, string) {
//	p := env.RootKeyPathForKeyGen()
//
//	content, err := os.ReadFile(p)
//	if err != nil {
//		log.Fatalf("Error reading file: %v", err)
//	}
//
//	trimmed := strings.TrimSpace(string(content))
//
//	return rootKeyTriplet(trimmed)
//}
