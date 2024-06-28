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
	"encoding/base64"
	"encoding/hex"
	"os"
	"path"
	"reflect"
	"strings"
	"testing"

	"github.com/vmware-tanzu/secrets-manager/core/constants/crypto"
	"github.com/vmware-tanzu/secrets-manager/core/constants/env"
)

func TestDecryptValue(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name    string
		args    args
		setup   func(*args)
		cleanup func()
		want    string
		wantErr bool
	}{
		{
			name: "Invalid base64 encoded value, should return error",
			args: args{
				value: "invalid-value",
			},
			want:    "",
			wantErr: true,
		},
		{
			name:    "Valid encrypted value with AES algorithm but no root key, should return error",
			want:    "",
			wantErr: true,
			setup: func(args *args) {
				rtk, err := NewRootKeyCollection()
				if err != nil {
					t.Errorf("failed to generate root key collection: %v", err)
				}
				SetRootKeyInMemory(rtk.PrivateKey + "\n" + rtk.PublicKey + "\n" + rtk.AesSeed)

				_ = os.Setenv(string(env.VSecMSafeFipsCompliant), "true")
				_ = os.Setenv(string(env.VSecMSafeIvInitializationInterval), "0")

				encrypted, err := EncryptValue("test-data")
				if err != nil {
					t.Errorf("failed to encrypt data: %v", err)
				}
				args.value = encrypted

				SetRootKeyInMemory("")
			},

			cleanup: func() {
				_ = os.Unsetenv(string(env.VSecMSafeFipsCompliant))
				_ = os.Unsetenv(string(env.VSecMSafeIvInitializationInterval))
			},
		},
		{
			name:    "Valid encrypted value with Age algorithm but no root key, should return error",
			want:    "",
			wantErr: true,
			setup: func(args *args) {
				rtk, err := NewRootKeyCollection()
				if err != nil {
					t.Errorf("failed to generate root key collection: %v", err)
				}
				SetRootKeyInMemory(rtk.PrivateKey + "\n" + rtk.PublicKey + "\n" + rtk.AesSeed)

				_ = os.Setenv(string(env.VSecMSafeIvInitializationInterval), "0")

				encrypted, err := EncryptValue("test-data")
				if err != nil {
					t.Errorf("failed to encrypt data: %v", err)
				}
				args.value = encrypted

				SetRootKeyInMemory("")
			},
			cleanup: func() { _ = os.Unsetenv(string(env.VSecMSafeIvInitializationInterval)) },
		},
		{
			name:    "Valid encrypted value and root-key-collection with AES algorithm, should return decrypted data",
			want:    "test-data",
			wantErr: false,
			setup: func(args *args) {
				rtk, err := NewRootKeyCollection()
				if err != nil {
					t.Errorf("failed to generate root key collection: %v", err)
				}
				SetRootKeyInMemory(rtk.PrivateKey + "\n" + rtk.PublicKey + "\n" + rtk.AesSeed)

				_ = os.Setenv(string(env.VSecMSafeFipsCompliant), "true")
				_ = os.Setenv(string(env.VSecMSafeIvInitializationInterval), "0")

				encrypted, err := EncryptValue("test-data")
				if err != nil {
					t.Errorf("failed to encrypt data: %v", err)
				}
				args.value = encrypted
			},
			cleanup: func() {
				_ = os.Unsetenv(string(env.VSecMSafeFipsCompliant))
				_ = os.Unsetenv(string(env.VSecMSafeIvInitializationInterval))
				SetRootKeyInMemory("")
			},
		},
		{
			name:    "Valid encrypted value and root-key-collection with Age algorithm, should return decrypted data",
			want:    "test-data",
			wantErr: false,
			setup: func(args *args) {
				rtk, err := NewRootKeyCollection()
				if err != nil {
					t.Errorf("failed to generate root key collection: %v", err)
				}
				SetRootKeyInMemory(rtk.PrivateKey + "\n" + rtk.PublicKey + "\n" + rtk.AesSeed)

				_ = os.Setenv(string(env.VSecMSafeIvInitializationInterval), "0")

				encrypted, err := EncryptValue("test-data")
				if err != nil {
					t.Errorf("failed to encrypt data: %v", err)
				}
				args.value = encrypted
			},
			cleanup: func() {
				_ = os.Unsetenv(string(env.VSecMSafeIvInitializationInterval))
				SetRootKeyInMemory("")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup(&tt.args)
			}
			got, err := DecryptValue(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecryptValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DecryptValue() = %v, want %v", got, tt.want)
			}
			if tt.cleanup != nil {
				tt.cleanup()
			}
		})
	}
}

