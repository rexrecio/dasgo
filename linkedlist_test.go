package dasgo

import (
	"testing"
)

var list = new(LinkedList)

func TestListEmpty(t *testing.T) {
	if !list.IsEmpty() {
		t.Error("Newly created list should be empty")
	}

}
func TestListAddAndCount(t *testing.T) {
	for i := 0; i < 10; i++ {
		testValue := new(int)
		*testValue = i
		list.Add(testValue)
	}
	if list.Count() != 10 {
		t.Error("Failure in adding elements to the list")
	}
}
func TestListDeleteHead(t *testing.T) {
	list.DeleteHead()
	if list.Count() != 9 {
		t.Errorf("Failed to delete the head, expecting 9 remaining elements but got %v", list.Count())
	}
}
func TestListDeleteTail(t *testing.T) {
	list.DeleteHead()
	if list.Count() != 8 {
		t.Errorf("Failed to delete the tail, expecting 8 remaining elements but got %v", list.Count())
	}
}
