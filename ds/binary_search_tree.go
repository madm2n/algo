package ds

import (
	"cmp"
)

// BinarySearchTree is a binary search tree of ordered values.
//
// Duplicate values are not stored more than once.
type BinarySearchTree[T cmp.Ordered] struct {
	root *binarySearchTreeNode[T]
	size int
}

type binarySearchTreeNode[T cmp.Ordered] struct {
	value T
	left  *binarySearchTreeNode[T]
	right *binarySearchTreeNode[T]
}

// NewBinarySearchTree returns an empty binary search tree.
func NewBinarySearchTree[T cmp.Ordered]() *BinarySearchTree[T] {
	return new(BinarySearchTree[T])
}

// Insert adds value to the tree.
func (t *BinarySearchTree[T]) Insert(value T) {
	if t.root == nil {
		t.root = t.newNode(value)
	} else {
		t.insert(t.root, value)
	}
}

func (t *BinarySearchTree[T]) newNode(value T) *binarySearchTreeNode[T] {
	t.size += 1
	return &binarySearchTreeNode[T]{
		value: value,
	}
}

func (t *BinarySearchTree[T]) insert(node *binarySearchTreeNode[T], value T) {
	if node.value == value {
		return
	}

	if node.value <= value {
		if node.right == nil {
			node.right = t.newNode(value)
		} else {
			t.insert(node.right, value)
		}
	} else if node.value >= value {
		if node.left == nil {
			node.left = t.newNode(value)
		} else {
			t.insert(node.left, value)
		}
	}
}

// Search reports whether value is present in the tree.
func (t *BinarySearchTree[T]) Search(value T) bool {
	return t.search(t.root, value)
}

func (t *BinarySearchTree[T]) search(node *binarySearchTreeNode[T], value T) bool {
	if node == nil {
		return false
	}

	if node.value == value {
		return true
	}

	if node.value <= value {
		return t.search(node.right, value)
	}

	if node.value >= value {
		return t.search(node.left, value)
	}

	return false
}

// Delete removes value from the tree if it is present.
func (t *BinarySearchTree[T]) Delete(value T) {
	panic("not implemented")
}

// InOrder returns the values in sorted order.
func (t *BinarySearchTree[T]) InOrder() []T {
	panic("not implemented")
}

// IsEmpty reports whether the tree contains no values.
func (t *BinarySearchTree[T]) IsEmpty() bool {
	return t.size == 0
}

// Size returns the number of values in the tree.
func (t *BinarySearchTree[T]) Size() int {
	return t.size
}
