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
	"crypto/rand"
	"errors"
	"github.com/vmware-tanzu/secrets-manager/lib/crypto"
	"testing"
)

func TestRandomString(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		setup   func()
		name    string
		args    args
		want    int
		wantErr error
		cleanup func()
	}{
		{
			name: "success_case",
			args: args{
				n: 8,
			},
			want:    8,
			wantErr: nil,
		},
		//{
		//	name: "failure_case",
		//	setup: func() {
		//		reader = func(b []byte) (n int, err error) {
		//			return 0, errors.New("failed during rand.Read() call")
		//		}
		//	},
		//	args: args{
		//		n: 8,
		//	},
		//	want:    0,
		//	wantErr: errors.New("failed during rand.Read() call"),
		//	cleanup: func() {
		//		reader = rand.Read
		//	},
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup()
				defer tt.cleanup()
			}
			got, err := crypto.RandomString(tt.args.n)
			if (err != nil) && err.Error() != tt.wantErr.Error() {
				t.Errorf("RandomString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.want {
				t.Errorf("RandomString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateAesSeed(t *testing.T) {
	tests := []struct {
		setup   func()
		name    string
		want    int
		wantErr error
		cleanup func()
	}{
		{
			name:    "success_case",
			want:    64,
			wantErr: nil,
		},
		{
			setup: func() {
				reader = func(b []byte) (n int, err error) {
					return 0, errors.New("failed during rand.Read() call")
				}
			},
			name:    "failure_case",
			want:    0,
			wantErr: errors.New("failed during rand.Read() call\ngenerateAesSeed: failed to generate random key"),
			cleanup: func() {
				reader = rand.Read
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup()
				defer tt.cleanup()
			}
			got, err := generateAesSeed()
			if (err != nil) && err.Error() != tt.wantErr.Error() {
				t.Errorf("generateAesSeed() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.want {
				t.Errorf("generateAesSeed() = %v, want %v", got, tt.want)
			}
		})
	}
}

//func TestGenerateKeys(t *testing.T) {
//	tests := []struct {
//		name    string
//		want    int
//		want1   int
//		want2   int
//		wantErr bool
//	}{
//		{
//			name:    "success_case",
//			want:    74,
//			want1:   62,
//			want2:   64,
//			wantErr: false,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, got1, got2, err := NewRootKeyCollection()
//			if (err != nil) != tt.wantErr {
//				t.Errorf("NewRootKeyCollection() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if len(got) != tt.want {
//				t.Errorf("NewRootKeyCollection() got = %v, want %v", got, tt.want)
//			}
//			if len(got1) != tt.want1 {
//				t.Errorf("NewRootKeyCollection() got1 = %v, want %v", got1, tt.want1)
//			}
//			if len(got2) != tt.want2 {
//				t.Errorf("NewRootKeyCollection() got2 = %v, want %v", got2, tt.want2)
//			}
//		})
//	}
//}

//func TestCombineKeys(t *testing.T) {
//	type args struct {
//		privateKey string
//		publicKey  string
//		aesSeed    string
//	}
//	tests := []struct {
//		name string
//		args args
//		want string
//	}{
//		{
//			name: "success_case",
//			args: args{
//				privateKey: "key-1",
//				publicKey:  "key-2",
//				aesSeed:    "key-3",
//			},
//			want: "key-1" + "\n" + "key-2" + "\n" + "key-3",
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := CombineKeys(tt.args.privateKey, tt.args.publicKey, tt.args.aesSeed); got != tt.want {
//				t.Errorf("CombineKeys() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
