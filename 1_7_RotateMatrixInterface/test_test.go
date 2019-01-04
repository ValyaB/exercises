package mycompression

import (
	"reflect"
	"testing"
)

func Test_mycompression(t *testing.T) {
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
			args: args{layers: []interface{}{1, 2, 3, 4}},
			want: []interface{}{4, 1, 2, 3},
		},
		{
			name: "test on array of Strings",
			args: args{layers: []interface{}{"1", "2", "3", "4"}},
			want: []interface{}{"4", "1", "2", "3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mycompression(tt.args.layers...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mycompression() = %v, want %v", got, tt.want)
			}
		})
	}
}
