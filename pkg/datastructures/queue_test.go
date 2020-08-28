package datastructures

import (
	"reflect"
	"testing"
)

func TestNewQueue(t *testing.T) {
	tests := []struct {
		name string
		want *Queue
	}{
		{
			name: "create a new empty queue",
			want: &Queue{&LinkedList{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewQueue(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_Enqueue(t *testing.T) {
	type args struct {
		data interface{}
	}
	tests := []struct {
		name     string
		initData []interface{}
		args     args
		wantSize int
	}{
		{
			name:     "add new node to empty queue",
			initData: []interface{}{},
			args:     args{data: 1},
			wantSize: 1,
		},
		{
			name:     "add new node to one-item queue",
			initData: []interface{}{1},
			args:     args{data: 2},
			wantSize: 2,
		},
		{
			name:     "add new node to two-item queue",
			initData: []interface{}{1, 2},
			args:     args{data: 3},
			wantSize: 3,
		},
		{
			name:     "add new node to populated queue",
			initData: []interface{}{1, 2, 3, 4},
			args:     args{data: 5},
			wantSize: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := buildLinkedList(tt.initData)
			q := &Queue{LinkedList: l}
			q.Enqueue(tt.args.data)

			if q.Tail.Data != tt.args.data {
				t.Errorf("Queue.Enqueue() Tail = %v, want = %v", q.Tail.Data, tt.args.data)
			}

			if q.Size != tt.wantSize {
				t.Errorf("Queue.Enqueue() Size = %v, want = %v", q.Size, tt.wantSize)
			}
		})
	}
}

func TestQueue_Dequeue(t *testing.T) {
	tests := []struct {
		name     string
		initData []interface{}
		want     interface{}
		wantSize int
	}{
		{
			name:     "dequeue from empty queue",
			initData: []interface{}{},
			want:     nil,
			wantSize: 0,
		},
		{
			name:     "dequeue from single-item queue",
			initData: []interface{}{1},
			want:     1,
			wantSize: 0,
		},
		{
			name:     "dequeue from two-item queue",
			initData: []interface{}{1, 2},
			want:     1,
			wantSize: 1,
		},
		{
			name:     "dequeue from populated queue",
			initData: []interface{}{1, 2, 3, 4},
			want:     1,
			wantSize: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := buildLinkedList(tt.initData)
			q := &Queue{LinkedList: l}

			if got := q.Dequeue(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Queue.Dequeue() data = %v, want = %v", got, tt.want)
			}

			if q.Size != tt.wantSize {
				t.Errorf("Queue.Dequeue() Size = %v, want = %v", q.Size, tt.wantSize)
			}
		})
	}
}

func TestQueue_Peek(t *testing.T) {
	tests := []struct {
		name     string
		initData []interface{}
		want     interface{}
		wantSize int
	}{
		{
			name:     "peek from empty queue",
			initData: []interface{}{},
			want:     nil,
			wantSize: 0,
		},
		{
			name:     "peek from populated queue",
			initData: []interface{}{1, 2, 3, 4},
			want:     1,
			wantSize: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := buildLinkedList(tt.initData)
			q := &Queue{LinkedList: l}

			if got := q.Peek(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Queue.Peek() data = %v, want = %v", got, tt.want)
			}

			if q.Size != tt.wantSize {
				t.Errorf("Queue.Peek() Size = %v, want = %v", q.Size, tt.wantSize)
			}
		})
	}
}
