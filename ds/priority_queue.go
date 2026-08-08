package ds

import (
	"cmp"
	"errors"
)

// ErrHeapUnderflow is returned when an element is requested from an empty queue.
var ErrHeapUnderflow = errors.New("heap underflow")

// NewPriorityQueue returns an empty max-priority queue.
func NewPriorityQueue[T cmp.Ordered]() *PriorityQueue[T] {
	return new(PriorityQueue[T])
}

// PriorityQueue is a max-heap backed by a slice.
type PriorityQueue[T cmp.Ordered] struct {
	data []T
}

// Maximum returns the largest element in the queue without removing it.
func (p *PriorityQueue[T]) Maximum() (T, error) {
	if p.IsEmpty() {
		var zero T
		return zero, ErrHeapUnderflow
	}

	return p.data[0], nil
}

// ExtractMax removes and returns the largest element in the queue.
func (p *PriorityQueue[T]) ExtractMax() (T, error) {
	if p.IsEmpty() {
		var zero T
		return zero, ErrHeapUnderflow
	}

	max := p.data[0]
	p.data[0] = p.data[len(p.data)-1]
	p.data = p.data[:len(p.data)-1]
	p.maxHeapify(0)

	return max, nil
}

// Insert adds an element to the queue.
func (p *PriorityQueue[T]) Insert(v T) {
	p.data = append(p.data, v)
	p.heapifyUp(len(p.data) - 1)
}

// IsEmpty reports whether the queue contains no elements.
func (p *PriorityQueue[T]) IsEmpty() bool {
	return len(p.data) == 0
}

// Size returns the number of elements in the queue.
func (p *PriorityQueue[T]) Size() int {
	return len(p.data)
}

func (p *PriorityQueue[T]) maxHeapify(i int) {
	largest := i

	if left := p.leftChild(i); left < len(p.data) && p.data[left] > p.data[largest] {
		largest = left
	}

	if right := p.rightChild(i); right < len(p.data) && p.data[right] > p.data[largest] {
		largest = right
	}

	if largest != i {
		p.data[i], p.data[largest] = p.data[largest], p.data[i]
		p.maxHeapify(largest)
	}
}

func (p *PriorityQueue[T]) heapifyUp(i int) {
	for i > 0 && p.data[i] > p.data[p.parent(i)] {
		p.data[i], p.data[p.parent(i)] = p.data[p.parent(i)], p.data[i]
		i = p.parent(i)
	}
}

func (p *PriorityQueue[T]) parent(i int) int {
	return (i - 1) / 2
}

func (p *PriorityQueue[T]) leftChild(i int) int {
	return 2*i + 1
}

func (p *PriorityQueue[T]) rightChild(i int) int {
	return 2*i + 2
}
