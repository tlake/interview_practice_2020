package datastructures

import (
	"reflect"
	"testing"
)

func setupDLL(initData []interface{}) *DLL {
	initNodes := setupNodes(initData)
	l := NewDLL(nil)
	l.Len = len(initNodes)
	if l.Len > 0 {
		l.Head, l.Tail = initNodes[0], initNodes[len(initNodes)-1]
	}

	return l
}

func TestNewDLLNode(t *testing.T) {
	type args struct {
		data interface{}
	}
	tests := []struct {
		name string
		args args
		want *DLLNode
	}{
		{
			name: "new DLL node",
			args: args{1},
			want: &DLLNode{Data: 1, Prev: nil, Next: nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDLLNode(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDLLNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDLL(t *testing.T) {
	type args struct {
		node *DLLNode
	}
	tests := []struct {
		name string
		args args
		want *DLL
	}{
		{
			name: "new DLL with nil node",
			args: args{nil},
			want: &DLL{Head: nil, Tail: nil, Len: 0},
		},
		{
			name: "new DLL with initialized node",
			args: args{&DLLNode{Data: 1}},
			want: &DLL{
				Head: &DLLNode{Data: 1},
				Tail: &DLLNode{Data: 1},
				Len:  1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDLL(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDLL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDLL_Push(t *testing.T) {
	type args struct {
		data interface{}
	}
	tests := []struct {
		name     string
		initData []interface{}
		args     args
		wantData []interface{}
		wantLen  int
	}{
		{
			name:     "push onto empty DLL",
			initData: []interface{}{},
			args:     args{data: 1},
			wantData: []interface{}{1},
			wantLen:  1,
		},
		{
			name:     "push onto single-item DLL",
			initData: []interface{}{1},
			args:     args{data: 2},
			wantData: []interface{}{2, 1},
			wantLen:  2,
		},
		{
			name:     "push onto two-item DLL",
			initData: []interface{}{1, 2},
			args:     args{data: 3},
			wantData: []interface{}{3, 1, 2},
			wantLen:  3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := setupDLL(tt.initData)
			l.Push(tt.args.data)

			gotData := []interface{}{}
			curr := l.Head
			for curr != nil {
				gotData = append(gotData, curr.Data)
				curr = curr.Next
			}
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("DLL.Push() Data = %v, want %v", gotData, tt.wantData)
			}

			if l.Len != tt.wantLen {
				t.Errorf("DLL.Push() Len = %v, want %v", l.Len, tt.wantLen)
			}

		})
	}
}

func TestDLL_Append(t *testing.T) {
	type args struct {
		data interface{}
	}
	tests := []struct {
		name     string
		initData []interface{}
		args     args
		wantData []interface{}
		wantLen  int
	}{
		{
			name:     "append onto empty DLL",
			initData: []interface{}{},
			args:     args{data: 1},
			wantData: []interface{}{1},
			wantLen:  1,
		},
		{
			name:     "append onto single-item DLL",
			initData: []interface{}{1},
			args:     args{data: 2},
			wantData: []interface{}{1, 2},
			wantLen:  2,
		},
		{
			name:     "append onto two-item DLL",
			initData: []interface{}{1, 2},
			args:     args{data: 3},
			wantData: []interface{}{1, 2, 3},
			wantLen:  3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := setupDLL(tt.initData)
			l.Append(tt.args.data)

			gotData := []interface{}{}
			curr := l.Head
			for curr != nil {
				gotData = append(gotData, curr.Data)
				curr = curr.Next
			}
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("DLL.Append() Data = %v, want %v", gotData, tt.wantData)
			}

			if l.Len != tt.wantLen {
				t.Errorf("DLL.Append() Len = %v, want %v", l.Len, tt.wantLen)
			}
		})
	}
}

func TestDLL_InsertBefore(t *testing.T) {
	type args struct {
		searchForGivenNode int
		data               interface{}
	}
	tests := []struct {
		name     string
		initData []interface{}
		args     args
		wantErr  bool
		wantData []interface{}
		wantLen  int
	}{
		{
			name:     "errors if given node is nil",
			initData: []interface{}{1, 2, 3},
			args: args{
				searchForGivenNode: 0,
			},
			wantErr:  true,
			wantData: []interface{}{1, 2, 3},
			wantLen:  3,
		},
		{
			name:     "insert before head",
			initData: []interface{}{1, 2, 3},
			args: args{
				searchForGivenNode: 1,
				data:               4,
			},
			wantErr:  false,
			wantData: []interface{}{4, 1, 2, 3},
			wantLen:  4,
		},
		{
			name:     "insert before middle",
			initData: []interface{}{1, 2, 3},
			args: args{
				searchForGivenNode: 2,
				data:               4,
			},
			wantErr:  false,
			wantData: []interface{}{1, 4, 2, 3},
			wantLen:  4,
		},
		{
			name:     "insert before end",
			initData: []interface{}{1, 2, 3},
			args: args{
				searchForGivenNode: 3,
				data:               4,
			},
			wantErr:  false,
			wantData: []interface{}{1, 2, 4, 3},
			wantLen:  4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := setupDLL(tt.initData)
			targetNode := l.Find(tt.args.searchForGivenNode)
			err := l.InsertBefore(targetNode, tt.args.data)

			if (err != nil) != tt.wantErr {
				t.Errorf("DLL.InsertBefore() error = %v, wantErr %v", err, tt.wantErr)
			}

			gotData := []interface{}{}
			curr := l.Head
			for curr != nil {
				gotData = append(gotData, curr.Data)
				curr = curr.Next
			}

			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("DLL.InsertBefore() Data = %v, want %v", gotData, tt.wantData)
			}

			if l.Len != tt.wantLen {
				t.Errorf("DLL.InsertBefore() Len = %v, want %v", l.Len, tt.wantLen)
			}

		})
	}
}

func TestDLL_InsertAfter(t *testing.T) {
	type args struct {
		searchForGivenNode int
		data               interface{}
	}
	tests := []struct {
		name     string
		initData []interface{}
		args     args
		wantErr  bool
		wantData []interface{}
		wantLen  int
	}{
		{
			name:     "errors if given node is nil",
			initData: []interface{}{1, 2, 3},
			args: args{
				searchForGivenNode: 0,
			},
			wantErr:  true,
			wantData: []interface{}{1, 2, 3},
			wantLen:  3,
		},
		{
			name:     "insert after head",
			initData: []interface{}{1, 2, 3},
			args: args{
				searchForGivenNode: 1,
				data:               4,
			},
			wantErr:  false,
			wantData: []interface{}{1, 4, 2, 3},
			wantLen:  4,
		},
		{
			name:     "insert after middle",
			initData: []interface{}{1, 2, 3},
			args: args{
				searchForGivenNode: 2,
				data:               4,
			},
			wantErr:  false,
			wantData: []interface{}{1, 2, 4, 3},
			wantLen:  4,
		},
		{
			name:     "insert after end",
			initData: []interface{}{1, 2, 3},
			args: args{
				searchForGivenNode: 3,
				data:               4,
			},
			wantErr:  false,
			wantData: []interface{}{1, 2, 3, 4},
			wantLen:  4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := setupDLL(tt.initData)
			targetNode := l.Find(tt.args.searchForGivenNode)
			err := l.InsertAfter(targetNode, tt.args.data)

			if (err != nil) != tt.wantErr {
				t.Errorf("DLL.InsertAfter() error = %v, wantErr %v", err, tt.wantErr)
			}

			gotData := []interface{}{}
			curr := l.Head
			for curr != nil {
				gotData = append(gotData, curr.Data)
				curr = curr.Next
			}

			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("DLL.InsertAfter() Data = %v, want %v", gotData, tt.wantData)
			}

			if l.Len != tt.wantLen {
				t.Errorf("DLL.InsertAfter() Len = %v, want %v", l.Len, tt.wantLen)
			}

		})
	}
}

func TestDLL_Delete(t *testing.T) {
	tests := []struct {
		name               string
		initData           []interface{}
		searchForGivenNode int
		wantErr            bool
		wantData           []interface{}
		wantLen            int
	}{
		{
			name:               "errors if given node is nil",
			initData:           []interface{}{1, 2, 3},
			searchForGivenNode: 0,
			wantErr:            true,
			wantData:           []interface{}{1, 2, 3},
			wantLen:            3,
		},
		{
			name:               "delete head",
			initData:           []interface{}{1, 2, 3},
			searchForGivenNode: 1,
			wantErr:            false,
			wantData:           []interface{}{2, 3},
			wantLen:            2,
		},
		{
			name:               "delete tail",
			initData:           []interface{}{1, 2, 3},
			searchForGivenNode: 3,
			wantErr:            false,
			wantData:           []interface{}{1, 2},
			wantLen:            2,
		},
		{
			name:               "delete middle",
			initData:           []interface{}{1, 2, 3},
			searchForGivenNode: 2,
			wantErr:            false,
			wantData:           []interface{}{1, 3},
			wantLen:            2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := setupDLL(tt.initData)
			targetNode := l.Find(tt.searchForGivenNode)
			err := l.Delete(targetNode)

			if (err != nil) != tt.wantErr {
				t.Errorf("DLL.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}

			gotData := []interface{}{}
			curr := l.Head
			for curr != nil {
				gotData = append(gotData, curr.Data)
				curr = curr.Next
			}

			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("DLL.Delete() Data = %v, want %v", gotData, tt.wantData)
			}

			if l.Len != tt.wantLen {
				t.Errorf("DLL.Delete() Len = %v, want %v", l.Len, tt.wantLen)
			}

		})
	}
}

func TestDLL_Find(t *testing.T) {
	type args struct {
		data interface{}
	}
	tests := []struct {
		name     string
		initData []interface{}
		args     args
		wantNode *DLLNode
	}{
		{
			name:     "find at head",
			initData: []interface{}{1, 2, 3},
			args:     args{data: 1},
			wantNode: &DLLNode{Data: 1},
		},
		{
			name:     "find at tail",
			initData: []interface{}{1, 2, 3},
			args:     args{data: 3},
			wantNode: &DLLNode{Data: 3},
		},
		{
			name:     "find in middle",
			initData: []interface{}{1, 2, 3},
			args:     args{data: 2},
			wantNode: &DLLNode{Data: 2},
		},
		{
			name:     "not found",
			initData: []interface{}{1, 2, 3},
			args:     args{data: 4},
			wantNode: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := setupDLL(tt.initData)
			foundNode := l.Find(tt.args.data)

			if tt.wantNode == nil {
				if !reflect.DeepEqual(tt.wantNode, foundNode) {
					t.Errorf("DLL.Find() Node = %v, want = %v", foundNode.Data, tt.wantNode.Data)
				}
			} else {
				if tt.wantNode.Data != foundNode.Data {
					t.Errorf("DLL.Find() Node = %v, want = %v", foundNode.Data, tt.wantNode.Data)
				}
			}
		})
	}
}
