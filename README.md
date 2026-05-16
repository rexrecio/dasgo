# dasgo

Fundamental data structures and algorithms for Go.

## Packages

- `github.com/rexrecio/dasgo/linkedlist`
- `github.com/rexrecio/dasgo/bst`
- `github.com/rexrecio/dasgo/avl`
- `github.com/rexrecio/dasgo/heap`
- `github.com/rexrecio/dasgo/stack`
- `github.com/rexrecio/dasgo/queue`

## Install

```bash
go get github.com/rexrecio/dasgo@latest
```

## Usage

### Linked List

```go
package main

import (
	"fmt"

	"github.com/rexrecio/dasgo/linkedlist"
)

func main() {
	list := linkedlist.New[string]()

	list.Append("banana")
	list.Append("cherry")
	list.Prepend("apple")

	fmt.Println("values:  ", list.Values())
	fmt.Println("len:     ", list.Len())
	fmt.Println("empty:   ", list.IsEmpty())
	_, foundCherry := list.Find("cherry")
	fmt.Println("find cherry:", foundCherry)

	list.Delete("banana")
	fmt.Println("after delete banana:", list.Values())
}
```

Linked list API highlights:

- `Find(value T)` and `Delete(value T)` for value-based behavior
- `FindFunc(match func(T) bool)` and `DeleteFunc(match func(T) bool)` for custom matching
- `Front()` and `PopFront()` for efficient head operations

### Binary Search Tree

```go
package main

import (
	"fmt"

	"github.com/rexrecio/dasgo/bst"
)

func main() {
	tree := bst.New[int]()

	for _, v := range []int{10, 4, 20, 1, 7, 15, 25} {
		tree.Insert(v)
	}

	fmt.Println("in-order:  ", tree.Values())
	fmt.Println("len:       ", tree.Len())
	_, found7 := tree.Find(7)
	_, found99 := tree.Find(99)
	fmt.Println("find 7:    ", found7)
	fmt.Println("find 99:   ", found99)

	tree.Delete(10)
	fmt.Println("after delete root 10:", tree.Values())
}
```

### AVL Tree

```go
package main

import (
	"fmt"

	"github.com/rexrecio/dasgo/avl"
)

func main() {
	tree := avl.New[int]()

	for _, v := range []int{10, 4, 20, 1, 7, 15, 25} {
		tree.Insert(v)
	}

	fmt.Println("in-order:  ", tree.Values())
	fmt.Println("len:       ", tree.Len())
	_, found7 := tree.Find(7)
	_, found99 := tree.Find(99)
	fmt.Println("find 7:    ", found7)
	fmt.Println("find 99:   ", found99)

	tree.Delete(10)
	fmt.Println("after delete root 10:", tree.Values())
}
```

The AVL tree is a self-balancing BST with the same API as `bst`. It guarantees
O(log n) insert, delete, and lookup even for sorted or adversarial input.

### Heap

```go
package main

import (
	"fmt"

	"github.com/rexrecio/dasgo/heap"
)

func main() {
	// Min-heap (default).
	h := heap.New[int]()
	h.Push(5)
	h.Push(1)
	h.Push(3)

	if v, ok := h.Peek(); ok {
		fmt.Println("peek:", v) // 1
	}
	for !h.IsEmpty() {
		v, _ := h.Pop()
		fmt.Println("pop:", v) // 1, 3, 5
	}

	// Max-heap.
	mh := heap.NewMax[int]()
	mh.Push(5)
	mh.Push(1)
	mh.Push(3)
	v, _ := mh.Pop()
	fmt.Println("max pop:", v) // 5

	// Custom ordering — min-heap by string length.
	sh := heap.NewFunc(func(a, b string) bool {
		return len(a) < len(b)
	})
	sh.Push("banana")
	sh.Push("fig")
	sh.Push("apple")
	v2, _ := sh.Pop()
	fmt.Println("shortest:", v2) // fig
}
```

Heap constructors:

- `New[T]()` — min-heap using natural ordering (`T` must satisfy `cmp.Ordered`)
- `NewMax[T]()` — max-heap using natural ordering
- `NewFunc(less func(a, b T) bool)` — custom ordering, works with any type

### Stack

```go
package main

import (
	"fmt"

	"github.com/rexrecio/dasgo/stack"
)

func main() {
	s := stack.New[int]()

	s.Push(1)
	s.Push(2)
	s.Push(3)

	fmt.Println("values:  ", s.Values())
	fmt.Println("len:     ", s.Len())

	if v, ok := s.Peek(); ok {
		fmt.Println("peek:    ", v)
	}

	for !s.IsEmpty() {
		v, _ := s.Pop()
		fmt.Println("pop:", v)
	}
}
```

### Queue

```go
package main

import (
	"fmt"

	"github.com/rexrecio/dasgo/queue"
)

func main() {
	q := queue.New[string]()

	q.Enqueue("first")
	q.Enqueue("second")
	q.Enqueue("third")

	fmt.Println("values:  ", q.Values())
	fmt.Println("len:     ", q.Len())

	if v, ok := q.Peek(); ok {
		fmt.Println("peek:    ", v)
	}

	for !q.IsEmpty() {
		v, _ := q.Dequeue()
		fmt.Println("dequeue:", v)
	}
}
```

## Local Demo

The demo app lives in the nested module at `cmd/demo` and shows usage of all included packages:

- `linkedlist`
- `bst`
- `avl`
- `heap`
- `stack`
- `queue`

Run it from the demo directory:

```bash
cd cmd/demo
go run .
```

The demo prints example output for:

- linked list insertion, lookup, and delete
- binary search tree insertion, traversal, lookup, and delete
- AVL tree insertion, traversal, lookup, and delete
- heap push, peek, and pop (min, max, and custom)
- stack push, peek, and pop
- queue enqueue, peek, and dequeue

## Test

```bash
go test ./...
```