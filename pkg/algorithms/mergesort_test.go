package algorithms

import (
	"reflect"
	"testing"
)

func TestMergesort(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "[6,2,1,9,7,4,0,3,8,5]",
			args: args{s: []int{6, 2, 1, 9, 7, 4, 0, 3, 8, 5}},
			want: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "[3,1,6,7,3,2]",
			args: args{s: []int{3, 1, 6, 7, 3, 2}},
			want: []int{1, 2, 3, 3, 6, 7},
		},
		{
			name: "[6,3,3,6,6,3,3,3,6,3,3,3,6,6,6,3,6,3,6,3,3]",
			args: args{s: []int{6, 3, 3, 6, 6, 3, 3, 3, 6, 3, 3, 3, 6, 6, 6, 3, 6, 3, 6, 3, 3}},
			want: []int{3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Mergesort(tt.args.s)
			if !reflect.DeepEqual(tt.args.s, tt.want) {
				t.Errorf("Mergesort() = %v, want = %v", tt.args.s, tt.want)
			}
		})
	}
}
