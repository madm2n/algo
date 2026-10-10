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
	value  T
	left   *binarySearchTreeNode[T]
	right  *binarySearchTreeNode[T]
	parent *binarySearchTreeNode[T]
}

// NewBinarySearchTree returns an empty binary search tree.
func NewBinarySearchTree[T cmp.Ordered]() *BinarySearchTree[T] {
	return new(BinarySearchTree[T])
}

// Insert adds value to the tree.
func (t *BinarySearchTree[T]) Insert(value T) {
	if t.root == nil {
		t.root = t.newNode(nil, value)
	} else {
		t.insert(t.root, value)
	}
}

func (t *BinarySearchTree[T]) newNode(parent *binarySearchTreeNode[T], value T) *binarySearchTreeNode[T] {
	t.size += 1
	return &binarySearchTreeNode[T]{
		parent: parent,
		value:  value,
	}
}

func (t *BinarySearchTree[T]) insert(node *binarySearchTreeNode[T], value T) {
	if node.value == value {
		return
	}

	if node.value <= value {
		if node.right == nil {
			node.right = t.newNode(node, value)
		} else {
			t.insert(node.right, value)
		}
	} else if node.value >= value {
		if node.left == nil {
			node.left = t.newNode(node, value)
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
	curr := node

	for curr != nil && value != curr.value {
		if value < curr.value {
			curr = curr.left
		} else {
			curr = curr.right
		}
	}

	return curr
}

func (t *BinarySearchTree[T]) successor(node *binarySearchTreeNode[T]) *binarySearchTreeNode[T] {
	if node.right != nil {
		curr := node.right
		for curr.left != nil {
			curr = curr.left
		}
		return curr
	}

	curr := node
	for curr.parent != nil && curr == curr.parent.right {
		curr = curr.parent
	}
	return curr.parent
}

// Delete removes value from the tree if it is present.
func (t *BinarySearchTree[T]) Delete(value T) {
	node := t.search(t.root, value)

	if node == nil {
		return
	}

	// Case 0: Deleting the leaf node.
	if node.left == nil && node.right == nil {
		if node.parent == nil {
			t.root = nil
		} else if node.parent.right == node {
			node.parent.right = nil
		} else if node.parent.left == node {
			node.parent.left = nil
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
		t.root.parent = nil
	}

	if node.left == nil && node.right != nil {
		if node.parent == nil {
			deleteOneChildRoot()
		} else if node.parent.right == node {
			node.parent.right = node.right
			node.right.parent = node.parent
		} else if node.parent.left == node {
			node.parent.left = node.right
			node.right.parent = node.parent
		}

		t.size -= 1
		return
	}

	if node.left != nil && node.right == nil {
		if node.parent == nil {
			deleteOneChildRoot()
		} else if node.parent.right == node {
			node.parent.right = node.left
			node.left.parent = node.parent
		} else if node.parent.left == node {
			node.parent.left = node.left
			node.left.parent = node.parent
		}

		t.size -= 1
		return
	}

	// Case 2: Deleating a node with two children.
	if node.left != nil && node.right != nil {
		succ := t.successor(node)
		node.value = succ.value

		if succ.parent.left == succ {
			succ.parent.left = succ.right
		} else {
			succ.parent.right = succ.right
		}

		if succ.right != nil {
			succ.right.parent = succ.parent
		}

		t.size -= 1
		return
	}
}

// InOrder returns the values in sorted order.
func (t *BinarySearchTree[T]) InOrder() []T {
	values := make([]T, 0, t.size)

	if t.root == nil {
		return values
	}

	curr := t.root
	for curr.left != nil {
		curr = curr.left
	}

	for curr != nil {
		values = append(values, curr.value)
		curr = t.successor(curr)
	}

	return values
}

// IsEmpty reports whether the tree contains no values.
func (t *BinarySearchTree[T]) IsEmpty() bool {
	return t.size == 0
}

// Size returns the number of values in the tree.
func (t *BinarySearchTree[T]) Size() int {
	return t.size
}
