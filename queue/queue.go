package queue

import "github.com/rexrecio/dasgo/linkedlist"

type Queue[T any] struct {
	list *linkedlist.SinglyLinkedList[T]
}

func New[T any]() *Queue[T] {
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
	return q.list.PopFront()
}

func (q *Queue[T]) Peek() (T, bool) {
	return q.list.Front()
}

func (q *Queue[T]) Values() []T {
	return q.list.Values()
}
