/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package filter

import (
	"errors"
	"reflect"
	"testing"
)

func TestValueFromPath(t *testing.T) {
	type args struct {
		data any
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    any
		wantErr error
	}{
		{
			name:    "empty path",
			args:    args{path: "", data: ""},
			want:    "",
			wantErr: nil,
		},
		{
			name:    "path without dotted notation",
			args:    args{path: "samplePath", data: "sampleData"},
			want:    "sampleData",
			wantErr: nil,
		},
		{
			name:    "path with dotted notation, valid data",
			args:    args{path: "path1.path2", data: map[string]interface{}{"path1": map[string]interface{}{"path2": "data"}}},
			want:    "data",
			wantErr: nil,
		},
		{
			name:    "path with dotted notation, invalid key",
			args:    args{path: "invalidPath.path2", data: map[string]interface{}{"path1": map[string]interface{}{"path2": "data"}}},
			want:    nil,
			wantErr: errors.New("key not found: invalidPath"),
		},
		{
			name:    "path with dotted notation, array type data",
			args:    args{path: "path1.path2", data: map[string]interface{}{"path1": []interface{}{"data"}}},
			want:    nil,
			wantErr: errors.New("arrays are not supported in path queries"),
		},
		{
			name:    "path with dotted notation, invalid data",
			args:    args{path: "path1.path2", data: map[string]interface{}{"path1": "data"}},
			want:    nil,
			wantErr: errors.New("cannot navigate further from data"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ValueFromPath(tt.args.data, tt.args.path)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("ValueFromPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValueFromPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
