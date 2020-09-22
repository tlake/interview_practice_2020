package datastructures

import (
	"reflect"
	"sync"
	"testing"
)

func TestNewBlockingQueue(t *testing.T) {
	type args struct {
		limit int
	}
	tests := []struct {
		name string
		args args
		want *BlockingQueue
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBlockingQueue(tt.args.limit); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBlockingQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBlockingQueue_Enqueue(t *testing.T) {
	type args struct {
		data interface{}
	}
	tests := []struct {
		name  string
		setup func(*BlockingQueue)
		args  args
		want  interface{}
	}{
		{
			name: "enqueues",
			setup: func(b *BlockingQueue) {
				return
			},
			args: args{data: 1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBlockingQueue(2)
			tt.setup(b)
			b.Enqueue(tt.args.data)
			if tt.want != b.Peek() {
				t.Errorf("BlockingQueue.Enqueue() = %v, want = %v", b.Peek(), tt.want)
			}
		})
	}
}

func TestBlockingQueue_Dequeue(t *testing.T) {
	tests := []struct {
		name     string
		setup    func(*BlockingQueue)
		want     interface{}
		wantPeek interface{}
	}{
		{
			name: "dequeue",
			setup: func(b *BlockingQueue) {
				b.Enqueue(1)
				b.Enqueue(2)
			},
			want:     1,
			wantPeek: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBlockingQueue(2)
			tt.setup(b)
			if got := b.Dequeue(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BlockingQueue.Dequeue() = %v, want %v", got, tt.want)
			}
			if b.Peek() != tt.wantPeek {
				t.Errorf("BlockingQueue.Dequeue() Peek() = %v, want = %v", b.Peek(), tt.wantPeek)
			}
		})
	}
}

func TestBlockingQueue_Peek(t *testing.T) {
	type fields struct {
		LinkedList *LinkedList
		limit      int
		mutex      sync.Mutex
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BlockingQueue{
				LinkedList: tt.fields.LinkedList,
				limit:      tt.fields.limit,
				mutex:      tt.fields.mutex,
			}
			if got := b.Peek(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BlockingQueue.Peek() = %v, want %v", got, tt.want)
			}
		})
	}
}
