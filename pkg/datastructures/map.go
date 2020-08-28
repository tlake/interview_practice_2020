package datastructures

// Map implements a hash table. Entries is a slice of linked lists, and each (key, value)
// pair is stored at the head of a unique linked list. In the event of collisions, overflowing
// (key, value) pairs are stored further within the linked list.
// The HashFunc can be overridden by the user.
type Map struct {
	Entries  []LinkedList
	HashFunc func(string) int
	Size     int
}

// defaultHashFunc is the hash function that new Maps are initialized with. It's very simple,
// and just converts the bytes of each character in 'key' into ints and then sums those ints.
func defaultHashFunc(key string) int {
	var sum int
	for _, b := range key {
		sum += int(b)
	}
	return sum
}

// NewMap returns a pointer to a Map, initialized to a length of 32 to start.
func NewMap() *Map {
	e := make([]LinkedList, 32)
	m := &Map{
		Entries:  e,
		HashFunc: defaultHashFunc,
		Size:     0,
	}

	return m
}

// DoHash takes a string and runs the map's stored hash function on it to produce an int
// which can be used as the index for entry into the map's Entries slice.
func (m *Map) DoHash(key string) int {
	hashedResult := m.HashFunc(key)
	return hashedResult % len(m.Entries)
}
