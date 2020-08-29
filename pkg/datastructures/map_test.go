package datastructures

import (
	"reflect"
	"testing"
)

func initEntries(initNodes []*mapNode) []*LinkedList {
	lists := make([]*LinkedList, 2)
	for i := range lists {
		lists[i] = NewLinkedList()
	}

	for _, n := range initNodes {
		lists[0].Append(&LLNode{Data: n})
	}

	return lists
}

func TestNewMap(t *testing.T) {
	tests := []struct {
		name         string
		wantNumLists int
	}{
		{
			name:         "create new empty map",
			wantNumLists: 32,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMap(); len(got.entries) != tt.wantNumLists {
				t.Errorf("NewMap() = %v, want %v", len(got.entries), tt.wantNumLists)
			}
		})
	}
}

func TestMap_SetHashFunc(t *testing.T) {
	type fields struct {
		entries  []*LinkedList
		hashFunc func(string) int
	}
	type args struct {
		hashFunc func(string) int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "replace defaultHashFunc",
			fields: fields{
				hashFunc: defaultHashFunc,
			},
			args: args{
				hashFunc: func(string) int { return 42 },
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Map{
				entries:  tt.fields.entries,
				hashFunc: tt.fields.hashFunc,
			}
			defaultResult := m.hashFunc("abcd")
			m.SetHashFunc(tt.args.hashFunc)
			replacedResult := m.hashFunc("abcd")

			if defaultResult == replacedResult {
				t.Errorf("Map.ReplaceHashFunc() default result = %v, replaced result = %v", defaultResult, replacedResult)
			}
		})
	}
}

