/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package validation

import (
	"testing"
)

func TestIsSentinel(t *testing.T) {
	type args struct {
		spiffeid string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "has_prefix",
			args: args{
				spiffeid: "spiffe://vsecm.com/workload/vsecm-sentinel/ns/vsecm-system/sa/vsecm-sentinel/n/test",
			},
			want: true,
		},
		{
			name: "does_not_have_prefix",
			args: args{
				spiffeid: "test",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSentinel(tt.args.spiffeid); got != tt.want {
				t.Errorf("IsSentinel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsSafe(t *testing.T) {
	type args struct {
		spiffeid string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "has_prefix",
			args: args{
				spiffeid: "spiffe://vsecm.com/workload/vsecm-safe/ns/vsecm-system/sa/vsecm-safe/n/test",
			},
			want: true,
		},
		{
			name: "does_not_have_prefix",
			args: args{
				spiffeid: "test",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSafe(tt.args.spiffeid); got != tt.want {
				t.Errorf("IsSafe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsWorkload(t *testing.T) {
	type args struct {
		spiffeid string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "has_prefix",
			args: args{
				spiffeid: "spiffe://vsecm.com/workload/test/ns/test/sa/test/n/test",
			},
			want: true,
		},
		{
			name: "does_not_have_prefix",
			args: args{
				spiffeid: "test",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsWorkload(tt.args.spiffeid); got != tt.want {
				t.Errorf("IsWorkload() = %v, want %v", got, tt.want)
			}
		})
	}
}
