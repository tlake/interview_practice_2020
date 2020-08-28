package datastructures

import (
	"reflect"
	"testing"
)

func TestLinkedList_Display(t *testing.T) {
	tests := []struct {
		name       string
		nodeValues []interface{}
		want       string
	}{
		{
			name:       "ZeroNodes",
			nodeValues: []interface{}{},
			want:       "",
		},
		{
			name:       "OneNode",
			nodeValues: []interface{}{"d"},
			want:       "d",
		},
		{
			name:       "TwoNodes",
			nodeValues: []interface{}{"a", "b"},
			want:       "a, b",
		},
		{
			name:       "TenNodes",
			nodeValues: []interface{}{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"},
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
		nodeValues []interface{}
		args       args
		want       string
		wantErr    bool
	}{
		{
			name:       "IndexOutOfBounds",
			nodeValues: []interface{}{"a"},
			args:       args{index: 2},
			want:       "",
			wantErr:    true,
		},
		{
			name:       "FirstNode",
			nodeValues: []interface{}{"a", "b", "c", "d", "e"},
			args:       args{index: 0},
			want:       "a",
			wantErr:    false,
		},
		{
			name:       "LastNode",
			nodeValues: []interface{}{"a", "b", "c", "d", "e"},
			args:       args{index: 4},
			want:       "e",
			wantErr:    false,
		},
		{
			name:       "InternalNode",
			nodeValues: []interface{}{"a", "b", "c", "d", "e"},
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
		Data string
	}
	tests := []struct {
		name       string
		nodeValues []interface{}
		args       args
		want       *LLNode
		wantErr    bool
	}{
		{
			name:       "ValueNotFound",
			nodeValues: []interface{}{"a", "b", "c", "d", "e"},
			args:       args{"f"},
			want:       nil,
			wantErr:    true,
		},
		{
			name:       "ValueFoundAtStart",
			nodeValues: []interface{}{"a", "b", "c", "d", "e"},
			args:       args{"a"},
			want:       &LLNode{Data: "a"},
			wantErr:    false,
		},
		{
			name:       "ValueFoundAtEnd",
			nodeValues: []interface{}{"a", "b", "c", "d", "e"},
			args:       args{"e"},
			want:       &LLNode{Data: "e"},
			wantErr:    false,
		},
		{
			name:       "ValueFoundInMiddle",
			nodeValues: []interface{}{"a", "b", "c", "d", "e"},
			args:       args{"c"},
			want:       &LLNode{Data: "c"},
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := buildLinkedList(tt.nodeValues)
			got, err := l.Search(tt.args.Data)
			if (err != nil) != tt.wantErr {
				t.Errorf("LinkedList.Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want == nil {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("LinkedList.Search() = %v, want %v", got, tt.want)
				}
			} else {
				if !reflect.DeepEqual(got.Data, tt.want.Data) {
					t.Errorf("LinkedList.Search() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestLinkedList_Prepend(t *testing.T) {
	type args struct {
		Node *LLNode
	}
	tests := []struct {
		name            string
		nodeValues      []interface{}
		args            args
		expectedDisplay string
		expectedLength  int
		expectedHead    *LLNode
		expectedTail    *LLNode
	}{
		{
			name:            "PrependToEmpty",
			nodeValues:      []interface{}{},
			args:            args{&LLNode{Data: "a"}},
			expectedDisplay: "a",
			expectedLength:  1,
			expectedHead:    &LLNode{Data: "a"},
			expectedTail:    &LLNode{Data: "a"},
		},
		{
			name:            "PrependToPopulated",
			nodeValues:      []interface{}{"a", "b", "c", "d", "e"},
			args:            args{&LLNode{Data: "f"}},
			expectedDisplay: "f, a, b, c, d, e",
			expectedLength:  6,
			expectedHead:    &LLNode{Data: "f"},
			expectedTail:    &LLNode{Data: "e"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := buildLinkedList(tt.nodeValues)
			// Because the nodes in a linked list aren't indexed, it's difficult to programmatically
			// ensure that the "Next" property points to the proper address in memory throughout
			// these table tests. It should be good enough to observe whether the desired node
			// contains a specific Data as intended.
			l.Prepend(tt.args.Node)

			// Is the Size property updated properly?
			if l.Size != tt.expectedLength {
				t.Errorf("LinkedList.Prepend() = [%d], want = [%d]", l.Size, tt.expectedLength)
			}
			// Check the bs of the list by comparing the string representations
			if d := l.Display(); d != tt.expectedDisplay {
				t.Errorf("LinkedList.Prepend() = [%s], want = [%s]", d, tt.expectedDisplay)
			}
			// Check the list's Head and Tail properties
			if l.Head.Data != tt.expectedHead.Data {
				t.Errorf("LinkedList.Prepend() Head = %v, want = %v", l.Head.Data, tt.expectedHead.Data)
			}
			if l.Tail.Data != tt.expectedTail.Data {
				t.Errorf("LinkedList.Prepend() Tail = %v, want = %v", l.Tail.Data, tt.expectedTail.Data)
			}
		})
	}
}

func TestLinkedList_Append(t *testing.T) {
	type args struct {
		Node *LLNode
	}
	tests := []struct {
		name            string
		nodeValues      []interface{}
		args            args
		expectedDisplay string
		expectedLength  int
		expectedHead    *LLNode
		expectedTail    *LLNode
	}{
		{
			name:            "AppendToEmpty",
			nodeValues:      []interface{}{},
			args:            args{&LLNode{Data: "a"}},
			expectedDisplay: "a",
			expectedLength:  1,
			expectedHead:    &LLNode{Data: "a"},
			expectedTail:    &LLNode{Data: "a"},
		},
		{
			name:            "AppendToPopulated",
			nodeValues:      []interface{}{"a", "b", "c", "d"},
			args:            args{&LLNode{Data: "e"}},
			expectedDisplay: "a, b, c, d, e",
			expectedLength:  5,
			expectedHead:    &LLNode{Data: "a"},
			expectedTail:    &LLNode{Data: "e"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := buildLinkedList(tt.nodeValues)
			// Because the nodes in a linked list aren't indexed, it's difficult to programmatically
			// ensure that the "Next" property points to the proper address in memory throughout
			// these table tests. It should be good enough to observe whether the desired node
			// contains a specific Data as intended.
			l.Append(tt.args.Node)

			// Is the Size property updated properly?
			if l.Size != tt.expectedLength {
				t.Errorf("LinkedList.Append() = [%d], want = [%d]", l.Size, tt.expectedLength)
			}
			// Check the bs of the list by comparing the string representations
			if d := l.Display(); d != tt.expectedDisplay {
				t.Errorf("LinkedList.Append() = [%s], want = [%s]", d, tt.expectedDisplay)
			}
			// Check the list's Head and Tail properties
			if l.Head.Data != tt.expectedHead.Data {
				t.Errorf("LinkedList.Append() Head = %v, want = %v", l.Head.Data, tt.expectedHead.Data)
			}
			if l.Tail.Data != tt.expectedTail.Data {
				t.Errorf("LinkedList.Append() Tail = %v, want = %v", l.Tail.Data, tt.expectedTail.Data)
			}
		})
	}
}

func TestLinkedList_Insert(t *testing.T) {
	type args struct {
		node  *LLNode
		index int
	}
	tests := []struct {
		name            string
		nodeValues      []interface{}
		args            args
		wantErr         bool
		expectedDisplay string
		expectedLength  int
		expectedHead    *LLNode
		expectedTail    *LLNode
	}{
		{
			name:       "InsertOutOfBounds",
			nodeValues: []interface{}{"a", "b"},
			args: args{
				node:  &LLNode{Data: "8"},
				index: 3,
			},
			wantErr:         true,
			expectedDisplay: "a, b",
			expectedLength:  2,
			expectedHead:    &LLNode{Data: "a"},
			expectedTail:    &LLNode{Data: "b"},
		},
		{
			name:       "InsertToEmpty",
			nodeValues: []interface{}{},
			args: args{
				node:  &LLNode{Data: "a"},
				index: 0,
			},
			wantErr:         false,
			expectedDisplay: "a",
			expectedLength:  1,
			expectedHead:    &LLNode{Data: "a"},
			expectedTail:    &LLNode{Data: "a"},
		},
		{
			name:       "InsertAtFrontIntoSingleItem",
			nodeValues: []interface{}{"a"},
			args: args{
				node:  &LLNode{Data: "b"},
				index: 0,
			},
			wantErr:         false,
			expectedDisplay: "b, a",
			expectedLength:  2,
			expectedHead:    &LLNode{Data: "b"},
			expectedTail:    &LLNode{Data: "a"},
		},
		{
			name:       "InsertAtEndIntoSingleItem",
			nodeValues: []interface{}{"a"},
			args: args{
				node:  &LLNode{Data: "b"},
				index: 1,
			},
			wantErr:         false,
			expectedDisplay: "a, b",
			expectedLength:  2,
			expectedHead:    &LLNode{Data: "a"},
			expectedTail:    &LLNode{Data: "b"},
		},
		{
			name:       "InsertAtFrontIntoTwoItems",
			nodeValues: []interface{}{"a", "b"},
			args: args{
				node:  &LLNode{Data: "c"},
				index: 0,
			},
			wantErr:         false,
			expectedDisplay: "c, a, b",
			expectedLength:  3,
			expectedHead:    &LLNode{Data: "c"},
			expectedTail:    &LLNode{Data: "b"},
		},
		{
			name:       "InsertAtEndIntoTwoItems",
			nodeValues: []interface{}{"a", "b"},
			args: args{
				node:  &LLNode{Data: "c"},
				index: 2,
			},
			wantErr:         false,
			expectedDisplay: "a, b, c",
			expectedLength:  3,
			expectedHead:    &LLNode{Data: "a"},
			expectedTail:    &LLNode{Data: "c"},
		},
		{
			name:       "InsertAtMiddleIntoTwoItems",
			nodeValues: []interface{}{"a", "b"},
			args: args{
				node:  &LLNode{Data: "c"},
				index: 1,
			},
			wantErr:         false,
			expectedDisplay: "a, c, b",
			expectedLength:  3,
			expectedHead:    &LLNode{Data: "a"},
			expectedTail:    &LLNode{Data: "b"},
		},
		{
			name:       "InsertAtMiddleIntoPopulated",
			nodeValues: []interface{}{"a", "b", "c", "d", "e"},
			args: args{
				node:  &LLNode{Data: "f"},
				index: 3,
			},
			wantErr:         false,
			expectedDisplay: "a, b, c, f, d, e",
			expectedLength:  6,
			expectedHead:    &LLNode{Data: "a"},
			expectedTail:    &LLNode{Data: "e"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := buildLinkedList(tt.nodeValues)
			err := l.Insert(tt.args.node, tt.args.index)

			// Do we get the error that we expect?
			if (err != nil) != tt.wantErr {
				t.Errorf("LinkedList.Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
			// Is the Size property updated properly?
			if l.Size != tt.expectedLength {
				t.Errorf("LinkedList.Insert() Size = [%d], want = [%d]", l.Size, tt.expectedLength)
			}
			// Check the bs of the list by comparing the string representations
			if d := l.Display(); d != tt.expectedDisplay {
				t.Errorf("LinkedList.Insert() display = [%s], want = [%s]", d, tt.expectedDisplay)
			}
			// Check the list's Head and Tail properties
			if l.Head.Data != tt.expectedHead.Data {
				t.Errorf("LinkedList.Insert() LinkedList Head = %s, want = %s", l.Head.Data, tt.expectedHead.Data)
			}
			if l.Tail.Data != tt.expectedTail.Data {
				t.Errorf("LinkedList.Insert() LinkedList Tail = %s, want = %s", l.Tail.Data, tt.expectedTail.Data)
			}
		})
	}
}

func TestLinkedList_Shift(t *testing.T) {
	tests := []struct {
		name            string
		nodeValues      []interface{}
		want            *LLNode
		expectedDisplay string
		expectedLength  int
		expectedHead    *LLNode
		expectedTail    *LLNode
	}{
		{
			name:            "ShiftEmptyList",
			nodeValues:      []interface{}{},
			want:            nil,
			expectedDisplay: "",
			expectedLength:  0,
			expectedHead:    nil,
			expectedTail:    nil,
		},
		{
			name:            "ShiftOneItemList",
			nodeValues:      []interface{}{"a"},
			want:            &LLNode{Data: "a"},
			expectedDisplay: "",
			expectedLength:  0,
			expectedHead:    nil,
			expectedTail:    nil,
		},
		{
			name:            "ShiftTwoItemList",
			nodeValues:      []interface{}{"a", "b"},
			want:            &LLNode{Data: "a"},
			expectedDisplay: "b",
			expectedLength:  1,
			expectedHead:    &LLNode{Data: "b"},
			expectedTail:    &LLNode{Data: "b"},
		},
		{
			name:            "ShiftThreeItemList",
			nodeValues:      []interface{}{"a", "b", "c"},
			want:            &LLNode{Data: "a"},
			expectedDisplay: "b, c",
			expectedLength:  2,
			expectedHead:    &LLNode{Data: "b"},
			expectedTail:    &LLNode{Data: "c"},
		},
		{
			name:            "ShiftPopulatedList",
			nodeValues:      []interface{}{"a", "b", "c", "d", "e", "f"},
			want:            &LLNode{Data: "a"},
			expectedDisplay: "b, c, d, e, f",
			expectedLength:  5,
			expectedHead:    &LLNode{Data: "b"},
			expectedTail:    &LLNode{Data: "f"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := buildLinkedList(tt.nodeValues)
			got := l.Shift()
			// Because the nodes in a linked list aren't indexed, it's difficult to programmatically
			// ensure that the "Next" property points to the proper address in memory throughout
			// these table tests. It should be good enough to observe whether the desired node is
			// either nil as intended (the case when the list is empty), or that it contains a
			// specific Data as intended (the case when the method is otherwise successful).
			if tt.want == nil {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("LinkedList.Shift() got = %v, want %v", got, tt.want)
				}
			} else {
				if got.Data != tt.want.Data {
					t.Errorf("LinkedList.Shift() got = %v, want %v", got, tt.want)
				}
			}
			// Is the Size property updated properly?
			if l.Size != tt.expectedLength {
				t.Errorf("LinkedList.Shift() Size = [%d], want = [%d]", l.Size, tt.expectedLength)
			}
			// Check the bs of the list by comparing the string representations
			if d := l.Display(); d != tt.expectedDisplay {
				t.Errorf("LinkedList.Shift() display = [%s], want = [%s]", d, tt.expectedDisplay)
			}
			// Check the list's Head and Tail properties; again, node struct comparisons make
			// inspection of an internal node difficult outside of the list.
			if tt.expectedHead == nil {
				if l.Head != nil {
					t.Errorf("LinkedList.Shift() Head = %v, want = %v", l.Head, tt.expectedHead)
				}
			} else {
				if l.Head.Data != tt.expectedHead.Data {
					t.Errorf("LinkedList.Shift() Head = %v, want = %v", l.Head.Data, tt.expectedHead.Data)
				}
			}
			if tt.expectedTail == nil {
				if l.Tail != nil {
					t.Errorf("LinkedList.Shift() Tail = %v, want = %v", l.Tail, tt.expectedTail)
				}
			} else {
				if l.Tail.Data != tt.expectedTail.Data {
					t.Errorf("LinkedList.Shift() Tail = %v, want = %v", l.Tail.Data, tt.expectedTail.Data)
				}
			}
		})
	}
}

func TestLinkedList_Pop(t *testing.T) {
	tests := []struct {
		name            string
		nodeValues      []interface{}
		want            *LLNode
		expectedDisplay string
		expectedLength  int
		expectedHead    *LLNode
		expectedTail    *LLNode
	}{
		{
			name:            "PopEmptyList",
			nodeValues:      []interface{}{},
			want:            nil,
			expectedDisplay: "",
			expectedLength:  0,
			expectedHead:    nil,
			expectedTail:    nil,
		},
		{
			name:            "PopOneItemList",
			nodeValues:      []interface{}{"a"},
			want:            &LLNode{Data: "a"},
			expectedDisplay: "",
			expectedLength:  0,
			expectedHead:    nil,
			expectedTail:    nil,
		},
		{
			name:            "PopTwoItemList",
			nodeValues:      []interface{}{"a", "b"},
			want:            &LLNode{Data: "b"},
			expectedDisplay: "a",
			expectedLength:  1,
			expectedHead:    &LLNode{Data: "a"},
			expectedTail:    &LLNode{Data: "a"},
		},
		{
			name:            "PopThreeItemList",
			nodeValues:      []interface{}{"a", "b", "c"},
			want:            &LLNode{Data: "c"},
			expectedDisplay: "a, b",
			expectedLength:  2,
			expectedHead:    &LLNode{Data: "a"},
			expectedTail:    &LLNode{Data: "b"},
		},
		{
			name:            "PopPopulatedList",
			nodeValues:      []interface{}{"a", "b", "c", "d", "e", "f"},
			want:            &LLNode{Data: "f"},
			expectedDisplay: "a, b, c, d, e",
			expectedLength:  5,
			expectedHead:    &LLNode{Data: "a"},
			expectedTail:    &LLNode{Data: "e"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := buildLinkedList(tt.nodeValues)
			got := l.Pop()
			// Because the nodes in a linked list aren't indexed, it's difficult to programmatically
			// ensure that the "Next" property points to the proper address in memory throughout
			// these table tests. It should be good enough to observe whether the desired node is
			// either nil as intended (the case when the list is empty), or that it contains a
			// specific Data as intended (the case when the method is otherwise successful).
			if tt.want == nil {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("LinkedList.Pop() got = %v, want %v", got, tt.want)
				}
			} else {
				if got.Data != tt.want.Data {
					t.Errorf("LinkedList.Pop() got = %v, want %v", got, tt.want)
				}
			}
			// Is the Size property updated properly?
			if l.Size != tt.expectedLength {
				t.Errorf("LinkedList.Pop() Size = [%d], want = [%d]", l.Size, tt.expectedLength)
			}
			// Check the bs of the list by comparing the string representations
			if d := l.Display(); d != tt.expectedDisplay {
				t.Errorf("LinkedList.Pop() display = [%s], want = [%s]", d, tt.expectedDisplay)
			}
			// Check the list's Head and Tail properties; again, node struct comparisons make
			// inspection of an internal node difficult outside of the list.
			if tt.expectedHead == nil {
				if l.Head != nil {
					t.Errorf("LinkedList.Pop() Head = %v, want = %v", l.Head, tt.expectedHead)
				}
			} else {
				if l.Head.Data != tt.expectedHead.Data {
					t.Errorf("LinkedList.Pop() Head = %v, want = %v", l.Head.Data, tt.expectedHead.Data)
				}
			}
			if tt.expectedTail == nil {
				if l.Tail != nil {
					t.Errorf("LinkedList.Pop() Tail = %v, want = %v", l.Tail, tt.expectedTail)
				}
			} else {
				if l.Tail.Data != tt.expectedTail.Data {
					t.Errorf("LinkedList.Pop() Tail = %v, want = %v", l.Tail.Data, tt.expectedTail.Data)
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
		nodeValues      []interface{}
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
			nodeValues:      []interface{}{"a", "b"},
			args:            args{index: 2},
			want:            nil,
			wantErr:         true,
			expectedDisplay: "a, b",
			expectedLength:  2,
			expectedHead:    &LLNode{Data: "a"},
			expectedTail:    &LLNode{Data: "b"},
		},
		{
			name:            "RemoveFromEmpty",
			nodeValues:      []interface{}{},
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
			nodeValues:      []interface{}{"a"},
			args:            args{index: 0},
			want:            &LLNode{Data: "a"},
			wantErr:         false,
			expectedDisplay: "",
			expectedLength:  0,
			expectedHead:    nil,
			expectedTail:    nil,
		},
		{
			name:            "RemoveAtFrontFromTwoItems",
			nodeValues:      []interface{}{"a", "b"},
			args:            args{index: 0},
			want:            &LLNode{Data: "a"},
			wantErr:         false,
			expectedDisplay: "b",
			expectedLength:  1,
			expectedHead:    &LLNode{Data: "b"},
			expectedTail:    &LLNode{Data: "b"},
		},
		{
			name:            "RemoveAtBackFromTwoItems",
			nodeValues:      []interface{}{"a", "b"},
			args:            args{index: 1},
			want:            &LLNode{Data: "b"},
			wantErr:         false,
			expectedDisplay: "a",
			expectedLength:  1,
			expectedHead:    &LLNode{Data: "a"},
			expectedTail:    &LLNode{Data: "a"},
		},
		{
			name:            "RemoveAtFrontFromThreeItems",
			nodeValues:      []interface{}{"a", "b", "c"},
			args:            args{index: 0},
			want:            &LLNode{Data: "a"},
			wantErr:         false,
			expectedDisplay: "b, c",
			expectedLength:  2,
			expectedHead:    &LLNode{Data: "b"},
			expectedTail:    &LLNode{Data: "c"},
		},
		{
			name:            "RemoveAtBackFromThreeItems",
			nodeValues:      []interface{}{"a", "b", "c"},
			args:            args{index: 2},
			want:            &LLNode{Data: "c"},
			wantErr:         false,
			expectedDisplay: "a, b",
			expectedLength:  2,
			expectedHead:    &LLNode{Data: "a"},
			expectedTail:    &LLNode{Data: "b"},
		},
		{
			name:            "RemoveAtMiddleFromThreeItems",
			nodeValues:      []interface{}{"a", "b", "c"},
			args:            args{index: 1},
			want:            &LLNode{Data: "b"},
			wantErr:         false,
			expectedDisplay: "a, c",
			expectedLength:  2,
			expectedHead:    &LLNode{Data: "a"},
			expectedTail:    &LLNode{Data: "c"},
		},
		{
			name:            "RemoveAtFrontFromPopulated",
			nodeValues:      []interface{}{"a", "b", "c", "d", "e", "f"},
			args:            args{index: 0},
			want:            &LLNode{Data: "a"},
			wantErr:         false,
			expectedDisplay: "b, c, d, e, f",
			expectedLength:  5,
			expectedHead:    &LLNode{Data: "b"},
			expectedTail:    &LLNode{Data: "f"},
		},
		{
			name:            "RemoveAtBackFromPopulated",
			nodeValues:      []interface{}{"a", "b", "c", "d", "e", "f"},
			args:            args{index: 5},
			want:            &LLNode{Data: "f"},
			wantErr:         false,
			expectedDisplay: "a, b, c, d, e",
			expectedLength:  5,
			expectedHead:    &LLNode{Data: "a"},
			expectedTail:    &LLNode{Data: "e"},
		},
		{
			name:            "RemoveAtMiddleFromPopulated",
			nodeValues:      []interface{}{"a", "b", "c", "d", "e", "f"},
			args:            args{index: 3},
			want:            &LLNode{Data: "d"},
			wantErr:         false,
			expectedDisplay: "a, b, c, e, f",
			expectedLength:  5,
			expectedHead:    &LLNode{Data: "a"},
			expectedTail:    &LLNode{Data: "f"},
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
			// ensure that the "Next" property points to the proper address in memory throughout
			// these table tests. It should be good enough to observe whether the desired node is
			// either nil as intended (the case when the method errors), or that it contains a
			// specific Data as intended (the case when the method is successful).
			if tt.want == nil {
				if got != nil {
					t.Errorf("LinkedList.Remove() got = %v, want %v", got, tt.want)
				}
			} else {
				if got.Data != tt.want.Data {
					t.Errorf("LinkedList.Remove() got = %v, want %v", got, tt.want)
				}
			}
			// Is the Size property updated properly?
			if l.Size != tt.expectedLength {
				t.Errorf("LinkedList.Remove() Size = [%d], want = [%d]", l.Size, tt.expectedLength)
			}
			// Check the bs of the list by comparing the string representations
			if d := l.Display(); d != tt.expectedDisplay {
				t.Errorf("LinkedList.Remove() display = [%s], want = [%s]", d, tt.expectedDisplay)
			}
			// Check the list's Head and Tail properties; again, node struct comparisons make
			// inspection of an internal node difficult outside of the list.
			if tt.expectedHead == nil {
				if l.Head != nil {
					t.Errorf("LinkedList.Remove() Head = %v, want = %v", l.Head, tt.expectedHead)
				}
			} else {
				if l.Head.Data != tt.expectedHead.Data {
					t.Errorf("LinkedList.Remove() Head = %v, want = %v", l.Head.Data, tt.expectedHead.Data)
				}
			}
			if tt.expectedTail == nil {
				if l.Tail != nil {
					t.Errorf("LinkedList.Remove() Tail = %v, want = %v", l.Tail, tt.expectedTail)
				}
			} else {
				if l.Tail.Data != tt.expectedTail.Data {
					t.Errorf("LinkedList.Remove() Tail = %v, want = %v", l.Tail.Data, tt.expectedTail.Data)
				}
			}
		})
	}
}

func TestLinkedList_Delete(t *testing.T) {
	type args struct {
		Data string
	}
	tests := []struct {
		name            string
		nodeValues      []interface{}
		args            args
		want            *LLNode
		expectedDisplay string
		expectedLength  int
		expectedHead    *LLNode
		expectedTail    *LLNode
	}{
		{
			name:            "DeleteFromEmpty",
			nodeValues:      []interface{}{},
			args:            args{Data: "a"},
			want:            nil,
			expectedDisplay: "",
			expectedLength:  0,
			expectedHead:    nil,
			expectedTail:    nil,
		},
		{
			name:            "DeleteExistingFromSingleItem",
			nodeValues:      []interface{}{"a"},
			args:            args{Data: "a"},
			want:            &LLNode{Data: "a"},
			expectedDisplay: "",
			expectedLength:  0,
			expectedHead:    nil,
			expectedTail:    nil,
		},
		{
			name:            "DeleteMisingFromSingleItem",
			nodeValues:      []interface{}{"a"},
			args:            args{Data: "b"},
			want:            nil,
			expectedDisplay: "a",
			expectedLength:  1,
			expectedHead:    &LLNode{Data: "a"},
			expectedTail:    &LLNode{Data: "a"},
		},
		{
			name:            "DeleteExistingFromTwoItemsFront",
			nodeValues:      []interface{}{"a", "b"},
			args:            args{Data: "a"},
			want:            &LLNode{Data: "a"},
			expectedDisplay: "b",
			expectedLength:  1,
			expectedHead:    &LLNode{Data: "b"},
			expectedTail:    &LLNode{Data: "b"},
		},
		{
			name:            "DeleteExistingFromTwoItemsBack",
			nodeValues:      []interface{}{"a", "b"},
			args:            args{Data: "b"},
			want:            &LLNode{Data: "b"},
			expectedDisplay: "a",
			expectedLength:  1,
			expectedHead:    &LLNode{Data: "a"},
			expectedTail:    &LLNode{Data: "a"},
		},
		{
			name:            "DeleteMissingFromTwoItems",
			nodeValues:      []interface{}{"a", "b"},
			args:            args{Data: "c"},
			want:            nil,
			expectedDisplay: "a, b",
			expectedLength:  2,
			expectedHead:    &LLNode{Data: "a"},
			expectedTail:    &LLNode{Data: "b"},
		},
		{
			name:            "DeleteExistingFromPopulatedFront",
			nodeValues:      []interface{}{"a", "b", "c", "d", "e", "f"},
			args:            args{Data: "a"},
			want:            &LLNode{Data: "a"},
			expectedDisplay: "b, c, d, e, f",
			expectedLength:  5,
			expectedHead:    &LLNode{Data: "b"},
			expectedTail:    &LLNode{Data: "f"},
		},
		{
			name:            "DeleteExistingFromPopulatedBack",
			nodeValues:      []interface{}{"a", "b", "c", "d", "e", "f"},
			args:            args{Data: "f"},
			want:            &LLNode{Data: "f"},
			expectedDisplay: "a, b, c, d, e",
			expectedLength:  5,
			expectedHead:    &LLNode{Data: "a"},
			expectedTail:    &LLNode{Data: "e"},
		},
		{
			name:            "DeleteExistingFromPopulatedMiddle",
			nodeValues:      []interface{}{"a", "b", "c", "d", "e", "f"},
			args:            args{Data: "c"},
			want:            &LLNode{Data: "c"},
			expectedDisplay: "a, b, d, e, f",
			expectedLength:  5,
			expectedHead:    &LLNode{Data: "a"},
			expectedTail:    &LLNode{Data: "f"},
		},
		{
			name:            "DeleteMissingFromPopulated",
			nodeValues:      []interface{}{"a", "b", "c", "d", "e", "f"},
			args:            args{Data: "z"},
			want:            nil,
			expectedDisplay: "a, b, c, d, e, f",
			expectedLength:  6,
			expectedHead:    &LLNode{Data: "a"},
			expectedTail:    &LLNode{Data: "f"},
		},
		{
			name:            "DeleteOnlyFirstOccurrenceFromPopulated",
			nodeValues:      []interface{}{"a", "b", "c", "d", "b", "f"},
			args:            args{Data: "b"},
			want:            &LLNode{Data: "b"},
			expectedDisplay: "a, c, d, b, f",
			expectedLength:  5,
			expectedHead:    &LLNode{Data: "a"},
			expectedTail:    &LLNode{Data: "f"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := buildLinkedList(tt.nodeValues)
			got := l.Delete(tt.args.Data)
			// Because the nodes in a linked list aren't indexed, it's difficult to programmatically
			// ensure that the "Next" property points to the proper address in memory throughout
			// these table tests. It should be good enough to observe whether the desired node is
			// either nil as intended (the case when the list is empty), or that it contains a
			// specific Data as intended (the case when the method is successful).
			if tt.want == nil {
				if got != nil {
					t.Errorf("LinkedList.Delete() got = %v, want %v", got, tt.want)
				}
			} else {
				if got.Data != tt.want.Data {
					t.Errorf("LinkedList.Delete() got = %v, want %v", got, tt.want)
				}
			}
			// Is the Size property updated properly?
			if l.Size != tt.expectedLength {
				t.Errorf("LinkedList.Delete() Size = [%d], want = [%d]", l.Size, tt.expectedLength)
			}
			// Check the bs of the list by comparing the string representations
			if d := l.Display(); d != tt.expectedDisplay {
				t.Errorf("LinkedList.Delete() display = [%s], want = [%s]", d, tt.expectedDisplay)
			}
			// Check the list's Head and Tail properties; again, node struct comparisons make
			// inspection of an internal node difficult outside of the list.
			if tt.expectedHead == nil {
				if l.Head != nil {
					t.Errorf("LinkedList.Delete() Head = %v, want = %v", l.Head, tt.expectedHead)
				}
			} else {
				if l.Head.Data != tt.expectedHead.Data {
					t.Errorf("LinkedList.Delete() Head = %v, want = %v", l.Head.Data, tt.expectedHead.Data)
				}
			}
			if tt.expectedTail == nil {
				if l.Tail != nil {
					t.Errorf("LinkedList.Delete() Tail = %v, want = %v", l.Tail, tt.expectedTail)
				}
			} else {
				if l.Tail.Data != tt.expectedTail.Data {
					t.Errorf("LinkedList.Delete() Tail = %v, want = %v", l.Tail.Data, tt.expectedTail.Data)
				}
			}
		})
	}
}
