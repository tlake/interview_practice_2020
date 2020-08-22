package doublylinkedlist

import (
	"reflect"
	"testing"
)

func setupNodes(initData []interface{}) []*DLLNode {
	var initNodes []*DLLNode

	if len(initData) > 0 {
		for i := 0; i < len(initData); i++ {
			newNode := &DLLNode{Data: initData[i]}

			if i > 0 {
				prevNode := initNodes[i-1]
				prevNode.Next = newNode
				newNode.Prev = prevNode
			}

			initNodes = append(initNodes, newNode)
		}
	}

	return initNodes
}

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
	type fields struct {
		Head *DLLNode
		Tail *DLLNode
		Len  int
	}
	type args struct {
		givenNode *DLLNode
		data      interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &DLL{
				Head: tt.fields.Head,
				Tail: tt.fields.Tail,
				Len:  tt.fields.Len,
			}
			if err := l.InsertBefore(tt.args.givenNode, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("DLL.InsertBefore() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDLL_InsertAfter(t *testing.T) {
	type fields struct {
		Head *DLLNode
		Tail *DLLNode
		Len  int
	}
	type args struct {
		givenNode *DLLNode
		data      interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &DLL{
				Head: tt.fields.Head,
				Tail: tt.fields.Tail,
				Len:  tt.fields.Len,
			}
			if err := l.InsertAfter(tt.args.givenNode, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("DLL.InsertAfter() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDLL_Delete(t *testing.T) {
	type fields struct {
		Head *DLLNode
		Tail *DLLNode
		Len  int
	}
	type args struct {
		givenNode *DLLNode
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &DLL{
				Head: tt.fields.Head,
				Tail: tt.fields.Tail,
				Len:  tt.fields.Len,
			}
			if err := l.Delete(tt.args.givenNode); (err != nil) != tt.wantErr {
				t.Errorf("DLL.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDLL_Find(t *testing.T) {
	type fields struct {
		Head *DLLNode
		Tail *DLLNode
		Len  int
	}
	type args struct {
		data interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *DLLNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &DLL{
				Head: tt.fields.Head,
				Tail: tt.fields.Tail,
				Len:  tt.fields.Len,
			}
			if got := l.Find(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DLL.Find() = %v, want %v", got, tt.want)
			}
		})
	}
}
