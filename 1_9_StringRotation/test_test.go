package exercises

import "testing"

func Test_stringRotation(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "if s2 is a rotation of s1",
			args: args{s1: "waterbottle", s2: "erbottlewat"},
			want: true,
		},
		{
			name: "if s2 is a rotation of s1: s2 empty",
			args: args{s1: "", s2: "dd"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringRotation(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("stringRotation() = %v, want %v", got, tt.want)
			}
		})
	}
}
