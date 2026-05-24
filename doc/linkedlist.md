# linkedlist

Package `linkedlist` provides a generic, thread-safe singly-linked list.

All public methods are safe for concurrent use by multiple goroutines.

> **Note:** The list has no tail-removal operation in O(1). `PopFront` is O(1), but removing from the back requires traversal and is O(n). If you need frequent back-removal, consider a different structure.

## Complexity

| Operation   | Time |
|-------------|------|
| Prepend     | O(1) |
| Append      | O(1) |
| Front       | O(1) |
| PopFront    | O(1) |
| Find        | O(n) |
| FindFunc    | O(n) |
| Delete      | O(n) |
| DeleteFunc  | O(n) |
| Len         | O(1) |
| Values      | O(n) |
| ForEach     | O(n) |

## Type

```go
type SinglyLinkedList[T any] struct { ... }
```

A generic, thread-safe singly-linked list. All operations are safe for concurrent use.

## Functions

### New

```go
func New[T any]() *SinglyLinkedList[T]
```

Returns an empty SinglyLinkedList.

```go
l := linkedlist.New[int]()
```

## Methods

### IsEmpty

```go
func (l *SinglyLinkedList[T]) IsEmpty() bool
```

Reports whether the list contains no elements.

### Len

```go
func (l *SinglyLinkedList[T]) Len() int
```

Returns the number of elements in the list.

### Prepend

```go
func (l *SinglyLinkedList[T]) Prepend(value T)
```

Inserts value at the front of the list.

```go
l.Prepend(1)
l.Prepend(2)
l.Values() // [2, 1]
```

### Append

```go
func (l *SinglyLinkedList[T]) Append(value T)
```

Inserts value at the back of the list.

```go
l.Append(1)
l.Append(2)
l.Values() // [1, 2]
```

### Front

```go
func (l *SinglyLinkedList[T]) Front() (T, bool)
```

Returns the first element without removing it, and `true`. Returns the zero value and `false` if the list is empty.

### PopFront

```go
func (l *SinglyLinkedList[T]) PopFront() (T, bool)
```

Removes and returns the first element and `true`. Returns the zero value and `false` if the list is empty.

```go
l.Append(10)
v, ok := l.PopFront() // 10, true
v, ok  = l.PopFront() // 0, false
```

### Find

```go
func (l *SinglyLinkedList[T]) Find(value T) (T, bool)
```

Returns the first element that compares equal to value using `reflect.DeepEqual`, and `true`. Returns the zero value and `false` if not found.

### FindFunc

```go
func (l *SinglyLinkedList[T]) FindFunc(match func(T) bool) (T, bool)
```

Returns the first element for which match returns `true`, and `true`. Returns the zero value and `false` if no element matches.

```go
v, ok := l.FindFunc(func(n int) bool { return n > 5 })
```

### Delete

```go
func (l *SinglyLinkedList[T]) Delete(value T) bool
```

Removes the first element that compares equal to value using `reflect.DeepEqual`. Returns `true` if an element was removed, `false` otherwise.

### DeleteFunc

```go
func (l *SinglyLinkedList[T]) DeleteFunc(match func(T) bool) bool
```

Removes the first element for which match returns `true`. Returns `true` if an element was removed, `false` otherwise.

```go
l.DeleteFunc(func(n int) bool { return n%2 == 0 }) // removes first even
```

### Values

```go
func (l *SinglyLinkedList[T]) Values() []T
```

Returns all elements in list order as a slice.

### ForEach

```go
func (l *SinglyLinkedList[T]) ForEach(fn func(T) bool)
```

Calls fn for each element in list order. Iteration stops early if fn returns `false`.
