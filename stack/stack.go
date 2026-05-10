package stack

import "github.com/rexrecio/dasgo/linkedlist"

type Stack[T comparable] struct {
	list *linkedlist.SinglyLinkedList[T]
}

func New[T comparable]() *Stack[T] {
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
	values := s.list.Values()
	if len(values) == 0 {
		var zero T
		return zero, false
	}

	value := values[0]
	s.list.Delete(value)
	return value, true
}

func (s *Stack[T]) Peek() (T, bool) {
	values := s.list.Values()
	if len(values) == 0 {
		var zero T
		return zero, false
	}
	return values[0], true
}

func (s *Stack[T]) Values() []T {
	return s.list.Values()
}
