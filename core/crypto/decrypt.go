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
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"io"
	"os"
	"path"

	"filippo.io/age"
	"github.com/pkg/errors"

	"github.com/vmware-tanzu/secrets-manager/core/env"
)

// DecryptValue takes a base64-encoded and encrypted string value and returns
// the original, decrypted string. If the decryption process encounters any
// error, it will return an empty string and the corresponding error.
func DecryptValue(value string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		return "", err
	}

	if env.FipsCompliantModeForSafe() {
		decrypted, err := DecryptBytesAes(decoded)
		if err != nil {
			return "", err
		}
		return string(decrypted), nil
	}

	decrypted, err := DecryptBytesAge(decoded)
	if err != nil {
		return "", err
	}

	return string(decrypted), nil
}

func DecryptBytesAge(data []byte) ([]byte, error) {
	rkt := RootKeyTriplet()
	privateKey := rkt.PrivateKey

	identity, err := age.ParseX25519Identity(privateKey)
	if err != nil {
		return []byte{}, errors.Wrap(
			err, "decryptBytes: failed to parse private key")
	}

	if len(data) == 0 {
		return []byte{}, errors.Wrap(
			err, "decryptBytes: file on disk appears to be empty")
	}

	out := &bytes.Buffer{}
	f := bytes.NewReader(data)

	r, err := age.Decrypt(f, identity)
	if err != nil {
		return []byte{}, errors.Wrap(
			err, "decryptBytes: failed to open encrypted file")
	}

	if _, err := io.Copy(out, r); err != nil {
		return []byte{}, errors.Wrap(
			err, "decryptBytes: failed to read encrypted file")
	}

	return out.Bytes(), nil
}

func DecryptBytesAes(data []byte) ([]byte, error) {
	rkt := RootKeyTriplet()
	aesKey := rkt.AesSeed

	aesKeyDecoded, err := hex.DecodeString(aesKey)
	if err != nil {
		return nil, errors.Wrap(err,
			"encryptToWriter: failed to decode AES key")
	}

	block, err := aes.NewCipher(aesKeyDecoded)
	if err != nil {
		return nil, errors.Wrap(err,
			"decryptBytes: failed to create AES cipher block")
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

func Decrypt(value []byte, algorithm Algorithm) (string, error) {
	decodedValue, err := base64.StdEncoding.DecodeString(string(value))
	if err != nil {
		return "", err
	}

	if algorithm == Age {
		res, err := DecryptBytesAge(decodedValue)

		if err != nil {
			return "", err
		}

		return string(res), nil
	}

	res, err := DecryptBytesAes(decodedValue)

	if err != nil {
		return "", err
	}

	return string(res), nil
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

	return DecryptBytesAge(data)
}
