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
	"errors"
	"io"
	"os"
	"path"

	"filippo.io/age"

	c "github.com/vmware-tanzu/secrets-manager/core/constants/crypto"
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

// DecryptBytesAge decrypts data using an X25519 private key extracted
// from memory. This function is intended to decrypt a slice of bytes that have
// been encrypted using the age encryption tool, specifically formatted for
// X25519 keys.
//
// Parameters:
//   - data: the byte slice containing the encrypted data.
//
// Returns:
//   - A byte slice containing the decrypted data if successful.
//   - An error if decryption fails at any stage, including issues with key
//     parsing, opening the encrypted content, or reading the decrypted data.
//
// Usage:
//
//	decryptedData, err := cryptography.DecryptBytesAge(encryptedBytes)
//	if err != nil {
//	    log.Fatalf("Failed to decrypt: %v", err)
//	}
//	fmt.Println("Decrypted data:", string(decryptedData))
func DecryptBytesAge(data []byte) ([]byte, error) {
	rkt := RootKeyCollectionFromMemory()
	privateKey := rkt.PrivateKey

	identity, err := age.ParseX25519Identity(privateKey)
	if err != nil {
		return []byte{}, errors.Join(
			err,
			errors.New("decryptBytes: failed to parse private key"),
		)
	}

	if len(data) == 0 {
		return []byte{}, errors.Join(
			err,
			errors.New("decryptBytes: file on disk appears to be empty"),
		)
	}

	out := &bytes.Buffer{}
	f := bytes.NewReader(data)

	r, err := age.Decrypt(f, identity)
	if err != nil {
		return []byte{}, errors.Join(
			err,
			errors.New("decryptBytes: failed to open encrypted file"),
		)
	}

	if _, err := io.Copy(out, r); err != nil {
		return []byte{}, errors.Join(
			err,
			errors.New("decryptBytes: failed to read encrypted file"),
		)
	}

	return out.Bytes(), nil
}

// DecryptBytesAes decrypts data that has been encrypted using AES encryption.
// This function assumes that the AES key is retrieved from a key-holder in
// memory and that the initial vector (IV) is prepended to the ciphertext. The
// function supports AES encryption modes that use CFB (Cipher Feedback Mode).
//
// Parameters:
//   - data: a byte slice containing the encrypted data, with the IV at the
//     beginning.
//
// Returns:
//   - A byte slice containing the decrypted data if the process is successful.
//   - An error if any step of the decryption process fails, including key
//     retrieval, key decoding, cipher block creation, or data processing.
//
// Usage:
//
//	decryptedData, err := crypto.DecryptBytesAes(encryptedData)
//	if err != nil {
//	    log.Fatalf("Failed to decrypt: %v", err)
//	}
//	fmt.Println("Decrypted data:", string(decryptedData))
func DecryptBytesAes(data []byte) ([]byte, error) {
	rkt := RootKeyCollectionFromMemory()
	aesKey := rkt.AesSeed

	aesKeyDecoded, err := hex.DecodeString(aesKey)
	if err != nil {
		return nil, errors.Join(
			err,
			errors.New("encryptToWriter: failed to decode AES key"),
		)
	}

	block, err := aes.NewCipher(aesKeyDecoded)
	if err != nil {
		return nil, errors.Join(
			err,
			errors.New("decryptBytes: failed to create AES cipher block"),
		)
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

// Decrypt decrypts the provided base64-encoded value using the specified
// algorithm. This function first decodes the base64-encoded input, and then
// depending on the algorithm specified, it either decrypts using age or AES.
//
// Parameters:
//   - value: a byte slice containing the base64-encoded data to be decrypted.
//   - algorithm: the Algorithm type specifying which decryption method to use.
//
// Returns:
//   - A string containing the decrypted data if successful.
//   - An error if decoding fails, or if the decryption process encounters an
//     issue.
//
// Usage:
//
//	decryptedText, err := crypto.Decrypt(encodedValue, cryptography.AES)
//	if err != nil {
//	    log.Fatalf("Failed to decrypt: %v", err)
//	}
//	fmt.Println("Decrypted text:", decryptedText)
func Decrypt(value []byte, algorithm c.Algorithm) (string, error) {
	decodedValue, err := base64.StdEncoding.DecodeString(string(value))
	if err != nil {
		return "", err
	}

	if algorithm == c.Age {
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
// associated with that key from the disk. The key is used to locate the data
// file, which is expected to have a ".age" extension
// and to be stored in a directory specified by the environment's data path for
// safe storage.
//
// Parameters:
//   - key (string): A string representing the unique identifier for the data
//     to be decrypted. The actual data file is expected to be named using this
//     key with a ".age" extension.
//
// Returns:
//   - ([]byte, error): This function returns two values. The first value is a
//     byte slice containing the decrypted data if the process is successful.
//     The second value is an error object that will be non-nil if any step of
//     the decryption process fails. Possible errors include the absence of the
//     target data file on disk and failures related to reading the file or the
//     decryption process itself.
func DecryptDataFromDisk(key string) ([]byte, error) {
	dataPath := path.Join(env.DataPathForSafe(), key+".age")

	if _, err := os.Stat(dataPath); os.IsNotExist(err) {
		return nil, errors.Join(
			err,
			errors.New("decryptDataFromDisk: No file at: "+dataPath),
		)
	}

	data, err := os.ReadFile(dataPath)
	if err != nil {
		return nil, errors.Join(
			err,
			errors.New("decryptDataFromDisk: Error reading file"),
		)
	}

	if env.FipsCompliantModeForSafe() {
		return DecryptBytesAes(data)
	}

	return DecryptBytesAge(data)
}
