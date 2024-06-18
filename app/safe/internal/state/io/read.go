/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package io

import (
	"encoding/json"
	"errors"

	"github.com/vmware-tanzu/secrets-manager/core/crypto"
	entity "github.com/vmware-tanzu/secrets-manager/core/entity/v1/data"
)

// ReadFromDisk retrieves and decrypts a secret stored on disk, identified by
// the provided key, and deserializes it into a SecretStored entity. This
// is critical for secure retrieval of persisted secrets, ensuring both
// confidentiality and integrity by decrypting and validating the secret's
// structure.
//
// Parameters:
//   - key (string): A unique identifier for the secret. This key is used to
//     locate the encrypted file on the disk which contains the secret's data.
//
// Returns:
//   - '(*entity.SecretStored, error)': This function returns a pointer to a
//     SecretStored entity if the operation is successful. The SecretStored
//     entity represents the decrypted and deserialized secret. If any error
//     occurs during the process, a nil pointer and an error object are
//     returned. The error provides context about the nature of the failure,
//     such as issues with decryption or data deserialization.
func ReadFromDisk(key string) (*entity.SecretStored, error) {
	contents, err := crypto.DecryptDataFromDisk(key)
	if err != nil {
		return nil, errors.Join(
			err,
			errors.New("readFromDisk: error decrypting file"),
		)
	}

	var secret entity.SecretStored
	err = json.Unmarshal(contents, &secret)
	if err != nil {
		return nil, errors.Join(
			err,
			errors.New("readFromDisk: Failed to unmarshal secret"),
		)
	}

	return &secret, nil
}
