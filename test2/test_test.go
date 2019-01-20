package solution

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		A int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{A: 130},
			want: 103,
		},
		{
			name: "test2",
			args: args{A: 123456},
			want: 162534,
		},
		{
			name: "test2",
			args: args{A: 0},
			want: 0,
		},
		{
			name: "test2",
			args: args{A: 100000000},
			want: 100000000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.A); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
