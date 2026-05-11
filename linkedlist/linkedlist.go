package linkedlist

import (
	"reflect"
	"sync"
)

type node[T any] struct {
	Value T
	Next  *node[T]
}

type SinglyLinkedList[T any] struct {
	mu   sync.RWMutex
	head *node[T]
	tail *node[T]
	size int
}

func New[T any]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{}
}

func (l *SinglyLinkedList[T]) IsEmpty() bool {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.size == 0
}

func (l *SinglyLinkedList[T]) Len() int {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.size
}

func (l *SinglyLinkedList[T]) Prepend(value T) {
	l.mu.Lock()
	defer l.mu.Unlock()

	node := &node[T]{Value: value, Next: l.head}
	l.head = node
	if l.tail == nil {
		l.tail = node
	}
	l.size++
}

func (l *SinglyLinkedList[T]) Append(value T) {
	l.mu.Lock()
	defer l.mu.Unlock()

	node := &node[T]{Value: value}
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

func (l *SinglyLinkedList[T]) Find(value T) (T, bool) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	return l.findFuncNoLock(func(current T) bool {
		return reflect.DeepEqual(current, value)
	})
}

func (l *SinglyLinkedList[T]) Delete(value T) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	return l.deleteFuncNoLock(func(current T) bool {
		return reflect.DeepEqual(current, value)
	})
}

func (l *SinglyLinkedList[T]) FindFunc(match func(T) bool) (T, bool) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	return l.findFuncNoLock(match)
}

func (l *SinglyLinkedList[T]) findFuncNoLock(match func(T) bool) (T, bool) {
	for curr := l.head; curr != nil; curr = curr.Next {
		if match(curr.Value) {
			return curr.Value, true
		}
	}
	var zero T
	return zero, false
}

func (l *SinglyLinkedList[T]) DeleteFunc(match func(T) bool) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	return l.deleteFuncNoLock(match)
}

func (l *SinglyLinkedList[T]) deleteFuncNoLock(match func(T) bool) bool {
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
	l.mu.RLock()
	defer l.mu.RUnlock()

	if l.head == nil {
		var zero T
		return zero, false
	}
	return l.head.Value, true
}

func (l *SinglyLinkedList[T]) PopFront() (T, bool) {
	l.mu.Lock()
	defer l.mu.Unlock()

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
	l.mu.RLock()
	defer l.mu.RUnlock()

	values := make([]T, 0, l.size)
	for curr := l.head; curr != nil; curr = curr.Next {
		values = append(values, curr.Value)
	}
	return values
}

// ForEach calls fn for each value in list order. If fn returns false, iteration stops.
func (l *SinglyLinkedList[T]) ForEach(fn func(T) bool) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	for curr := l.head; curr != nil; curr = curr.Next {
		if !fn(curr.Value) {
			return
		}
	}
}
