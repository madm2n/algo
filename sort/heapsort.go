package sort

// HeapSort implementation of heap sort
// for the slice of integers.
func HeapSort(a []int) []int {
	size := len(a)
	buildMaxHeap(a, size)

	for i := size - 1; i >= 0; i-- {
		a[0], a[i] = a[i], a[0]
		maxHeapify(a, 0, i)
	}

	return a
}

func buildMaxHeap(a []int, size int) {
	for i := size/2 - 1; i >= 0; i-- {
		maxHeapify(a, i, size)
	}
}

func maxHeapify(a []int, i int, size int) {
	left := left(i)
	right := right(i)
	largest := i

	if left < size && a[left] > a[i] {
		largest = left
	}

	if right < size && a[right] > a[largest] {
		largest = right
	}

	if largest != i {
		a[i], a[largest] = a[largest], a[i]
		maxHeapify(a, largest, size)
	}
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}
