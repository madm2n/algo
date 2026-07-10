package recurrences

// MatrixMultiplyNaive implements the naive O(n³) matrix multiplication algorithm.
func MatrixMultiplyNaive(a [][]int, b [][]int) [][]int {
	m := len(a)
	n := len(b[0])
	k := len(b)
	c := make([][]int, m)
	for i := range m {
		c[i] = make([]int, n)
		for j := range n {
			for l := range k {
				c[i][j] += a[i][l] * b[l][j]
			}
		}
	}
	return c
}
