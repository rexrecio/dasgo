package main

import (
	"fmt"

	"github.com/rexrecio/dasgo/linkedlist"
)

func main() {
	list := linkedlist.New[string]()

	list.Append("apple")
	list.Append("banana")
	list.Prepend("avocado")

	fmt.Printf("list=%v len=%d empty=%v\n", list.Values(), list.Len(), list.IsEmpty())

	deleted := list.Delete("apple")
	fmt.Printf("deleted apple=%v list=%v\n", deleted, list.Values())

	found := list.Find("banana") != nil
	fmt.Printf("found banana=%v\n", found)
}
