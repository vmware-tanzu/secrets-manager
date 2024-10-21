package json_test

import (
	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/server/route/base/json"
	"testing"
)

func TestUnmarshalSecretUpsertRequest(t *testing.T) {
	tests := []struct {
		name    string
		body    []byte
		wantErr bool
	}{
		{
			name:    "Valid SecretUpsertRequest",
			body:    []byte(`{"FieldName": "some_value"}`), // Replace FieldName with actual fields
			wantErr: false,
		},
		{
			name:    "Invalid JSON",
			body:    []byte(`{"FieldName": "some_value"`), // Invalid JSON
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := json.UnmarshalSecretUpsertRequest(tt.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalSecretUpsertRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got == nil {
				t.Errorf("UnmarshalSecretUpsertRequest() got = nil, expected non-nil result")
			}
		})
	}
}

func TestUnmarshalKeyInputRequest(t *testing.T) {
	tests := []struct {
		name    string
		body    []byte
		wantErr bool
	}{
		{
			name:    "Valid KeyInputRequest",
			body:    []byte(`{"KeyFieldName": "key_value"}`), // Replace KeyFieldName with actual fields
			wantErr: false,
		},
		{
			name:    "Invalid JSON",
			body:    []byte(`{"KeyFieldName": "key_value"`), // Invalid JSON
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := json.UnmarshalKeyInputRequest(tt.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalKeyInputRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got == nil {
				t.Errorf("UnmarshalKeyInputRequest() got = nil, expected non-nil result")
			}
		})
	}
}
