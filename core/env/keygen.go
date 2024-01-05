package env

import (
	"os"
)

// KeyGenRootKeyPath returns the root key path. Root key is used to decrypt
// VSecM-encrypted secrets.
// It reads the environment variable VSECM_KEYGEN_ROOT_KEY_PATH to determine
// the path.
// If the environment variable is not set, it defaults to "/opt/vsecm/keys.txt".
//
// Returns:
//
//	string: The path to the root key.
func KeyGenRootKeyPath() string {
	p := os.Getenv("VSECM_KEYGEN_ROOT_KEY_PATH")
	if p == "" {
		return "/opt/vsecm/keys.txt"
	}
	return p
}

// KeyGenExportedSecretPath returns the path where the exported secrets are stored.
// It reads the environment variable VSECM_KEYGEN_EXPORTED_SECRET_PATH to determine
// the path.
// If the environment variable is not set, it defaults to "/opt/vsecm/secrets.json".
//
// Returns:
//
//	string: The path to the exported secrets.
func KeyGenExportedSecretPath() string {
	p := os.Getenv("VSECM_KEYGEN_EXPORTED_SECRET_PATH")
	if p == "" {
		return "/opt/vsecm/secrets.json"
	}
	return p
}

// KeyGenDecrypt determines if the decryption process should be executed.
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
	p := os.Getenv("VSECM_KEYGEN_DECRYPT")
	return p == "true"
}
