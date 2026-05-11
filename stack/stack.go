// Package stack provides a generic, thread-safe LIFO stack backed by a singly-linked list.
// All public methods are safe for concurrent use by multiple goroutines.
package stack

import "github.com/rexrecio/dasgo/linkedlist"

type Stack[T any] struct {
	list *linkedlist.SinglyLinkedList[T]
}

func New[T any]() *Stack[T] {
	return &Stack[T]{
		list: linkedlist.New[T](),
	}
}

func (s *Stack[T]) IsEmpty() bool {
	return s.list.IsEmpty()
}

func (s *Stack[T]) Len() int {
	return s.list.Len()
}

func (s *Stack[T]) Push(value T) {
	s.list.Prepend(value)
}

func (s *Stack[T]) Pop() (T, bool) {
	return s.list.PopFront()
}

func (s *Stack[T]) Peek() (T, bool) {
	return s.list.Front()
}

func (s *Stack[T]) Values() []T {
	return s.list.Values()
}
