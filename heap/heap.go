// Package heap provides a generic, thread-safe binary heap (priority queue).
// Use [New] for a min-heap, [NewMax] for a max-heap, or [NewFunc] for a
// custom ordering. All public methods are safe for concurrent use by multiple goroutines.
package heap

import (
	"cmp"
	"sync"
)

// Heap is a binary heap backed by a slice.
// The root is always the element for which less(root, x) holds for every other x.
type Heap[T any] struct {
	mu   sync.RWMutex
	data []T
	less func(a, b T) bool
}

// New returns a min-heap ordered by the natural ordering of T.
func New[T cmp.Ordered]() *Heap[T] {
	return NewFunc(func(a, b T) bool { return a < b })
}

// NewMax returns a max-heap ordered by the natural ordering of T.
func NewMax[T cmp.Ordered]() *Heap[T] {
	return NewFunc(func(a, b T) bool { return a > b })
}

// NewFunc returns a heap ordered by the provided less function.
// The root will be the element for which less(root, x) is true for all other x.
func NewFunc[T any](less func(a, b T) bool) *Heap[T] {
	return &Heap[T]{less: less}
}

func (h *Heap[T]) Len() int {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return len(h.data)
}

func (h *Heap[T]) IsEmpty() bool {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return len(h.data) == 0
}

// Push adds value to the heap.
func (h *Heap[T]) Push(value T) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.data = append(h.data, value)
	h.siftUp(len(h.data) - 1)
}

// Pop removes and returns the top element (minimum for a min-heap, maximum for a max-heap).
func (h *Heap[T]) Pop() (T, bool) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if len(h.data) == 0 {
		var zero T
		return zero, false
	}
	top := h.data[0]
	last := len(h.data) - 1
	h.data[0] = h.data[last]
	h.data = h.data[:last]
	if len(h.data) > 0 {
		h.siftDown(0)
	}
	return top, true
}

// Peek returns the top element without removing it.
func (h *Heap[T]) Peek() (T, bool) {
	h.mu.RLock()
	defer h.mu.RUnlock()
	if len(h.data) == 0 {
		var zero T
		return zero, false
	}
	return h.data[0], true
}

func (h *Heap[T]) siftUp(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if h.less(h.data[i], h.data[parent]) {
			h.data[i], h.data[parent] = h.data[parent], h.data[i]
			i = parent
		} else {
			break
		}
	}
}

func (h *Heap[T]) siftDown(i int) {
	n := len(h.data)
	for {
		top := i
		left := 2*i + 1
		right := 2*i + 2
		if left < n && h.less(h.data[left], h.data[top]) {
			top = left
		}
		if right < n && h.less(h.data[right], h.data[top]) {
			top = right
		}
		if top == i {
			break
		}
		h.data[i], h.data[top] = h.data[top], h.data[i]
		i = top
	}
}
