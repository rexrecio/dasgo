# avl

Package `avl` provides a generic, thread-safe AVL tree (self-balancing binary search tree).

Keys must satisfy `cmp.Ordered`. All public methods are safe for concurrent use by multiple goroutines.

## Complexity

| Operation | Time     |
|-----------|----------|
| Insert    | O(log n) |
| Find      | O(log n) |
| Delete    | O(log n) |
| Values    | O(n)     |
| ForEach   | O(n)     |

## Type

```go
type AVLTree[T cmp.Ordered] struct { ... }
```

A generic self-balancing binary search tree. The tree rebalances itself after every insert and delete, keeping operations at O(log n) even in adversarial insertion order. For an unbalanced alternative, see the [bst](bst.md) package.

## Functions

### New

```go
func New[T cmp.Ordered]() *AVLTree[T]
```

Returns an empty AVLTree.

```go
t := avl.New[int]()
```

## Methods

### IsEmpty

```go
func (t *AVLTree[T]) IsEmpty() bool
```

Reports whether the tree contains no elements.

### Len

```go
func (t *AVLTree[T]) Len() int
```

Returns the number of elements in the tree.

### Insert

```go
func (t *AVLTree[T]) Insert(value T) bool
```

Adds value to the tree. Returns `true` if the value was inserted, `false` if it was already present (duplicates are rejected).

```go
t.Insert(10) // true
t.Insert(10) // false
```

### Find

```go
func (t *AVLTree[T]) Find(value T) (T, bool)
```

Returns the stored value equal to value and `true`, or the zero value and `false` if not found.

```go
v, ok := t.Find(10) // 10, true
v, ok  = t.Find(99) // 0, false
```

### Delete

```go
func (t *AVLTree[T]) Delete(value T) bool
```

Removes value from the tree. Returns `true` if the value was found and removed, `false` otherwise.

```go
t.Delete(10) // true
t.Delete(10) // false
```

### Values

```go
func (t *AVLTree[T]) Values() []T
```

Returns all elements in ascending order as a slice.

```go
t.Insert(3); t.Insert(1); t.Insert(2)
t.Values() // [1, 2, 3]
```

### ForEach

```go
func (t *AVLTree[T]) ForEach(fn func(T) bool)
```

Calls fn for each element in ascending order. Iteration stops early if fn returns `false`.

```go
t.ForEach(func(v int) bool {
    fmt.Println(v)
    return v < 5 // stop after 5
})
```
