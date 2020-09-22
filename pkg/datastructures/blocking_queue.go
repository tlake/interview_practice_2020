package datastructures

import (
	"sync"
	"time"
)

// BlockingQueue implements a blocking queue.
type BlockingQueue struct {
	*LinkedList
	limit int
	mutex sync.Mutex
}

// NewBlockingQueue returns a pointer to an initialized blocking queue.
func NewBlockingQueue(limit int) *BlockingQueue {
	l := NewLinkedList()
	return &BlockingQueue{
		l,
		limit,
		sync.Mutex{},
	}
}

// Enqueue creates a linked list node containing the supplied data and adds it to the
// tail of the list. If the queue is already at its limit, this method will block until
// an item is removed from the queue.
func (b *BlockingQueue) Enqueue(data interface{}) {
	for b.Size >= b.limit {
		time.Sleep(10 * time.Millisecond)
	}

	b.mutex.Lock()
	defer b.mutex.Unlock()

	b.Append((&LLNode{Data: data}))
}

// Dequeue removes the node at the head of the linked list and returns the data contained
// therein. If the queue is empty, this method will block until an item is added to the queue.
func (b *BlockingQueue) Dequeue() interface{} {
	for b.Size < 1 {
		time.Sleep(10 * time.Millisecond)
	}

	b.mutex.Lock()
	defer b.mutex.Unlock()

	node := b.Shift()
	if node == nil {
		return nil
	}
	return node.Data
}

// Peek returns the data contained by the node at the head of the queue without modifying
// the queue.
func (b *BlockingQueue) Peek() interface{} {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	node := b.Head
	if node == nil {
		return nil
	}

	return node.Data
}
