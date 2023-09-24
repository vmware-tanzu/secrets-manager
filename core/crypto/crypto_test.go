/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware, Inc.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package crypto

import (
	"errors"
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
	}{
		{
			name: "success_case",
			args: args{
				n: 8,
			},
			want:    8,
			wantErr: nil,
		},
		{
			name: "failure_case",
			setup: func() {
				reader = func(b []byte) (n int, err error) {
					return 0, errors.New("failed during rand.Read() call")
				}
			},
			args: args{
				n: 8,
			},
			want:    0,
			wantErr: errors.New("failed during rand.Read() call"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup()
			}
			got, err := RandomString(tt.args.n)
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
