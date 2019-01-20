package linkedlist

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		first *Node
	}
	tests := []struct {
		name string
		args args
		want *Feed
	}{
		{
			name: "create empty",
			args: args{first: &Node{}},
			want: &Feed{0, nil, nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.first); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
