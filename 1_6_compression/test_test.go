package mycompression

import "testing"

func Test_mycompression(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "same",
			args: args{s: "same"},
			want: "s1a1m1e1",
		},
		{
			name: "compression",
			args: args{s: "ssess "},
			want: "s2e1s2 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mycompression(tt.args.s); got != tt.want {
				t.Errorf("mycompression() = %v, want %v", got, tt.want)
			}
		})
	}
}
