package exercises_test

import (
	"algo/exercises"
	"testing"
)

type HasIntSumTestCase struct {
	Input  []int
	Sum    int
	Output bool
}

func TestHasIntSum(t *testing.T) {
	testCases := []HasIntSumTestCase{
		{
			Input:  []int{},
			Sum:    5,
			Output: false,
		},
		{
			Input:  []int{5},
			Sum:    10,
			Output: false,
		},
		{
			Input:  []int{3, 7},
			Sum:    10,
			Output: true,
		},
		{
			Input:  []int{1, 2, 3, 4},
			Sum:    7,
			Output: true,
		},
		{
			Input:  []int{1, 2, 3, 4},
			Sum:    8,
			Output: false,
		},
		{
			Input:  []int{5, 5},
			Sum:    10,
			Output: true,
		},
		{
			Input:  []int{-3, 1, 2, 4},
			Sum:    -2,
			Output: true,
		},
		{
			Input:  []int{0, 0, 1},
			Sum:    0,
			Output: true,
		},
	}

	for _, testCase := range testCases {
		result := exercises.HasIntSum(testCase.Input, testCase.Sum)
		if result != testCase.Output {
			t.Errorf(
				`HasIntSum(%v, %d) = %t, want %t`,
				testCase.Input,
				testCase.Sum,
				result,
				testCase.Output,
			)
		}
	}
}
