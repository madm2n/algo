package exercises

// HasIntSum reports whether any two
// distinct elements sum to the target value.
func HasIntSum(inp []int, sum int) bool {
	lkp := map[int]int{}

	for i, val := range inp {
		lkp[val] = i
	}

	for i, val := range inp {
		com := sum - val

		if j, ok := lkp[com]; ok && j != i {
			return true
		}
	}

	return false
}
