package dasgo

import (
	"testing"
)

var s = new(Stack)

func TestStackEmpty(t *testing.T) {
	if !s.IsEmpty() {
		t.Error("Newly created stack should be empty")
	}

}
func TestStackPushAndCount(t *testing.T) {
	for i := 0; i < 10; i++ {
		testValue := new(int)
		*testValue = i
		s.Push(testValue)
	}
	if s.Count() != 10 {
		t.Error("Failure in adding elements to the s")
	}
}
func TestStackPeekAndCount(t *testing.T) {
	returnvalue := s.Peek()

	if s.Count() != 10 {
		t.Errorf("Failed to peek, expecting a total 10  but got %v", s.Count())
	}
	if *returnvalue.(*int) != 9 {
		//Use assertion to read value from interfcae{}
		t.Errorf("Failed to peek a value, expecting 10 but got %v", *returnvalue.(*int))
	}
}
func TestStackPopAndCount(t *testing.T) {
	returnvalue := s.Pop()

	if s.Count() != 9 {
		t.Errorf("Failed to delete the head, expecting the total 9  but got %v", s.Count())
	}
	if *returnvalue.(*int) != 9 {
		//Use assertion to read value from interfcae{}
		t.Errorf("Failed to pop a value, expecting 9 but got %v", *returnvalue.(*int))
	}
}
