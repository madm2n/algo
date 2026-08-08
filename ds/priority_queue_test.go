package ds_test

import (
	"algo/ds"
	"errors"
	"slices"
	"testing"
)

type InsertTestCase[T any] struct {
	Name   string
	Input  []T
	Output []T
}

func TestInsertAndExtractMax(t *testing.T) {
	testCases := []InsertTestCase[int]{
		{
			Name:   "single element",
			Input:  []int{5},
			Output: []int{5},
		},
		{
			Name:   "sorted input",
			Input:  []int{1, 2, 3, 4, 5},
			Output: []int{5, 4, 3, 2, 1},
		},
		{
			Name:   "reverse sorted input",
			Input:  []int{5, 4, 3, 2, 1},
			Output: []int{5, 4, 3, 2, 1},
		},
		{
			Name:   "duplicate values",
			Input:  []int{3, 1, 4, 1, 5, 9, 2, 6, 5},
			Output: []int{9, 6, 5, 5, 4, 3, 2, 1, 1},
		},
		{
			Name:   "negative values",
			Input:  []int{-3, -1, -2},
			Output: []int{-1, -2, -3},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			queue := ds.NewPriorityQueue[int]()

			for _, v := range testCase.Input {
				queue.Insert(v)
			}

			if queue.Size() != len(testCase.Input) {
				t.Errorf("Size() = %d, want %d", queue.Size(), len(testCase.Input))
			}

			got := make([]int, 0, len(testCase.Input))
			for range testCase.Input {
				v, err := queue.ExtractMax()
				if err != nil {
					t.Fatalf("ExtractMax() returned unexpected error: %v", err)
				}
				got = append(got, v)
			}

			if !slices.Equal(got, testCase.Output) {
				t.Errorf(
					"ExtractMax() sequence = %v, want %v for input %v",
					got,
					testCase.Output,
					testCase.Input,
				)
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	testCases := []InsertTestCase[int]{
		{
			Name:   "single element",
			Input:  []int{5},
			Output: []int{5},
		},
		{
			Name:   "multiple elements",
			Input:  []int{3, 1, 4, 1, 5},
			Output: []int{5},
		},
		{
			Name:   "duplicate maximum",
			Input:  []int{7, 7, 3},
			Output: []int{7},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			queue := ds.NewPriorityQueue[int]()

			for _, v := range testCase.Input {
				queue.Insert(v)
			}

			got, err := queue.Maximum()
			if err != nil {
				t.Fatalf("Maximum() returned unexpected error: %v", err)
			}

			if got != testCase.Output[0] {
				t.Errorf("Maximum() = %d, want %d", got, testCase.Output[0])
			}

			if queue.Size() != len(testCase.Input) {
				t.Errorf(
					"Size() after Maximum() = %d, want %d",
					queue.Size(),
					len(testCase.Input),
				)
			}
		})
	}
}

func TestUnderflow(t *testing.T) {
	t.Run("Maximum", func(t *testing.T) {
		queue := ds.NewPriorityQueue[int]()

		_, err := queue.Maximum()
		if !errors.Is(err, ds.ErrHeapUnderflow) {
			t.Errorf("Maximum() error = %v, want %v", err, ds.ErrHeapUnderflow)
		}
	})

	t.Run("ExtractMax", func(t *testing.T) {
		queue := ds.NewPriorityQueue[int]()

		_, err := queue.ExtractMax()
		if !errors.Is(err, ds.ErrHeapUnderflow) {
			t.Errorf("ExtractMax() error = %v, want %v", err, ds.ErrHeapUnderflow)
		}
	})

	t.Run("after all extracted", func(t *testing.T) {
		queue := ds.NewPriorityQueue[int]()
		queue.Insert(1)
		queue.Insert(2)

		for queue.Size() > 0 {
			_, _ = queue.ExtractMax()
		}

		_, err := queue.ExtractMax()
		if !errors.Is(err, ds.ErrHeapUnderflow) {
			t.Errorf("ExtractMax() error = %v, want %v", err, ds.ErrHeapUnderflow)
		}
	})
}

func TestIsEmpty(t *testing.T) {
	testCases := []struct {
		Name     string
		Input    []int
		Expected bool
	}{
		{Name: "empty queue", Input: []int{}, Expected: true},
		{Name: "single element", Input: []int{1}, Expected: false},
		{Name: "multiple elements", Input: []int{1, 2, 3}, Expected: false},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			queue := ds.NewPriorityQueue[int]()

			for _, v := range testCase.Input {
				queue.Insert(v)
			}

			if got := queue.IsEmpty(); got != testCase.Expected {
				t.Errorf("IsEmpty() = %t, want %t", got, testCase.Expected)
			}
		})
	}
}

func TestPriorityQueueWithStrings(t *testing.T) {
	queue := ds.NewPriorityQueue[string]()
	queue.Insert("apple")
	queue.Insert("cherry")
	queue.Insert("banana")

	got := make([]string, 0, 3)
	for !queue.IsEmpty() {
		v, err := queue.ExtractMax()
		if err != nil {
			t.Fatalf("ExtractMax() returned unexpected error: %v", err)
		}
		got = append(got, v)
	}

	expected := []string{"cherry", "banana", "apple"}
	if !slices.Equal(got, expected) {
		t.Errorf("ExtractMax() sequence = %v, want %v", got, expected)
	}
}
