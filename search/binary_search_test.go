package search_test

import (
	"algo/search"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	testCases := []SearchTestCase{
		{
			Input:  []int{1, 2, 3, 4, 5, 6},
			Target: 1,
			Output: 0,
		},
		{
			Input:  []int{1, 2, 3, 4, 5, 6},
			Target: 6,
			Output: 5,
		},
		{
			Input:  []int{1, 2, 3, 4, 5, 6},
			Target: 4,
			Output: 3,
		},
		{
			Input:  []int{1, 3, 5, 7, 9},
			Target: 2,
			Output: -1,
		},
		{
			Input:  []int{},
			Target: 1,
			Output: -1,
		},
		{
			Input:  []int{10},
			Target: 10,
			Output: 0,
		},
		{
			Input:  []int{10},
			Target: 9,
			Output: -1,
		},
	}

	for _, testCase := range testCases {
		result := search.BinarySearch(testCase.Input, testCase.Target)
		if result != testCase.Output {
			t.Errorf(
				`BinarySearch(%v, %d) = %d, want %d`,
				testCase.Input,
				testCase.Target,
				result,
				testCase.Output,
			)
		}
	}
}