func TestDecryptBytesAes(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		setup   func(*args)
		cleanup func()
		want    []byte
		wantErr bool
	}{
		{
			name:    "Invalid AES key, should return error",
			want:    nil,
			wantErr: true,
			setup: func(*args) {
				rtk, err := NewRootKeyCollection()
				if err != nil {
					t.Errorf("failed to generate root key collection: %v", err)
				}
				SetRootKeyInMemory(rtk.PrivateKey + "\n" + rtk.PublicKey + "\n" + "invalid-aes-key")
			},
			cleanup: func() { SetRootKeyInMemory("") },
		},
		{
			name:    "Invalid decoded AES key length, should return error",
			want:    nil,
			wantErr: true,
			setup: func(*args) {
				aesKey := "must-16-24-32-bytes"
				aesSeed := make([]byte, hex.EncodedLen(len(aesKey)))
				_ = hex.Encode(aesSeed, []byte(aesKey))
				SetRootKeyInMemory("private-key\npublic-key\n" + string(aesSeed))
			},
			cleanup: func() { SetRootKeyInMemory("") },
		},
		{
			name:    "Data length less than AES block size, should return error",
			want:    nil,
			wantErr: true,
			setup: func(args *args) {
				rtk, err := NewRootKeyCollection()
				if err != nil {
					t.Errorf("failed to generate root key collection: %v", err)
				}
				SetRootKeyInMemory(rtk.PrivateKey + "\n" + rtk.PublicKey + "\n" + rtk.AesSeed)

				args.data = []byte("invalid-data")
			},
			cleanup: func() { SetRootKeyInMemory("") },
		},
		{
			name:    "Valid AES key, length and data, should return decrypted data",
			want:    []byte("some-texts"),
			wantErr: false,
			setup: func(args *args) {
				rkt, err := NewRootKeyCollection()
				if err != nil {
					t.Errorf("failed to generate root key collection: %v", err)
				}
				SetRootKeyInMemory(rkt.PrivateKey + "\n" + rkt.PublicKey + "\n" + rkt.AesSeed)

				_ = os.Setenv(string(env.VSecMSafeFipsCompliant), "true")
				_ = os.Setenv(string(env.VSecMSafeIvInitializationInterval), "0")

				encrypted, err := EncryptValue("some-texts")
				if err != nil {
					t.Errorf("failed to encrypt data: %v", err)
				}
				decodedValue, err := base64.StdEncoding.DecodeString(encrypted)
				if err != nil {
					t.Errorf("failed to decode data: %v", err)
				}

				args.data = decodedValue
			},
			cleanup: func() {
				_ = os.Unsetenv(string(env.VSecMSafeFipsCompliant))
				_ = os.Unsetenv(string(env.VSecMSafeIvInitializationInterval))
				SetRootKeyInMemory("")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup(&tt.args)
			}
			got, err := DecryptBytesAes(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecryptBytesAes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecryptBytesAes() got = %v, want %v", got, tt.want)
			}
			if tt.cleanup != nil {
				tt.cleanup()
			}
		})
	}
}

