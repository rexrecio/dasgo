// Package stack provides a generic, thread-safe LIFO stack backed by a singly-linked list.
// All public methods are safe for concurrent use by multiple goroutines.
package stack

import "github.com/rexrecio/dasgo/linkedlist"

// Stack is a generic, thread-safe LIFO stack.
// All operations are safe for concurrent use.
type Stack[T any] struct {
	list *linkedlist.SinglyLinkedList[T]
}

// New returns an empty Stack.
func New[T any]() *Stack[T] {
	return &Stack[T]{
		list: linkedlist.New[T](),
	}
}

// IsEmpty reports whether the stack contains no elements.
func (s *Stack[T]) IsEmpty() bool {
	return s.list.IsEmpty()
}

// Len returns the number of elements in the stack.
func (s *Stack[T]) Len() int {
	return s.list.Len()
}

// Push adds value to the top of the stack.
func (s *Stack[T]) Push(value T) {
	s.list.Prepend(value)
}

// Pop removes and returns the top element and true.
// Returns the zero value and false if the stack is empty.
func (s *Stack[T]) Pop() (T, bool) {
	return s.list.PopFront()
}

// Peek returns the top element without removing it, and true.
// Returns the zero value and false if the stack is empty.
func (s *Stack[T]) Peek() (T, bool) {
	return s.list.Front()
}

// Values returns all elements in top-to-bottom order as a slice.
func (s *Stack[T]) Values() []T {
	return s.list.Values()
}
