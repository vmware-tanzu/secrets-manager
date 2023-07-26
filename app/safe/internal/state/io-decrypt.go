/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware, Inc.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package state

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"filippo.io/age"
	"github.com/pkg/errors"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	"io"
	"os"
	"path"
)

func decryptBytes(data []byte) ([]byte, error) {
	privateKey, _, _ := ageKeyTriplet()

	identity, err := age.ParseX25519Identity(privateKey)
	if err != nil {
		return []byte{}, errors.Wrap(err, "decryptBytes: failed to parse private key")
	}

	if len(data) == 0 {
		return []byte{}, errors.Wrap(err, "decryptBytes: file on disk appears to be empty")
	}

	out := &bytes.Buffer{}
	f := bytes.NewReader(data)

	r, err := age.Decrypt(f, identity)
	if err != nil {
		return []byte{}, errors.Wrap(err, "decryptBytes: failed to open encrypted file")
	}

	if _, err := io.Copy(out, r); err != nil {
		return []byte{}, errors.Wrap(err, "decryptBytes: failed to read encrypted file")
	}

	return out.Bytes(), nil
}

func decryptBytesAes(data []byte) ([]byte, error) {
	_, _, aesKey := ageKeyTriplet()
	aesKeyDecoded, err := hex.DecodeString(aesKey)
	if err != nil {
		return nil, errors.Wrap(err, "encryptToWriter: failed to decode AES key")
	}

	block, err := aes.NewCipher(aesKeyDecoded)
	if err != nil {
		return nil, errors.Wrap(err, "decryptBytes: failed to create AES cipher block")
	}

	// The IV is included in the beginning of the ciphertext.
	if len(data) < aes.BlockSize {
		return nil, errors.New("decryptBytes: ciphertext too short")
	}
	iv := data[:aes.BlockSize]
	data = data[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(data, data)

	return data, nil
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
