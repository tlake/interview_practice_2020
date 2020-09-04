# interview prep 2020

Solutions implemented with Golang 1.15.

## About

Welp, it's time for another job hunt!
That means I've gotta work the rust off of my interviewing skills and brush up on my fundamentals.
This repository will serve as a compilation of all the code-based preparation I do for this round of job acquisition.

## Implementations

For each implentation:

- add basic tests (like success cases)
- list non-basic tests (like error cases and edge cases)
- create a template/formulaic approach for using the implemenation during an interview
- narrate out loud what you're doing and why you're doing it
- explain how to implement them (verbally)
- explain the runtime and space complexity of each operation using big-O notation
- explain what each is useful for

### Data Structures

- [x] linked list
- [x] doubly-linked list
- [x] stack
- [x] queue
- [x] map
- [x] binary search tree
  - [ ] add balancing?
- [ ] graph
- [ ] heap
- [ ] array

### Algorithms

- [x] quicksort
- [x] merge sort
- [x] bubble sort
- [x] breadth-first search of a binary search tree
- [x] depth-first (pre-order) search of a binary search tree
- [x] breadth-first traversal of a binary search tree
- [x] in-order traversal of a binary search tree
- [x] pre-order traversal of a binary search tree
- [x] post-order traversal of a binary search tree
- [x] binary search on an array

### Concurrency

- [ ] blocking queue
- [ ] thread-safe versions of all the above data structures

### Other

- [x] LRU Cache
  - eviction policy is "least recently used"
  - operations: `Put(key, value)`, `Get(key)`, `Remove(key)`
  - the cache has a capacity. when cache is at cap and Put() is called with a new key, one element is removed/evicted from the cache according to the eviction policy.
- [ ] Given an in-order traversal and a pre-order traversal of a tree, rebuild the tree
- [ ] a DNS service
