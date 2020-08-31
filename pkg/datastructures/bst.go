package datastructures

// BSTNode implements an individual element within a BST.
type BSTNode struct {
	data int
}

// BST implements a binary search tree.
type BST struct{}

// NewBSTNode returns an new BST node initialized with the given data.
func NewBSTNode(data int) *BSTNode {
	return &BSTNode{data}
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

// Search walks the tree looking for the given data. If found, Search returns the node;
// if not, Search returns nil.
func (b *BST) Search(data int) *BSTNode {
	return nil
}

// Delete removes the given node from the tree.
func (b *BST) Delete(node *BSTNode) {}

// insert is called by Insert, and does the work of adding the given node to the tree.
func (b *BST) insert(node *BSTNode) {}
