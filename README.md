# godas
Fundamental data structures and algorithms in Go.
## Run unit tests
`go test -v`
## Run sample app
```
package main

import (
	"fmt"
	"github.com/rexrecio/dasgo"
)

func main() {
	var s = new(dasgo.Stack)
	for i := 0; i < 10; i++ {
		testValue := new(int)
		*testValue = i
		fmt.Println(*testValue, "added to stack")
		s.Push(testValue)
	}
	fmt.Println("The stack size is ", s.Count())
	var returnvalue interface{}  //we'll use interface{} until generics come to Go
	for !s.IsEmpty() {
		returnvalue = s.Pop()
		fmt.Println(*returnvalue.(*int), "popped out of the stack")
	}
	fmt.Println("The stack size is ", s.Count())
}
```
