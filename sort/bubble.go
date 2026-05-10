package sort

// BubbleSort sorts the input slice in place
// using the bubble sort algorithm.
func BubbleSort(inp []int) {
	for i := 0; i < len(inp)-1; i++ {
		swapped := false

		for j := 0; j < len(inp)-i-1; j++ {
			if inp[j] > inp[j+1] {
				inp[j], inp[j+1] = inp[j+1], inp[j]
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}
}
