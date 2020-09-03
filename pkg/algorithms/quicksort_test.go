package algorithms

import (
	"reflect"
	"testing"
)

func TestQuicksort(t *testing.T) {
	type args struct {
		inputSlice []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "[6,2,1,9,7,4,0,3,8,5]",
			args: args{inputSlice: []int{6, 2, 1, 9, 7, 4, 0, 3, 8, 5}},
			want: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "[3,1,6,7,3,2]",
			args: args{inputSlice: []int{3, 1, 6, 7, 3, 2}},
			want: []int{1, 2, 3, 3, 6, 7},
		},
		{
			name: "[6,3,3,6,6,3,3,3,6,3,3,3,6,6,6,3,6,3,6,3,3]",
			args: args{inputSlice: []int{6, 3, 3, 6, 6, 3, 3, 3, 6, 3, 3, 3, 6, 6, 6, 3, 6, 3, 6, 3, 3}},
			want: []int{3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Quicksort(tt.args.inputSlice)
			if !reflect.DeepEqual(tt.args.inputSlice, tt.want) {
				t.Errorf("Quicksort() = %v, want = %v", tt.args.inputSlice, tt.want)
			}
		})
	}
}

func Test_partition(t *testing.T) {
	type args struct {
		s []int
		p int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_swapAtIndices(t *testing.T) {
	type args struct {
		s []int
		a int
		b int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			swapAtIndices(tt.args.s, tt.args.a, tt.args.b)
		})
	}
}