func TestMap_DoHash(t *testing.T) {
	type fields struct {
		entries  []*LinkedList
		hashFunc func(string) int
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "360 % 32 = 1",
			fields: fields{
				entries:  make([]*LinkedList, 32),
				hashFunc: func(string) int { return 360 },
			},
			args: args{""},
			want: 8,
		},
		{
			name: "360 % 8 = 2",
			fields: fields{
				entries:  make([]*LinkedList, 8),
				hashFunc: func(string) int { return 360 },
			},
			args: args{""},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Map{
				entries:  tt.fields.entries,
				hashFunc: tt.fields.hashFunc,
			}
			if got := m.DoHash(tt.args.key); got != tt.want {
				t.Errorf("Map.DoHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMap_Insert(t *testing.T) {
	type fields struct {
		entries         []*LinkedList
		hashFunc        func(string) int
		resizeThreshold float32
	}
	type args struct {
		key   string
		value interface{}
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantSize int
		wantLen  int
	}{
		{
			// init: [{("a" 1)}, {}]
			// want: [{("a" 1) ("b" 2)}, {}, {}, {}]
			name: "insert prompts resize",
			fields: fields{
				entries: initEntries([]*mapNode{
					{key: "a", val: 1},
				}),
				hashFunc:        func(string) int { return 0 },
				resizeThreshold: 0.4, // ratio during insert: 1/2 = 0.5
			},
			args:     args{key: "b", value: 2},
			wantSize: 2,
			wantLen:  4,
		},
		{
			// init: [{("a" 1)}, {}]
			// want: [{("a" 1) ("b" 2)}, {}]
			name: "insert does not prompt resize",
			fields: fields{
				entries: initEntries([]*mapNode{
					{key: "a", val: 1},
				}),
				hashFunc:        func(string) int { return 0 },
				resizeThreshold: 0.7, // ratio during insert: 1/2 = 0.5
			},
			args:     args{key: "b", value: 2},
			wantSize: 2,
			wantLen:  2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Map{
				entries:         tt.fields.entries,
				hashFunc:        tt.fields.hashFunc,
				resizeThreshold: tt.fields.resizeThreshold,
			}
			m.Insert(tt.args.key, tt.args.value)

			if m.Size() != tt.wantSize {
				t.Errorf("Map.Insert() Size = %v, want = %v", m.Size(), tt.wantSize)
			}
			if len(m.entries) != tt.wantLen {
				t.Errorf("Map.Insert() len(entries) = %v, want = %v", len(m.entries), tt.wantLen)
			}
		})
	}
}

func TestMap_Delete(t *testing.T) {
	type fields struct {
		entries  []*LinkedList
		hashFunc func(string) int
	}
	type args struct {
		key string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantSize int
	}{
		{
			// init: [{}, {}]
			// want: [{}, {}]
			name: "delete from empty",
			fields: fields{
				entries:  initEntries([]*mapNode{}),
				hashFunc: func(string) int { return 0 },
			},
			args:     args{key: "a"},
			wantSize: 0,
		},
		{
			// init: [{("a" 1)}, {}]
			// want: [{}, {}]
			name: "delete with no collisions",
			fields: fields{
				entries: initEntries([]*mapNode{
					{key: "a", val: 1},
				}),
				hashFunc: func(string) int { return 0 },
			},
			args:     args{key: "a"},
			wantSize: 0,
		},
		{
			// init: [{("a" 1) ("b" 1) ("c" 1)}, {}]
			// want: [{("a" 1) ("c" 1)}, {}]
			name: "delete with collisions",
			fields: fields{
				entries: initEntries([]*mapNode{
					{key: "a", val: 1},
					{key: "b", val: 1},
					{key: "c", val: 1},
				}),
				hashFunc: func(string) int { return 0 },
			},
			args:     args{key: "b"},
			wantSize: 2,
		},
		{
			// init: [{("a" 1) ("b" 1) ("c" 1)}, {}]
			// want: [{("a" 1) ("b" 1)}, {}]
			name: "delete at tail with collisions",
			fields: fields{
				entries: initEntries([]*mapNode{
					{key: "a", val: 1},
					{key: "b", val: 1},
					{key: "c", val: 1},
				}),
				hashFunc: func(string) int { return 0 },
			},
			args:     args{key: "c"},
			wantSize: 2,
		},
		{
			// init: [{("a" 1) ("b" 1)}, {}]
			// want: [{("a" 1) ("b" 1)}, {}]
			name: "delete key not found",
			fields: fields{
				entries: initEntries([]*mapNode{
					{key: "a", val: 1},
					{key: "b", val: 1},
				}),
				hashFunc: func(string) int { return 0 },
			},
			args:     args{key: "c"},
			wantSize: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Map{
				entries:  tt.fields.entries,
				hashFunc: tt.fields.hashFunc,
			}
			m.Delete(tt.args.key)

			if m.Size() != tt.wantSize {
				t.Errorf("Map.Delete() Size = %v, want %v", m.Size(), tt.wantSize)
			}
		})
	}
}

func TestMap_Get(t *testing.T) {
	type fields struct {
		entries  []*LinkedList
		hashFunc func(string) int
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
	}{
		{
			name: "get from empty map",
			fields: fields{
				entries:  initEntries(nil),
				hashFunc: func(string) int { return 0 },
			},
			args: args{key: "a"},
			want: nil,
		},
		{
			name: "key found at head",
			fields: fields{
				entries: initEntries([]*mapNode{
					{key: "a", val: 1},
				}),
				hashFunc: func(string) int { return 0 },
			},
			args: args{key: "a"},
			want: 1,
		},
		{
			name: "key found in collisions",
			fields: fields{
				entries: initEntries([]*mapNode{
					{key: "a", val: 1},
					{key: "b", val: 2},
					{key: "c", val: 3},
				}),
				hashFunc: func(string) int { return 0 },
			},
			args: args{key: "c"},
			want: 3,
		},
		{
			name: "key not found, with collisions",
			fields: fields{
				entries: initEntries([]*mapNode{
					{key: "a", val: 1},
					{key: "b", val: 2},
					{key: "c", val: 3},
				}),
				hashFunc: func(string) int { return 0 },
			},
			args: args{key: "d"},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Map{
				entries:  tt.fields.entries,
				hashFunc: tt.fields.hashFunc,
			}
			if got := m.Get(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMap_Size(t *testing.T) {
	type fields struct {
		entries  []*LinkedList
		hashFunc func(string) int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "empty map",
			fields: fields{
				entries: []*LinkedList{},
			},
			want: 0,
		},
		{
			name: "single-item linked lists",
			fields: fields{
				entries: []*LinkedList{
					{Size: 1},
					{Size: 1},
				},
			},
			want: 2,
		},
		{
			name: "multi-item linked lists",
			fields: fields{
				entries: []*LinkedList{
					{Size: 3},
					{Size: 2},
					{Size: 1},
				},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Map{
				entries:  tt.fields.entries,
				hashFunc: tt.fields.hashFunc,
			}
			if got := m.Size(); got != tt.want {
				t.Errorf("Map.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_defaultHashFunc(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "abcd",
			args: args{key: "abcd"},
			want: 394,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := defaultHashFunc(tt.args.key); got != tt.want {
				t.Errorf("defaultHashFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMap_needsResize(t *testing.T) {
	type fields struct {
		entries         []*LinkedList
		hashFunc        func(string) int
		resizeThreshold float32
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "does not need resize",
			fields: fields{
				entries:         initEntries(nil),
				resizeThreshold: 0.5, // ratio: 0/2 = 0
			},
			want: false,
		},
		{
			name: "needs resize",
			fields: fields{
				entries: initEntries([]*mapNode{
					{key: "a", val: 1},
					{key: "b", val: 2},
				}),
				resizeThreshold: 0.5, // ratio: 2/2 = 1
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Map{
				entries:  tt.fields.entries,
				hashFunc: tt.fields.hashFunc,
			}
			if got := m.needsResize(); got != tt.want {
				t.Errorf("Map.needsResize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMap_doResize(t *testing.T) {
	type fields struct {
		entries         []*LinkedList
		hashFunc        func(string) int
		resizeThreshold float32
	}
	tests := []struct {
		name        string
		fields      fields
		wantLen     int
		wantEntries []*LinkedList
	}{
		{
			// init: [
			// 	{ ("b" 1) ("d" 1) },
			// 	{ ("a" 1) ("c" 1) },
			// ]
			//
			// want: [
			//  { ("d" 1) },
			// 	{ ("a" 1) },
			//  { ("b" 1) },
			//  { ("c" 1) },
			// ]
			name: "resize",
			fields: fields{
				entries: initEntries([]*mapNode{
					{key: "a", val: 1},
					{key: "b", val: 1},
					{key: "c", val: 1},
					{key: "d", val: 1},
				}),
				hashFunc: defaultHashFunc,
				// a =  97 %2: 1, %4: 1
				// b =  98 %2: 0, %4: 2
				// c =  99 %2: 1, %4: 3
				// d = 100 %2: 0, %4: 0

			},
			wantLen: 4,
			wantEntries: []*LinkedList{
				{Head: &LLNode{Data: &mapNode{key: "d", val: 1}}},
				{Head: &LLNode{Data: &mapNode{key: "a", val: 1}}},
				{Head: &LLNode{Data: &mapNode{key: "b", val: 1}}},
				{Head: &LLNode{Data: &mapNode{key: "c", val: 1}}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Map{
				entries:  tt.fields.entries,
				hashFunc: tt.fields.hashFunc,
			}
			m.doResize()

			if m.Size() != tt.wantLen {
				t.Errorf("Map.doResize() Size = %v, want = %v", m.Size(), tt.wantLen)
			}

			for i := range m.entries {
				if !reflect.DeepEqual(m.entries[i].Head.Data, tt.wantEntries[i].Head.Data) {
					t.Errorf("Map.doResize() entries[%v].Head.Data = %v, want = %v", i, m.entries[i].Head.Data, tt.wantEntries[i].Head.Data)
				}
			}
		})
	}
}

func TestMap_insertNode(t *testing.T) {
	type fields struct {
		entries  []*LinkedList
		hashFunc func(string) int
	}
	type args struct {
		mn *mapNode
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantSize int
	}{
		{
			// init: [{}, {}]
			// want: [{("a" 1)}, {}]
			name: "insert with no collision",
			fields: fields{
				entries:  initEntries(nil),
				hashFunc: func(string) int { return 0 },
			},
			args:     args{mn: &mapNode{key: "a", val: 1}},
			wantSize: 1,
		},
		{
			// init: [{("a" 1)}, {}]
			// want: [{("a" 1) ("b" 2)}, {}]
			name: "insert with collision, new value",
			fields: fields{
				entries: initEntries([]*mapNode{
					{key: "a", val: 1},
				}),
				hashFunc: func(string) int { return 0 },
			},
			args:     args{mn: &mapNode{key: "b", val: 2}},
			wantSize: 2,
		},
		{
			// init: [{("a" 1) ("b" 1) ("c" 1)}, {}]
			// want: [{("a" 1) ("b" 2) ("c" 1)}, {}]
			name: "insert with chained collision, new value",
			fields: fields{
				entries: initEntries([]*mapNode{
					{key: "a", val: 1},
					{key: "b", val: 1},
					{key: "c", val: 1},
				}),
				hashFunc: func(string) int { return 0 },
			},
			args:     args{mn: &mapNode{key: "b", val: 2}},
			wantSize: 3,
		},
		{
			// init: [{("a" 1)}, {}]
			// want: [{("a" 2)}, {}]
			name: "insert with collision, overwrite value",
			fields: fields{
				entries: initEntries([]*mapNode{
					{key: "a", val: 1},
				}),
				hashFunc: func(string) int { return 0 },
			},
			args:     args{mn: &mapNode{key: "a", val: 2}},
			wantSize: 1,
		},
		{
			// init: [{("a" 1) ("b" 1) ("c" 1)}, {}]
			// want: [{("a" 1) ("b" 2) ("c" 1)}, {}]
			name: "insert with chained collisions, overwrite value",
			fields: fields{
				entries: initEntries([]*mapNode{
					{key: "a", val: 1},
					{key: "b", val: 1},
					{key: "c", val: 1},
				}),
				hashFunc: func(string) int { return 0 },
			},
			args:     args{mn: &mapNode{key: "b", val: 2}},
			wantSize: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Map{
				entries:  tt.fields.entries,
				hashFunc: tt.fields.hashFunc,
			}
			m.insertNode(tt.args.mn)

			if m.Size() != tt.wantSize {
				t.Errorf("Map.insertNode() Size = %v, want = %v", m.Size(), tt.wantSize)
			}
		})
	}
}
