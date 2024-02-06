/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package state

import (
	"github.com/pkg/errors"
	"github.com/vmware-tanzu/secrets-manager/core/crypto"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	"os"
	"path"
)

func decryptBytes(data []byte) ([]byte, error) {
	privateKey, _, _ := ageKeyTriplet()
	return crypto.DecryptBytesAge(data, privateKey)
}

func decryptBytesAes(data []byte) ([]byte, error) {
	_, _, aesKey := ageKeyTriplet()
	return crypto.DecryptBytesAes(data, aesKey)
}

func decryptDataFromDisk(key string) ([]byte, error) {
	dataPath := path.Join(env.SafeDataPath(), key+".age")

	if _, err := os.Stat(dataPath); os.IsNotExist(err) {
		return nil, errors.Wrap(err, "decryptDataFromDisk: No file at: "+dataPath)
	}

	data, err := os.ReadFile(dataPath)
	if err != nil {
		return nil, errors.Wrap(err, "decryptDataFromDisk: Error reading file")
	}

	if env.SafeFipsCompliant() {
		return decryptBytesAes(data)
	}

	return decryptBytes(data)
}
