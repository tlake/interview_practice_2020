package doublylinkedlist

import (
	"errors"
)

// DLLNode is a node in a doubly-linked list.
type DLLNode struct {
	Data interface{}
	Prev *DLLNode
	Next *DLLNode
}

// NewDLLNode creates a new DLLNode with the given data and returns a pointer to the node.
func NewDLLNode(data interface{}) *DLLNode {
	return &DLLNode{Data: data}
}

// DLL is a doubly-linked list.
type DLL struct {
	Head *DLLNode
	Tail *DLLNode
	Len  int
}

// NewDLL returns a points to a DLL initialized with a the given *DLLNode.
func NewDLL(node *DLLNode) *DLL {
	len := 0
	if node != nil {
		len = 1
	}

	return &DLL{
		Head: node,
		Tail: node,
		Len:  len,
	}
}

// Push creates a new node containing the given data at the front of the DLL.
func (l *DLL) Push(data interface{}) {
	node := NewDLLNode(data)

	if l.Head == nil {
		l.Head, l.Tail = node, node
	} else {
		node.Next = l.Head
		l.Head.Prev = node
		l.Head = node
	}

	l.Len++
}

// Append creates a new node containing the given data at the end of the DLL.
func (l *DLL) Append(data interface{}) {
	node := NewDLLNode(data)

	if l.Tail == nil {
		l.Head, l.Tail = node, node
	} else {
		node.Prev = l.Tail
		l.Tail.Next = node
		l.Tail = node
	}

	l.Len++
}

// InsertBefore creates a new node in the list situated before the given node.
func (l *DLL) InsertBefore(givenNode *DLLNode, data interface{}) error {
	if givenNode == nil {
		return errors.New("given node cannot be nil")
	}

	node := NewDLLNode(data)
	node.Prev = givenNode.Prev
	node.Next = givenNode

	if givenNode == l.Head {
		l.Head = node
	} else {
		givenNode.Prev.Next = node
	}
	givenNode.Prev = node

	l.Len++
	return nil
}

// InsertAfter creates a new node in the list situated after the given node.
func (l *DLL) InsertAfter(givenNode *DLLNode, data interface{}) error {
	if givenNode == nil {
		return errors.New("given node cannot be nil")
	}

	node := NewDLLNode(data)
	node.Next = givenNode.Next
	node.Prev = givenNode
	givenNode.Next = node

	if givenNode == l.Tail {
		l.Tail = node
	}

	l.Len++
	return nil
}

// Delete removes the given node from the doubly-linked list.
func (l *DLL) Delete(givenNode *DLLNode) error {
	if givenNode == nil {
		return errors.New("given node cannot be nil")
	}

	if l.Head == givenNode {
		l.Head = givenNode.Next
	}

	if l.Tail == givenNode {
		l.Tail = givenNode.Prev
	}

	prev, next := givenNode.Prev, givenNode.Next

	if prev != nil {
		prev.Next = next
	}

	if next != nil {
		next.Prev = prev
	}

	givenNode.Prev, givenNode.Next = nil, nil
	l.Len--

	return nil
}

// Find searches through the DLL and returns a pointer to the first node that contains
// the given data. If no such node is found, Find returns nil.
func (l *DLL) Find(data interface{}) *DLLNode {
	if l.Head.Data == data {
		return l.Head
	}

	curr := l.Head
	for curr.Next != nil {
		curr = curr.Next
		if curr.Data == data {
			return curr
		}
	}

	return nil
}
