package datastructures

// Stack implements a LIFO stack using an embedded linked list.
type Stack struct {
	*LinkedList
}

// NewStack initializes a new stack and returns a pointer to it.
func NewStack() *Stack {
	list := NewLinkedList()
	return &Stack{
		LinkedList: list,
	}
}

// Push creates a new node containing 'value' and adds it to the top of the stack.
func (s *Stack) Push(value interface{}) {
	newNode := &LLNode{Data: value}
	s.Prepend(newNode)
}

// Pop removes the top node of the stack and returns its data value.
func (s *Stack) Pop() interface{} {
	node := s.Shift()
	if node == nil {
		return nil
	}

	return node.Data
}

// Peek returns the data value of the top node of the stack without modifying the stack.
func (s *Stack) Peek() interface{} {
	node := s.Head
	if node == nil {
		return nil
	}

	return node.Data
}
