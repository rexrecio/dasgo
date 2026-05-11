// Package avl provides a generic, thread-safe AVL tree (self-balancing binary search tree).
// Keys must satisfy [cmp.Ordered]. Insert, Find, and Delete all run in O(log n) time.
// All public methods are safe for concurrent use by multiple goroutines.
package avl

import (
	"cmp"
	"sync"
)

type node[T cmp.Ordered] struct {
	Value  T
	Left   *node[T]
	Right  *node[T]
	height int
}

// AVLTree is a generic self-balancing binary search tree.
// All operations are safe for concurrent use.
type AVLTree[T cmp.Ordered] struct {
	mu   sync.RWMutex
	root *node[T]
	size int
}

func New[T cmp.Ordered]() *AVLTree[T] {
	return &AVLTree[T]{}
}

func (t *AVLTree[T]) IsEmpty() bool {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.size == 0
}

func (t *AVLTree[T]) Len() int {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.size
}

func nodeHeight[T cmp.Ordered](n *node[T]) int {
	if n == nil {
		return 0
	}
	return n.height
}

func balanceFactor[T cmp.Ordered](n *node[T]) int {
	if n == nil {
		return 0
	}
	return nodeHeight(n.Left) - nodeHeight(n.Right)
}

func updateHeight[T cmp.Ordered](n *node[T]) {
	n.height = 1 + max(nodeHeight(n.Left), nodeHeight(n.Right))
}

func rotateRight[T cmp.Ordered](y *node[T]) *node[T] {
	x := y.Left
	y.Left = x.Right
	x.Right = y
	updateHeight(y)
	updateHeight(x)
	return x
}

func rotateLeft[T cmp.Ordered](x *node[T]) *node[T] {
	y := x.Right
	x.Right = y.Left
	y.Left = x
	updateHeight(x)
	updateHeight(y)
	return y
}

func rebalance[T cmp.Ordered](n *node[T]) *node[T] {
	updateHeight(n)
	bf := balanceFactor(n)
	if bf > 1 {
		if balanceFactor(n.Left) < 0 {
			n.Left = rotateLeft(n.Left)
		}
		return rotateRight(n)
	}
	if bf < -1 {
		if balanceFactor(n.Right) > 0 {
			n.Right = rotateRight(n.Right)
		}
		return rotateLeft(n)
	}
	return n
}

func insertNode[T cmp.Ordered](n *node[T], value T) (*node[T], bool) {
	if n == nil {
		return &node[T]{Value: value, height: 1}, true
	}
	var inserted bool
	switch {
	case value < n.Value:
		n.Left, inserted = insertNode(n.Left, value)
	case value > n.Value:
		n.Right, inserted = insertNode(n.Right, value)
	default:
		return n, false // duplicate
	}
	if !inserted {
		return n, false
	}
	return rebalance(n), true
}

// detachMin removes the minimum node from the subtree and returns
// the new subtree root and the removed value.
func detachMin[T cmp.Ordered](n *node[T]) (*node[T], T) {
	if n.Left == nil {
		return n.Right, n.Value
	}
	var minVal T
	n.Left, minVal = detachMin(n.Left)
	return rebalance(n), minVal
}

func deleteNode[T cmp.Ordered](n *node[T], value T) (*node[T], bool) {
	if n == nil {
		return nil, false
	}
	var deleted bool
	switch {
	case value < n.Value:
		n.Left, deleted = deleteNode(n.Left, value)
	case value > n.Value:
		n.Right, deleted = deleteNode(n.Right, value)
	default:
		if n.Left == nil {
			return n.Right, true
		}
		if n.Right == nil {
			return n.Left, true
		}
		// Two children: replace with in-order successor.
		n.Right, n.Value = detachMin(n.Right)
		deleted = true
	}
	if !deleted {
		return n, false
	}
	return rebalance(n), true
}

func (t *AVLTree[T]) Insert(value T) bool {
	t.mu.Lock()
	defer t.mu.Unlock()
	var inserted bool
	t.root, inserted = insertNode(t.root, value)
	if inserted {
		t.size++
	}
	return inserted
}

func (t *AVLTree[T]) Find(value T) (T, bool) {
	t.mu.RLock()
	defer t.mu.RUnlock()
	curr := t.root
	for curr != nil {
		switch {
		case value < curr.Value:
			curr = curr.Left
		case value > curr.Value:
			curr = curr.Right
		default:
			return curr.Value, true
		}
	}
	var zero T
	return zero, false
}

func (t *AVLTree[T]) Delete(value T) bool {
	t.mu.Lock()
	defer t.mu.Unlock()
	var deleted bool
	t.root, deleted = deleteNode(t.root, value)
	if deleted {
		t.size--
	}
	return deleted
}

// Values returns all values in ascending order.
func (t *AVLTree[T]) Values() []T {
	t.mu.RLock()
	defer t.mu.RUnlock()
	values := make([]T, 0, t.size)
	var walk func(n *node[T])
	walk = func(n *node[T]) {
		if n == nil {
			return
		}
		walk(n.Left)
		values = append(values, n.Value)
		walk(n.Right)
	}
	walk(t.root)
	return values
}

// ForEach calls fn for each value in ascending order. If fn returns false, iteration stops.
func (t *AVLTree[T]) ForEach(fn func(T) bool) {
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
