package search

// BinarySearch search a slice of integers using
// binary search algorithm.
func BinarySearch(inp []int, target int) int {
	start := 0
	end := len(inp)
	return binarySearch(inp, target, start, end)
}

func binarySearch(inp []int, target int, start, end int) int {
	if start >= end {
		return -1
	}

	idx := start + (end-start)/2
	val := inp[idx]

	if val == target {
		return idx
	}

	if target < val {
		return binarySearch(inp, target, start, idx)
	}

	return binarySearch(inp, target, idx+1, end)
}
