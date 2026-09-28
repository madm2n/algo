package ds_test

import (
	"algo/ds"
	"errors"
	"slices"
	"testing"
)

type RingBufferTestCase[T any] struct {
	Name     string
	Capacity int
	Input    []T
	Output   []T
}

func TestWriteAndRead(t *testing.T) {
	testCases := []RingBufferTestCase[int]{
		{
			Name:     "single element",
			Capacity: 1,
			Input:    []int{5},
			Output:   []int{5},
		},
		{
			Name:     "fifo order",
			Capacity: 3,
			Input:    []int{1, 2, 3},
			Output:   []int{1, 2, 3},
		},
		{
			Name:     "negative values",
			Capacity: 3,
			Input:    []int{-3, -1, -2},
			Output:   []int{-3, -1, -2},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			buffer := ds.NewRingBuffer[int](testCase.Capacity)

			for _, v := range testCase.Input {
				if err := buffer.Write(v); err != nil {
					t.Fatalf("Write() returned unexpected error: %v", err)
				}
			}

			if buffer.Size() != len(testCase.Output) {
				t.Errorf("Size() = %d, want %d", buffer.Size(), len(testCase.Output))
			}

			got := make([]int, 0, len(testCase.Output))
			for range testCase.Output {
				v, err := buffer.Read()
				if err != nil {
					t.Fatalf("Read() returned unexpected error: %v", err)
				}
				got = append(got, v)
			}

			if !slices.Equal(got, testCase.Output) {
				t.Errorf(
					"Read() sequence = %v, want %v for input %v",
					got,
					testCase.Output,
					testCase.Input,
				)
			}
		})
	}
}

func TestWrapAround(t *testing.T) {
	buffer := ds.NewRingBuffer[int](3)

	for _, v := range []int{1, 2, 3} {
		if err := buffer.Write(v); err != nil {
			t.Fatalf("Write() returned unexpected error: %v", err)
		}
	}

	first, err := buffer.Read()
	if err != nil {
		t.Fatalf("Read() returned unexpected error: %v", err)
	}
	if first != 1 {
		t.Errorf("Read() = %d, want 1", first)
	}

	if err := buffer.Write(4); err != nil {
		t.Fatalf("Write() returned unexpected error: %v", err)
	}

	got := make([]int, 0, 3)
	for !buffer.IsEmpty() {
		v, err := buffer.Read()
		if err != nil {
			t.Fatalf("Read() returned unexpected error: %v", err)
		}
		got = append(got, v)
	}

	expected := []int{2, 3, 4}
	if !slices.Equal(got, expected) {
		t.Errorf("Read() sequence = %v, want %v", got, expected)
	}
}

func TestReadEmpty(t *testing.T) {
	buffer := ds.NewRingBuffer[int](1)

	_, err := buffer.Read()
	if !errors.Is(err, ds.ErrBufferEmpty) {
		t.Errorf("Read() error = %v, want %v", err, ds.ErrBufferEmpty)
	}
}

func TestWriteFull(t *testing.T) {
	buffer := ds.NewRingBuffer[int](2)

	for _, v := range []int{1, 2} {
		if err := buffer.Write(v); err != nil {
			t.Fatalf("Write() returned unexpected error: %v", err)
		}
	}

	if !buffer.IsFull() {
		t.Errorf("IsFull() = false, want true")
	}

	if err := buffer.Write(3); !errors.Is(err, ds.ErrBufferFull) {
		t.Errorf("Write() error = %v, want %v", err, ds.ErrBufferFull)
	}
}

func TestRingBufferIsEmpty(t *testing.T) {
	testCases := []struct {
		Name     string
		Input    []int
		Expected bool
	}{
		{Name: "empty buffer", Input: []int{}, Expected: true},
		{Name: "single element", Input: []int{1}, Expected: false},
		{Name: "multiple elements", Input: []int{1, 2, 3}, Expected: false},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			buffer := ds.NewRingBuffer[int](3)

			for _, v := range testCase.Input {
				if err := buffer.Write(v); err != nil {
					t.Fatalf("Write() returned unexpected error: %v", err)
				}
			}

			if got := buffer.IsEmpty(); got != testCase.Expected {
				t.Errorf("IsEmpty() = %t, want %t", got, testCase.Expected)
			}
		})
	}
}

func TestCapacity(t *testing.T) {
	buffer := ds.NewRingBuffer[int](4)

	if got := buffer.Capacity(); got != 4 {
		t.Errorf("Capacity() = %d, want 4", got)
	}
}

func TestRingBufferWithStrings(t *testing.T) {
	buffer := ds.NewRingBuffer[string](2)
	buffer.Write("apple")
	buffer.Write("cherry")

	first, err := buffer.Read()
	if err != nil {
		t.Fatalf("Read() returned unexpected error: %v", err)
	}
	if first != "apple" {
		t.Errorf("Read() = %q, want %q", first, "apple")
	}

	buffer.Write("banana")

	got := make([]string, 0, 2)
	for !buffer.IsEmpty() {
		v, err := buffer.Read()
		if err != nil {
			t.Fatalf("Read() returned unexpected error: %v", err)
		}
		got = append(got, v)
	}

	expected := []string{"cherry", "banana"}
	if !slices.Equal(got, expected) {
		t.Errorf("Read() sequence = %v, want %v", got, expected)
	}
}
