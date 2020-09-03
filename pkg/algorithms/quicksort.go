package algorithms

/*
	algorithm:
	- check exit condition len(slice) < 2
	- pick a pivot
	- partition slice into three parts: less than pivot, pivot itself, greater than pivot
		- swap the pivot to end of slice
		- two pointers left, right = 0, 0
		- for right < len(slice)-1
			- as long as right item >= pivot, right++
			- if right < pivot
				- swap right and left
				- left++, right++
		- once all the swapping is finished, swap the pivot with
		  the element at left (= the first item larger than the pivot)
	- quicksort the left and right partitions
	  we'll just keep passing in smaller and smaller slices, which are just views
	  into the underlying inputSlice and as such will be modifying the original array
*/

/*
time complexity: O(n log n) on average, O(n^2) in the worst case
space complexity: O(log n)
	we're creating new slices with every recursive call to quicksort(), but this fits a visualization
	of a depth-first traversal, so "branches" of the algorithm finish before all "branches" are loaded.
	(plus, slices are super lightweight.)
*/

// Quicksort sorts the given input slice of ints using the quicksort algorithm.
func Quicksort(s []int) {
	if len(s) < 2 {
		return
	}

	pivotIdx := partition(s, len(s)/2)
	Quicksort(s[:pivotIdx])
	Quicksort(s[pivotIdx+1:])
}

func partition(s []int, pivotIdx int) int {
	pivotValue := s[pivotIdx]
	swapAtIndices(s, pivotIdx, len(s)-1)
	l, r := 0, 0
	for r < len(s)-1 {
		if s[r] < pivotValue {
			swapAtIndices(s, l, r)
			l++
		}
		r++
	}
	swapAtIndices(s, l, len(s)-1)
	return l
}

func swapAtIndices(s []int, a, b int) {
	s[a], s[b] = s[b], s[a]
}
