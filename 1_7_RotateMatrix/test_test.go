package mycompression

import (
	"reflect"
	"testing"
)

func Test_mycompression(t *testing.T) {
	type args struct {
		layers []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test",
			args: args{[]string{"a", "b", "c", "d"}},
			want: []string{"d", "a", "b", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMatrix := mycompression(tt.args.layers...); !reflect.DeepEqual(gotMatrix, tt.want) {
				t.Errorf("mycompression() = %v, want %v", gotMatrix, tt.want)
			}
		})
	}
}
