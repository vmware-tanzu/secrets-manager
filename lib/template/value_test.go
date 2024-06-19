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

import "testing"

func TestGenerateValue(t *testing.T) {
	type args struct {
		template string
	}
	tests := []struct {
		name        string
		args        args
		wantErr     bool
		errorOutput string
	}{
		{
			name: "Success case for alphanumeric",
			args: args{
				template: `football[\w]{8}bartender`,
			},
			wantErr:     false,
			errorOutput: "",
		},
		{
			name: "Success case for alphanumeric and symbol",
			args: args{
				template: `football[\x]{4}bartender`,
			},
			wantErr:     false,
			errorOutput: "",
		},
		{
			name: "Fail case for invalid regex",
			args: args{
				template: `pass[z-a]{8}`,
			},
			wantErr:     true,
			errorOutput: "invalid range specified: z-a",
		},
		{
			name: "Fail case for invalid range",
			args: args{
				template: `pass[a-Z]{8`,
			},
			wantErr:     true,
			errorOutput: "no generator expressions found",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Value(tt.args.template)
			if (err != nil) != tt.wantErr {
				t.Errorf("Value() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err.Error() != tt.errorOutput {
					t.Errorf("Value() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
			t.Logf("Value() = %v", got)
		})
	}
}
