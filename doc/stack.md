# stack

Package `stack` provides a generic, thread-safe LIFO stack backed by a singly-linked list.

All public methods are safe for concurrent use by multiple goroutines.

## Complexity

| Operation | Time |
|-----------|------|
| Push      | O(1) |
| Pop       | O(1) |
| Peek      | O(1) |
| Len       | O(1) |
| Values    | O(n) |

## Type

```go
type Stack[T any] struct { ... }
```

A generic, thread-safe LIFO stack. The last element pushed is the first one returned by Pop.

## Functions

### New

```go
func New[T any]() *Stack[T]
```

Returns an empty Stack.

```go
s := stack.New[int]()
```

## Methods

### IsEmpty

```go
func (s *Stack[T]) IsEmpty() bool
```

Reports whether the stack contains no elements.

### Len

```go
func (s *Stack[T]) Len() int
```

Returns the number of elements in the stack.

### Push

```go
func (s *Stack[T]) Push(value T)
```

Adds value to the top of the stack.

```go
s.Push(1)
s.Push(2)
s.Push(3)
```

### Pop

```go
func (s *Stack[T]) Pop() (T, bool)
```

Removes and returns the top element and `true`. Returns the zero value and `false` if the stack is empty.

```go
s.Push(1); s.Push(2)
v, ok := s.Pop() // 2, true
v, ok  = s.Pop() // 1, true
v, ok  = s.Pop() // 0, false
```

### Peek

```go
func (s *Stack[T]) Peek() (T, bool)
```

Returns the top element without removing it, and `true`. Returns the zero value and `false` if the stack is empty.

```go
s.Push(42)
v, ok := s.Peek() // 42, true
v, ok  = s.Pop()  // 42, true — still there
```

### Values

```go
func (s *Stack[T]) Values() []T
```

Returns all elements in top-to-bottom order as a slice.

```go
s.Push(1); s.Push(2); s.Push(3)
s.Values() // [3, 2, 1]
```
