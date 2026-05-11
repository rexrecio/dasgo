package bst

import (
	"cmp"
	"sync"
)

type node[T cmp.Ordered] struct {
	Value T
	Left  *node[T]
	Right *node[T]
}

type BinarySearchTree[T cmp.Ordered] struct {
	mu   sync.RWMutex
	root *node[T]
	size int
}

func New[T cmp.Ordered]() *BinarySearchTree[T] {
	return &BinarySearchTree[T]{}
}

func (t *BinarySearchTree[T]) IsEmpty() bool {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.size == 0
}

func (t *BinarySearchTree[T]) Len() int {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.size
}

func (t *BinarySearchTree[T]) Insert(value T) bool {
	t.mu.Lock()
	defer t.mu.Unlock()

	if t.root == nil {
		t.root = &node[T]{Value: value}
		t.size++
		return true
	}

	curr := t.root
	for {
		if value < curr.Value {
			if curr.Left == nil {
				curr.Left = &node[T]{Value: value}
				t.size++
				return true
			}
			curr = curr.Left
			continue
		}

		if value > curr.Value {
			if curr.Right == nil {
				curr.Right = &node[T]{Value: value}
				t.size++
				return true
			}
			curr = curr.Right
			continue
		}

		return false
	}
}

func (t *BinarySearchTree[T]) Find(value T) (T, bool) {
	t.mu.RLock()
	defer t.mu.RUnlock()

	curr := t.root
	for curr != nil {
		if value < curr.Value {
			curr = curr.Left
			continue
		}
		if value > curr.Value {
			curr = curr.Right
			continue
		}
		return curr.Value, true
	}
	var zero T
	return zero, false
}

func (t *BinarySearchTree[T]) Delete(value T) bool {
	t.mu.Lock()
	defer t.mu.Unlock()

	var deleted bool
	t.root, deleted = deleteNode(t.root, value)
	if deleted {
		t.size--
	}
	return deleted
}

func deleteNode[T cmp.Ordered](node *node[T], value T) (*node[T], bool) {
	if node == nil {
		return nil, false
	}

	if value < node.Value {
		left, deleted := deleteNode(node.Left, value)
		node.Left = left
		return node, deleted
	}
	if value > node.Value {
		right, deleted := deleteNode(node.Right, value)
		node.Right = right
		return node, deleted
	}

	if node.Left == nil {
		return node.Right, true
	}
	if node.Right == nil {
		return node.Left, true
	}

	successor := node.Right
	for successor.Left != nil {
		successor = successor.Left
	}

	node.Value = successor.Value
	right, _ := deleteNode(node.Right, successor.Value)
	node.Right = right
	return node, true
}

func (t *BinarySearchTree[T]) Values() []T {
	t.mu.RLock()
	defer t.mu.RUnlock()

	values := make([]T, 0, t.size)
	var walk func(node *node[T])
	walk = func(node *node[T]) {
		if node == nil {
			return
		}
		walk(node.Left)
		values = append(values, node.Value)
		walk(node.Right)
	}
	walk(t.root)
	return values
}

// ForEach calls fn for each value in ascending order. If fn returns false, iteration stops.
func (t *BinarySearchTree[T]) ForEach(fn func(T) bool) {
	t.mu.RLock()
	defer t.mu.RUnlock()

	var walk func(n *node[T]) bool
	walk = func(n *node[T]) bool {
		if n == nil {
			return true
		}
		return walk(n.Left) && fn(n.Value) && walk(n.Right)
	}
	walk(t.root)
}
