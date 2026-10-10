package search

import "cmp"

// BinarySearch search a slice using
// binary search algorithm.
func BinarySearch[T cmp.Ordered](inp []T, target T) int {
	start := 0
	end := len(inp)
	return binarySearch(inp, target, start, end)
}

func binarySearch[T cmp.Ordered](inp []T, target T, start, end int) int {
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
