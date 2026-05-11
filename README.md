# dasgo

Fundamental data structures and algorithms for Go.

## Packages

- `github.com/rexrecio/dasgo/linkedlist`
- `github.com/rexrecio/dasgo/bst`
- `github.com/rexrecio/dasgo/avl`
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
	list.Append("alpha")
	list.Prepend("zero")
	front, _ := list.Front()
	fmt.Println(front)         // zero
	fmt.Println(list.Values()) // [zero alpha]

	_ = list.FindFunc(func(v string) bool {
		return len(v) > 4
	})
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
	tree.Insert(10)
	tree.Insert(3)
	tree.Insert(20)
	fmt.Println(tree.Values()) // [3 10 20]
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
	tree.Insert(10)
	tree.Insert(3)
	tree.Insert(20)
	fmt.Println(tree.Values()) // [3 10 20]
}
```

The AVL tree is a self-balancing BST with the same API as `bst`. It guarantees
O(log n) insert, delete, and lookup even for sorted or adversarial input.

### Stack

```go
package main

import (
	"fmt"

	"github.com/rexrecio/dasgo/stack"
)

func main() {
	// Stack supports any type parameter (T any).
	s := stack.New[string]()
	s.Push("a")
	s.Push("b")
	v, _ := s.Pop()
	fmt.Println(v) // b
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
	// Queue supports any type parameter (T any).
	q := queue.New[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	v, _ := q.Dequeue()
	fmt.Println(v) // 1
}
```

## Local Demo

The demo app lives in the nested module at `cmd/demo` and shows usage of all included packages:

- `linkedlist`
- `bst`
- `avl`
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
- stack push, peek, and pop
- queue enqueue, peek, and dequeue

## Test

```bash
go test ./...
```

## Release

Tag a version so consumers can pin dependencies:

```bash
git tag v1.1.0
git push origin v1.1.0
```
