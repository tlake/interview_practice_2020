package algorithms

import "testing"

func TestBinarySearch(t *testing.T) {
	type args struct {
		s []int
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "value in right subtree",
			args: args{
				s: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				x: 8,
			},
			want: 8,
		},
		{
			name: "value in left subtree",
			args: args{
				s: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				x: 2,
			},
			want: 2,
		},
		{
			name: "value not in slice",
			args: args{
				s: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				x: 44,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearch(tt.args.s, tt.args.x); got != tt.want {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
