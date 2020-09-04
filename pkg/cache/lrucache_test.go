package cache

import (
	"reflect"
	"testing"
)

func TestNewLRUCache(t *testing.T) {
	type args struct {
		c int
	}
	tests := []struct {
		name string
		args args
		want *LRUCache
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLRUCache(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLRUCache() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLRUCache_Put(t *testing.T) {
	type args struct {
		key   string
		value interface{}
	}
	type initData struct {
		key   string
		value interface{}
	}
	type init struct {
		capacity int
		data     []initData
	}
	tests := []struct {
		name    string
		init    init
		args    args
		wantErr bool
	}{
		{
			name: "put into empty",
			init: init{
				capacity: 2,
			},
			args: args{
				key: "a", value: 1,
			},
			wantErr: false,
		},
		{
			name: "put into not-empty",
			init: init{
				capacity: 2,
				data: []initData{
					{key: "a", value: 1},
				},
			},
			args: args{
				key: "b", value: 2,
			},
			wantErr: false,
		},
		{
			name: "put overwrites",
			init: init{
				capacity: 2,
				data: []initData{
					{key: "a", value: 1},
				},
			},
			args: args{
				key: "a", value: 2,
			},
			wantErr: false,
		},
		{
			name: "put at capacity",
			init: init{
				capacity: 2,
				data: []initData{
					{key: "a", value: 1},
					{key: "b", value: 2},
				},
			},
			args: args{
				key: "c", value: 3,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewLRUCache(tt.init.capacity)

			for _, p := range tt.init.data {
				l.Put(p.key, p.value)
			}

			if err := l.Put(tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("LRUCache.Put() error = \"%v\", wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLRUCache_Get(t *testing.T) {
	type args struct {
		key string
	}
	type initData struct {
		key   string
		value interface{}
	}
	type init struct {
		capacity int
		data     []initData
	}
	tests := []struct {
		name    string
		init    init
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "it just works :tm:",
			init: init{
				capacity: 2,
				data: []initData{
					{key: "a", value: 1},
				},
			},
			args:    args{"a"},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewLRUCache(tt.init.capacity)

			for _, p := range tt.init.data {
				l.Put(p.key, p.value)
			}

			got, err := l.Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("LRUCache.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LRUCache.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLRUCache_Remove(t *testing.T) {
	type args struct {
		key string
	}
	type initData struct {
		key   string
		value interface{}
	}
	type init struct {
		capacity int
		data     []initData
	}
	tests := []struct {
		name     string
		init     init
		args     args
		wantErr  bool
		wantSize int
	}{
		{
			name: "it just works :tm:",
			init: init{
				capacity: 2,
				data: []initData{
					{key: "a", value: 1},
				},
			},
			args:     args{key: "a"},
			wantErr:  false,
			wantSize: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewLRUCache(tt.init.capacity)
			if err := l.Remove(tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("LRUCache.Remove() error = %v, wantErr %v", err, tt.wantErr)
			}

			if size := l.size(); size != tt.wantSize {
				t.Errorf("LRUCache.Remove() size = %v, want = %v", size, tt.wantSize)
			}
		})
	}
}

// func TestLRUCache_delete(t *testing.T) {
// 	type args struct {
// 		key string
// 	}
// 	type initData struct {
// 		key   string
// 		value interface{}
// 	}
// 	type init struct {
// 		capacity int
// 		data     []initData
// 	}
// 	tests := []struct {
// 		name    string
// 		init    init
// 		args    args
// 		wantErr bool
// 	}{
// 		{
// 			name: "it just works",
// 			init: init{
// 				capacity: 2,
// 				data: []initData{
// 					{key: "a", value: 1},
// 				},
// 			},
// 			args:    args{key: "a"},
// 			wantErr: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 				l := NewLRUCache(tt.init.capacity)
// 				for _, p := range tt.init.data {
// 					l.Put(p.key, p.value)
// 				}
// 				fmt.Println(l.dll)
// 				fmt.Println(l.m)
//
// 				got, _ := l.Get(tt.args.key)
// 				fmt.Println(got)
// 				dllNode, ok := got.(*datastructures.DLLNode)
// 				if !ok {
// 					t.Errorf("fuck off, got type %T instead of *DLLNode, goddammit %v", dllNode, ok)
// 				}
//
// 				if err := l.delete(dllNode); (err != nil) != tt.wantErr {
// 					t.Errorf("LRUCache.delete() error = %v, wantErr %v", err, tt.wantErr)
// 				}
// 		})
// 	}
// }

// func TestLRUCache_evict(t *testing.T) {
// 	type fields struct {
// 		capacity int
// 		dll      *datastructures.DLL
// 		m        map[string]*datastructures.DLLNode
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			l := &LRUCache{
// 				capacity: tt.fields.capacity,
// 				dll:      tt.fields.dll,
// 				m:        tt.fields.m,
// 			}
// 			if err := l.evict(); (err != nil) != tt.wantErr {
// 				t.Errorf("LRUCache.evict() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }

// func TestLRUCache_size(t *testing.T) {
// 	type fields struct {
// 		capacity int
// 		dll      *datastructures.DLL
// 		m        map[string]*datastructures.DLLNode
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		want   int
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			l := &LRUCache{
// 				capacity: tt.fields.capacity,
// 				dll:      tt.fields.dll,
// 				m:        tt.fields.m,
// 			}
// 			if got := l.size(); got != tt.want {
// 				t.Errorf("LRUCache.size() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
