package ds_test

import (
	"algo/ds"
	"slices"
	"testing"
)

func TestBinarySearchTreeStartsEmpty(t *testing.T) {
	tree := ds.NewBinarySearchTree[int]()
	if !tree.IsEmpty() {
		t.Error("new tree should be empty")
	}
	if tree.Size() != 0 {
		t.Errorf("Size() = %d, want 0", tree.Size())
	}
	if got := tree.InOrder(); len(got) != 0 {
		t.Errorf("InOrder() = %v, want an empty slice", got)
	}
}

func TestBinarySearchTreeInsertAndSearch(t *testing.T) {
	tree := ds.NewBinarySearchTree[int]()
	for _, value := range []int{5, 3, 7, 2, 4, 6, 8} {
		tree.Insert(value)
	}
	for _, value := range []int{2, 3, 4, 5, 6, 7, 8} {
		if !tree.Search(value) {
			t.Errorf("Search(%d) = false, want true", value)
		}
	}
	if tree.Search(1) || tree.Search(9) {
		t.Error("Search() found a value that was not inserted")
	}
}

func TestBinarySearchTreeInOrder(t *testing.T) {
	tree := ds.NewBinarySearchTree[int]()
	for _, value := range []int{5, 3, 7, 2, 4, 6, 8} {
		tree.Insert(value)
	}

	want := []int{2, 3, 4, 5, 6, 7, 8}
	if got := tree.InOrder(); !slices.Equal(got, want) {
		t.Errorf("InOrder() = %v, want %v", got, want)
	}
}

func TestBinarySearchTreeIgnoresDuplicates(t *testing.T) {
	tree := ds.NewBinarySearchTree[int]()
	for _, value := range []int{5, 3, 5, 3, 5} {
		tree.Insert(value)
	}

	if tree.Size() != 2 {
		t.Errorf("Size() = %d, want 2", tree.Size())
	}
	if got, want := tree.InOrder(), []int{3, 5}; !slices.Equal(got, want) {
		t.Errorf("InOrder() = %v, want %v", got, want)
	}
}

func TestBinarySearchTreeDelete(t *testing.T) {
	testCases := []struct {
		name  string
		input []int
		value int
		want  []int
	}{
		{name: "leaf", input: []int{5, 3, 7, 2, 4, 6, 8}, value: 2, want: []int{3, 4, 5, 6, 7, 8}},
		{name: "one child", input: []int{5, 3, 7, 2, 8}, value: 3, want: []int{2, 5, 7, 8}},
		{name: "two children", input: []int{5, 3, 7, 2, 4, 6, 8}, value: 5, want: []int{2, 3, 4, 6, 7, 8}},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			tree := ds.NewBinarySearchTree[int]()
			for _, value := range testCase.input {
				tree.Insert(value)
			}

			tree.Delete(testCase.value)
			if got := tree.InOrder(); !slices.Equal(got, testCase.want) {
				t.Errorf("InOrder() after Delete(%d) = %v, want %v", testCase.value, got, testCase.want)
			}
			if tree.Search(testCase.value) {
				t.Errorf("Search(%d) = true after deletion", testCase.value)
			}
		})
	}
}
