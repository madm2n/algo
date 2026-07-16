package sort_test

import (
	"algo/sort"
	"slices"
	"testing"
)

func TestHeapSort(t *testing.T) {
	testCases := []SortTestCase[int]{
		{
			Input:  []int{2, 4, 3, 1, 5, 6},
			Output: []int{1, 2, 3, 4, 5, 6},
		},
		{
			Input:  []int{3, 2, 1, 5, 2},
			Output: []int{1, 2, 2, 3, 5},
		},
		{
			Input:  []int{3, 1, 4, 1, 5, 9, 2, 6},
			Output: []int{1, 1, 2, 3, 4, 5, 6, 9},
		},
	}

	for _, testCase := range testCases {
		original := make([]int, len(testCase.Input))
		copy(original, testCase.Input)

		result := sort.HeapSort(testCase.Input)

		if !slices.Equal(result, testCase.Output) {
			t.Errorf(
				`Slice: %v is not equal to: %v, original is: %v`,
				result,
				testCase.Output,
				original,
			)
		}
	}
}
