package linkedlist

import "reflect"

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

type SinglyLinkedList[T any] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

func New[T any]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{}
}

func (l *SinglyLinkedList[T]) IsEmpty() bool {
	return l.size == 0
}

func (l *SinglyLinkedList[T]) Len() int {
	return l.size
}

func (l *SinglyLinkedList[T]) Prepend(value T) {
	node := &Node[T]{Value: value, Next: l.head}
	l.head = node
	if l.tail == nil {
		l.tail = node
	}
	l.size++
}

func (l *SinglyLinkedList[T]) Append(value T) {
	node := &Node[T]{Value: value}
	if l.tail == nil {
		l.head = node
		l.tail = node
		l.size++
		return
	}

	l.tail.Next = node
	l.tail = node
	l.size++
}

func (l *SinglyLinkedList[T]) Find(value T) *Node[T] {
	return l.FindFunc(func(current T) bool {
		return reflect.DeepEqual(current, value)
	})
}

func (l *SinglyLinkedList[T]) Delete(value T) bool {
	return l.DeleteFunc(func(current T) bool {
		return reflect.DeepEqual(current, value)
	})
}

func (l *SinglyLinkedList[T]) FindFunc(match func(T) bool) *Node[T] {
	for curr := l.head; curr != nil; curr = curr.Next {
		if match(curr.Value) {
			return curr
		}
	}
	return nil
}

func (l *SinglyLinkedList[T]) DeleteFunc(match func(T) bool) bool {
	if l.head == nil {
		return false
	}

	if match(l.head.Value) {
		l.head = l.head.Next
		if l.head == nil {
			l.tail = nil
		}
		l.size--
		return true
	}

	prev := l.head
	for curr := l.head.Next; curr != nil; curr = curr.Next {
		if match(curr.Value) {
			prev.Next = curr.Next
			if curr == l.tail {
				l.tail = prev
			}
			l.size--
			return true
		}
		prev = curr
	}

	return false
}

func (l *SinglyLinkedList[T]) Front() (T, bool) {
	if l.head == nil {
		var zero T
		return zero, false
	}
	return l.head.Value, true
}

func (l *SinglyLinkedList[T]) PopFront() (T, bool) {
	if l.head == nil {
		var zero T
		return zero, false
	}

	value := l.head.Value
	l.head = l.head.Next
	if l.head == nil {
		l.tail = nil
	}
	l.size--
	return value, true
}

func (l *SinglyLinkedList[T]) Values() []T {
	values := make([]T, 0, l.size)
	for curr := l.head; curr != nil; curr = curr.Next {
		values = append(values, curr.Value)
	}
	return values
}
