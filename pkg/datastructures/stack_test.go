package datastructures

import (
	"reflect"
	"testing"
)

func TestNewStack(t *testing.T) {
	tests := []struct {
		name string
		want *Stack
	}{
		{
			name: "creates a new, empty stack",
			want: &Stack{&LinkedList{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStack(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Push(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name     string
		initData []interface{}
		args     args
		wantSize int
	}{
		{
			name:     "add new node to empty stack",
			initData: []interface{}{},
			args:     args{value: 1},
			wantSize: 1,
		},
		{
			name:     "add new node to head",
			initData: []interface{}{1, 2, 2},
			args:     args{value: 4},
			wantSize: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := buildLinkedList(tt.initData)
			initialHead := l.Head

			s := &Stack{LinkedList: l}
			s.Push(tt.args.value)

			if !reflect.DeepEqual(s.Head.Next, initialHead) {
				t.Errorf("Stack.Pop() second node = %v, want = %v", s.Head.Next, initialHead)
			}

			if s.Head.Data != tt.args.value {
				t.Errorf("Stack.Pop() Head = %v, want = %v", s.Head.Data, tt.args.value)
			}

			if s.Size != tt.wantSize {
				t.Errorf("Stack.Pop() Size = %v, want = %v", s.Size, tt.wantSize)
			}
		})
	}
}

func TestStack_Pop(t *testing.T) {
	tests := []struct {
		name     string
		initData []interface{}
		want     interface{}
		wantSize int
	}{
		{
			name:     "pop from empty stack",
			initData: []interface{}{},
			want:     nil,
			wantSize: 0,
		},
		{
			name:     "pop from populated stack",
			initData: []interface{}{1, 2, 3},
			want:     1,
			wantSize: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := buildLinkedList(tt.initData)
			s := &Stack{LinkedList: l}

			if got := s.Pop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Stack.Pop() data = %v, want = %v", got, tt.want)
			}

			if s.Size != tt.wantSize {
				t.Errorf("Stack.Pop() Size = %v, want = %v", s.Size, tt.wantSize)
			}
		})
	}
}

func TestStack_Peek(t *testing.T) {
	tests := []struct {
		name     string
		initData []interface{}
		want     interface{}
		wantSize int
	}{
		{
			name:     "peek from empty stack",
			initData: []interface{}{},
			want:     nil,
			wantSize: 0,
		},
		{
			name:     "peek from populated stack",
			initData: []interface{}{1, 2, 3},
			want:     1,
			wantSize: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := buildLinkedList(tt.initData)
			s := &Stack{LinkedList: l}

			if got := s.Peek(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Stack.Peek() Data = %v, want = %v", got, tt.want)
			}

			if s.Size != tt.wantSize {
				t.Errorf("Stack.Peek() Size = %v, want = %v", s.Size, tt.wantSize)
			}
		})
	}
}