func TestDecryptBytesAge(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		setup   func(*args)
		cleanup func()
		want    []byte
		wantErr bool
	}{
		{
			name:    "Malformed, mixed case private key, should return error",
			want:    nil,
			wantErr: true,
			setup: func(*args) {
				rkt, err := NewRootKeyCollection()
				if err != nil {
					t.Errorf("failed to generate root key collection: %v", err)
				}
				rkt.PrivateKey = strings.ToLower(rkt.PrivateKey[:len(rkt.PrivateKey)/2]) +
					strings.ToUpper(rkt.PrivateKey[len(rkt.PrivateKey)/2:])

				SetRootKeyInMemory(rkt.PrivateKey + "\n" + rkt.PublicKey + "\n" + rkt.AesSeed)
			},
			cleanup: func() {
				SetRootKeyInMemory("")
			},
		},
		{
			name: "Empty data, should return error",
			args: args{
				data: []byte(""),
			},
			want:    nil,
			wantErr: true,
			setup: func(*args) {
				rkt, err := NewRootKeyCollection()
				if err != nil {
					t.Errorf("failed to generate root key collection: %v", err)
				}
				SetRootKeyInMemory(rkt.PrivateKey + "\n" + rkt.PublicKey + "\n" + rkt.AesSeed)
			},
			cleanup: func() {
				SetRootKeyInMemory("")
			},
		},
		{
			name: "Invalid encrypted data, should return error",
			args: args{
				data: []byte("invalid-data"),
			},
			want:    nil,
			wantErr: true,
			setup: func(*args) {
				rkt, err := NewRootKeyCollection()
				if err != nil {
					t.Errorf("failed to generate root key collection: %v", err)
				}
				SetRootKeyInMemory(rkt.PrivateKey + "\n" + rkt.PublicKey + "\n" + rkt.AesSeed)
			},
			cleanup: func() { SetRootKeyInMemory("") },
		},
		{
			name:    "Valid encrypted data, should return decrypted data",
			args:    args{},
			want:    []byte("success-story"),
			wantErr: false,
			setup: func(args *args) {
				rkt, err := NewRootKeyCollection()
				if err != nil {
					t.Errorf("failed to generate root key collection: %v", err)
				}
				SetRootKeyInMemory(rkt.PrivateKey + "\n" + rkt.PublicKey + "\n" + rkt.AesSeed)

				_ = os.Setenv(string(env.VSecMSafeIvInitializationInterval), "0")

				encrypted, err := EncryptValue("success-story")
				if err != nil {
					t.Errorf("failed to encrypt data: %v", err)
				}
				decodedValue, err := base64.StdEncoding.DecodeString(encrypted)
				if err != nil {
					t.Errorf("failed to decode data: %v", err)
				}

				args.data = decodedValue
			},
			cleanup: func() {
				_ = os.Unsetenv(string(env.VSecMSafeIvInitializationInterval))
				SetRootKeyInMemory("")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup(&tt.args)
			}
			got, err := DecryptBytesAge(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecryptBytesAge() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecryptBytesAge() got = %v, want %v", got, tt.want)
			}
			if tt.cleanup != nil {
				tt.cleanup()
			}
		})
	}
}

