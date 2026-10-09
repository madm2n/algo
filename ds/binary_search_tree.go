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
	return t.search(t.root, value) != nil
}

func (t *BinarySearchTree[T]) search(node *binarySearchTreeNode[T], value T) *binarySearchTreeNode[T] {
	if node == nil {
		return nil
	}

	if node.value == value {
		return node
	}

	if node.value <= value {
		return t.search(node.right, value)
	}

	if node.value >= value {
		return t.search(node.left, value)
	}

	return nil
}

func (t *binarySearchTreeNode[T]) findNode(node, parent *binarySearchTreeNode[T], value T) (*binarySearchTreeNode[T], *binarySearchTreeNode[T]) {
	if node == nil {
		return nil, parent
	}

	if node.value == value {
		return node, parent
	}

	if node.value <= value {
		return t.findNode(node.right, node, value)
	}

	if node.value >= value {
		return t.findNode(node.left, node, value)
	}

	return nil, nil
}

// Delete removes value from the tree if it is present.
func (t *BinarySearchTree[T]) Delete(value T) {
	node, parent := t.root.findNode(t.root, nil, value)

	if node == nil {
		return
	}

	// Case 0: Deleting the leaf node.
	if node.left == nil && node.right == nil {
		if parent == nil {
			t.root = nil
		} else if parent.right == node {
			parent.right = nil
		} else if parent.left == node {
			parent.left = nil
		}

		t.size -= 1
		return
	}

	// Case 1: Deleating the node with a single child.
	deleteOneChildRoot := func() {
		if t.root.left != nil {
			t.root = t.root.left
		} else if t.root.right != nil {
			t.root = t.root.right
		}
	}

	if node.left == nil && node.right != nil {
		if parent == nil {
			deleteOneChildRoot()
		} else if parent.right == node {
			parent.right = node.right
		} else if parent.left == node {
			parent.left = node.right
		}

		t.size -= 1
		return
	}

	if node.left != nil && node.right == nil {
		if parent == nil {
			deleteOneChildRoot()
		} else if parent.right == node {
			parent.right = node.left
		} else if parent.left == node {
			parent.left = node.left
		}

		t.size -= 1
		return
	}

	// Case 2: Deleating a node with two children.
	if node.left != nil && node.right != nil {
		t.size -= 1
		return
	}
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
