/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package template

import (
	"testing"
)

func Test_rangesAndLength(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		length  int
		wantErr bool
	}{
		{
			name: "Success with Integer length",
			args: args{
				"example-input{3}",
			},
			want:    "example-input",
			length:  3,
			wantErr: false,
		},
		{
			name: "Fail with Index out of range",
			args: args{
				"example-input{3",
			},
			want:    "example-input",
			length:  0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := rangesAndLength(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("rangesAndLength() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("rangesAndLength() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.length {
				t.Errorf("rangesAndLength() got1 = %v, want %v", got1, tt.length)
			}
		})
	}
}

func Test_parseLength(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Success with Integer length",
			args: args{
				"example-input{3}",
			},
			want:    3,
			wantErr: false,
		},
		{
			name: "Fail with Index out of range",
			args: args{
				"example-input{3",
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "Fail with String length",
			args: args{
				"example-input{three}",
			},

			want:    0,
			wantErr: true,
		},
		{
			name: "Fail with Special Character with Integer length",
			args: args{
				"example-input{3!}",
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "Fail with Float length",
			args: args{
				"example-input{3.0}",
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseLength(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseLength() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parseLength() got = %v, want %v", got, tt.want)
			}
		})
	}
}
