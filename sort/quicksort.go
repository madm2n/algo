package sort

// QuickSort sorts the input slice in place.
func QuickSort(A []int) {
	p := 0
	r := len(A) - 1
	quickSort(A, p, r)
}

func quickSort(A []int, p, r int) {
	if p >= r {
		return
	}
	q := partition(A, p, r)

	quickSort(A, p, q-1)
	quickSort(A, q+1, r)
}

func partition(A []int, p, r int) int {
	x := A[r]
	i := p - 1

	for j := p; j < r; j++ {
		if A[j] <= x {
			i++
			A[i], A[j] = A[j], A[i]
		}
	}

	A[i+1], A[r] = A[r], A[i+1]

	return i + 1
}
