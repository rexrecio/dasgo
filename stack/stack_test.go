package stack

import (
	"reflect"
	"testing"
)

func TestNewStackIsEmpty(t *testing.T) {
	s := New[int]()

	if !s.IsEmpty() {
		t.Fatal("expected new stack to be empty")
	}
	if s.Len() != 0 {
		t.Fatalf("expected length 0, got %d", s.Len())
	}
}

func TestPushPopLIFO(t *testing.T) {
	s := New[int]()
	s.Push(10)
	s.Push(20)
	s.Push(30)

	if got, want := s.Values(), []int{30, 20, 10}; !reflect.DeepEqual(got, want) {
		t.Fatalf("unexpected values: got %v want %v", got, want)
	}

	v, ok := s.Pop()
	if !ok || v != 30 {
		t.Fatalf("expected pop 30,true got %d,%v", v, ok)
	}
	v, ok = s.Pop()
	if !ok || v != 20 {
		t.Fatalf("expected pop 20,true got %d,%v", v, ok)
	}
	v, ok = s.Pop()
	if !ok || v != 10 {
		t.Fatalf("expected pop 10,true got %d,%v", v, ok)
	}
	if _, ok = s.Pop(); ok {
		t.Fatal("expected pop on empty stack to fail")
	}
}

func TestPeek(t *testing.T) {
	s := New[int]()
	if _, ok := s.Peek(); ok {
		t.Fatal("expected peek on empty stack to fail")
	}

	s.Push(1)
	s.Push(2)

	v, ok := s.Peek()
	if !ok || v != 2 {
		t.Fatalf("expected peek 2,true got %d,%v", v, ok)
	}
	if s.Len() != 2 {
		t.Fatalf("expected length 2 after peek, got %d", s.Len())
	}
}

func TestDuplicateValues(t *testing.T) {
	s := New[string]()
	s.Push("a")
	s.Push("b")
	s.Push("a")

	v, ok := s.Pop()
	if !ok || v != "a" {
		t.Fatalf("expected first pop to return top duplicate a,true got %q,%v", v, ok)
	}
	if got, want := s.Values(), []string{"b", "a"}; !reflect.DeepEqual(got, want) {
		t.Fatalf("unexpected values after pop: got %v want %v", got, want)
	}
}
