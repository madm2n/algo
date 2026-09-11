package search_test

import (
	"algo/search"
	"testing"
)

func TestQuickSelect(t *testing.T) {
	testCases := []SearchTestCase{
		{
			Input:  []int{3, 2, 1, 5, 4},
			Target: 0,
			Output: 1,
		},
		{
			Input:  []int{3, 2, 1, 5, 4},
			Target: 4,
			Output: 5,
		},
		{
			Input:  []int{3, 2, 1, 5, 4},
			Target: 2,
			Output: 3,
		},
		{
			Input:  []int{3, 2, 1, 5, 2},
			Target: 1,
			Output: 2,
		},
		{
			Input:  []int{42},
			Target: 0,
			Output: 42,
		},
	}

	for _, testCase := range testCases {
		result := search.QuickSelect(testCase.Input, testCase.Target)
		if result != testCase.Output {
			t.Errorf(
				`QuickSelect(%v, %d) = %d, want %d`,
				testCase.Input,
				testCase.Target,
				result,
				testCase.Output,
			)
		}
	}
}
