package datastructures

import (
	"reflect"
)

// Map implements a hash table. Entries is a slice of linked lists, and each (key, value)
// pair is stored at the head of a unique linked list. In the event of collisions, overflowing
// (key, value) pairs are stored further within the linked list.
// The hashFunc can be overridden by the user.
type Map struct {
	entries         []*LinkedList
	hashFunc        func(string) int
	resizeThreshold float32
}

// mapNode implements a (key, value) pair to be stored within a linked list node.
type mapNode struct {
	key string
	val interface{}
}

// NewMap returns a pointer to a map, initialized to a length of 32 to start.
func NewMap() *Map {
	e := make([]*LinkedList, 32)
	for i := range e {
		e[i] = NewLinkedList()
	}

	m := &Map{
		entries:         e,
		hashFunc:        defaultHashFunc,
		resizeThreshold: 0.75,
	}

	return m
}

// SetHashFunc replaces the map's hash function with the given function.
func (m *Map) SetHashFunc(hashFunc func(string) int) {
	m.hashFunc = hashFunc
}

// DoHash takes a string and runs the map's stored hash function on it to produce an int
// which can be used as the index for entry into the map's Entries slice.
func (m *Map) DoHash(key string) int {
	hashedResult := m.hashFunc(key)
	return hashedResult % len(m.entries)
}

func (m *Map) insertNode(mn *mapNode) {
	key := mn.key
	entriesIndex := m.DoHash(key)
	entryList := m.entries[entriesIndex] // entryList is *LinkedList
	ln := &LLNode{Data: mn}

	if entryList.Size > 0 { // if we have a hash collision
		curr := entryList.Head
		for curr != nil {
			if curr.Data.(*mapNode).key == key { // if the key already exists, overwrite it
				curr.Data = mn
				return
			}
			curr = curr.Next
		}
		entryList.Append(ln) // if the key doesn't already exist, just tack it onto the end

	} else { // if we don't have a hash collision
		entryList.Prepend(ln)
	}
}

// Insert stores the value at the hashed location of the key. If the key already exists
// within the map, the existing value will be overwritten with the new value.
func (m *Map) Insert(key string, value interface{}) {
	if m.needsResize() {
		m.doResize()
	}

	mn := &mapNode{key: key, val: value}
	m.insertNode(mn)
}

// Delete removes the given key and its associated value from the map.
func (m *Map) Delete(key string) {
	entriesIndex := m.DoHash(key)
	entryList := m.entries[entriesIndex] // entryList is *LinkedList

	switch entryList.Size {
	case 0:
		return
	case 1:
		mn := entryList.Head.Data.(*mapNode)
		if mn.key == key {
			_ = entryList.Shift()
			return
		}

	default:
		curr := entryList.Head
		for curr.Next != nil {
			check := curr.Next // check is *LLNode
			mn := check.Data.(*mapNode)
			if mn.key == key {
				if reflect.DeepEqual(entryList.Tail, check) {
					entryList.Tail = curr
				}
				curr.Next = check.Next
				check.Next = nil
				entryList.Size--
				return
			}

			curr = curr.Next
		}
	}
}

// Get returns the value stored in the map at the associated key.
func (m *Map) Get(key string) interface{} {
	entriesIndex := m.DoHash(key)
	entryList := m.entries[entriesIndex]

	switch entryList.Size {
	case 0:
		return nil
	default:
		curr := entryList.Head
		for curr != nil {
			mn := curr.Data.(*mapNode)
			if mn.key == key {
				return mn.val
			}
			curr = curr.Next
		}
		return nil
	}
}

// Size counts the number of k/v pairs in the map and returns that number.
func (m *Map) Size() int {
	var sum int
	for _, e := range m.entries {
		sum += e.Size
	}
	return sum
}

// defaultHashFunc is the hash function that new maps are initialized with. It's very simple,
// and just converts the bytes of each character in 'key' into ints and then sums those ints.
// If you don't like it, use m.SetHashFunc() to supply your own.
func defaultHashFunc(key string) int {
	var sum int
	for _, b := range key {
		sum += int(b)
	}
	return sum
}

// needsResize determines whether or not the map contains too many items to work efficiently
// with its given size.
func (m *Map) needsResize() bool {
	ratio := float32(m.Size()) / float32(len(m.entries))
	return bool(ratio > m.resizeThreshold)
}

// doResize creates a new entries slice of adequate size, replaces the old entries slice
// with the new one, and re-inserts the existing k/v pairs into the new entries slice.,
func (m *Map) doResize() {
	l := len(m.entries)
	existingEntries := m.entries

	resizedEntries := make([]*LinkedList, l*2)
	for i := range resizedEntries {
		resizedEntries[i] = NewLinkedList()
	}
	m.entries = resizedEntries

	for _, entryList := range existingEntries {
		curr := entryList.Head
		for curr != nil {
			mn := curr.Data.(*mapNode)
			m.insertNode(mn)

			curr = curr.Next
		}
	}
}
