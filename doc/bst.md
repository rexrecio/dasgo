# bst

Package `bst` provides a generic, thread-safe unbalanced binary search tree.

Keys must satisfy `cmp.Ordered`. All public methods are safe for concurrent use by multiple goroutines.

> **Note:** Because the tree is unbalanced, worst-case time for Insert, Find, and Delete is O(n) when keys are inserted in sorted order. For guaranteed O(log n) operations, use the [avl](avl.md) package.

## Complexity

| Operation | Average  | Worst |
|-----------|----------|-------|
| Insert    | O(log n) | O(n)  |
| Find      | O(log n) | O(n)  |
| Delete    | O(log n) | O(n)  |
| Values    | O(n)     | O(n)  |
| ForEach   | O(n)     | O(n)  |

## Type

```go
type BinarySearchTree[T cmp.Ordered] struct { ... }
```

A generic unbalanced binary search tree. All operations are safe for concurrent use.

## Functions

### New

```go
func New[T cmp.Ordered]() *BinarySearchTree[T]
```

Returns an empty BinarySearchTree.

```go
t := bst.New[string]()
```

## Methods

### IsEmpty

```go
func (t *BinarySearchTree[T]) IsEmpty() bool
```

Reports whether the tree contains no elements.

### Len

```go
func (t *BinarySearchTree[T]) Len() int
```

Returns the number of elements in the tree.

### Insert

```go
func (t *BinarySearchTree[T]) Insert(value T) bool
```

Adds value to the tree. Returns `true` if inserted, `false` if value was already present (duplicates are rejected).

```go
t.Insert("b") // true
t.Insert("b") // false
```

### Find

```go
func (t *BinarySearchTree[T]) Find(value T) (T, bool)
```

Returns the stored value equal to value and `true`, or the zero value and `false` if not found.

```go
v, ok := t.Find("b") // "b", true
v, ok  = t.Find("z") // "", false
```

### Delete

```go
func (t *BinarySearchTree[T]) Delete(value T) bool
```

Removes value from the tree. Returns `true` if the value was found and removed, `false` otherwise.

### Values

```go
func (t *BinarySearchTree[T]) Values() []T
```

Returns all elements in ascending order as a slice.

```go
t.Insert("c"); t.Insert("a"); t.Insert("b")
t.Values() // ["a", "b", "c"]
```

### ForEach

```go
func (t *BinarySearchTree[T]) ForEach(fn func(T) bool)
```

Calls fn for each element in ascending order. Iteration stops early if fn returns `false`.

```go
t.ForEach(func(v string) bool {
    fmt.Println(v)
    return true
})
```
