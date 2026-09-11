package search

// QuickSelect returns the k-th smallest element of the input slice,
// where k is zero-based. The input slice is modified in place.
func QuickSelect(A []int, k int) int {
	p := 0
	r := len(A) - 1
	return quickSelect(A, p, r, k)
}

func quickSelect(A []int, p, r, k int) int {
	if p == r {
		return A[p]
	}

	q := partition(A, p, r)

	if k == q {
		return A[q]
	}

	if k < q {
		return quickSelect(A, p, q-1, k)
	}

	return quickSelect(A, q+1, r, k)
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
