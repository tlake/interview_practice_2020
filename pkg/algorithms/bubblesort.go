package algorithms

/*
algorithm:
	- create a swap := true to track whether the slice is sorted
	- if swap = true
		- set swap = false
		- create a pair of pointers l, r := 0, 1
		- walk the pointers together across the slice
			- if ever s[l] > s[r], swap them, then set swap = true
	- when swap = false, the slice is sorted
*/

/*
time complexity: O(n^2) average and worst-case - bubblesort must traverse the slice a number
	of times roughly on-par with the number of items in the slice. in the event that the slice
	is already sorted, however, bubblesort has to walk the slice only once, providing a best-
	case performance of O(n).

space complexity: O(1) - bubblesort is an in-place method, and only creates three variables
	across its entire operation.
*/

// Bubblesort sorts the given slice by comparing pairs of items. It modifies the given
// input slice, and does the sort completely in-situ.
func Bubblesort(s []int) {
	swap := true
	for swap == true {
		swap = false
		l, r := 0, 1
		for r < len(s) {
			if s[l] > s[r] {
				s[l], s[r] = s[r], s[l]
				swap = true
			}
			l++
			r++
		}
	}
}
