package heap

import (
	"testing"
)

func TestNewHeapIsEmpty(t *testing.T) {
	h := New[int]()
	if !h.IsEmpty() {
		t.Fatal("expected new heap to be empty")
	}
	if h.Len() != 0 {
		t.Fatalf("expected length 0, got %d", h.Len())
	}
}

func TestPeekAndPopOnEmpty(t *testing.T) {
	h := New[int]()
	if _, ok := h.Peek(); ok {
		t.Fatal("expected peek on empty heap to return false")
	}
	if _, ok := h.Pop(); ok {
		t.Fatal("expected pop on empty heap to return false")
	}
}

func TestMinHeapOrder(t *testing.T) {
	h := New[int]()
	for _, v := range []int{5, 1, 4, 2, 3} {
		h.Push(v)
	}
	if h.Len() != 5 {
		t.Fatalf("expected length 5, got %d", h.Len())
	}

	v, ok := h.Peek()
	if !ok || v != 1 {
		t.Fatalf("expected peek 1, got %d ok=%v", v, ok)
	}

	want := []int{1, 2, 3, 4, 5}
	for _, w := range want {
		got, ok := h.Pop()
		if !ok {
			t.Fatal("expected pop to succeed")
		}
		if got != w {
			t.Fatalf("expected pop %d, got %d", w, got)
		}
	}
	if !h.IsEmpty() {
		t.Fatal("expected heap to be empty after all pops")
	}
}

func TestMaxHeapOrder(t *testing.T) {
	h := NewMax[int]()
	for _, v := range []int{3, 1, 5, 2, 4} {
		h.Push(v)
	}

	v, ok := h.Peek()
	if !ok || v != 5 {
		t.Fatalf("expected peek 5, got %d ok=%v", v, ok)
	}

	want := []int{5, 4, 3, 2, 1}
	for _, w := range want {
		got, ok := h.Pop()
		if !ok {
			t.Fatal("expected pop to succeed")
		}
		if got != w {
			t.Fatalf("expected pop %d, got %d", w, got)
		}
	}
}

func TestNewFuncCustomOrder(t *testing.T) {
	// Min-heap by string length.
	h := NewFunc(func(a, b string) bool {
		return len(a) < len(b)
	})
	h.Push("banana")
	h.Push("fig")
	h.Push("apple")
	h.Push("kiwi")

	got, ok := h.Pop()
	if !ok || got != "fig" {
		t.Fatalf("expected shortest string fig, got %q ok=%v", got, ok)
	}
}

func TestPushPopSingleElement(t *testing.T) {
	h := New[string]()
	h.Push("only")

	v, ok := h.Peek()
	if !ok || v != "only" {
		t.Fatalf("expected peek only, got %q ok=%v", v, ok)
	}
	v, ok = h.Pop()
	if !ok || v != "only" {
		t.Fatalf("expected pop only, got %q ok=%v", v, ok)
	}
	if !h.IsEmpty() {
		t.Fatal("expected heap to be empty")
	}
}

func TestHeapPropertyAfterManyPushPop(t *testing.T) {
	h := New[int]()
	// Push 0..99 in reverse order.
	for i := 99; i >= 0; i-- {
		h.Push(i)
	}
	prev, _ := h.Pop()
	for !h.IsEmpty() {
		v, _ := h.Pop()
		if v < prev {
			t.Fatalf("heap order violated: popped %d after %d", v, prev)
		}
		prev = v
	}
}

func TestDuplicateValues(t *testing.T) {
	h := New[int]()
	h.Push(3)
	h.Push(1)
	h.Push(1)
	h.Push(2)

	if h.Len() != 4 {
		t.Fatalf("expected length 4, got %d", h.Len())
	}
	v, _ := h.Pop()
	if v != 1 {
		t.Fatalf("expected first pop 1, got %d", v)
	}
	v, _ = h.Pop()
	if v != 1 {
		t.Fatalf("expected second pop 1, got %d", v)
	}
}
