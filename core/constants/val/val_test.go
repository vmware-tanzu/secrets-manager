package val

import "testing"

func TestNever(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "never",
			args: args{
				s: "never",
			},
			want: true,
		},
		{
			name: "Never",
			args: args{
				s: "Never",
			},
			want: true,
		},
		{
			name: "never with space",
			args: args{
				s: " never ",
			},
			want: true,
		},
		{
			name: "never with space and caps",
			args: args{
				s: " NeVeR ",
			},
			want: true,
		},
		{
			name: "not never",
			args: args{
				s: "not never",
			},
			want: false,
		},
		{
			name: "empty",
			args: args{
				s: "",
			},
			want: false,
		},
		{
			name: "space",
			args: args{
				s: " ",
			},
			want: false,
		},
		{
			name: "never with space",
			args: args{
				s: " never",
			},
			want: true,
		},
		{
			name: "never with space",
			args: args{
				s: "never ",
			},
			want: true,
		},
		{
			name: "never with space",
			args: args{
				s: " never ",
			},
			want: true,
		},
		{
			name: "never with space",
			args: args{
				s: " NeVeR ",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Never(tt.args.s); got != tt.want {
				t.Errorf("Never() = %v, want %v", got, tt.want)
			}
		})
	}
}
