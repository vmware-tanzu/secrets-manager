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
			want: []byte(fmt.Sprintf("\"%s\"", timeNow.Format(time.RFC3339))),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := tt.tr.MarshalJSON(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonTime.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJsonTimeUnmarshalJSON(t *testing.T) {
	testCases := []struct {
		name     string
		jsonData []byte
		wantErr  bool
	}{
		{
			name:     "invalid_json_data",
			jsonData: []byte(`"invalid_date"`),
			wantErr:  true,
		},
		{
			name:     "valid_json_data",
			jsonData: []byte(`"2024-03-01T12:00:00Z"`),
			wantErr:  false,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			var jsonTime JsonTime
			err := jsonTime.UnmarshalJSON(tt.jsonData)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestJsonTimeString(t *testing.T) {
	jsonTime := JsonTime(time.Now())

	expectedTimeStr := time.Now().Format(time.RFC3339)

	result := jsonTime.String()

	if result != expectedTimeStr {
		t.Errorf("Expected: %s, Got: %s", expectedTimeStr, result)
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
			name: "error_invalid_json_value",
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
		{
			name: "empty_template",
			args: args{
				secret: SecretStored{
					Name:             "test",
					Values:           []string{"{\"username\":\"admin\",\"password\":\"VSecMRocks\"}"},
					ValueTransformed: "",
					Meta: SecretMeta{
						Template: "",
					},
					Created: timeNow,
					Updated: timeUpdated,
				},
			},
			want:    map[string]string{"username": "admin", "password": "VSecMRocks"},
			wantErr: nil,
		},
		{
			name: "error_incorrect_template",
			args: args{
				secret: SecretStored{
					Name:             "test",
					Values:           []string{"{\"username\":\"admin\",\"password\":\"VSecMRocks\"}"},
					ValueTransformed: "",
					Meta: SecretMeta{
						Template: "template-1",
					},
					Created: timeNow,
					Updated: timeUpdated,
				},
			},
			want:    map[string]string{},
			wantErr: errors.New("invalid character 'e' in literal true (expecting 'r')"),
		},
		{
			name: "error_unparsable_template",
			args: args{
				secret: SecretStored{
					Name:             "test",
					Values:           []string{"{\"username\":\"admin\",\"password-1\":\"VSecMRocks\"}"},
					ValueTransformed: "",
					Meta: SecretMeta{
						Template: "{\"USER\":\"{{.username}}\", \"PASS-1\":\"{{.password-1}}\"}",
						Format:   "yaml",
					},
					Created: timeNow,
					Updated: timeUpdated,
				},
			},
			want:    map[string]string{"username": "admin", "password-1": "VSecMRocks"},
			wantErr: errors.New("template: secret:1: bad character U+002D '-'"),
		},

		{
			name: "error_executing_template",
			args: args{
				secret: SecretStored{
					Name:             "test",
					Values:           []string{"{\"user\":\"pass\"}"},
					ValueTransformed: "",
					Meta: SecretMeta{
						Template: "{{.USER .PASS}}",
						Format:   "yaml",
					},
					Created: timeNow,
					Updated: timeUpdated,
				},
			},
			want:    map[string]string{"user": "pass"},
			wantErr: errors.New("template: secret:1:2: executing \"secret\" at <.USER>: USER is not a method but has arguments"),
		},
		{
			name: "success",
			args: args{
				secret: SecretStored{
					Name:             "test",
					Values:           []string{"{\"username\":\"admin\",\"password\":\"VSecMRocks\"}"},
					ValueTransformed: "",
					Meta: SecretMeta{
						Template: "{\"USER\":\"{{.username}}\", \"PASS\":\"{{.password}}\"}",
						Format:   "yaml",
					},
					Created: timeNow,
					Updated: timeUpdated,
				},
			},
			want:    map[string]string{"USER": "admin", "PASS": "VSecMRocks"},
			wantErr: nil,
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

func TestSecretStored_ToMapForK8s(t *testing.T) {
	type fields struct {
		Name             string
		Values           []string
		ValueTransformed string
		Meta             SecretMeta
		Created          time.Time
		Updated          time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string][]byte
	}{
		{
			name: "empty_values_list",
			fields: fields{
				Values: []string{},
			},
			want: make(map[string][]byte),
		},
		//{
		//	name: "empty_template_failed_to_unmarshal_value",
		//	fields: fields{
		//		Values: []string{"secret"},
		//	},
		//	want: map[string][]byte{"VALUE": []byte("secret")},
		//},
		{
			name: "empty_template_valid_value",
			fields: fields{
				Values: []string{"{\"username\":\"admin\",\"password\":\"VSecMRocks\"}"},
			},
			want: map[string][]byte{"username": []byte("admin"), "password": []byte("VSecMRocks")},
		},
		{
			name: "valid_value_invalid_template",
			fields: fields{
				Values: []string{"{\"pass\":\"secret\"}"},
				Meta: SecretMeta{
					Template: "template-1",
					Format:   "json",
				},
			},
			want: map[string][]byte{"VALUE": []byte("{\"pass\":\"secret\"}")},
		},
		{
			name: "success_case",
			fields: fields{
				Values: []string{"{\"pass\":\"secret\"}"},
				Meta: SecretMeta{
					Template: "{\"PASS\":\"{{.pass}}\"}",
				},
			},
			want: map[string][]byte{"PASS": []byte("secret")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			secret := SecretStored{
				Name:             tt.fields.Name,
				Values:           tt.fields.Values,
				ValueTransformed: tt.fields.ValueTransformed,
				Meta:             tt.fields.Meta,
				Created:          tt.fields.Created,
				Updated:          tt.fields.Updated,
			}
			if got := secret.ToMapForK8s(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SecretStored.ToMapForK8s() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSecretStored_ToMap(t *testing.T) {
	type fields struct {
		Name             string
		Values           []string
		ValueTransformed string
		Meta             SecretMeta
		Created          time.Time
		Updated          time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]any
	}{
		{
			name: "success_case",
			fields: fields{
				Name:    "test_name",
				Values:  []string{"test_values"},
				Created: timeNow,
				Updated: timeUpdated,
			},
			want: map[string]any{
				"Name":    "test_name",
				"Values":  []string{"test_values"},
				"Created": timeNow,
				"Updated": timeUpdated,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			secret := SecretStored{
				Name:             tt.fields.Name,
				Values:           tt.fields.Values,
				ValueTransformed: tt.fields.ValueTransformed,
				Meta:             tt.fields.Meta,
				Created:          tt.fields.Created,
				Updated:          tt.fields.Updated,
			}
			if got := secret.ToMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SecretStored.ToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

//func Test_transform(t *testing.T) {
//	type args struct {
//		secret SecretStored
//		value  string
//	}
//	tests := []struct {
//		name    string
//		args    args
//		want    string
//		wantErr bool
//		err     error
//	}{
//		{
//			name: "unsupported_format",
//			args: args{
//				secret: SecretStored{
//					Meta: SecretMeta{
//						Format: "RAML",
//					},
//				},
//			},
//			want:    "",
//			wantErr: true,
//			err:     errors.New("unknown format: RAML"),
//		},
//		{
//			name: "invalid_json_template",
//			args: args{
//				value: "{\"pass\":\"secret\"}",
//				secret: SecretStored{
//					Meta: SecretMeta{
//						Template: "\"PASS\":\"{{.pass}}\"",
//						Format:   "json",
//					},
//				},
//			},
//			want:    "{\"pass\":\"secret\"}",
//			wantErr: false,
//			err:     nil,
//		},
//		{
//			name: "valid_json_template",
//			args: args{
//				value: "{\"pass\":\"secret\"}",
//				secret: SecretStored{
//					Meta: SecretMeta{
//						Template: "{\"PASS\":\"{{.pass}}\"}",
//						Format:   "json",
//					},
//				},
//			},
//			want:    "{\"PASS\":\"secret\"}",
//			wantErr: false,
//			err:     nil,
//		},
//		{
//			name: "valid_json_template_invalid_value",
//			args: args{
//				value: "\"pass\":\"secret\"",
//				secret: SecretStored{
//					Meta: SecretMeta{
//						Template: "{\"PASS\":\"{{.pass}}\"}",
//						Format:   "json",
//					},
//				},
//			},
//			want:    "\"pass\":\"secret\"",
//			wantErr: false,
//			err:     nil,
//		},
//		{
//			name: "invalid_json_template_invalid_value",
//			args: args{
//				value: "\"pass\":\"secret\"",
//				secret: SecretStored{
//					Meta: SecretMeta{
//						Template: "\"PASS\":\"{{.pass}}\"",
//						Format:   "json",
//					},
//				},
//			},
//			want:    "\"pass\":\"secret\"",
//			wantErr: false,
//			err:     nil,
//		},
//		{
//			name: "invalid_yaml_template",
//			args: args{
//				value: "{\"pass\":\"secret\"}",
//				secret: SecretStored{
//					Meta: SecretMeta{
//						Template: "\"PASS\":\"{{.pass}}\"",
//						Format:   "yaml",
//					},
//				},
//			},
//			want:    "\"PASS\":\"secret\"",
//			wantErr: false,
//			err:     nil,
//		},
//		{
//			name: "valid_yaml_template",
//			args: args{
//				value: "{\"pass\":\"secret\"}",
//				secret: SecretStored{
//					Meta: SecretMeta{
//						Template: "\"PASS\":\"{{.pass}}\"",
//						Format:   "yaml",
//					},
//				},
//			},
//			want:    "\"PASS\":\"secret\"",
//			wantErr: false,
//			err:     nil,
//		},
//		{
//			name: "valid_values_json_passed_as_yaml",
//			args: args{
//				value: "{\"pass\":\"secret\"}",
//				secret: SecretStored{
//					Meta: SecretMeta{
//						Template: "{\"PASS\":\"{{.pass}}\"}",
//						Format:   "yaml",
//					},
//				},
//			},
//			want:    "PASS: secret\n",
//			wantErr: false,
//			err:     nil,
//		},
//		{
//			name: "valid_yaml_template_invalid_value",
//			args: args{
//				value: "\"pass\":\"secret\"",
//				secret: SecretStored{
//					Meta: SecretMeta{
//						Template: "\"PASS\":\"{{.pass}}\"",
//						Format:   "yaml",
//					},
//				},
//			},
//			want:    "\"pass\":\"secret\"",
//			wantErr: false,
//			err:     nil,
//		},
//		{
//			name: "invalid_yaml_template_invalid_value",
//			args: args{
//				value: "\"pass\":\"secret\"",
//				secret: SecretStored{
//					Meta: SecretMeta{
//						Template: "{\"PASS\":\"{{.pass}}\"}",
//						Format:   "yaml",
//					},
//				},
//			},
//			want:    "\"pass\":\"secret\"",
//			wantErr: false,
//			err:     nil,
//		},
//		{
//			name: "raw_format",
//			args: args{
//				value: "This is a raw text",
//				secret: SecretStored{
//					Meta: SecretMeta{
//						Format: "raw",
//					},
//				},
//			},
//			want:    "This is a raw text",
//			wantErr: false,
//			err:     nil,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, gotErr := transform(tt.args.secret, tt.args.value)
//			if tt.wantErr == true && tt.err.Error() != gotErr.Error() {
//				t.Errorf("transform() error = %v, wantErr %v", gotErr, tt.err)
//				return
//			}
//			if got != tt.want {
//				t.Errorf("transform() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func TestSecretStored_Parse(t *testing.T) {
	type fields struct {
		Name             string
		Values           []string
		ValueTransformed string
		Meta             SecretMeta
		Created          time.Time
		Updated          time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
		err     error
	}{
		{
			name: "empty_secret_value",
			fields: fields{
				Name:   "test",
				Values: []string{},
			},
			wantErr: true,
			err:     errors.New("no values found for secret test"),
		},
		{
			name: "valid_values",
			fields: fields{
				Name:   "test-2",
				Values: []string{"{\"pass\":\"secret\"}"},
				Meta: SecretMeta{
					Template: "\"PASS\":\"{{.pass}}\"",
					Format:   "yaml",
				},
			},
			want:    "\"PASS\":\"secret\"",
			wantErr: false,
			err:     nil,
		},
		{
			name: "valid_values_json",
			fields: fields{
				Name:   "test-2",
				Values: []string{"{\"pass\":\"secret\"}"},
				Meta: SecretMeta{
					Template: "{\"PASS\":\"{{.pass}}\"}",
					Format:   "json",
				},
			},
			want:    "{\"PASS\":\"secret\"}",
			wantErr: false,
			err:     nil,
		},
		{
			name: "multiple_valid_values",
			fields: fields{
				Name:   "test-2",
				Values: []string{"{\"pass\":\"secret\"}", "{\"pass\":\"secret-2\"}"},
				Meta: SecretMeta{
					Template: "PASS:{{.pass}}",
					Format:   "yaml",
				},
			},
			want:    "[\"PASS:secret\",\"PASS:secret-2\"]",
			wantErr: false,
			err:     nil,
		},
		{
			name: "multiple_templates",
			fields: fields{
				Name:   "test-2",
				Values: []string{"{\"pass\":\"secret\"}", "{\"pass_2\":\"secret-2\"}"},
				Meta: SecretMeta{
					Template: "PASS:{{.pass}},PASS_2:{{.pass_2}}",
					Format:   "yaml",
				},
			},
			want:    "[\"PASS:secret\",\"PASS_2:secret-2\"]",
			wantErr: false,
			err:     nil,
		},
		{
			name: "multiple_selected_template",
			fields: fields{
				Name:   "test-3",
				Values: []string{"{\"pass\":\"secret\"}", "{\"pass_2\":\"secret-2\"}", "{\"pass_3\":\"secret-3\"}"},
				Meta: SecretMeta{
					Template: "PASS:{{.pass}},PASS_3:{{.pass_3}}",
					Format:   "yaml",
				},
			},
			want:    "[\"PASS:secret\",\"PASS_3:secret-3\"]",
			wantErr: false,
			err:     nil,
		},
		{
			name: "unsupported_format",
			fields: fields{
				Name:   "test-3",
				Values: []string{"{\"pass\":\"secret\"}"},
				Meta: SecretMeta{
					Template: "\"PASS\":\"{{.pass}}\"",
					Format:   "RAML",
				},
			},
			wantErr: true,
			err:     errors.New("failed to parse secret test-3"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			secret := SecretStored{
				Name:             tt.fields.Name,
				Values:           tt.fields.Values,
				ValueTransformed: tt.fields.ValueTransformed,
				Meta:             tt.fields.Meta,
				Created:          tt.fields.Created,
				Updated:          tt.fields.Updated,
			}
			got, gotErr := secret.Parse()
			if tt.wantErr == true && gotErr.Error() != tt.err.Error() {
				t.Errorf("SecretStored.Parse() error = %v, wantErr %v", gotErr, tt.err)
				return
			}
			if got != tt.want {
				t.Errorf("SecretStored.Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
