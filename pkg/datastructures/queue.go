package datastructures

// Queue implements a FIFO queue using an embedded singly-linked list.
type Queue struct {
	*LinkedList
}

// NewQueue initializes a new queue and returns a pointer to it.
func NewQueue() *Queue {
	l := NewLinkedList()
	return &Queue{
		LinkedList: l,
	}
}

// Enqueue creates an LLNode containing the given data and appends it to the tail of the queue.
func (q *Queue) Enqueue(data interface{}) {
	q.Append(&LLNode{Data: data})
}

// Dequeue removes the node at the head of the queue and returns its containing data.
func (q *Queue) Dequeue() interface{} {
	node := q.Shift()
	if node == nil {
		return nil
	}

	return node.Data
}

// Peek returns the data value at the head of of the queue without modifying the queue.
func (q *Queue) Peek() interface{} {
	node := q.Head
	if node == nil {
		return nil
	}

	return q.Head.Data
}
