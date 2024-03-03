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
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"
	"math"
	"time"

	"filippo.io/age"
	"github.com/pkg/errors"

	"github.com/vmware-tanzu/secrets-manager/core/env"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
)

// EncryptToWriterAge encrypts the provided data using a public key and writes
// the encrypted data to the specified io.Writer. The encryption is performed
// using the age encryption protocol, which is designed for simple, secure, and
// modern encryption.
//
// Parameters:
//   - out (io.Writer): The writer interface where the encrypted data will be written.
//     This could be a file, a network connection, or any other type that implements
//     the io.Writer interface.
//   - data (string): The plaintext data that needs to be encrypted. This function
//     converts the string to a byte slice and encrypts it.
//
// Returns:
//   - error: If any step in the encryption or writing process fails, an error is returned.
//     Possible errors include issues with the public key (such as being empty or unparseable)
//     and failures related to the encryption process or writing to the writer interface.
func EncryptToWriterAge(out io.Writer, data string, rootKeyTriplet []string) error {
	_, publicKey, _ := rootKeyTriplet[0], rootKeyTriplet[1], rootKeyTriplet[2]

	if publicKey == "" {
		return errors.New("encryptToWriterAge: no public key")
	}

	recipient, err := age.ParseX25519Recipient(publicKey)
	if err != nil {
		return errors.Wrap(err, "encryptToWriterAge: failed to parse public key")
	}

	wrappedWriter, err := age.Encrypt(out, recipient)
	if err != nil {
		return errors.Wrap(err, "encryptToWriterAge: failed to create encrypted file")
	}

	defer func() {
		err := wrappedWriter.Close()
		if err != nil {
			id := "AEGIIOCL"
			log.InfoLn(&id, "encryptToWriterAge: problem closing stream", err.Error())
		}
	}()

	if _, err := io.WriteString(wrappedWriter, data); err != nil {
		return errors.Wrap(err, "encryptToWriterAge: failed to write to encrypted file")
	}

	return nil
}

var lastEncryptToWriterAesCall time.Time

// EncryptToWriterAes encrypts the given data string using the AES encryption standard
// and writes the encrypted data to the specified io.Writer. This function emphasizes
// secure encryption practices, including the management of the Initialization
// Vector (IV) and the AES key.
//
// Parameters:
//   - out (io.Writer): The output writer where the encrypted data will be written.
//     This writer can represent various types of data sinks, such as files, network
//     connections, or in-memory buffers.
//   - data (string): The plaintext data to be encrypted. This data is converted to
//     a byte slice and then encrypted using AES.
//
// Returns:
//   - error: An error is returned if any step of the encryption or writing process
//     encounters an issue. This includes errors related to call frequency, key
//     management, encryption initialization, and data writing.
func EncryptToWriterAes(out io.Writer, data string, rootKeyTriplet []string) error {
	// Calling this method too frequently can result in a less-than random IV,
	// which can be used to break the encryption when combined with other
	// attack vectors. Therefore, we throttle calls to this method.
	if time.Since(lastEncryptToWriterAesCall) < time.Millisecond*time.Duration(
		env.IvInitializationIntervalForSafe(),
	) {
		return errors.New("Calls too frequent")
	}

	lastEncryptToWriterAesCall = time.Now()

	_, _, aesKey := rootKeyTriplet[0], rootKeyTriplet[1], rootKeyTriplet[2]

	if aesKey == "" {
		return errors.New("encryptToWriter: no AES key")
	}

	aesKeyDecoded, err := hex.DecodeString(aesKey)
	defer func() {
		// Clear the key from memory for security reasons.
		for i := range aesKeyDecoded {
			aesKeyDecoded[i] = 0
		}
	}()

	if err != nil {
		return errors.Wrap(err, "encryptToWriter: failed to decode AES key")
	}

	block, err := aes.NewCipher(aesKeyDecoded)
	if err != nil {
		return errors.Wrap(err, "encryptToWriter: failed to create AES cipher block")
	}

	totalSize := uint64(aes.BlockSize) + uint64(len(data))
	if totalSize > uint64(math.MaxInt64) {
		return errors.New("encryptToWriter: data too large")
	}

	// The IV needs to be unique, but not secure. Therefore, it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, totalSize)

	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(data))

	_, err = out.Write(ciphertext)
	if err != nil {
		return errors.Wrap(err, "encryptToWriter: failed to write to encrypted file")
	}

	return nil
}
