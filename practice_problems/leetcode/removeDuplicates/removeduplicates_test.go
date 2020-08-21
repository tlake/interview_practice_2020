package removeduplicates

import (
	"strconv"
	"strings"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		array string
	}{
		{
			name:  "[]",
			args:  args{nums: []int{}},
			want:  0,
			array: "",
		},
		{
			name:  "[1]",
			args:  args{nums: []int{1}},
			want:  1,
			array: "1",
		},
		{
			name:  "[1,1,2]",
			args:  args{nums: []int{1, 1, 2}},
			want:  2,
			array: "1, 2",
		},
		{
			name:  "[1,2,2,2,3,4,4,4,4,5,5,6,7,7,8,8,8,9]",
			args:  args{nums: []int{1, 2, 2, 2, 3, 4, 4, 4, 4, 5, 5, 6, 7, 7, 8, 8, 8, 9}},
			want:  9,
			array: "1, 2, 3, 4, 5, 6, 7, 8, 9",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeDuplicates(tt.args.nums)
			if got != tt.want {
				t.Errorf("removeDuplicates() result = %v, want %v", got, tt.want)
			}

			var stringArray []string
			for i := range tt.args.nums[:got] {
				stringArray = append(stringArray, strconv.Itoa(tt.args.nums[i]))
			}
			resultArray := strings.Join(stringArray, ", ")
			if resultArray != tt.array {
				t.Errorf("removeDuplicates() array = %v, want %v", resultArray, tt.array)
			}
		})
	}
}
