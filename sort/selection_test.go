package sort_test

import (
	"algo/sort"
	"slices"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	testCases := []SortTestCase[int]{
		{
			Input:  []int{2, 4, 3, 1, 5, 6},
			Output: []int{1, 2, 3, 4, 5, 6},
		},
		{
			Input:  []int{3, 1, 4, 1, 5, 9, 2, 6},
			Output: []int{1, 1, 2, 3, 4, 5, 6, 9},
		},
	}

	for _, testCase := range testCases {
		original := make([]int, len(testCase.Input))
		copy(original, testCase.Input)

		sort.SelectionSort(testCase.Input)

		if !slices.Equal(testCase.Input, testCase.Output) {
			t.Errorf(
				`Slice: %v is not equal to: %v, original is: %v`,
				testCase.Input,
				testCase.Output,
				original,
			)
		}
	}
}
