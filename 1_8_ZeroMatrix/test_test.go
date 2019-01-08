package exercises

import (
	"reflect"
	"testing"
)

func Test_zeroMatrix(t *testing.T) {
	type args struct {
		matrix [][]interface{}
	}
	tests := []struct {
		name string
		args args
		want [][]interface{}
	}{
		{
			name: "test on 2D array of INTs",
			args: args{matrix: [][]interface{}{{1, 2, 3, 4}, {5, 0, 7, 8}, {9, 10, 11, 12}}},
			want: [][]interface{}{{1, 0, 3, 4}, {0, 0, 0, 0}, {9, 0, 11, 12}},
		},
		{
			name: "test on 2D array of INTs",
			args: args{matrix: [][]interface{}{{1, 2, 3, 0}, {5, 0, 7, 8}, {9, 10, 11, 12}}},
			want: [][]interface{}{{0, 0, 0, 0}, {0, 0, 0, 0}, {9, 0, 11, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := zeroMatrix(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("zeroMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
