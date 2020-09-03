package algorithms

/*
algorithm:
	- base case len(s) < 2: return s
	- divide s in half
	- mergesort each half
	- make copies of each half
	- overwrite the values in s one at a time choosing the lowest from the copies each time
*/

/*
time complexity: O(n log n) - we have to touch every element multiple times, but we only have
	to do it a binary-search-amount of times

space complexity: O(n) - the final completion of the merge makes a copy of all the elements
	in the slice
*/

// Mergesort sorts the given slice by combining sorted subslices. It modifies the original
// input, but is not an in-situ sort.
func Mergesort(s []int) {
	// base case
	if len(s) < 2 {
		return
	}

	// divide
	midpoint := len(s) / 2

	// sort
	Mergesort(s[:midpoint])
	Mergesort(s[midpoint:])

	// copy
	l, r := make([]int, len(s[:midpoint])), make([]int, len(s[midpoint:]))
	copy(l, s[:midpoint])
	copy(r, s[midpoint:])

	// merge
	for i := range s {
		switch {
		case len(l) == 0:
			s[i] = r[0]
			r = r[1:]
		case len(r) == 0:
			s[i] = l[0]
			l = l[1:]
		default:
			if r[0] < l[0] {
				s[i] = r[0]
				r = r[1:]
			} else {
				s[i] = l[0]
				l = l[1:]
			}
		}
	}
}
