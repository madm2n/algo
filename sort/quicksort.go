package sort

import "cmp"

// QuickSort sorts the input slice in place.
func QuickSort[T cmp.Ordered](A []T) {
	p := 0
	r := len(A) - 1
	quickSort(A, p, r)
}

func quickSort[T cmp.Ordered](A []T, p, r int) {
	if p >= r {
		return
	}
	q := partition(A, p, r)

	quickSort(A, p, q-1)
	quickSort(A, q+1, r)
}

func partition[T cmp.Ordered](A []T, p, r int) int {
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
