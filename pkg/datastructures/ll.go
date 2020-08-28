package datastructures

import (
	"fmt"
	"strings"
)

// LinkedList implements a linked list where the items are composed of LLNode structs.
type LinkedList struct {
	Head *LLNode
	Tail *LLNode
	Size int
}

// LLNode implements the individual elements managed by a LinkedList struct.
type LLNode struct {
	Data interface{}
	Next *LLNode
}

// NewLinkedList initializes an empty linked list and returns a pointer to it.
func NewLinkedList() *LinkedList {
	return &LinkedList{
		Head: nil,
		Tail: nil,
		Size: 0,
	}
}

// Display returns a string representation of the linked list.
func (l LinkedList) Display() string {
	s := make([]string, l.Size)
	currentNode := l.Head

	for i := 0; i < l.Size; i++ {
		s[i] = fmt.Sprintf("%v", currentNode.Data)
		currentNode = currentNode.Next
	}

	return strings.Join(s, ", ")
}

// ValueAt returns the data of the node at the given index.
func (l LinkedList) ValueAt(index int) (interface{}, error) {
	if index >= l.Size {
		return "", fmt.Errorf("Index %d out of range", index)
	}

	var val interface{}
	node := l.Head

	for i := 0; i < l.Size; i++ {
		if i == index {
			val = node.Data
			break
		}

		node = node.Next
	}

	return val, nil
}

// Search looks through the linked list for the given data.
// If found, it returns the first node containing that data.
func (l *LinkedList) Search(data interface{}) (*LLNode, error) {
	if l.Head.Data == data {
		return l.Head, nil
	}

	if l.Tail.Data == data {
		return l.Tail, nil
	}

	node := l.Head
	for l.Tail != node {
		if node.Data == data {
			return node, nil
		}
		node = node.Next
	}

	return nil, fmt.Errorf("Could not find %v in list", data)
}

// Prepend adds a new node with the given data to the start of the linked list.
func (l *LinkedList) Prepend(node *LLNode) {
	node.Next = l.Head
	l.Head = node

	if l.Tail == nil {
		l.Tail = node
	}

	l.Size++
}

// Append adds a new node with the given data to the end of the linked list.
func (l *LinkedList) Append(node *LLNode) {
	if l.Tail == nil {
		l.Head = node
		l.Tail = node
	} else {
		l.Tail.Next = node
		l.Tail = node
	}

	l.Size++
}

// Insert adds a new node with the given data at the given index.
func (l *LinkedList) Insert(node *LLNode, index int) error {
	switch {
	case index > l.Size:
		return fmt.Errorf("Index %d out of range", index)

	case index == 0:
		l.Prepend(node)
		return nil

	case index == l.Size:
		l.Append(node)
		return nil

	default:
		currentNode, NextNode := l.Head, l.Head.Next
		for i := 0; i < l.Size; i++ {
			if i+1 == index {
				node.Next = NextNode
				currentNode.Next = node
				break
			}

			currentNode = NextNode
			NextNode = NextNode.Next
		}

		l.Size++

		return nil
	}
}

// Shift removes the first node from the linked list and returns it.
func (l *LinkedList) Shift() *LLNode {
	var shiftedNode *LLNode

	if l.Head == nil {
		return nil
	}

	shiftedNode = l.Head
	if l.Head == l.Tail {
		l.Head, l.Tail = nil, nil
	} else {
		l.Head = l.Head.Next
	}

	shiftedNode.Next = nil
	l.Size--

	return shiftedNode
}

// Pop removes the last node from the linked list and returns it.
func (l *LinkedList) Pop() *LLNode {
	if l.Tail == nil {
		return nil
	}

	poppedNode := l.Tail

	if l.Head == l.Tail {
		l.Head, l.Tail = nil, nil
	} else {
		currentNode := l.Head
		for i := 0; i < l.Size-1; i++ {
			if currentNode.Next == l.Tail {
				l.Tail = currentNode
				l.Tail.Next = nil
			}
			currentNode = currentNode.Next
		}
	}

	l.Size--

	return poppedNode
}

// Remove removes the node at the given index from the linked list.
func (l *LinkedList) Remove(index int) (*LLNode, error) {
	switch {
	case index > l.Size-1:
		return nil, fmt.Errorf("Index %d out of range", index)

	case index == 0:
		return l.Shift(), nil

	case index == l.Size-1:
		return l.Pop(), nil

	default:
		var removedNode *LLNode
		currentNode, NextNode := l.Head, l.Head.Next
		for i := 0; i < l.Size-1; i++ {
			if i+1 == index {
				removedNode = NextNode
				currentNode.Next = NextNode.Next
				break
			}

			currentNode = NextNode
			NextNode = NextNode.Next
		}

		l.Size--

		return removedNode, nil
	}
}

// Delete removes the first node with the given data from the linked list.
func (l *LinkedList) Delete(data interface{}) *LLNode {
	var deletedNode *LLNode

	switch l.Size {
	case 0:
		return deletedNode
	case 1:
		if l.Head.Data == data {
			deletedNode = l.Head
			l.Head, l.Tail = nil, nil
		}
	default:
		if l.Head.Data == data {
			deletedNode = l.Head
			l.Head = l.Head.Next
		} else {
			currentNode, NextNode := l.Head, l.Head.Next
			for NextNode != nil {
				if NextNode.Data == data {
					deletedNode = NextNode
					currentNode.Next = NextNode.Next
					break
				}

				currentNode = NextNode
				NextNode = NextNode.Next
			}

			if deletedNode == l.Tail {
				l.Tail = currentNode
			}
		}
	}

	if deletedNode != nil {
		l.Size--
	}

	return deletedNode
}
