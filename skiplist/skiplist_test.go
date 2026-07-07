package skiplist

import (
	"reflect"
	"testing"
)

func TestNewSkipListIsEmpty(t *testing.T) {
	list := New[int]()

	if !list.IsEmpty() {
		t.Fatal("expected new skip list to be empty")
	}
	if list.Len() != 0 {
		t.Fatalf("expected length 0, got %d", list.Len())
	}
	if got := list.Values(); len(got) != 0 {
		t.Fatalf("expected empty values, got %v", got)
	}
}

func TestInsertAndValuesInOrder(t *testing.T) {
	list := New[int]()
	inputs := []int{10, 5, 20, 15, 25, 3, 7}
	for _, v := range inputs {
		if inserted := list.Insert(v); !inserted {
			t.Fatalf("expected insert %d to succeed", v)
		}
	}

	if list.IsEmpty() {
		t.Fatal("expected non-empty skip list")
	}
	if list.Len() != len(inputs) {
		t.Fatalf("expected length %d, got %d", len(inputs), list.Len())
	}
	if got, want := list.Values(), []int{3, 5, 7, 10, 15, 20, 25}; !reflect.DeepEqual(got, want) {
		t.Fatalf("unexpected values: got %v want %v", got, want)
	}
}

func TestInsertDuplicate(t *testing.T) {
	list := New[int]()
	if inserted := list.Insert(8); !inserted {
		t.Fatal("expected first insert to succeed")
	}
	if inserted := list.Insert(8); inserted {
		t.Fatal("expected duplicate insert to fail")
	}
	if list.Len() != 1 {
		t.Fatalf("expected length 1, got %d", list.Len())
	}
}

func TestFind(t *testing.T) {
	list := New[int]()
	list.Insert(10)
	list.Insert(4)
	list.Insert(12)

	v, ok := list.Find(4)
	if !ok {
		t.Fatal("expected to find value 4")
	}
	if v != 4 {
		t.Fatalf("expected value 4, got %d", v)
	}
	if _, ok := list.Find(999); ok {
		t.Fatal("expected find of missing value to return false")
	}
}

func TestDeleteFromEmptyList(t *testing.T) {
	list := New[int]()
	if deleted := list.Delete(1); deleted {
		t.Fatal("expected delete on empty skip list to return false")
	}
}

func TestDeleteExistingValue(t *testing.T) {
	list := New[int]()
	for _, v := range []int{10, 5, 20, 15, 25, 3, 7} {
		list.Insert(v)
	}

	if deleted := list.Delete(20); !deleted {
		t.Fatal("expected delete to return true")
	}
	if got, want := list.Values(), []int{3, 5, 7, 10, 15, 25}; !reflect.DeepEqual(got, want) {
		t.Fatalf("unexpected values after delete: got %v want %v", got, want)
	}
	if list.Len() != 6 {
		t.Fatalf("expected length 6, got %d", list.Len())
	}
}

func TestDeleteMissingValue(t *testing.T) {
	list := New[int]()
	for _, v := range []int{1, 2, 3} {
		list.Insert(v)
	}

	if deleted := list.Delete(99); deleted {
		t.Fatal("expected delete for missing value to return false")
	}
	if got, want := list.Values(), []int{1, 2, 3}; !reflect.DeepEqual(got, want) {
		t.Fatalf("unexpected mutation after missing delete: got %v want %v", got, want)
	}
}

func TestDeleteRootToEmpty(t *testing.T) {
	list := New[int]()
	list.Insert(42)

	if deleted := list.Delete(42); !deleted {
		t.Fatal("expected delete to return true")
	}
	if !list.IsEmpty() {
		t.Fatal("expected skip list to be empty")
	}
	if list.Len() != 0 {
		t.Fatalf("expected length 0, got %d", list.Len())
	}
}

func TestStringValues(t *testing.T) {
	list := New[string]()
	for _, v := range []string{"pear", "apple", "orange", "banana"} {
		list.Insert(v)
	}

	if got, want := list.Values(), []string{"apple", "banana", "orange", "pear"}; !reflect.DeepEqual(got, want) {
		t.Fatalf("unexpected string values: got %v want %v", got, want)
	}
	if _, ok := list.Find("banana"); !ok {
		t.Fatal("expected to find banana")
	}
	if deleted := list.Delete("orange"); !deleted {
		t.Fatal("expected delete to return true")
	}
	if got, want := list.Values(), []string{"apple", "banana", "pear"}; !reflect.DeepEqual(got, want) {
		t.Fatalf("unexpected values after delete: got %v want %v", got, want)
	}
}

func TestForEach(t *testing.T) {
	list := New[int]()
	for _, v := range []int{10, 5, 20, 3, 7, 15, 25} {
		list.Insert(v)
	}

	var got []int
	list.ForEach(func(v int) bool {
		got = append(got, v)
		return true
	})
	if want := []int{3, 5, 7, 10, 15, 20, 25}; !reflect.DeepEqual(got, want) {
		t.Fatalf("ForEach full: got %v want %v", got, want)
	}

	got = nil
	list.ForEach(func(v int) bool {
		got = append(got, v)
		return v < 10
	})
	if want := []int{3, 5, 7, 10}; !reflect.DeepEqual(got, want) {
		t.Fatalf("ForEach early stop: got %v want %v", got, want)
	}
}
