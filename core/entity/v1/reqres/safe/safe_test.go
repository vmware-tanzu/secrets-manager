/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package safe

import (
	"encoding/json"
	"github.com/spiffe/vsecm-sdk-go/core/entity/v1/reqres/safe"
	"github.com/vmware-tanzu/secrets-manager/core/constants/crypto"
	"github.com/vmware-tanzu/secrets-manager/core/entity/v1/data"
	"github.com/vmware-tanzu/secrets-manager/lib/entity"
	"reflect"
	"testing"
	"time"
)

func TestSecretUpsertRequest_JSONMarshalling(t *testing.T) {
	expected := safe.SecretUpsertRequest{
		WorkloadIds: []string{"workload1", "workload2"},
		Namespaces:  []string{"namespace1", "namespace2"},
		Value:       "secretValue",
		Template:    "template1",
		Format:      "json",
		Encrypt:     true,
		// AppendValue: false,
		NotBefore: "2024-01-01T00:00:00Z",
		Expires:   "2024-12-31T23:59:59Z",
	}

	bytes, err := json.Marshal(expected)
	if err != nil {
		t.Fatalf("Failed to marshal SecretUpsertRequest: %v", err)
	}

	var actual safe.SecretUpsertRequest
	if err := json.Unmarshal(bytes, &actual); err != nil {
		t.Fatalf("Failed to unmarshal SecretUpsertRequest: %v", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %+v, but got %+v", expected, actual)
	}
}

func TestKeyInputRequest_JSONMarshalling(t *testing.T) {
	expected := safe.KeyInputRequest{
		AgeSecretKey: "secretKey",
		AgePublicKey: "publicKey",
		AesCipherKey: "cipherKey",
	}

	bytes, err := json.Marshal(expected)
	if err != nil {
		t.Fatalf("Failed to marshal KeyInputRequest: %v", err)
	}

	var actual safe.KeyInputRequest
	if err := json.Unmarshal(bytes, &actual); err != nil {
		t.Fatalf("Failed to unmarshal KeyInputRequest: %v", err)
	}

	if expected != actual {
		t.Errorf("Expected %+v, but got %+v", expected, actual)
	}
}

// Similarly, you can add tests for other structs like SentinelInitCompleteRequest, SecretFetchRequest, etc.

func TestSecretListResponse_JSONMarshalling(t *testing.T) {
	createdTime := entity.JsonTime(time.Date(2024, time.January, 1, 12, 0, 0, 0, time.UTC))
	updatedTime := entity.JsonTime(time.Date(2024, time.January, 2, 12, 0, 0, 0, time.UTC))
	notBeforeTime := entity.JsonTime(time.Date(2024, time.January, 3, 12, 0, 0, 0, time.UTC))
	expiresAfterTime := entity.JsonTime(time.Date(2025, time.January, 1, 12, 0, 0, 0, time.UTC))

	//[]data.Secret{
	//		{
	//			Name:         "secret1",
	//			Created:      createdTime,
	//			Updated:      updatedTime,
	//			NotBefore:    notBeforeTime,
	//			ExpiresAfter: expiresAfterTime,
	//		},
	//		{
	//			Name:         "secret2",
	//			Created:      createdTime,
	//			Updated:      updatedTime,
	//			NotBefore:    notBeforeTime,
	//			ExpiresAfter: expiresAfterTime,
	//		},
	//	}

	expected := SecretListResponse{
		Secrets: []data.Secret{
			{
				Name:         "secret1",
				Created:      createdTime,
				Updated:      updatedTime,
				NotBefore:    notBeforeTime,
				ExpiresAfter: expiresAfterTime,
			},
			{
				Name:         "secret2",
				Created:      createdTime,
				Updated:      updatedTime,
				NotBefore:    notBeforeTime,
				ExpiresAfter: expiresAfterTime,
			},
		},
		Err: "example error",
	}

	bytes, err := json.Marshal(expected)
	if err != nil {
		t.Fatalf("Failed to marshal SecretListResponse: %v", err)
	}

	// Unmarshal the JSON back into a struct
	var actual SecretListResponse
	if err := json.Unmarshal(bytes, &actual); err != nil {
		t.Fatalf("Failed to unmarshal SecretListResponse: %v", err)
	}

	// Use reflect.DeepEqual to compare the structs
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %+v, but got %+v", expected, actual)
	}

}

func TestSecretEncrypted_JSONMarshalling(t *testing.T) {
	// Example times for testing
	createdTime := entity.JsonTime(time.Date(2024, time.January, 1, 12, 0, 0, 0, time.UTC))
	updatedTime := entity.JsonTime(time.Date(2024, time.January, 2, 12, 0, 0, 0, time.UTC))
	notBeforeTime := entity.JsonTime(time.Date(2024, time.January, 3, 12, 0, 0, 0, time.UTC))
	expiresAfterTime := entity.JsonTime(time.Date(2025, time.January, 1, 12, 0, 0, 0, time.UTC))

	expected := SecretEncryptedListResponse{
		Secrets: []data.SecretEncrypted{
			{
				Name:           "secret1",
				EncryptedValue: "encrypted_value_1",
				Created:        createdTime,
				Updated:        updatedTime,
				NotBefore:      notBeforeTime,
				ExpiresAfter:   expiresAfterTime,
			},
			{
				Name:           "secret2",
				EncryptedValue: "encrypted_value_3",
				Created:        createdTime,
				Updated:        updatedTime,
				NotBefore:      notBeforeTime,
				ExpiresAfter:   expiresAfterTime,
			},
		},
		Algorithm: crypto.Aes,
		Err:       "example error",
	}
	// Marshal the struct to JSON
	bytes, err := json.Marshal(expected)
	if err != nil {
		t.Fatalf("Failed to marshal SecretEncryptedListResponse: %v", err)
	}

	// Unmarshal the JSON back into a struct
	var actual SecretEncryptedListResponse
	if err := json.Unmarshal(bytes, &actual); err != nil {
		t.Fatalf("Failed to unmarshal SecretEncryptedListResponse: %v", err)
	}

	// Use reflect.DeepEqual to compare the structs
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %+v, but got %+v", expected, actual)
	}

}
