package sort

// SelectionSort sorts the input slice in place using selection sort.
func SelectionSort(inp []int) {
	for i, cur := range inp {
		smi := i
		smv := cur

		for j := i; j < len(inp); j++ {
			cvl := inp[j]

			if cvl < smv {
				smv = cvl
				smi = j
			}
		}

		inp[smi] = cur
		inp[i] = smv
	}
}
