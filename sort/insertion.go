package sort

// InsertionSort sorts the input slice in place.
func InsertionSort(inp []int) {
	for i := 1; i < len(inp); i++ {
		cur := inp[i]
		j := i - 1

		for j >= 0 && inp[j] > cur {
			inp[j+1] = inp[j]
			j--
		}

		inp[j+1] = cur
	}
}
