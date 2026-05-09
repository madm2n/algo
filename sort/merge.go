package sort

// MergeSort sorts the input slice in place.
func MergeSort(A []int) {
	p := 0
	r := len(A) - 1
	mergeSort(A, p, r)
}

func mergeSort(A []int, p, r int) {
	if p >= r {
		return
	}
	q := (p + r) / 2

	mergeSort(A, p, q)
	mergeSort(A, q+1, r)
	merge(A, p, q, r)
}

func merge(A []int, p, q, r int) {
	lN := q - p + 1
	L := make([]int, lN)

	for i := range lN {
		L[i] = A[p+i]
	}

	rN := r - q
	R := make([]int, rN)
	for j := range rN {
		R[j] = A[q+j+1]
	}

	i, j, k := 0, 0, p

	for i < lN && j < rN {
		if L[i] <= R[j] {
			A[k] = L[i]
			i++
		} else {
			A[k] = R[j]
			j++
		}
		k++
	}

	for i < lN {
		A[k] = L[i]
		i++
		k++
	}

	for j < rN {
		A[k] = R[j]
		j++
		k++
	}
}
