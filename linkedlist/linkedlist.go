package linkedlist

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

type SinglyLinkedList[T comparable] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

func New[T comparable]() *SinglyLinkedList[T] {
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
	for curr := l.head; curr != nil; curr = curr.Next {
		if curr.Value == value {
			return curr
		}
	}
	return nil
}

func (l *SinglyLinkedList[T]) Delete(value T) bool {
	if l.head == nil {
		return false
	}

	if l.head.Value == value {
		l.head = l.head.Next
		if l.head == nil {
			l.tail = nil
		}
		l.size--
		return true
	}

	prev := l.head
	for curr := l.head.Next; curr != nil; curr = curr.Next {
		if curr.Value == value {
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

func (l *SinglyLinkedList[T]) Values() []T {
	values := make([]T, 0, l.size)
	for curr := l.head; curr != nil; curr = curr.Next {
		values = append(values, curr.Value)
	}
	return values
}
