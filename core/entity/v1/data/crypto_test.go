/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package data

import (
	"reflect"
	"testing"
)

func TestRootKeyCollection_Combine(t *testing.T) {
	type fields struct {
		PrivateKey string
		PublicKey  string
		AesSeed    string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "empty rkt keys",
			fields: fields{
				PrivateKey: "",
				PublicKey:  "",
				AesSeed:    "",
			},
			want: "",
		},
		{
			name: "initialized rkt keys",
			fields: fields{
				PrivateKey: "test-pvt-key",
				PublicKey:  "test-pub-key",
				AesSeed:    "test-seed",
			},
			want: "test-pvt-key\ntest-pub-key\ntest-seed",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rkt := &RootKeyCollection{
				PrivateKey: tt.fields.PrivateKey,
				PublicKey:  tt.fields.PublicKey,
				AesSeed:    tt.fields.AesSeed,
			}
			if got := rkt.Combine(); got != tt.want {
				t.Errorf("RootKeyCollection.Combine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRootKeyCollection_Empty(t *testing.T) {
	type fields struct {
		PrivateKey string
		PublicKey  string
		AesSeed    string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "empty rkt keys",
			fields: fields{
				PrivateKey: "",
				PublicKey:  "",
				AesSeed:    "",
			},
			want: true,
		},
		{
			name: "non empty private key",
			fields: fields{
				PrivateKey: "pvt-key",
				PublicKey:  "",
				AesSeed:    "",
			},
			want: false,
		},
		{
			name: "non empty public key",
			fields: fields{
				PrivateKey: "",
				PublicKey:  "pub-key",
				AesSeed:    "",
			},
			want: false,
		},
		{
			name: "non empty aes seed",
			fields: fields{
				PrivateKey: "",
				PublicKey:  "",
				AesSeed:    "aes-seed",
			},
			want: false,
		},
		{
			name: "non empty private and public keys",
			fields: fields{
				PrivateKey: "pvt-key",
				PublicKey:  "pub-key",
				AesSeed:    "",
			},
			want: false,
		},
		{
			name: "non empty private key and aes seed",
			fields: fields{
				PrivateKey: "pvt-key",
				PublicKey:  "",
				AesSeed:    "aes-seed",
			},
			want: false,
		},
		{
			name: "non empty public key and aes seed",
			fields: fields{
				PrivateKey: "",
				PublicKey:  "pub-key",
				AesSeed:    "aes-seed",
			},
			want: false,
		},
		{
			name: "non empty public, private key, and aes seed",
			fields: fields{
				PrivateKey: "pvt-key",
				PublicKey:  "pub-key",
				AesSeed:    "aes-seed",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rkt := &RootKeyCollection{
				PrivateKey: tt.fields.PrivateKey,
				PublicKey:  tt.fields.PublicKey,
				AesSeed:    tt.fields.AesSeed,
			}
			if got := rkt.Empty(); got != tt.want {
				t.Errorf("RootKeyCollection.Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRootKeyCollection_UpdateFromSerialized(t *testing.T) {
	tests := []struct {
		name       string
		serialized string
		want       *RootKeyCollection
	}{
		{
			name:       "empty serialized",
			serialized: "",
			want: &RootKeyCollection{
				PrivateKey: "",
				PublicKey:  "",
				AesSeed:    "",
			},
		},
		{
			name:       "serialized with no separator",
			serialized: "pvt-key",
			want: &RootKeyCollection{
				PrivateKey: "",
				PublicKey:  "",
				AesSeed:    "",
			},
		},
		{
			name:       "serialized with one separator",
			serialized: "pvt-key\npub-key",
			want: &RootKeyCollection{
				PrivateKey: "",
				PublicKey:  "",
				AesSeed:    "",
			},
		},
		{
			name:       "serialized with two separator",
			serialized: "pvt-key\npub-key\naes-seed",
			want: &RootKeyCollection{
				PrivateKey: "pvt-key",
				PublicKey:  "pub-key",
				AesSeed:    "aes-seed",
			},
		},
		{
			name:       "serialized with four separator",
			serialized: "pvt-key\npub-key\naes-seed\n",
			want: &RootKeyCollection{
				PrivateKey: "pvt-key",
				PublicKey:  "pub-key",
				AesSeed:    "aes-seed",
			},
		},
		{
			name:       "serialized with three keys and with invalid separator",
			serialized: "pvt-key\tpub-key\taes-seed",
			want: &RootKeyCollection{
				PrivateKey: "",
				PublicKey:  "",
				AesSeed:    "",
			},
		},
		{
			name:       "empty serialized with valid separator",
			serialized: "\n\n",
			want: &RootKeyCollection{
				PrivateKey: "",
				PublicKey:  "",
				AesSeed:    "",
			},
		},
		{
			name:       "empty serialized with invalid separator",
			serialized: "\t\t",
			want: &RootKeyCollection{
				PrivateKey: "",
				PublicKey:  "",
				AesSeed:    "",
			},
		},
		{
			name:       "empty serialized with invalid separator",
			serialized: "\t\t",
			want: &RootKeyCollection{
				PrivateKey: "",
				PublicKey:  "",
				AesSeed:    "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rkt := &RootKeyCollection{}
			rkt.UpdateFromSerialized(tt.serialized)
			if !reflect.DeepEqual(rkt, tt.want) {
				t.Errorf("RootKeyCollection.UpdateFromSerialized() = %v, want %v", rkt, tt.want)
			}
		})
	}
}
