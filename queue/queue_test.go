package queue

import (
	"reflect"
	"testing"
)

func TestNewQueueIsEmpty(t *testing.T) {
	q := New[int]()

	if !q.IsEmpty() {
		t.Fatal("expected new queue to be empty")
	}
	if q.Len() != 0 {
		t.Fatalf("expected length 0, got %d", q.Len())
	}
}

func TestEnqueueDequeueFIFO(t *testing.T) {
	q := New[int]()
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)

	if got, want := q.Values(), []int{10, 20, 30}; !reflect.DeepEqual(got, want) {
		t.Fatalf("unexpected values: got %v want %v", got, want)
	}

	v, ok := q.Dequeue()
	if !ok || v != 10 {
		t.Fatalf("expected dequeue 10,true got %d,%v", v, ok)
	}
	v, ok = q.Dequeue()
	if !ok || v != 20 {
		t.Fatalf("expected dequeue 20,true got %d,%v", v, ok)
	}
	v, ok = q.Dequeue()
	if !ok || v != 30 {
		t.Fatalf("expected dequeue 30,true got %d,%v", v, ok)
	}
	if _, ok = q.Dequeue(); ok {
		t.Fatal("expected dequeue on empty queue to fail")
	}
}

func TestPeek(t *testing.T) {
	q := New[int]()
	if _, ok := q.Peek(); ok {
		t.Fatal("expected peek on empty queue to fail")
	}

	q.Enqueue(1)
	q.Enqueue(2)

	v, ok := q.Peek()
	if !ok || v != 1 {
		t.Fatalf("expected peek 1,true got %d,%v", v, ok)
	}
	if q.Len() != 2 {
		t.Fatalf("expected length 2 after peek, got %d", q.Len())
	}
}

func TestDuplicateValues(t *testing.T) {
	q := New[string]()
	q.Enqueue("a")
	q.Enqueue("b")
	q.Enqueue("a")

	v, ok := q.Dequeue()
	if !ok || v != "a" {
		t.Fatalf("expected first dequeue to return front duplicate a,true got %q,%v", v, ok)
	}
	if got, want := q.Values(), []string{"b", "a"}; !reflect.DeepEqual(got, want) {
		t.Fatalf("unexpected values after dequeue: got %v want %v", got, want)
	}
}
