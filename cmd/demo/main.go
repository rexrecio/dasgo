package main

import (
	"fmt"

	"github.com/rexrecio/dasgo/bst"
	"github.com/rexrecio/dasgo/linkedlist"
	"github.com/rexrecio/dasgo/queue"
	"github.com/rexrecio/dasgo/stack"
)

func main() {
	linkedListDemo()
	bstDemo()
	stackDemo()
	queueDemo()
}

func linkedListDemo() {
	list := linkedlist.New[string]()

	fmt.Println("=== Linked List ===")
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
	fmt.Println()
}

func bstDemo() {
	fmt.Println("=== Binary Search Tree ===")
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
	fmt.Println()
}

func stackDemo() {
	fmt.Println("=== Stack (LIFO) ===")
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
	fmt.Println()
}

func queueDemo() {
	fmt.Println("=== Queue (FIFO) ===")
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
	fmt.Println()
}