func TestDecrypt(t *testing.T) {
	type args struct {
		value     []byte
		algorithm crypto.Algorithm
	}
	tests := []struct {
		name    string
		args    args
		setup   func(*args)
		cleanup func()
		want    string
		wantErr bool
	}{
		{
			name: "Value is not base64 encoded, should return error",
			args: args{
				value: []byte("invalid-value"),
			},
			want:    "",
			wantErr: true,
		},
		{
			name:    "Use Age algorithm with no root key, should return error",
			want:    "",
			wantErr: true,
			setup: func(args *args) {
				rkt, err := NewRootKeyCollection()
				if err != nil {
					t.Errorf("failed to generate root key collection: %v", err)
				}
				SetRootKeyInMemory(rkt.PrivateKey + "\n" + rkt.PublicKey + "\n" + rkt.AesSeed)

				_ = os.Setenv(string(env.VSecMSafeIvInitializationInterval), "0")

				args.algorithm = crypto.Age
				encrypted, err := EncryptValue("test-data")
				if err != nil {
					t.Errorf("failed to encrypt data: %v", err)
				}
				args.value = []byte(encrypted)

				SetRootKeyInMemory("")
			},
			cleanup: func() { _ = os.Unsetenv(string(env.VSecMSafeIvInitializationInterval)) },
		},
		{
			name:    "Use AES algorithm with no root key, should return error",
			want:    "",
			wantErr: true,
			setup: func(args *args) {
				rkt, err := NewRootKeyCollection()
				if err != nil {
					t.Errorf("failed to generate root key collection: %v", err)
				}
				SetRootKeyInMemory(rkt.PrivateKey + "\n" + rkt.PublicKey + "\n" + rkt.AesSeed)

				_ = os.Setenv(string(env.VSecMSafeIvInitializationInterval), "0")

				args.algorithm = crypto.Aes
				encrypted, err := EncryptValue("test-data")
				if err != nil {
					t.Errorf("failed to encrypt data: %v", err)
				}
				args.value = []byte(encrypted)

				SetRootKeyInMemory("")
			},
			cleanup: func() { _ = os.Unsetenv(string(env.VSecMSafeIvInitializationInterval)) },
		},
		{
			name:    "Valid encrypted data with AES algorithm, should return decrypted data",
			want:    "test-data",
			wantErr: false,
			setup: func(args *args) {
				rkt, err := NewRootKeyCollection()
				if err != nil {
					t.Errorf("failed to generate root key collection: %v", err)
				}
				SetRootKeyInMemory(rkt.PrivateKey + "\n" + rkt.PublicKey + "\n" + rkt.AesSeed)

				_ = os.Setenv(string(env.VSecMSafeFipsCompliant), "true")
				_ = os.Setenv(string(env.VSecMSafeIvInitializationInterval), "0")

				args.algorithm = crypto.Aes
				encrypted, err := EncryptValue("test-data")
				if err != nil {
					t.Errorf("failed to encrypt data: %v", err)
				}

				args.value = []byte(encrypted)
			},
			cleanup: func() {
				_ = os.Unsetenv(string(env.VSecMSafeFipsCompliant))
				_ = os.Unsetenv(string(env.VSecMSafeIvInitializationInterval))
				SetRootKeyInMemory("")
			},
		},
		{
			name:    "Valid encrypted data with Age algorithm, should return decrypted data",
			want:    "test-data",
			wantErr: false,
			setup: func(args *args) {
				rkt, err := NewRootKeyCollection()
				if err != nil {
					t.Errorf("failed to generate root key collection: %v", err)
				}
				SetRootKeyInMemory(rkt.PrivateKey + "\n" + rkt.PublicKey + "\n" + rkt.AesSeed)

				_ = os.Setenv(string(env.VSecMSafeIvInitializationInterval), "0")

				args.algorithm = crypto.Age
				encrypted, err := EncryptValue("test-data")
				if err != nil {
					t.Errorf("failed to encrypt data: %v", err)
				}
				args.value = []byte(encrypted)
			},
			cleanup: func() {
				_ = os.Unsetenv(string(env.VSecMSafeIvInitializationInterval))
				SetRootKeyInMemory("")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup(&tt.args)
			}
			got, err := Decrypt(tt.args.value, tt.args.algorithm)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Decrypt() = %v, want %v", got, tt.want)
			}
			if tt.cleanup != nil {
				tt.cleanup()
			}
		})
	}
}

