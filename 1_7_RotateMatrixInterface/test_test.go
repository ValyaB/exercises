package exercises

import (
	"reflect"
	"testing"
)

func Test_rotateRow(t *testing.T) {
	type args struct {
		layers []interface{}
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		{
			name: "nil",
			args: args{},
			want: nil,
		},
		{
			name: "test on array of INTs",
			args: args{layers: []interface{}{1, 2, 3, 4, 5, 6, 7}},
			want: []interface{}{7, 1, 2, 3, 4, 5, 6},
		},
		{
			name: "test on array of Strings",
			args: args{layers: []interface{}{"1", "2", "3", "4"}},
			want: []interface{}{"4", "1", "2", "3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateRow(tt.args.layers...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotateRow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transposeMatrix(t *testing.T) {
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
			args: args{matrix: [][]interface{}{{1, 2, 3, 4}, {1, 2, 3, 4}, {1, 2, 3, 4}, {1, 2, 3, 4}}},
			want: [][]interface{}{{1, 1, 1, 1}, {2, 2, 2, 2}, {3, 3, 3, 3}, {4, 4, 4, 4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := transposeMatrix(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("transposeMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_verticalReflactionMatrix(t *testing.T) {
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
			args: args{matrix: [][]interface{}{{1, 2, 3, 4}, {1, 2, 3, 4}, {1, 2, 3, 4}, {1, 2, 3, 4}}},
			want: [][]interface{}{{4, 3, 2, 1}, {4, 3, 2, 1}, {4, 3, 2, 1}, {4, 3, 2, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := verticalReflactionMatrix(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("verticalReflactionMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_horizontalReflactionMatrix(t *testing.T) {
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
			args: args{matrix: [][]interface{}{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}},
			want: [][]interface{}{{13, 14, 15, 16}, {9, 10, 11, 12}, {5, 6, 7, 8}, {1, 2, 3, 4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := horizontalReflactionMatrix(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("horizontalReflactionMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rotateMatrix90(t *testing.T) {
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
			args: args{matrix: [][]interface{}{{1, 6, 11, 16, 21}, {2, 7, 12, 17, 22}, {3, 8, 13, 18, 23}, {4, 9, 14, 19, 24}, {5, 10, 15, 20, 25}}},
			want: [][]interface{}{{21, 22, 23, 24, 25}, {16, 17, 18, 19, 20}, {11, 12, 13, 14, 15}, {6, 7, 8, 9, 10}, {1, 2, 3, 4, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateMatrix90(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotateMatrix90() = %v, want %v", got, tt.want)
			}
		})
	}
}
