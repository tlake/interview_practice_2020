package cache

import (
	"fmt"

	"github.com/tlake/interview_prep_2020/pkg/datastructures"
)

/*
LRU Cache
(Least Recently Used)

The cache has a notion of a capacity. When the cache is at capacity and Put() is called
with a new key, one element is remove/evicted from the cache according to the eviction policy.

Eviction Policy: Evict the key which has been touched least recently.

Operations:
- Put(key, value)
- Get(key)
- Remove(key)

#######################

implmentation plan:

- store key in map, where value contains a pointer to a doubly-linked list node
- dll node contains value
- store access "time" by position in dll with tail pointer
	- every Put creates a new node at the head of the dll
	- every (successful) Get moves the accessed node to the head of the dll
	- removing least recently used is as simple as lopping off the tail of the dll
*/

// LRUCache implements a Least Recently Used cache.
type LRUCache struct {
	capacity int
	dll      *datastructures.DLL
	m        map[string]*datastructures.DLLNode
}

// kvPair implements a pairing of the key and value stored in the cache. This struct is what
// gets stored within the DLL, and is important during removal to eliminate the need to
// iterate over the map's keys to find the corresponding node when we remove it from the DLL.
type kvPair struct {
	key   string
	value interface{}
}

// NewLRUCache returns a pointer to an LRUCache struct, initialized with capacity = c,
// an initialized doubly-linked list, and an initialized map.
func NewLRUCache(c int) *LRUCache {
	return &LRUCache{
		capacity: c,
		dll:      datastructures.NewDLL(nil),
		m:        make(map[string]*datastructures.DLLNode, c),
	}
}

// Put creates a new entry in the cache, storing the key in the map and the value
// in a DLLNode at the head of the DLL. If the key already exists within the
// cache, it is overwritten with the new value. If the cache is at capacity, it
// removes the least recently used item from the cache before creating the new item.
func (l *LRUCache) Put(key string, value interface{}) error {
	kvp := &kvPair{key, value}

	// if exists, overwrite and return
	if _, ok := l.m[key]; ok {
		dllNode := l.m[key]
		dllNode.Data = kvp
		l.dll.Delete(dllNode)
		l.dll.PushNode(dllNode)
		return nil
	}

	// if at capacity, evict first
	if l.size() >= l.capacity {
		return l.evict()
	}

	// add new item
	dllNode := datastructures.NewDLLNode(kvp)
	l.dll.PushNode(dllNode)
	l.m[key] = dllNode
	return nil
}

// Get looks up the given key in the cache's map. If the key is found, it returns
// the value stored in the associated DLLNode, then moves that node to the head
// of the DLL. If the key is not found, it returns nil.
func (l *LRUCache) Get(key string) (interface{}, error) {
	if _, ok := l.m[key]; !ok {
		return nil, nil
	}

	dllNode := l.m[key]
	l.dll.Delete(dllNode)
	l.dll.PushNode(dllNode)

	kvp, ok := dllNode.Data.(*kvPair)
	if !ok {
		return nil, fmt.Errorf("could not convert %T to *kvPair", kvp)
	}
	return kvp.value, nil
}

// Remove looks for a given key in the cache's map. If the key is found, it removes
// the associated DLLNodes from the DLL, adjusting the DLL's structure as necessary.
// It then deletes the key from the map.
func (l *LRUCache) Remove(key string) error {
	if _, ok := l.m[key]; !ok {
		return nil
	}

	return l.delete(l.m[key])
}

// delete removes the given DLLNode from the DLL and its associated key from the map.
func (l *LRUCache) delete(dllNode *datastructures.DLLNode) error {
	kvp, ok := dllNode.Data.(*kvPair)
	if !ok {
		return fmt.Errorf("expected *cache.kvPair, got %v (type %T)", kvp, kvp)
	}

	l.dll.Delete(dllNode)
	delete(l.m, kvp.key)

	return nil
}

// evict calls l.delete() with the DLL's tail node.
func (l *LRUCache) evict() error {
	return l.delete(l.dll.Tail)
}

// size returns the length of the underlying DLL.
func (l *LRUCache) size() int {
	return l.dll.Len
}
