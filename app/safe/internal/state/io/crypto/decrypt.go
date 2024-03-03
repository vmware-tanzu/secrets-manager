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
	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/state/io/key"
	"os"
	"path"

	"github.com/pkg/errors"

	"github.com/vmware-tanzu/secrets-manager/core/crypto"
	"github.com/vmware-tanzu/secrets-manager/core/env"
)

// DecryptBytes decrypts a byte slice using the Age encryption algorithm. It is a
// wrapper function that retrieves the private key from a root key triplet and t
// hen calls a specific decryption function from the crypto package to perform
// the decryption.
//
// Parameters:
// - data ([]byte): The encrypted data as a byte slice that needs to be decrypted.
//
// Returns:
//   - ([]byte, error): The function returns two values. The first is a byte slice
//     containing the decrypted data. The second is an error object, which will be
//     non-nil if any errors occurred during the decryption process.
func DecryptBytes(data []byte) ([]byte, error) {
	privateKey, _, _ := key.RootKeyTriplet()
	return crypto.DecryptBytesAge(data, privateKey)
}

// DecryptBytesAes decrypts a byte slice using the AES encryption algorithm.
// Similar to DecryptBytes, this function retrieves the AES key from a root key
// triplet and utilizes a specific function from the crypto package to carry out
// the decryption.
//
// Parameters:
//   - data ([]byte): The encrypted data in the form of a byte slice that is to be
//     decrypted.
//
// Returns:
//   - ([]byte, error): This function also returns a byte slice containing the
//     decrypted data and an error object. The error will be non-nil if the
//     decryption process encounters any issues.
func DecryptBytesAes(data []byte) ([]byte, error) {
	_, _, aesKey := key.RootKeyTriplet()
	return crypto.DecryptBytesAes(data, aesKey)
}

// DecryptDataFromDisk takes a key as input and attempts to decrypt the data
// associated with that key from the disk. The key is used to locate the data file,
// which is expected to have a ".age" extension
// and to be stored in a directory specified by the environment's data path for safe storage.
//
// Parameters:
//   - key (string): A string representing the unique identifier for the data to be
//     decrypted. The actual data file is expected to be named using this key with a
//     ".age" extension.
//
// Returns:
//   - ([]byte, error): This function returns two values. The first value is a byte
//     slice containing the decrypted data if the process is successful. The second value
//     is an error object that will be non-nil if any step of the decryption process fails.
//     Possible errors include the absence of the target data file on disk and failures
//     related to reading the file or the decryption process itself.
func DecryptDataFromDisk(key string) ([]byte, error) {
	dataPath := path.Join(env.DataPathForSafe(), key+".age")

	if _, err := os.Stat(dataPath); os.IsNotExist(err) {
		return nil, errors.Wrap(err, "decryptDataFromDisk: No file at: "+dataPath)
	}

	data, err := os.ReadFile(dataPath)
	if err != nil {
		return nil, errors.Wrap(err, "decryptDataFromDisk: Error reading file")
	}

	if env.FipsCompliantModeForSafe() {
		return DecryptBytesAes(data)
	}

	return DecryptBytes(data)
}
