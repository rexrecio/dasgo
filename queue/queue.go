// Package queue provides a generic, thread-safe FIFO queue backed by a singly-linked list.
// All public methods are safe for concurrent use by multiple goroutines.
package queue

import "github.com/rexrecio/dasgo/linkedlist"

// Queue is a generic, thread-safe FIFO queue.
// All operations are safe for concurrent use.
type Queue[T any] struct {
	list *linkedlist.SinglyLinkedList[T]
}

// New returns an empty Queue.
func New[T any]() *Queue[T] {
	return &Queue[T]{
		list: linkedlist.New[T](),
	}
}

// IsEmpty reports whether the queue contains no elements.
func (q *Queue[T]) IsEmpty() bool {
	return q.list.IsEmpty()
}

// Len returns the number of elements in the queue.
func (q *Queue[T]) Len() int {
	return q.list.Len()
}

// Enqueue adds value to the back of the queue.
func (q *Queue[T]) Enqueue(value T) {
	q.list.Append(value)
}

// Dequeue removes and returns the front element and true.
// Returns the zero value and false if the queue is empty.
func (q *Queue[T]) Dequeue() (T, bool) {
	return q.list.PopFront()
}

// Peek returns the front element without removing it, and true.
// Returns the zero value and false if the queue is empty.
func (q *Queue[T]) Peek() (T, bool) {
	return q.list.Front()
}

// Values returns all elements in front-to-back order as a slice.
func (q *Queue[T]) Values() []T {
	return q.list.Values()
}
