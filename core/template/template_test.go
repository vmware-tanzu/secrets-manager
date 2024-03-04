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

func TestValidJSON(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid_json",
			args: args{
				s: "{\"key\":\"value\"}",
			},
			want: true,
		},
		{
			name: "invalid_json",
			args: args{
				s: "{\"value\"}",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidJSON(tt.args.s); got != tt.want {
				t.Errorf("ValidJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJsonToYaml(t *testing.T) {
	type args struct {
		js string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr string
	}{
		{
			name: "valid_json",
			args: args{
				js: "{\"key\":\"value\"}",
			},
			want:    "key: value\n",
			wantErr: "",
		},
		{
			name: "invalid_json",
			args: args{
				js: "{\"value\"}",
			},
			want:    "",
			wantErr: "invalid character '}' after object key",
		},
		{
			name: "invalid_json",
			args: args{
				js: "{\"value\"}",
			},
			want:    "",
			wantErr: "invalid character '}' after object key",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := JsonToYaml(tt.args.js)
			if (tt.wantErr == "" && err != nil) ||
				(tt.wantErr != "" && err == nil) ||
				(tt.wantErr != "" && err != nil && tt.wantErr != err.Error()) {
				t.Errorf("JsonToYaml() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("JsonToYaml() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTryParse(t *testing.T) {
	type args struct {
		tmpStr string
		jason  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "valid_template_valid_json",
			args: args{
				tmpStr: "{\"USER\":\"{{.username}}\", \"PASS\":\"{{.password}}\"}",
				jason:  "{\"username\":\"admin\",\"password\":\"VSecMRocks\"}",
			},
			want: "{\"USER\":\"admin\", \"PASS\":\"VSecMRocks\"}",
		},
		{
			name: "valid_template_invalid_json",
			args: args{
				tmpStr: "{\"USER\":\"{{.username}}\", \"PASS\":\"{{.password}}\"}",
				jason:  "{username}",
			},
			want: "{username}",
		},
		{
			name: "invalid_template_valid_json",
			args: args{
				tmpStr: "{\"USER\":\"{{}}\", \"PASS\":\"{{.password}}\"}",
				jason:  "{\"username\":\"admin\",\"password\":\"VSecMRocks\"}",
			},
			want: "{\"username\":\"admin\",\"password\":\"VSecMRocks\"}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TryParse(tt.args.tmpStr, tt.args.jason); got != tt.want {
				t.Errorf("TryParse() = %v, want %v", got, tt.want)
			}
		})
	}
}
