package crypto

import "testing"

func TestGenerateValue(t *testing.T) {
	type args struct {
		template string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Success case for alphanumeric",
			args: args{
				template: `foo[\w]{8}bar`,
			},
			wantErr: false,
		},
		{
			name: "Fail case for invalid regex",
			args: args{
				template: `pass[z-a]{8}`,
			},
			wantErr: true,
		},
		{
			name: "Fail case for invalid range",
			args: args{
				template: `pass[a-Z]{8`,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateValue(tt.args.template)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("GenerateValue() = %v", got)
		})
	}
}
