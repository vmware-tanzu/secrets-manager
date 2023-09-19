/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware, Inc.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package v1

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"
)

var (
	timeNow     = time.Now()
	timeUpdated = time.Now().Add(5 * time.Minute)
)

func TestJsonTime_MarshalJSON(t *testing.T) {
	tests := []struct {
		name string
		tr   JsonTime
		want []byte
	}{
		{
			name: "success_case",
			tr:   JsonTime(timeNow),
			want: []byte(fmt.Sprintf("\"%s\"", time.Time(timeNow).Format(time.RubyDate))),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.MarshalJSON(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonTime.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseForK8sSecret(t *testing.T) {
	type args struct {
		secret SecretStored
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]string
		wantErr error
	}{
		{
			name: "error_empty_value",
			args: args{
				secret: SecretStored{
					Name:             "test",
					Values:           []string{},
					ValueTransformed: "",
					Meta:             SecretMeta{},
					Created:          timeNow,
					Updated:          timeUpdated,
				},
			},
			want:    map[string]string{},
			wantErr: errors.New("no values found for secret test"),
		},
		{
			name: "error_empty_template",
			args: args{
				secret: SecretStored{
					Name:             "test",
					Values:           []string{"value-1"},
					ValueTransformed: "",
					Meta: SecretMeta{
						Template: "",
					},
					Created: timeNow,
					Updated: timeUpdated,
				},
			},
			want:    map[string]string{},
			wantErr: errors.New("invalid character 'v' looking for beginning of value"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseForK8sSecret(tt.args.secret)
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("parseForK8sSecret() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseForK8sSecret() = %v, want %v", got, tt.want)
			}
		})
	}
}
