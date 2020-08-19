// Package linkedlist implements a Linked List.
package linkedlist

import (
	"reflect"
	"testing"
)

func buildLinkedList(nodeValues []string) LinkedList {
	l := LinkedList{
		length: len(nodeValues),
	}

	var lastNode *LLNode

	if len(nodeValues) > 0 {
		lastNode = &LLNode{value: nodeValues[0]}
		l.head = lastNode
		l.tail = lastNode
	}

	for i := 1; i < len(nodeValues); i++ {
		newNode := &LLNode{value: nodeValues[i]}
		lastNode.next = newNode
		l.tail = newNode
		lastNode = newNode
	}

	return l
}

func TestLinkedList_Display(t *testing.T) {
	tests := []struct {
		name       string
		nodeValues []string
		want       string
	}{
		{
			name:       "ZeroNodes",
			nodeValues: []string{},
			want:       "",
		},
		{
			name:       "OneNode",
			nodeValues: []string{"d"},
			want:       "d",
		},
		{
			name:       "TwoNodes",
			nodeValues: []string{"a", "b"},
			want:       "a, b",
		},
		{
			name:       "TenNodes",
			nodeValues: []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"},
			want:       "a, b, c, d, e, f, g, h, i, j",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := buildLinkedList(tt.nodeValues)

			if got := l.Display(); got != tt.want {
				t.Errorf("LinkedList.Display() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_ValueAt(t *testing.T) {
	type args struct {
		index int
	}

	tests := []struct {
		name       string
		nodeValues []string
		args       args
		want       string
		wantErr    bool
	}{
		{
			name:       "IndexOutOfBounds",
			nodeValues: []string{"a"},
			args:       args{index: 2},
			want:       "",
			wantErr:    true,
		},
		{
			name:       "FirstNode",
			nodeValues: []string{"a", "b", "c", "d", "e"},
			args:       args{index: 0},
			want:       "a",
			wantErr:    false,
		},
		{
			name:       "LastNode",
			nodeValues: []string{"a", "b", "c", "d", "e"},
			args:       args{index: 4},
			want:       "e",
			wantErr:    false,
		},
		{
			name:       "InternalNode",
			nodeValues: []string{"a", "b", "c", "d", "e"},
			args:       args{index: 2},
			want:       "c",
			wantErr:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := buildLinkedList(tt.nodeValues)
			got, err := l.ValueAt(tt.args.index)

			if (err != nil) != tt.wantErr {
				t.Errorf("LinkedList.ValueAt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("LinkedList.ValueAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_Search(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name       string
		nodeValues []string
		args       args
		want       *LLNode
		wantErr    bool
	}{
		{
			name:       "ValueNotFound",
			nodeValues: []string{"a", "b", "c", "d", "e"},
			args:       args{"f"},
			want:       nil,
			wantErr:    true,
		},
		{
			name:       "ValueFoundAtStart",
			nodeValues: []string{"a", "b", "c", "d", "e"},
			args:       args{"a"},
			want:       &LLNode{value: "a"},
			wantErr:    false,
		},
		{
			name:       "ValueFoundAtEnd",
			nodeValues: []string{"a", "b", "c", "d", "e"},
			args:       args{"e"},
			want:       &LLNode{value: "e"},
			wantErr:    false,
		},
		{
			name:       "ValueFoundInMiddle",
			nodeValues: []string{"a", "b", "c", "d", "e"},
			args:       args{"c"},
			want:       &LLNode{value: "c"},
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := buildLinkedList(tt.nodeValues)
			got, err := l.Search(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("LinkedList.Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want == nil {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("LinkedList.Search() = %v, want %v", got, tt.want)
				}
			} else {
				if !reflect.DeepEqual(got.value, tt.want.value) {
					t.Errorf("LinkedList.Search() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestLinkedList_Prepend(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name            string
		nodeValues      []string
		args            args
		expectedDisplay string
		expectedLength  int
		expectedHead    *LLNode
		expectedTail    *LLNode
	}{
		{
			name:            "PrependToEmpty",
			nodeValues:      []string{},
			args:            args{"a"},
			expectedDisplay: "a",
			expectedLength:  1,
			expectedHead:    &LLNode{value: "a"},
			expectedTail:    &LLNode{value: "a"},
		},
		{
			name:            "PrependToPopulated",
			nodeValues:      []string{"a", "b", "c", "d", "e"},
			args:            args{"f"},
			expectedDisplay: "f, a, b, c, d, e",
			expectedLength:  6,
			expectedHead:    &LLNode{value: "f"},
			expectedTail:    &LLNode{value: "e"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := buildLinkedList(tt.nodeValues)
			// Because the nodes in a linked list aren't indexed, it's difficult to programmatically
			// ensure that the "next" property points to the proper address in memory throughout
			// these table tests. It should be good enough to observe whether the desired node
			// contains a specific value as intended.
			if got := l.Prepend(tt.args.value); got.value != tt.args.value {
				t.Errorf("LinkedList.Prepend() = %s, want = %s", got.value, tt.args.value)
			}
			// Is the length property updated properly?
			if l.length != tt.expectedLength {
				t.Errorf("LinkedList.Append() = [%d], want = [%d]", l.length, tt.expectedLength)
			}
			// Check the internals of the list by comparing the string representations
			if d := l.Display(); d != tt.expectedDisplay {
				t.Errorf("LinkedList.Append() = [%s], want = [%s]", d, tt.expectedDisplay)
			}
			// Check the list's head and tail properties
			if l.head.value != tt.expectedHead.value {
				t.Errorf("LinkedList.Append() head = %v, want = %v", l.head.value, tt.expectedHead.value)
			}
			if l.tail.value != tt.expectedTail.value {
				t.Errorf("LinkedList.Append() tail = %v, want = %v", l.tail.value, tt.expectedTail.value)
			}
		})
	}
}

func TestLinkedList_Append(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name            string
		nodeValues      []string
		args            args
		expectedDisplay string
		expectedLength  int
		expectedHead    *LLNode
		expectedTail    *LLNode
	}{
		{
			name:            "AppendToEmpty",
			nodeValues:      []string{},
			args:            args{"a"},
			expectedDisplay: "a",
			expectedLength:  1,
			expectedHead:    &LLNode{value: "a"},
			expectedTail:    &LLNode{value: "a"},
		},
		{
			name:            "AppendToPopulated",
			nodeValues:      []string{"a", "b", "c", "d"},
			args:            args{"e"},
			expectedDisplay: "a, b, c, d, e",
			expectedLength:  5,
			expectedHead:    &LLNode{value: "a"},
			expectedTail:    &LLNode{value: "e"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := buildLinkedList(tt.nodeValues)
			// Because the nodes in a linked list aren't indexed, it's difficult to programmatically
			// ensure that the "next" property points to the proper address in memory throughout
			// these table tests. It should be good enough to observe whether the desired node
			// contains a specific value as intended.
			if got := l.Append(tt.args.value); got.value != tt.args.value {
				t.Errorf("LinkedList.Append() = %s, want = %s", got.value, tt.args.value)
			}
			// Is the length property updated properly?
			if l.length != tt.expectedLength {
				t.Errorf("LinkedList.Append() = [%d], want = [%d]", l.length, tt.expectedLength)
			}
			// Check the internals of the list by comparing the string representations
			if d := l.Display(); d != tt.expectedDisplay {
				t.Errorf("LinkedList.Append() = [%s], want = [%s]", d, tt.expectedDisplay)
			}
			// Check the list's head and tail properties
			if l.head.value != tt.expectedHead.value {
				t.Errorf("LinkedList.Append() head = %v, want = %v", l.head.value, tt.expectedHead.value)
			}
			if l.tail.value != tt.expectedTail.value {
				t.Errorf("LinkedList.Append() tail = %v, want = %v", l.tail.value, tt.expectedTail.value)
			}
		})
	}
}

func TestLinkedList_Insert(t *testing.T) {
	type args struct {
		value string
		index int
	}
	tests := []struct {
		name            string
		nodeValues      []string
		args            args
		want            *LLNode
		wantErr         bool
		expectedDisplay string
		expectedLength  int
		expectedHead    *LLNode
		expectedTail    *LLNode
	}{
		{
			name:            "InsertOutOfBounds",
			nodeValues:      []string{"a", "b"},
			args:            args{value: "8", index: 3},
			want:            nil,
			wantErr:         true,
			expectedDisplay: "a, b",
			expectedLength:  2,
			expectedHead:    &LLNode{value: "a"},
			expectedTail:    &LLNode{value: "b"},
		},
		{
			name:            "InsertToEmpty",
			nodeValues:      []string{},
			args:            args{value: "a", index: 0},
			want:            &LLNode{value: "a"},
			wantErr:         false,
			expectedDisplay: "a",
			expectedLength:  1,
			expectedHead:    &LLNode{value: "a"},
			expectedTail:    &LLNode{value: "a"},
		},
		{
			name:            "InsertAtFrontIntoSingleItem",
			nodeValues:      []string{"a"},
			args:            args{value: "b", index: 0},
			want:            &LLNode{value: "b"},
			wantErr:         false,
			expectedDisplay: "b, a",
			expectedLength:  2,
			expectedHead:    &LLNode{value: "b"},
			expectedTail:    &LLNode{value: "a"},
		},
		{
			name:            "InsertAtEndIntoSingleItem",
			nodeValues:      []string{"a"},
			args:            args{value: "b", index: 1},
			want:            &LLNode{value: "b"},
			wantErr:         false,
			expectedDisplay: "a, b",
			expectedLength:  2,
			expectedHead:    &LLNode{value: "a"},
			expectedTail:    &LLNode{value: "b"},
		},
		{
			name:            "InsertAtFrontIntoTwoItems",
			nodeValues:      []string{"a", "b"},
			args:            args{value: "c", index: 0},
			want:            &LLNode{value: "c"},
			wantErr:         false,
			expectedDisplay: "c, a, b",
			expectedLength:  3,
			expectedHead:    &LLNode{value: "c"},
			expectedTail:    &LLNode{value: "b"},
		},
		{
			name:            "InsertAtEndIntoTwoItems",
			nodeValues:      []string{"a", "b"},
			args:            args{value: "c", index: 2},
			want:            &LLNode{value: "c"},
			wantErr:         false,
			expectedDisplay: "a, b, c",
			expectedLength:  3,
			expectedHead:    &LLNode{value: "a"},
			expectedTail:    &LLNode{value: "c"},
		},
		{
			name:            "InsertAtMiddleIntoTwoItems",
			nodeValues:      []string{"a", "b"},
			args:            args{value: "c", index: 1},
			want:            &LLNode{value: "c"},
			wantErr:         false,
			expectedDisplay: "a, c, b",
			expectedLength:  3,
			expectedHead:    &LLNode{value: "a"},
			expectedTail:    &LLNode{value: "b"},
		},
		{
			name:            "InsertAtMiddleIntoPopulated",
			nodeValues:      []string{"a", "b", "c", "d", "e"},
			args:            args{value: "f", index: 3},
			want:            &LLNode{value: "f"},
			wantErr:         false,
			expectedDisplay: "a, b, c, f, d, e",
			expectedLength:  6,
			expectedHead:    &LLNode{value: "a"},
			expectedTail:    &LLNode{value: "e"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := buildLinkedList(tt.nodeValues)
			got, err := l.Insert(tt.args.value, tt.args.index)
			// Because the nodes in a linked list aren't indexed, it's difficult to programmatically
			// ensure that the "next" property points to the proper address in memory throughout
			// these table tests. It should be good enough to observe whether the desired node is
			// either nil as intended (the case when the method errors), or that it contains a
			// specific value as intended (the case when the method is successful).
			if tt.want == nil {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("LinkedList.Insert() got = %v, want %v", got, tt.want)
				}
			} else {
				if !reflect.DeepEqual(got.value, tt.want.value) {
					t.Errorf("LinkedList.Insert() got = %v, want %v", got, tt.want)
				}
			}
			// Do we get the error that we expect?
			if (err != nil) != tt.wantErr {
				t.Errorf("LinkedList.Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
			// Is the length property updated properly?
			if l.length != tt.expectedLength {
				t.Errorf("LinkedList.Insert() length = [%d], want = [%d]", l.length, tt.expectedLength)
			}
			// Check the internals of the list by comparing the string representations
			if d := l.Display(); d != tt.expectedDisplay {
				t.Errorf("LinkedList.Insert() display = [%s], want = [%s]", d, tt.expectedDisplay)
			}
			// Check the list's head and tail properties
			if l.head.value != tt.expectedHead.value {
				t.Errorf("LinkedList.Insert() LinkedList head = %s, want = %s", l.head.value, tt.expectedHead.value)
			}
			if l.tail.value != tt.expectedTail.value {
				t.Errorf("LinkedList.Insert() LinkedList tail = %s, want = %s", l.tail.value, tt.expectedTail.value)
			}
		})
	}
}

func TestLinkedList_Shift(t *testing.T) {
	tests := []struct {
		name            string
		nodeValues      []string
		want            *LLNode
		expectedDisplay string
		expectedLength  int
		expectedHead    *LLNode
		expectedTail    *LLNode
	}{
		{
			name:            "ShiftEmptyList",
			nodeValues:      []string{},
			want:            nil,
			expectedDisplay: "",
			expectedLength:  0,
			expectedHead:    nil,
			expectedTail:    nil,
		},
		{
			name:            "ShiftOneItemList",
			nodeValues:      []string{"a"},
			want:            &LLNode{value: "a"},
			expectedDisplay: "",
			expectedLength:  0,
			expectedHead:    nil,
			expectedTail:    nil,
		},
		{
			name:            "ShiftTwoItemList",
			nodeValues:      []string{"a", "b"},
			want:            &LLNode{value: "a"},
			expectedDisplay: "b",
			expectedLength:  1,
			expectedHead:    &LLNode{value: "b"},
			expectedTail:    &LLNode{value: "b"},
		},
		{
			name:            "ShiftThreeItemList",
			nodeValues:      []string{"a", "b", "c"},
			want:            &LLNode{value: "a"},
			expectedDisplay: "b, c",
			expectedLength:  2,
			expectedHead:    &LLNode{value: "b"},
			expectedTail:    &LLNode{value: "c"},
		},
		{
			name:            "ShiftPopulatedList",
			nodeValues:      []string{"a", "b", "c", "d", "e", "f"},
			want:            &LLNode{value: "a"},
			expectedDisplay: "b, c, d, e, f",
			expectedLength:  5,
			expectedHead:    &LLNode{value: "b"},
			expectedTail:    &LLNode{value: "f"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := buildLinkedList(tt.nodeValues)
			got := l.Shift()
			// Because the nodes in a linked list aren't indexed, it's difficult to programmatically
			// ensure that the "next" property points to the proper address in memory throughout
			// these table tests. It should be good enough to observe whether the desired node is
			// either nil as intended (the case when the list is empty), or that it contains a
			// specific value as intended (the case when the method is otherwise successful).
			if tt.want == nil {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("LinkedList.Shift() got = %v, want %v", got, tt.want)
				}
			} else {
				if got.value != tt.want.value {
					t.Errorf("LinkedList.Shift() got = %v, want %v", got, tt.want)
				}
			}
			// Is the length property updated properly?
			if l.length != tt.expectedLength {
				t.Errorf("LinkedList.Shift() length = [%d], want = [%d]", l.length, tt.expectedLength)
			}
			// Check the internals of the list by comparing the string representations
			if d := l.Display(); d != tt.expectedDisplay {
				t.Errorf("LinkedList.Shift() display = [%s], want = [%s]", d, tt.expectedDisplay)
			}
			// Check the list's head and tail properties; again, node struct comparisons make
			// inspection of an internal node difficult outside of the list.
			if tt.expectedHead == nil {
				if l.head != nil {
					t.Errorf("LinkedList.Shift() head = %v, want = %v", l.head, tt.expectedHead)
				}
			} else {
				if l.head.value != tt.expectedHead.value {
					t.Errorf("LinkedList.Shift() head = %v, want = %v", l.head.value, tt.expectedHead.value)
				}
			}
			if tt.expectedTail == nil {
				if l.tail != nil {
					t.Errorf("LinkedList.Shift() tail = %v, want = %v", l.tail, tt.expectedTail)
				}
			} else {
				if l.tail.value != tt.expectedTail.value {
					t.Errorf("LinkedList.Shift() tail = %v, want = %v", l.tail.value, tt.expectedTail.value)
				}
			}
		})
	}
}

func TestLinkedList_Pop(t *testing.T) {
	tests := []struct {
		name            string
		nodeValues      []string
		want            *LLNode
		expectedDisplay string
		expectedLength  int
		expectedHead    *LLNode
		expectedTail    *LLNode
	}{
		{
			name:            "PopEmptyList",
			nodeValues:      []string{},
			want:            nil,
			expectedDisplay: "",
			expectedLength:  0,
			expectedHead:    nil,
			expectedTail:    nil,
		},
		{
			name:            "PopOneItemList",
			nodeValues:      []string{"a"},
			want:            &LLNode{value: "a"},
			expectedDisplay: "",
			expectedLength:  0,
			expectedHead:    nil,
			expectedTail:    nil,
		},
		{
			name:            "PopTwoItemList",
			nodeValues:      []string{"a", "b"},
			want:            &LLNode{value: "b"},
			expectedDisplay: "a",
			expectedLength:  1,
			expectedHead:    &LLNode{value: "a"},
			expectedTail:    &LLNode{value: "a"},
		},
		{
			name:            "PopThreeItemList",
			nodeValues:      []string{"a", "b", "c"},
			want:            &LLNode{value: "c"},
			expectedDisplay: "a, b",
			expectedLength:  2,
			expectedHead:    &LLNode{value: "a"},
			expectedTail:    &LLNode{value: "b"},
		},
		{
			name:            "PopPopulatedList",
			nodeValues:      []string{"a", "b", "c", "d", "e", "f"},
			want:            &LLNode{value: "f"},
			expectedDisplay: "a, b, c, d, e",
			expectedLength:  5,
			expectedHead:    &LLNode{value: "a"},
			expectedTail:    &LLNode{value: "e"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := buildLinkedList(tt.nodeValues)
			got := l.Pop()
			// Because the nodes in a linked list aren't indexed, it's difficult to programmatically
			// ensure that the "next" property points to the proper address in memory throughout
			// these table tests. It should be good enough to observe whether the desired node is
			// either nil as intended (the case when the list is empty), or that it contains a
			// specific value as intended (the case when the method is otherwise successful).
			if tt.want == nil {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("LinkedList.Pop() got = %v, want %v", got, tt.want)
				}
			} else {
				if got.value != tt.want.value {
					t.Errorf("LinkedList.Pop() got = %v, want %v", got, tt.want)
				}
			}
			// Is the length property updated properly?
			if l.length != tt.expectedLength {
				t.Errorf("LinkedList.Pop() length = [%d], want = [%d]", l.length, tt.expectedLength)
			}
			// Check the internals of the list by comparing the string representations
			if d := l.Display(); d != tt.expectedDisplay {
				t.Errorf("LinkedList.Pop() display = [%s], want = [%s]", d, tt.expectedDisplay)
			}
			// Check the list's head and tail properties; again, node struct comparisons make
			// inspection of an internal node difficult outside of the list.
			if tt.expectedHead == nil {
				if l.head != nil {
					t.Errorf("LinkedList.Pop() head = %v, want = %v", l.head, tt.expectedHead)
				}
			} else {
				if l.head.value != tt.expectedHead.value {
					t.Errorf("LinkedList.Pop() head = %v, want = %v", l.head.value, tt.expectedHead.value)
				}
			}
			if tt.expectedTail == nil {
				if l.tail != nil {
					t.Errorf("LinkedList.Pop() tail = %v, want = %v", l.tail, tt.expectedTail)
				}
			} else {
				if l.tail.value != tt.expectedTail.value {
					t.Errorf("LinkedList.Pop() tail = %v, want = %v", l.tail.value, tt.expectedTail.value)
				}
			}
		})
	}
}

func TestLinkedList_Remove(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name            string
		nodeValues      []string
		args            args
		want            *LLNode
		wantErr         bool
		expectedDisplay string
		expectedLength  int
		expectedHead    *LLNode
		expectedTail    *LLNode
	}{
		{
			name:            "RemoveOutOfBounds",
			nodeValues:      []string{"a", "b"},
			args:            args{index: 2},
			want:            nil,
			wantErr:         true,
			expectedDisplay: "a, b",
			expectedLength:  2,
			expectedHead:    &LLNode{value: "a"},
			expectedTail:    &LLNode{value: "b"},
		},
		{
			name:            "RemoveFromEmpty",
			nodeValues:      []string{},
			args:            args{index: 0},
			want:            nil,
			wantErr:         true,
			expectedDisplay: "",
			expectedLength:  0,
			expectedHead:    nil,
			expectedTail:    nil,
		},
		{
			name:            "RemoveFromOneItem",
			nodeValues:      []string{"a"},
			args:            args{index: 0},
			want:            &LLNode{value: "a"},
			wantErr:         false,
			expectedDisplay: "",
			expectedLength:  0,
			expectedHead:    nil,
			expectedTail:    nil,
		},
		{
			name:            "RemoveAtFrontFromTwoItems",
			nodeValues:      []string{"a", "b"},
			args:            args{index: 0},
			want:            &LLNode{value: "a"},
			wantErr:         false,
			expectedDisplay: "b",
			expectedLength:  1,
			expectedHead:    &LLNode{value: "b"},
			expectedTail:    &LLNode{value: "b"},
		},
		{
			name:            "RemoveAtBackFromTwoItems",
			nodeValues:      []string{"a", "b"},
			args:            args{index: 1},
			want:            &LLNode{value: "b"},
			wantErr:         false,
			expectedDisplay: "a",
			expectedLength:  1,
			expectedHead:    &LLNode{value: "a"},
			expectedTail:    &LLNode{value: "a"},
		},
		{
			name:            "RemoveAtFrontFromThreeItems",
			nodeValues:      []string{"a", "b", "c"},
			args:            args{index: 0},
			want:            &LLNode{value: "a"},
			wantErr:         false,
			expectedDisplay: "b, c",
			expectedLength:  2,
			expectedHead:    &LLNode{value: "b"},
			expectedTail:    &LLNode{value: "c"},
		},
		{
			name:            "RemoveAtBackFromThreeItems",
			nodeValues:      []string{"a", "b", "c"},
			args:            args{index: 2},
			want:            &LLNode{value: "c"},
			wantErr:         false,
			expectedDisplay: "a, b",
			expectedLength:  2,
			expectedHead:    &LLNode{value: "a"},
			expectedTail:    &LLNode{value: "b"},
		},
		{
			name:            "RemoveAtMiddleFromThreeItems",
			nodeValues:      []string{"a", "b", "c"},
			args:            args{index: 1},
			want:            &LLNode{value: "b"},
			wantErr:         false,
			expectedDisplay: "a, c",
			expectedLength:  2,
			expectedHead:    &LLNode{value: "a"},
			expectedTail:    &LLNode{value: "c"},
		},
		{
			name:            "RemoveAtFrontFromPopulated",
			nodeValues:      []string{"a", "b", "c", "d", "e", "f"},
			args:            args{index: 0},
			want:            &LLNode{value: "a"},
			wantErr:         false,
			expectedDisplay: "b, c, d, e, f",
			expectedLength:  5,
			expectedHead:    &LLNode{value: "b"},
			expectedTail:    &LLNode{value: "f"},
		},
		{
			name:            "RemoveAtBackFromPopulated",
			nodeValues:      []string{"a", "b", "c", "d", "e", "f"},
			args:            args{index: 5},
			want:            &LLNode{value: "f"},
			wantErr:         false,
			expectedDisplay: "a, b, c, d, e",
			expectedLength:  5,
			expectedHead:    &LLNode{value: "a"},
			expectedTail:    &LLNode{value: "e"},
		},
		{
			name:            "RemoveAtMiddleFromPopulated",
			nodeValues:      []string{"a", "b", "c", "d", "e", "f"},
			args:            args{index: 3},
			want:            &LLNode{value: "d"},
			wantErr:         false,
			expectedDisplay: "a, b, c, e, f",
			expectedLength:  5,
			expectedHead:    &LLNode{value: "a"},
			expectedTail:    &LLNode{value: "f"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := buildLinkedList(tt.nodeValues)
			got, err := l.Remove(tt.args.index)
			// Do we get the error that we expect?
			if (err != nil) != tt.wantErr {
				t.Errorf("LinkedList.Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
			// Because the nodes in a linked list aren't indexed, it's difficult to programmatically
			// ensure that the "next" property points to the proper address in memory throughout
			// these table tests. It should be good enough to observe whether the desired node is
			// either nil as intended (the case when the method errors), or that it contains a
			// specific value as intended (the case when the method is successful).
			if tt.want == nil {
				if got != nil {
					t.Errorf("LinkedList.Remove() got = %v, want %v", got, tt.want)
				}
			} else {
				if got.value != tt.want.value {
					t.Errorf("LinkedList.Remove() got = %v, want %v", got, tt.want)
				}
			}
			// Is the length property updated properly?
			if l.length != tt.expectedLength {
				t.Errorf("LinkedList.Remove() length = [%d], want = [%d]", l.length, tt.expectedLength)
			}
			// Check the internals of the list by comparing the string representations
			if d := l.Display(); d != tt.expectedDisplay {
				t.Errorf("LinkedList.Remove() display = [%s], want = [%s]", d, tt.expectedDisplay)
			}
			// Check the list's head and tail properties; again, node struct comparisons make
			// inspection of an internal node difficult outside of the list.
			if tt.expectedHead == nil {
				if l.head != nil {
					t.Errorf("LinkedList.Remove() head = %v, want = %v", l.head, tt.expectedHead)
				}
			} else {
				if l.head.value != tt.expectedHead.value {
					t.Errorf("LinkedList.Remove() head = %v, want = %v", l.head.value, tt.expectedHead.value)
				}
			}
			if tt.expectedTail == nil {
				if l.tail != nil {
					t.Errorf("LinkedList.Remove() tail = %v, want = %v", l.tail, tt.expectedTail)
				}
			} else {
				if l.tail.value != tt.expectedTail.value {
					t.Errorf("LinkedList.Remove() tail = %v, want = %v", l.tail.value, tt.expectedTail.value)
				}
			}
		})
	}
}

func TestLinkedList_Delete(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name            string
		nodeValues      []string
		args            args
		want            *LLNode
		expectedDisplay string
		expectedLength  int
		expectedHead    *LLNode
		expectedTail    *LLNode
	}{
		{
			name:            "DeleteFromEmpty",
			nodeValues:      []string{},
			args:            args{value: "a"},
			want:            nil,
			expectedDisplay: "",
			expectedLength:  0,
			expectedHead:    nil,
			expectedTail:    nil,
		},
		{
			name:            "DeleteExistingFromSingleItem",
			nodeValues:      []string{"a"},
			args:            args{value: "a"},
			want:            &LLNode{value: "a"},
			expectedDisplay: "",
			expectedLength:  0,
			expectedHead:    nil,
			expectedTail:    nil,
		},
		{
			name:            "DeleteMisingFromSingleItem",
			nodeValues:      []string{"a"},
			args:            args{value: "b"},
			want:            nil,
			expectedDisplay: "a",
			expectedLength:  1,
			expectedHead:    &LLNode{value: "a"},
			expectedTail:    &LLNode{value: "a"},
		},
		{
			name:            "DeleteExistingFromTwoItemsFront",
			nodeValues:      []string{"a", "b"},
			args:            args{value: "a"},
			want:            &LLNode{value: "a"},
			expectedDisplay: "b",
			expectedLength:  1,
			expectedHead:    &LLNode{value: "b"},
			expectedTail:    &LLNode{value: "b"},
		},
		{
			name:            "DeleteExistingFromTwoItemsBack",
			nodeValues:      []string{"a", "b"},
			args:            args{value: "b"},
			want:            &LLNode{value: "b"},
			expectedDisplay: "a",
			expectedLength:  1,
			expectedHead:    &LLNode{value: "a"},
			expectedTail:    &LLNode{value: "a"},
		},
		{
			name:            "DeleteMissingFromTwoItems",
			nodeValues:      []string{"a", "b"},
			args:            args{value: "c"},
			want:            nil,
			expectedDisplay: "a, b",
			expectedLength:  2,
			expectedHead:    &LLNode{value: "a"},
			expectedTail:    &LLNode{value: "b"},
		},
		{
			name:            "DeleteExistingFromPopulatedFront",
			nodeValues:      []string{"a", "b", "c", "d", "e", "f"},
			args:            args{value: "a"},
			want:            &LLNode{value: "a"},
			expectedDisplay: "b, c, d, e, f",
			expectedLength:  5,
			expectedHead:    &LLNode{value: "b"},
			expectedTail:    &LLNode{value: "f"},
		},
		{
			name:            "DeleteExistingFromPopulatedBack",
			nodeValues:      []string{"a", "b", "c", "d", "e", "f"},
			args:            args{value: "f"},
			want:            &LLNode{value: "f"},
			expectedDisplay: "a, b, c, d, e",
			expectedLength:  5,
			expectedHead:    &LLNode{value: "a"},
			expectedTail:    &LLNode{value: "e"},
		},
		{
			name:            "DeleteExistingFromPopulatedMiddle",
			nodeValues:      []string{"a", "b", "c", "d", "e", "f"},
			args:            args{value: "c"},
			want:            &LLNode{value: "c"},
			expectedDisplay: "a, b, d, e, f",
			expectedLength:  5,
			expectedHead:    &LLNode{value: "a"},
			expectedTail:    &LLNode{value: "f"},
		},
		{
			name:            "DeleteMissingFromPopulated",
			nodeValues:      []string{"a", "b", "c", "d", "e", "f"},
			args:            args{value: "z"},
			want:            nil,
			expectedDisplay: "a, b, c, d, e, f",
			expectedLength:  6,
			expectedHead:    &LLNode{value: "a"},
			expectedTail:    &LLNode{value: "f"},
		},
		{
			name:            "DeleteOnlyFirstOccurrenceFromPopulated",
			nodeValues:      []string{"a", "b", "c", "d", "b", "f"},
			args:            args{value: "b"},
			want:            &LLNode{value: "b"},
			expectedDisplay: "a, c, d, b, f",
			expectedLength:  5,
			expectedHead:    &LLNode{value: "a"},
			expectedTail:    &LLNode{value: "f"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := buildLinkedList(tt.nodeValues)
			got := l.Delete(tt.args.value)
			// Because the nodes in a linked list aren't indexed, it's difficult to programmatically
			// ensure that the "next" property points to the proper address in memory throughout
			// these table tests. It should be good enough to observe whether the desired node is
			// either nil as intended (the case when the list is empty), or that it contains a
			// specific value as intended (the case when the method is successful).
			if tt.want == nil {
				if got != nil {
					t.Errorf("LinkedList.Delete() got = %v, want %v", got, tt.want)
				}
			} else {
				if got.value != tt.want.value {
					t.Errorf("LinkedList.Delete() got = %v, want %v", got, tt.want)
				}
			}
			// Is the length property updated properly?
			if l.length != tt.expectedLength {
				t.Errorf("LinkedList.Delete() length = [%d], want = [%d]", l.length, tt.expectedLength)
			}
			// Check the internals of the list by comparing the string representations
			if d := l.Display(); d != tt.expectedDisplay {
				t.Errorf("LinkedList.Delete() display = [%s], want = [%s]", d, tt.expectedDisplay)
			}
			// Check the list's head and tail properties; again, node struct comparisons make
			// inspection of an internal node difficult outside of the list.
			if tt.expectedHead == nil {
				if l.head != nil {
					t.Errorf("LinkedList.Delete() head = %v, want = %v", l.head, tt.expectedHead)
				}
			} else {
				if l.head.value != tt.expectedHead.value {
					t.Errorf("LinkedList.Delete() head = %v, want = %v", l.head.value, tt.expectedHead.value)
				}
			}
			if tt.expectedTail == nil {
				if l.tail != nil {
					t.Errorf("LinkedList.Delete() tail = %v, want = %v", l.tail, tt.expectedTail)
				}
			} else {
				if l.tail.value != tt.expectedTail.value {
					t.Errorf("LinkedList.Delete() tail = %v, want = %v", l.tail.value, tt.expectedTail.value)
				}
			}
		})
	}
}
