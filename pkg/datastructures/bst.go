package datastructures

import (
	"fmt"
	"strings"
)

// BSTNode implements an individual element within a BST.
type BSTNode struct {
	data         int
	parent       *BSTNode
	greaterChild *BSTNode
	lesserChild  *BSTNode
}

// BST implements a binary search tree.
type BST struct {
	root *BSTNode
	size int
}

// NewBSTNode returns an new BST node initialized with the given data.
func NewBSTNode(data int) *BSTNode {
	return &BSTNode{data: data}
}

// NewBST returns a new binary search tree.
func NewBST() *BST {
	return &BST{}
}

// Insert creates a BST node with the given data and inserts the node into the tree.
func (b *BST) Insert(data int) {
	node := NewBSTNode(data)
	b.insert(node)
}

// insert is called by Insert, and does the work of adding the given node to the tree.
func (b *BST) insert(newNode *BSTNode) {
	if b.root == nil {
		b.root = newNode
		b.size = 1
		return
	}

	curr, next := b.root, b.root

	for next != nil {
		curr = next
		if newNode.data < curr.data {
			next = curr.lesserChild
			continue
		}
		next = curr.greaterChild
	}

	// at this point, we've reached a point where next is nil
	// and curr is the parent of that nil
	if newNode.data == curr.data {
		return
	}
	if newNode.data < curr.data {
		curr.lesserChild = newNode
		newNode.parent = curr
	}
	if curr.data < newNode.data {
		curr.greaterChild = newNode
		newNode.parent = curr
	}
	b.size++
}

// Search walks the tree in a binary search looking for the given data.
// If found, it returns a pointer to the node; if not, it returns nil.
func (b *BST) Search(data int) *BSTNode {
	curr := b.root

	for curr != nil {
		switch {
		case data < curr.data:
			curr = curr.lesserChild
		case data > curr.data:
			curr = curr.greaterChild
		default:
			return curr
		}
	}

	return nil
}

// BreadthFirstSearch walks the tree in a breadth-first traversal looking for the given data.
// If found, it returns a pointer to the node; if not, it returns nil.
func (b *BST) BreadthFirstSearch(data int) *BSTNode {
	q := NewQueue()
	q.Enqueue(b.root)
	for q.Peek() != nil {
		node := q.Pop().Data.(*BSTNode)
		if node.data == data {
			return node
		}
		if node.lesserChild != nil {
			q.Enqueue(node.lesserChild)
		}
		if node.greaterChild != nil {
			q.Enqueue(node.greaterChild)
		}
	}
	return nil
}

// DepthFirstSearch walks the tree in a pre-order traversal looking for the given data.
// If found, it returns a pointer to the node; if not, it returns nil.
func (b *BST) DepthFirstSearch(data int) *BSTNode {
	s := NewStack()
	s.Push(b.root)
	for s.Peek() != nil {
		node := s.Pop().(*BSTNode)
		if node.data == data {
			return node
		}
		if node.greaterChild != nil {
			s.Push(node.greaterChild)
		}
		if node.lesserChild != nil {
			s.Push(node.lesserChild)
		}
	}
	return nil
}

// Delete removes the given node from the tree.
func (b *BST) Delete(node *BSTNode) {
	switch {
	// no children
	case (node.lesserChild == nil && node.greaterChild == nil):
		p := node.parent
		if p.lesserChild == node {
			p.lesserChild = nil
		} else {
			p.greaterChild = nil
		}
		node.parent = nil
		b.size--
		return

	// one child
	case (node.lesserChild != nil && node.greaterChild == nil):
		if node.parent.lesserChild == node {
			node.parent.lesserChild = node.lesserChild
		} else {
			node.parent.greaterChild = node.lesserChild
		}
		node.lesserChild.parent = node.parent
		b.size--
		return

	case (node.greaterChild != nil && node.lesserChild == nil):
		if node.parent.lesserChild == node {
			node.parent.lesserChild = node.greaterChild
		} else {
			node.parent.greaterChild = node.greaterChild
		}
		node.greaterChild.parent = node.parent
		b.size--
		return

	// two children
	default:
		nextNode := b.getLeast(node.greaterChild)
		node.data = nextNode.data
		b.Delete(nextNode)
		return
	}
}

// BreadthFirst returns a string representation of a breadth-first traversal of the tree.
func (b *BST) BreadthFirst() string {
	nodeValues, i := make([]string, b.size), 0
	q := NewQueue()
	curr := b.root
	for curr != nil {
		nodeValues[i] = fmt.Sprintf("%d", curr.data)
		if curr.lesserChild != nil {
			q.Enqueue(curr.lesserChild)
		}
		if curr.greaterChild != nil {
			q.Enqueue(curr.greaterChild)
		}

		next := q.Dequeue()
		if next != nil {
			curr = next.(*BSTNode)
			i++
		} else {
			curr = nil
		}
	}

	return strings.Join(nodeValues, ", ")
}

// getLeast returns the least node that is a child of the given node
func (b *BST) getLeast(node *BSTNode) *BSTNode {
	if node == nil {
		return nil
	}
	for node.lesserChild != nil {
		node = node.lesserChild
	}
	return node
}

// InOrder returns a string representation of an in-order traversal of the tree.
// (lesser, root, greater)
func (b *BST) InOrder() string {
	getNext := func(input *BSTNode) *BSTNode {
		curr := input
		// if curr has a greater child, return that child's least terminus
		if curr.greaterChild != nil {
			curr = curr.greaterChild
			for curr.lesserChild != nil {
				curr = curr.lesserChild
			}
			return curr
		}
		// otherwise, return the first parent for which the traversal towards the root
		// approaches the parent from the lesser side (curr.parent.lesserChild == curr)
		for curr.parent != nil {
			if curr.parent.lesserChild == curr {
				return curr.parent
			}
			curr = curr.parent
		}

		return nil
	}

	if b.root == nil {
		return ""
	}

	nodeValues, i := make([]string, b.size), 0
	curr := b.getLeast(b.root)

	for curr != nil {
		nodeValues[i] = fmt.Sprintf("%d", curr.data)
		curr = getNext(curr)
		i++
	}

	return strings.Join(nodeValues, ", ")
}

// PreOrder returns a string representation of a pre-order traversal of the tree.
// (root, lesser, greater)
func (b *BST) PreOrder() string {
	if b.root == nil {
		return ""
	}

	nodeValues, i := make([]string, b.size), 0
	s := NewStack()
	s.Push(b.root)

	for s.Peek() != nil {
		node := s.Pop().(*BSTNode)
		nodeValues[i] = fmt.Sprintf("%d", node.data)

		if node.greaterChild != nil {
			s.Push(node.greaterChild)
		}
		if node.lesserChild != nil {
			s.Push(node.lesserChild)
		}

		i++
	}

	return strings.Join(nodeValues, ", ")
}

// PostOrder returns a string representation of a post-order traversal of the tree.
// (lesser, greater, root)
func (b *BST) PostOrder() string {
	if b.root == nil {
		return ""
	}

	nodeValues, i := make([]string, b.size), 0
	s1, s2 := NewStack(), NewStack()
	s1.Push(b.root)

	for s1.Peek() != nil {
		node := s1.Pop().(*BSTNode)
		s2.Push(node)

		if node.lesserChild != nil {
			s1.Push(node.lesserChild)
		}
		if node.greaterChild != nil {
			s1.Push(node.greaterChild)
		}
	}

	for s2.Peek() != nil {
		node := s2.Pop().(*BSTNode)
		nodeValues[i] = fmt.Sprintf("%d", node.data)
		i++
	}

	return strings.Join(nodeValues, ", ")
}
