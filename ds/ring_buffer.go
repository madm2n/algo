package ds

import "errors"

var (
	// ErrBufferEmpty is returned when an element is read from an empty buffer.
	ErrBufferEmpty = errors.New("buffer is empty")

	// ErrBufferFull is returned when an element is written to a full buffer.
	ErrBufferFull = errors.New("buffer is full")
)

// NewRingBuffer returns an empty ring buffer with the given capacity.
func NewRingBuffer[T any](capacity int) *RingBuffer[T] {
	return &RingBuffer[T]{
		data:     make([]T, capacity),
		capacity: capacity,
	}
}

// RingBuffer is a fixed-size circular buffer backed by a slice.
type RingBuffer[T any] struct {
	data     []T
	head     int
	tail     int
	size     int
	capacity int
}

// Write adds an element to the buffer.
func (r *RingBuffer[T]) Write(v T) error {
	if r.IsFull() {
		return ErrBufferFull
	}

	r.data[r.tail] = v
	r.tail = (r.tail + 1) % r.capacity
	r.size++

	return nil
}

// Read removes and returns the oldest element in the buffer.
func (r *RingBuffer[T]) Read() (T, error) {
	if r.IsEmpty() {
		var zero T
		return zero, ErrBufferEmpty
	}

	v := r.data[r.head]
	r.head = (r.head + 1) % r.capacity
	r.size--

	return v, nil
}

// IsEmpty reports whether the buffer contains no elements.
func (r *RingBuffer[T]) IsEmpty() bool {
	return r.size == 0
}

// IsFull reports whether the buffer cannot accept more elements.
func (r *RingBuffer[T]) IsFull() bool {
	return r.size == r.capacity
}

// Size returns the number of elements in the buffer.
func (r *RingBuffer[T]) Size() int {
	return r.size
}

// Capacity returns the maximum number of elements the buffer can hold.
func (r *RingBuffer[T]) Capacity() int {
	return r.capacity
}
