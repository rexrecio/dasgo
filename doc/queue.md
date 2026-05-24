# queue

Package `queue` provides a generic, thread-safe FIFO queue backed by a singly-linked list.

All public methods are safe for concurrent use by multiple goroutines.

## Complexity

| Operation | Time |
|-----------|------|
| Enqueue   | O(1) |
| Dequeue   | O(1) |
| Peek      | O(1) |
| Len       | O(1) |
| Values    | O(n) |

## Type

```go
type Queue[T any] struct { ... }
```

A generic, thread-safe FIFO queue. Elements are dequeued in the same order they were enqueued.

## Functions

### New

```go
func New[T any]() *Queue[T]
```

Returns an empty Queue.

```go
q := queue.New[string]()
```

## Methods

### IsEmpty

```go
func (q *Queue[T]) IsEmpty() bool
```

Reports whether the queue contains no elements.

### Len

```go
func (q *Queue[T]) Len() int
```

Returns the number of elements in the queue.

### Enqueue

```go
func (q *Queue[T]) Enqueue(value T)
```

Adds value to the back of the queue.

```go
q.Enqueue("a")
q.Enqueue("b")
q.Enqueue("c")
```

### Dequeue

```go
func (q *Queue[T]) Dequeue() (T, bool)
```

Removes and returns the front element and `true`. Returns the zero value and `false` if the queue is empty.

```go
q.Enqueue("a")
q.Enqueue("b")
v, ok := q.Dequeue() // "a", true
v, ok  = q.Dequeue() // "b", true
v, ok  = q.Dequeue() // "", false
```

### Peek

```go
func (q *Queue[T]) Peek() (T, bool)
```

Returns the front element without removing it, and `true`. Returns the zero value and `false` if the queue is empty.

```go
q.Enqueue("a")
v, ok := q.Peek()    // "a", true
v, ok  = q.Dequeue() // "a", true — still there
```

### Values

```go
func (q *Queue[T]) Values() []T
```

Returns all elements in front-to-back order as a slice.

```go
q.Enqueue("a"); q.Enqueue("b")
q.Values() // ["a", "b"]
```
