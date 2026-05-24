# heap

Package `heap` provides a generic, thread-safe binary heap (priority queue).

Use `New` for a min-heap, `NewMax` for a max-heap, or `NewFunc` for a custom ordering. All public methods are safe for concurrent use by multiple goroutines.

## Complexity

| Operation | Time     |
|-----------|----------|
| Push      | O(log n) |
| Pop       | O(log n) |
| Peek      | O(1)     |
| Len       | O(1)     |

## Type

```go
type Heap[T any] struct { ... }
```

A binary heap backed by a slice. The root is always the element for which `less(root, x)` holds for every other element x. For a min-heap that is the smallest element; for a max-heap it is the largest.

## Functions

### New

```go
func New[T cmp.Ordered]() *Heap[T]
```

Returns a min-heap ordered by the natural ordering of T. The smallest element is at the top.

```go
h := heap.New[int]()
h.Push(3); h.Push(1); h.Push(2)
h.Pop() // 1
```

### NewMax

```go
func NewMax[T cmp.Ordered]() *Heap[T]
```

Returns a max-heap ordered by the natural ordering of T. The largest element is at the top.

```go
h := heap.NewMax[int]()
h.Push(3); h.Push(1); h.Push(2)
h.Pop() // 3
```

### NewFunc

```go
func NewFunc[T any](less func(a, b T) bool) *Heap[T]
```

Returns a heap ordered by the provided less function. The element for which `less(element, x)` is true for all other x will be at the top. Use this for custom types or orderings.

```go
type Task struct {
    Priority int
    Name     string
}

h := heap.NewFunc(func(a, b Task) bool {
    return a.Priority < b.Priority
})
h.Push(Task{2, "low"})
h.Push(Task{1, "high"})
h.Pop() // Task{1, "high"}
```

## Methods

### IsEmpty

```go
func (h *Heap[T]) IsEmpty() bool
```

Reports whether the heap contains no elements.

### Len

```go
func (h *Heap[T]) Len() int
```

Returns the number of elements in the heap.

### Push

```go
func (h *Heap[T]) Push(value T)
```

Adds value to the heap.

### Pop

```go
func (h *Heap[T]) Pop() (T, bool)
```

Removes and returns the top element (minimum for a min-heap, maximum for a max-heap) and `true`. Returns the zero value and `false` if the heap is empty.

```go
h := heap.New[int]()
h.Push(5); h.Push(2); h.Push(8)
h.Pop() // 2, true
h.Pop() // 5, true
h.Pop() // 8, true
h.Pop() // 0, false
```

### Peek

```go
func (h *Heap[T]) Peek() (T, bool)
```

Returns the top element without removing it, and `true`. Returns the zero value and `false` if the heap is empty.

```go
h.Push(3)
v, ok := h.Peek() // 3, true — heap unchanged
```
