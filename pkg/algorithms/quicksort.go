package algorithms

func Quicksort(array []int) []int {
	/*
		algorithm:
		- pick a pivot (last, first, middle, random, median-of-three, whatever)
		- two pointers, one at each end
		- do swaps:
			- while left pointer < right pointer:
			- move left pointer to the right until it finds an item > pivot
			- then move right pointer to the left until it finds an item < pivot
			- swap those two items, then restart the loop at moving the left pointer to the right
		- after swaps completed, execute the algorithm on the two sub-arrays: the array
		  from 0 to left pointer, and the array from right pointer to the end

		what's the base case? when do we not kick off a new round of recursion?
		- when the leftmost index pointer is not less than the rightmost index pointer
	*/

	getPivot := func(a []int) int {
		return len(a) / 2
	}

	p := getPivot(array)

	return array
}
