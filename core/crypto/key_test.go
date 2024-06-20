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
	"github.com/vmware-tanzu/secrets-manager/core/entity/v1/data"
	"testing"
)

func TestSetRootKeyInMemory(t *testing.T) {
	tests := []struct {
		name     string
		key      string
		expected string
	}{
		{"Set valid root key", "test-private-key\ntest-public-key\ntest-aes-seed", "test-private-key\ntest-public-key\ntest-aes-seed"},
		{"Set empty root key", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetRootKeyInMemory(tt.key)

			RootKeyLock.RLock()
			defer RootKeyLock.RUnlock()

			if RootKey != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, RootKey)
			}
		})
	}
}

func TestRootKeySetInMemory(t *testing.T) {
	SetRootKeyInMemory("test-key")
	if !RootKeySetInMemory() {
		t.Errorf("expected true, got false")
	}

	SetRootKeyInMemory("")
	if RootKeySetInMemory() {
		t.Errorf("expected false, got true")
	}
}

func TestRootKeyCollection_Combine(t *testing.T) {
	rkt := data.RootKeyCollection{
		PrivateKey: "private-key",
		PublicKey:  "public-key",
		AesSeed:    "aes-seed",
	}

	expected := "private-key\npublic-key\naes-seed"
	result := rkt.Combine()
	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestRootKeyCollectionFromMemory(t *testing.T) {
	SetRootKeyInMemory("private-key\npublic-key\naes-seed")
	rkt := RootKeyCollectionFromMemory()

	expected := data.RootKeyCollection{
		PrivateKey: "private-key",
		PublicKey:  "public-key",
		AesSeed:    "aes-seed",
	}

	if rkt != expected {
		t.Errorf("expected %v, got %v", expected, rkt)
	}

	SetRootKeyInMemory("")
	rkt = RootKeyCollectionFromMemory()
	expected = data.RootKeyCollection{}

	if rkt != expected {
		t.Errorf("expected empty RootKeyCollection, got %v", rkt)
	}
}
