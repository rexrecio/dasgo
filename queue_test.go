package dasgo

import (
	"testing"
)

var q = new(Queue)

func TestQueueEmpty(t *testing.T) {
	if !s.IsEmpty() {
		t.Error("Newly created Queue should be empty")
	}

}
func TestQueueEnqueueAndCount(t *testing.T) {
	for i := 0; i < 10; i++ {
		testValue := new(int)
		*testValue = i
		q.Enqueue(testValue)
	}
	if q.Count() != 10 {
		t.Error("Failure in adding elements to the s")
	}
}
func TestQueuePeekAndCount(t *testing.T) {
	returnvalue := q.Peek()

	if q.Count() != 10 {
		t.Errorf("Failed to peek, expecting a total 10  but got %v", s.Count())
	}
	if *returnvalue.(*int) != 0 {
		//Use assertion to read value from interfcae{}
		t.Errorf("Failed to peek a value, expecting 0 but got %v", *returnvalue.(*int))
	}
}
func TestQueueDequeueAndCount(t *testing.T) {
	returnvalue := q.Dequeue()

	if q.Count() != 9 {
		t.Errorf("Failed to dequeue, expecting the total 9  but got %v", s.Count())
	}
	if *returnvalue.(*int) != 0 {
		//Use assertion to read value from interfcae{}
		t.Errorf("Failed to deque a value, expecting 0 but got %v", *returnvalue.(*int))
	}
}
