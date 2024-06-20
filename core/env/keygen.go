/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package env

import (
	"github.com/vmware-tanzu/secrets-manager/core/constants/env"
	"github.com/vmware-tanzu/secrets-manager/core/constants/val"
)

// RootKeyPathForKeyGen returns the root key path. Root key is used to decrypt
// VSecM-encrypted secrets.
// It reads the environment variable VSECM_KEYGEN_ROOT_KEY_PATH to determine
// the path.
// If the environment variable is not set, it defaults to "/opt/vsecm/keys.txt".
//
// Returns:
//
//	string: The path to the root key.
func RootKeyPathForKeyGen() string {
	p := env.Value(env.VSecMKeygenRootKeyPath)
	if p == "" {
		return string(env.VSecMKeygenRootKeyPathDefault)
	}
	return p
}

// ExportedSecretPathForKeyGen returns the path where the exported secrets are
// stored. It reads the environment variable VSECM_KEYGEN_EXPORTED_SECRET_PATH
// to determine the path.
//
// If the environment variable is not set, it defaults to
// "/opt/vsecm/secrets.json".
//
// Returns:
//
//	string: The path to the exported secrets.
func ExportedSecretPathForKeyGen() string {
	p := env.Value(env.VSecMKeygenExportedSecretPath)
	if p == "" {
		return string(env.VSecMKeygenExportedSecretPathDefault)
	}
	return p
}

// KeyGenDecrypt determines if VSecM Keygen should decrypt the secrets json
// file instead of generating a new root key (which is its default behavior).
//
// It reads the environment variable VSECM_KEYGEN_DECRYPT and checks if it is
// set to "true".
//
// If this value is `false`, VSecM Keygen will generate a new root key.
//
// If this value is `true`, VSecM Keygen will attempt to decrypt the secrets
// provided to it.
//
// Returns:
//
//	bool: True if decryption should proceed, false otherwise.
func KeyGenDecrypt() bool {
	p := env.Value(env.VSecMKeygenDecrypt)
	return val.True(p)
}
