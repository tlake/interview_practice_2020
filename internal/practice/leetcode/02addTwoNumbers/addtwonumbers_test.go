package addtwonumbers

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "SimpleCase",
			// [1, 2, 3] [2, 3, 4] -> [3, 5, 7]
			args: args{
				l1: &ListNode{
					Val: 1, Next: &ListNode{
						Val: 2, Next: &ListNode{
							Val: 3,
						},
					},
				},
				l2: &ListNode{
					Val: 2, Next: &ListNode{
						Val: 3, Next: &ListNode{
							Val: 4,
						},
					},
				},
			},
			want: &ListNode{
				Val: 3, Next: &ListNode{
					Val: 5, Next: &ListNode{
						Val: 7,
					},
				},
			},
		},
		{
			name: "CarryOver",
			// [3, 5, 4] [4, 8, 3] -> [7, 3, 8]
			args: args{
				l1: &ListNode{
					Val: 3, Next: &ListNode{
						Val: 5, Next: &ListNode{
							Val: 4,
						},
					},
				},
				l2: &ListNode{
					Val: 4, Next: &ListNode{
						Val: 8, Next: &ListNode{
							Val: 3,
						},
					},
				},
			},
			want: &ListNode{
				Val: 7, Next: &ListNode{
					Val: 3, Next: &ListNode{
						Val: 8,
					},
				},
			},
		},
		{
			name: "AsymmetricalInput",
			// [5, 2] [0] -> [5, 2]
			args: args{
				l1: &ListNode{
					Val: 5, Next: &ListNode{
						Val: 2,
					},
				},
				l2: &ListNode{Val: 0},
			},
			want: &ListNode{
				Val: 5, Next: &ListNode{
					Val: 2,
				},
			},
		},
		{
			name: "EndStageCarryOver",
			// [5] [5] -> [0, 1]
			args: args{
				l1: &ListNode{Val: 5},
				l2: &ListNode{Val: 5},
			},
			want: &ListNode{
				Val: 0, Next: &ListNode{
					Val: 1,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