func TestDecryptDataFromDisk(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		setup   func(*args)
		cleanup func()
		want    []byte
		wantErr bool
	}{
		{
			name: "File does not exist, should return error",
			args: args{
				key: "non-existent-file",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "File can't be read, should return error",
			args: args{
				key: "unreadable-file",
			},
			want:    nil,
			wantErr: true,
			setup: func(*args) {
				pwd, err := os.Getwd()
				if err != nil {
					t.Errorf("failed to get current working directory: %v", err)
				}
				tmpDir := path.Join(pwd, "tmp")
				_ = os.Setenv(string(env.VSecMSafeDataPath), tmpDir)

				if err := os.MkdirAll(env.Value(env.VSecMSafeDataPath), 0o777); err != nil {
					t.Errorf("failed to create directory: %v", err)
				}
				if err := os.WriteFile(path.Join(env.Value(env.VSecMSafeDataPath), "unreadable-file.age"),
					[]byte("test-data"),
					0o000,
				); err != nil {
					t.Errorf("failed to write file: %v", err)
				}
			},
			cleanup: func() {
				if err := os.RemoveAll(env.Value(env.VSecMSafeDataPath)); err != nil {
					t.Errorf("failed to remove directory: %v", err)
				}
				_ = os.Unsetenv(string(env.VSecMSafeDataPath))
			},
		},
		{
			name: "Valid file with FIPS compliance enabled, should return AES decrypted data",
			args: args{
				key: "test-file",
			},
			want:    []byte("test-data"),
			wantErr: false,
			setup: func(*args) {
				rkt, err := NewRootKeyCollection()
				if err != nil {
					t.Errorf("failed to generate root key collection: %v", err)
				}
				SetRootKeyInMemory(rkt.PrivateKey + "\n" + rkt.PublicKey + "\n" + rkt.AesSeed)

				pwd, err := os.Getwd()
				if err != nil {
					t.Errorf("failed to get current working directory: %v", err)
				}
				tmpDir := path.Join(pwd, "tmp")
				_ = os.Setenv(string(env.VSecMSafeDataPath), tmpDir)
				_ = os.Setenv(string(env.VSecMSafeFipsCompliant), "true")
				_ = os.Setenv(string(env.VSecMSafeIvInitializationInterval), "0")

				encrypted, err := EncryptValue("test-data")
				if err != nil {
					t.Errorf("failed to encrypt data: %v", err)
				}
				decodedValue, err := base64.StdEncoding.DecodeString(encrypted)
				if err != nil {
					t.Errorf("failed to decode data: %v", err)
				}

				if err := os.MkdirAll(env.Value(env.VSecMSafeDataPath), 0o777); err != nil {
					t.Errorf("failed to create directory: %v", err)
				}
				if err := os.WriteFile(path.Join(env.Value(env.VSecMSafeDataPath),
					"test-file.age"),
					decodedValue,
					0o777,
				); err != nil {
					t.Errorf("failed to write file: %v", err)
				}
			},
			cleanup: func() {
				_ = os.RemoveAll(env.Value(env.VSecMSafeDataPath))
				_ = os.Unsetenv(string(env.VSecMSafeDataPath))
				_ = os.Unsetenv(string(env.VSecMSafeFipsCompliant))
				_ = os.Unsetenv(string(env.VSecMSafeIvInitializationInterval))
				SetRootKeyInMemory("")
			},
		},
		{
			name: "Valid file with FIPS compliance disabled, should return Age decrypted data",
			args: args{
				key: "test-file",
			},
			want:    []byte("EncryptedAgeData"),
			wantErr: false,
			setup: func(args *args) {
				rkt, err := NewRootKeyCollection()
				if err != nil {
					t.Errorf("failed to generate root key collection: %v", err)
				}
				SetRootKeyInMemory(rkt.PrivateKey + "\n" + rkt.PublicKey + "\n" + rkt.AesSeed)

				pwd, err := os.Getwd()
				if err != nil {
					t.Errorf("failed to get current working directory: %v", err)
				}
				tmpDir := path.Join(pwd, "tmp")
				_ = os.Setenv(string(env.VSecMSafeDataPath), tmpDir)
				_ = os.Setenv(string(env.VSecMSafeIvInitializationInterval), "0")

				encrypted, err := EncryptValue("EncryptedAgeData")
				if err != nil {
					t.Errorf("failed to encrypt data: %v", err)
				}
				decodedValue, err := base64.StdEncoding.DecodeString(encrypted)
				if err != nil {
					t.Errorf("failed to decode data: %v", err)
				}

				if err := os.MkdirAll(env.Value(env.VSecMSafeDataPath), 0o777); err != nil {
					t.Errorf("failed to create directory: %v", err)
				}
				if err := os.WriteFile(path.Join(env.Value(env.VSecMSafeDataPath),
					"test-file.age"),
					decodedValue,
					0o777,
				); err != nil {
					t.Errorf("failed to write file: %v", err)
				}
			},
			cleanup: func() {
				if err := os.RemoveAll(env.Value(env.VSecMSafeDataPath)); err != nil {
					t.Errorf("failed to remove directory: %v", err)
				}
				_ = os.Unsetenv(string(env.VSecMSafeDataPath))
				_ = os.Unsetenv(string(env.VSecMSafeIvInitializationInterval))
				SetRootKeyInMemory("")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup(&tt.args)
			}
			got, err := DecryptDataFromDisk(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecryptDataFromDisk() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecryptDataFromDisk() got = %s, want %s", got, tt.want)
			}
			if tt.cleanup != nil {
				tt.cleanup()
			}
		})
	}
}
