package algorithms

/*
algorithm:
	- if len(s) == 0, return -1
	- check the value at the midpoint of the slice
	- if that's the value we're looking for, return it
	- if len(s) == 1, return -1
	- if the value we want is less than the value at the midpoint, recurse on the left half
	- if the value we want is greater than the value at the midpoint, recurse on the right half
*/

/*
time complexity: O(log n) - every iteration cuts the problem space in half

space complexity: O(log n) - every recursion creates a new frame, but the recursion happens
	a very small number of times = time complexity
*/

// BinarySearch searches a given slice s (the slice must already be sorted, in ascending
// order) for a given value x by constantly halving the problem space. If the value is
// found, it returns the value; if not, it returns -1.
func BinarySearch(s []int, x int) int {
	if len(s) == 0 {
		return -1
	}

	midpoint := len(s) / 2
	if s[midpoint] == x {
		return x
	}

	if len(s) == 1 {
		return -1
	}

	if x < s[midpoint] {
		return BinarySearch(s[:midpoint], x)
	}
	return BinarySearch(s[midpoint:], x)
}
