package solution

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{S: "We test coders. Give us a try?"},
			want: 4,
		},
		{
			name: "test2",
			args: args{S: "Forget  CVs..Save time . x x"},
			want: 2,
		},
		{
			name: "test3",
			args: args{S: "Fbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb! bbbb! b! bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbborget  CVs..Save time . x x"},
			want: 1,
		},
		{
			name: "test4",
			args: args{S: ""},
			want: 0,
		},
		{
			name: "test5",
			args: args{S: "!b"},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.S); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
