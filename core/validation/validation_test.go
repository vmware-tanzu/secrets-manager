/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware, Inc.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package validation

import (
	"testing"
)

func TestIsSentinel(t *testing.T) {
	type args struct {
		svid string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "has_prefix",
			args: args{
				svid: "spiffe://vsecm.com/workload/vsecm-sentinel/ns/vsecm-system/sa/vsecm-sentinel/n/test",
			},
			want: true,
		},
		{
			name: "does_not_have_prefix",
			args: args{
				svid: "test",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSentinel(tt.args.svid); got != tt.want {
				t.Errorf("IsSentinel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsSafe(t *testing.T) {
	type args struct {
		svid string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "has_prefix",
			args: args{
				svid: "spiffe://vsecm.com/workload/vsecm-safe/ns/vsecm-system/sa/vsecm-safe/n/test",
			},
			want: true,
		},
		{
			name: "does_not_have_prefix",
			args: args{
				svid: "test",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSafe(tt.args.svid); got != tt.want {
				t.Errorf("IsSafe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsNotary(t *testing.T) {
	type args struct {
		svid string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "has_prefix",
			args: args{
				svid: "spiffe://vsecm.com/workload/vsecm-notary/ns/vsecm-system/sa/vsecm-notary/n/test",
			},
			want: true,
		},
		{
			name: "does_not_have_prefix",
			args: args{
				svid: "test",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNotary(tt.args.svid); got != tt.want {
				t.Errorf("IsNotary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsWorkload(t *testing.T) {
	type args struct {
		svid string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "has_prefix",
			args: args{
				svid: "spiffe://vsecm.com/workload/test",
			},
			want: true,
		},
		{
			name: "does_not_have_prefix",
			args: args{
				svid: "test",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsWorkload(tt.args.svid); got != tt.want {
				t.Errorf("IsWorkload() = %v, want %v", got, tt.want)
			}
		})
	}
}
