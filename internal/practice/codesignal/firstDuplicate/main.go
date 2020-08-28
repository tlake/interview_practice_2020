package main

func main() {}

func firstDuplicate(a []int) int {
	type Duplicate struct {
		Index int
		Value int
	}

	duplicate := Duplicate{Value: -1}
	seen := make(map[int]Duplicate)

	for i := 0; i < len(a); i++ {
		// if we've found at least one duplicate, we don't need to keep
		// looking through the array for new potential duplicates past
		// the index of a discovered duplicate
		if duplicate.Value >= 0 && i > duplicate.Index {
			break
		}

		value := a[i]

		if (seen[value] != Duplicate{}) {
			duplicate.Value = value
			duplicate.Index = i
		}

		seen[value] = Duplicate{Index: i, Value: value}
	}

	return duplicate.Value
}
