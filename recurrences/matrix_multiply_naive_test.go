package recurrences_test

import (
	"algo/recurrences"
	"slices"
	"testing"
)

type MatrixMultiplyTestCase struct {
	A      [][]int
	B      [][]int
	Output [][]int
}

func TestMatrixMultiplyNaive(t *testing.T) {
	testCases := []MatrixMultiplyTestCase{
		{
			A:      [][]int{{2}},
			B:      [][]int{{3}},
			Output: [][]int{{6}},
		},
		{
			A:      [][]int{{1, 2}, {3, 4}},
			B:      [][]int{{5, 6}, {7, 8}},
			Output: [][]int{{19, 22}, {43, 50}},
		},
		{
			A:      [][]int{{1, 0}, {0, 1}},
			B:      [][]int{{5, 6}, {7, 8}},
			Output: [][]int{{5, 6}, {7, 8}},
		},
		{
			A:      [][]int{{0, 0}, {0, 0}},
			B:      [][]int{{5, 6}, {7, 8}},
			Output: [][]int{{0, 0}, {0, 0}},
		},
		{
			A:      [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			B:      [][]int{{9, 8, 7}, {6, 5, 4}, {3, 2, 1}},
			Output: [][]int{{30, 24, 18}, {84, 69, 54}, {138, 114, 90}},
		},
		{
			A:      [][]int{{1, 2, 3}, {4, 5, 6}},
			B:      [][]int{{7, 8}, {9, 10}, {11, 12}},
			Output: [][]int{{58, 64}, {139, 154}},
		},
		{
			A:      [][]int{{7, 8}, {9, 10}, {11, 12}},
			B:      [][]int{{1, 2, 3}, {4, 5, 6}},
			Output: [][]int{{39, 54, 69}, {49, 68, 87}, {59, 82, 105}},
		},
	}

	for _, testCase := range testCases {
		result := recurrences.MatrixMultiplyNaive(testCase.A, testCase.B)

		for i := range result {
			if !slices.Equal(result[i], testCase.Output[i]) {
				t.Errorf(
					`MatrixMultiplyNaive(%v, %v)[%d] = %v, want %v`,
					testCase.A,
					testCase.B,
					i,
					result,
					testCase.Output,
				)
				break
			}
		}
	}
}
