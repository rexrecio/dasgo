package queue

import "github.com/rexrecio/dasgo/linkedlist"

type Queue[T comparable] struct {
	list *linkedlist.SinglyLinkedList[T]
}

func New[T comparable]() *Queue[T] {
	return &Queue[T]{
		list: linkedlist.New[T](),
	}
}

func (q *Queue[T]) IsEmpty() bool {
	return q.list.IsEmpty()
}

func (q *Queue[T]) Len() int {
	return q.list.Len()
}

func (q *Queue[T]) Enqueue(value T) {
	q.list.Append(value)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	values := q.list.Values()
	if len(values) == 0 {
		var zero T
		return zero, false
	}

	value := values[0]
	q.list.Delete(value)
	return value, true
}

func (q *Queue[T]) Peek() (T, bool) {
	values := q.list.Values()
	if len(values) == 0 {
		var zero T
		return zero, false
	}
	return values[0], true
}

func (q *Queue[T]) Values() []T {
	return q.list.Values()
}
