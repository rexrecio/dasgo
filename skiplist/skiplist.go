// Package skiplist provides a generic, thread-safe skip list.
// Keys must satisfy [cmp.Ordered]. Insert, Find, and Delete run in expected
// O(log n) time. All public methods are safe for concurrent use by multiple goroutines.
package skiplist

import (
	"cmp"
	"math/rand"
	"sync"
	"time"
)

const (
	defaultMaxLevel = 16
	defaultP        = 0.5
)

type node[T cmp.Ordered] struct {
	value   T
	forward []*node[T]
}

// SkipList is a generic probabilistic ordered set backed by a skip list.
// All operations are safe for concurrent use.
type SkipList[T cmp.Ordered] struct {
	mu    sync.RWMutex
	head  *node[T]
	level int
	size  int
	rnd   *rand.Rand
}

// New returns an empty SkipList.
func New[T cmp.Ordered]() *SkipList[T] {
	return &SkipList[T]{
		head:  &node[T]{forward: make([]*node[T], defaultMaxLevel)},
		level: 1,
		rnd:   rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// IsEmpty reports whether the skip list contains no elements.
func (s *SkipList[T]) IsEmpty() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.size == 0
}

// Len returns the number of elements in the skip list.
func (s *SkipList[T]) Len() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.size
}

func (s *SkipList[T]) randomLevel() int {
	level := 1
	for level < defaultMaxLevel && s.rnd.Float64() < defaultP {
		level++
	}
	return level
}

// Insert adds value to the skip list and returns true. If value is already present
// it returns false and the skip list is unchanged.
func (s *SkipList[T]) Insert(value T) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	update := make([]*node[T], defaultMaxLevel)
	curr := s.head
	for i := s.level - 1; i >= 0; i-- {
		for curr.forward[i] != nil && curr.forward[i].value < value {
			curr = curr.forward[i]
		}
		update[i] = curr
	}

	curr = curr.forward[0]
	if curr != nil && curr.value == value {
		return false
	}

	newLevel := s.randomLevel()
	if newLevel > s.level {
		for i := s.level; i < newLevel; i++ {
			update[i] = s.head
		}
		s.level = newLevel
	}

	newNode := &node[T]{
		value:   value,
		forward: make([]*node[T], newLevel),
	}
	for i := 0; i < newLevel; i++ {
		newNode.forward[i] = update[i].forward[i]
		update[i].forward[i] = newNode
	}

	s.size++
	return true
}

// Find returns the stored value equal to value and true, or the zero value and
// false if not found.
func (s *SkipList[T]) Find(value T) (T, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	curr := s.head
	for i := s.level - 1; i >= 0; i-- {
		for curr.forward[i] != nil && curr.forward[i].value < value {
			curr = curr.forward[i]
		}
	}
	curr = curr.forward[0]
	if curr != nil && curr.value == value {
		return curr.value, true
	}

	var zero T
	return zero, false
}

// Delete removes value from the skip list and returns true. If value is not present
// it returns false and the skip list is unchanged.
func (s *SkipList[T]) Delete(value T) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	update := make([]*node[T], defaultMaxLevel)
	curr := s.head
	for i := s.level - 1; i >= 0; i-- {
		for curr.forward[i] != nil && curr.forward[i].value < value {
			curr = curr.forward[i]
		}
		update[i] = curr
	}

	target := curr.forward[0]
	if target == nil || target.value != value {
		return false
	}

	for i := 0; i < s.level; i++ {
		if update[i].forward[i] != target {
			continue
		}
		update[i].forward[i] = target.forward[i]
	}

	for s.level > 1 && s.head.forward[s.level-1] == nil {
		s.level--
	}

	s.size--
	return true
}

// Values returns all values in ascending order.
func (s *SkipList[T]) Values() []T {
	s.mu.RLock()
	defer s.mu.RUnlock()

	values := make([]T, 0, s.size)
	for curr := s.head.forward[0]; curr != nil; curr = curr.forward[0] {
		values = append(values, curr.value)
	}
	return values
}

// ForEach calls fn for each value in ascending order. If fn returns false, iteration stops.
func (s *SkipList[T]) ForEach(fn func(T) bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for curr := s.head.forward[0]; curr != nil; curr = curr.forward[0] {
		if !fn(curr.value) {
			return
		}
	}
}
