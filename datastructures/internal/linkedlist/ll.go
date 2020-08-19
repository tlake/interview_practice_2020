// Package linkedlist implements a Linked List.
package linkedlist

import (
	"fmt"
	"strings"
)

// LinkedList implements a linked list where the items are composed of LLNode structs.
type LinkedList struct {
	head   *LLNode
	tail   *LLNode
	length int
}

// LLNode implements the individual elements managed by a LinkedList struct.
type LLNode struct {
	value string
	next  *LLNode
}

// Display returns a string representation of the linked list.
func (l LinkedList) Display() string {
	s := make([]string, l.length)
	currentNode := l.head

	for i := 0; i < l.length; i++ {
		s[i] = currentNode.value
		currentNode = currentNode.next
	}

	return strings.Join(s, ", ")
}

// ValueAt returns the value of the node at the given index.
func (l LinkedList) ValueAt(index int) (string, error) {
	if index >= l.length {
		return "", fmt.Errorf("Index %d out of range", index)
	}

	var val string
	node := l.head

	for i := 0; i < l.length; i++ {
		if i == index {
			val = node.value
			break
		}

		node = node.next
	}

	return val, nil
}

// Search looks through the linked list for the given value.
// If found, it returns the first node containing that value.
func (l *LinkedList) Search(value string) (*LLNode, error) {
	if l.head.value == value {
		return l.head, nil
	}

	if l.tail.value == value {
		return l.tail, nil
	}

	node := l.head
	for l.tail != node {
		if node.value == value {
			return node, nil
		}
		node = node.next
	}

	return nil, fmt.Errorf("Could not find %v in list", value)
}

// Prepend adds a new node with the given value to the start of the linked list.
func (l *LinkedList) Prepend(value string) *LLNode {
	newNode := &LLNode{value: value, next: l.head}
	l.head = newNode

	if l.tail == nil {
		l.tail = newNode
	}

	l.length++
	return l.head
}

// Append adds a new node with the given value to the end of the linked list.
func (l *LinkedList) Append(value string) *LLNode {
	newNode := &LLNode{value: value}
	if l.tail == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.next = newNode
		l.tail = newNode
	}

	l.length++
	return l.tail
}

// Insert adds a new node with the given value at the given index.
func (l *LinkedList) Insert(value string, index int) (*LLNode, error) {
	switch {
	case index > l.length:
		return nil, fmt.Errorf("Index %d out of range", index)

	case index == 0:
		return l.Prepend(value), nil

	case index == l.length:
		return l.Append(value), nil

	default:
		newNode := &LLNode{value: value}
		currentNode, nextNode := l.head, l.head.next
		for i := 0; i < l.length; i++ {
			if i+1 == index {
				newNode.next = nextNode
				currentNode.next = newNode
				break
			}

			currentNode = nextNode
			nextNode = nextNode.next
		}

		l.length++

		return newNode, nil
	}
}

// Shift removes the first node from the linked list and returns it.
func (l *LinkedList) Shift() *LLNode {
	var shiftedNode *LLNode

	if l.head == nil {
		return nil
	}

	shiftedNode = l.head
	if l.head == l.tail {
		l.head, l.tail = nil, nil
	} else {
		l.head = l.head.next
	}

	l.length--

	return shiftedNode
}

// Pop removes the last node from the linked list and returns it.
func (l *LinkedList) Pop() *LLNode {
	if l.tail == nil {
		return nil
	}

	poppedNode := l.tail

	if l.head == l.tail {
		l.head, l.tail = nil, nil
	} else {
		currentNode := l.head
		for i := 0; i < l.length-1; i++ {
			if currentNode.next == l.tail {
				l.tail = currentNode
				l.tail.next = nil
			}
			currentNode = currentNode.next
		}
	}

	l.length--

	return poppedNode
}

// Remove removes the node at the given index from the linked list.
func (l *LinkedList) Remove(index int) (*LLNode, error) {
	switch {
	case index > l.length-1:
		return nil, fmt.Errorf("Index %d out of range", index)

	case index == 0:
		return l.Shift(), nil

	case index == l.length-1:
		return l.Pop(), nil

	default:
		var removedNode *LLNode
		currentNode, nextNode := l.head, l.head.next
		for i := 0; i < l.length-1; i++ {
			if i+1 == index {
				removedNode = nextNode
				currentNode.next = nextNode.next
				break
			}

			currentNode = nextNode
			nextNode = nextNode.next
		}

		l.length--

		return removedNode, nil
	}
}

// Delete removes the first node with the given value from the linked list.
func (l *LinkedList) Delete(value string) *LLNode {
	var deletedNode *LLNode

	switch l.length {
	case 0:
		return deletedNode
	case 1:
		if l.head.value == value {
			deletedNode = l.head
			l.head, l.tail = nil, nil
		}
	default:
		if l.head.value == value {
			deletedNode = l.head
			l.head = l.head.next
		} else {
			currentNode, nextNode := l.head, l.head.next
			for nextNode != nil {
				if nextNode.value == value {
					deletedNode = nextNode
					currentNode.next = nextNode.next
					break
				}

				currentNode = nextNode
				nextNode = nextNode.next
			}

			if deletedNode == l.tail {
				l.tail = currentNode
			}
		}
	}

	if deletedNode != nil {
		l.length--
	}

	return deletedNode
}
