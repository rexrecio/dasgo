package linkedlist

import (
	"reflect"
	"testing"
)

func TestNewListIsEmpty(t *testing.T) {
	list := New[int]()

	if !list.IsEmpty() {
		t.Fatal("expected new list to be empty")
	}
	if list.Len() != 0 {
		t.Fatalf("expected length 0, got %d", list.Len())
	}
	if got := list.Values(); len(got) != 0 {
		t.Fatalf("expected empty values, got %v", got)
	}
}

func TestAppendAndPrependOrder(t *testing.T) {
	list := New[int]()
	list.Append(10)
	list.Append(20)
	list.Prepend(5)

	want := []int{5, 10, 20}
	if got := list.Values(); !reflect.DeepEqual(got, want) {
		t.Fatalf("unexpected values: got %v want %v", got, want)
	}
	if list.Len() != 3 {
		t.Fatalf("expected length 3, got %d", list.Len())
	}
	if list.IsEmpty() {
		t.Fatal("expected non-empty list")
	}
}

func TestFind(t *testing.T) {
	list := New[int]()
	list.Append(7)
	list.Append(14)

	v, ok := list.Find(14)
	if !ok {
		t.Fatal("expected to find value 14")
	}
	if v != 14 {
		t.Fatalf("expected value 14, got %d", v)
	}

	if _, ok := list.Find(99); ok {
		t.Fatal("expected find of missing value to return false")
	}
}

func TestDeleteFromEmptyList(t *testing.T) {
	list := New[int]()

	if deleted := list.Delete(1); deleted {
		t.Fatal("expected delete on empty list to return false")
	}
	if list.Len() != 0 {
		t.Fatalf("expected length 0, got %d", list.Len())
	}
}

func TestDeleteHeadMiddleAndTail(t *testing.T) {
	list := New[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)

	if deleted := list.Delete(1); !deleted {
		t.Fatal("expected head delete to return true")
	}
	if got, want := list.Values(), []int{2, 3, 4}; !reflect.DeepEqual(got, want) {
		t.Fatalf("after head delete: got %v want %v", got, want)
	}

	if deleted := list.Delete(3); !deleted {
		t.Fatal("expected middle delete to return true")
	}
	if got, want := list.Values(), []int{2, 4}; !reflect.DeepEqual(got, want) {
		t.Fatalf("after middle delete: got %v want %v", got, want)
	}

	if deleted := list.Delete(4); !deleted {
		t.Fatal("expected tail delete to return true")
	}
	if got, want := list.Values(), []int{2}; !reflect.DeepEqual(got, want) {
		t.Fatalf("after tail delete: got %v want %v", got, want)
	}

	if list.Len() != 1 {
		t.Fatalf("expected length 1, got %d", list.Len())
	}
}

func TestDeleteOnlyElementResetsList(t *testing.T) {
	list := New[int]()
	list.Append(42)

	if deleted := list.Delete(42); !deleted {
		t.Fatal("expected delete to return true")
	}
	if !list.IsEmpty() {
		t.Fatal("expected list to be empty after deleting only element")
	}
	if list.Len() != 0 {
		t.Fatalf("expected length 0, got %d", list.Len())
	}
	if got := list.Values(); len(got) != 0 {
		t.Fatalf("expected empty values, got %v", got)
	}
}

func TestDeleteMissingValue(t *testing.T) {
	list := New[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	if deleted := list.Delete(99); deleted {
		t.Fatal("expected delete for missing value to return false")
	}
	if got, want := list.Values(), []int{1, 2, 3}; !reflect.DeepEqual(got, want) {
		t.Fatalf("unexpected mutation after missing delete: got %v want %v", got, want)
	}
	if list.Len() != 3 {
		t.Fatalf("expected length 3, got %d", list.Len())
	}
}

func TestStringValues(t *testing.T) {
	list := New[string]()
	list.Append("alpha")
	list.Append("beta")
	list.Prepend("zero")

	if got, want := list.Values(), []string{"zero", "alpha", "beta"}; !reflect.DeepEqual(got, want) {
		t.Fatalf("unexpected string values: got %v want %v", got, want)
	}
	if _, ok := list.Find("alpha"); !ok {
		t.Fatal("expected to find alpha")
	}
	if deleted := list.Delete("beta"); !deleted {
		t.Fatal("expected to delete beta")
	}
	if got, want := list.Values(), []string{"zero", "alpha"}; !reflect.DeepEqual(got, want) {
		t.Fatalf("unexpected values after delete: got %v want %v", got, want)
	}
}

func TestFront(t *testing.T) {
	list := New[int]()

	if _, ok := list.Front(); ok {
		t.Fatal("expected empty list front to fail")
	}

	list.Append(10)
	list.Prepend(5)

	v, ok := list.Front()
	if !ok || v != 5 {
		t.Fatalf("expected front 5,true got %d,%v", v, ok)
	}
}

func TestPopFront(t *testing.T) {
	list := New[int]()

	if _, ok := list.PopFront(); ok {
		t.Fatal("expected pop front on empty list to fail")
	}

	list.Append(1)
	list.Append(2)
	list.Append(3)

	v, ok := list.PopFront()
	if !ok || v != 1 {
		t.Fatalf("expected first pop front 1,true got %d,%v", v, ok)
	}
	v, ok = list.PopFront()
	if !ok || v != 2 {
		t.Fatalf("expected second pop front 2,true got %d,%v", v, ok)
	}
	v, ok = list.PopFront()
	if !ok || v != 3 {
		t.Fatalf("expected third pop front 3,true got %d,%v", v, ok)
	}
	if _, ok = list.PopFront(); ok {
		t.Fatal("expected pop front on empty list to fail")
	}

	if list.Len() != 0 {
		t.Fatalf("expected length 0, got %d", list.Len())
	}
	if !list.IsEmpty() {
		t.Fatal("expected list to be empty")
	}
}

func TestFindFunc(t *testing.T) {
	type user struct {
		ID   int
		Name string
	}

	list := New[user]()
	list.Append(user{ID: 1, Name: "alpha"})
	list.Append(user{ID: 2, Name: "beta"})

	v, ok := list.FindFunc(func(value user) bool {
		return value.ID == 2
	})
	if !ok {
		t.Fatal("expected to find user with ID 2")
	}
	if v.Name != "beta" {
		t.Fatalf("expected matched user beta, got %q", v.Name)
	}
}

func TestDeleteFunc(t *testing.T) {
	type user struct {
		ID   int
		Name string
	}

	list := New[user]()
	list.Append(user{ID: 1, Name: "alpha"})
	list.Append(user{ID: 2, Name: "beta"})
	list.Append(user{ID: 3, Name: "gamma"})

	deleted := list.DeleteFunc(func(value user) bool {
		return value.ID == 2
	})
	if !deleted {
		t.Fatal("expected delete func to return true")
	}

	if got, want := list.Values(), []user{{ID: 1, Name: "alpha"}, {ID: 3, Name: "gamma"}}; !reflect.DeepEqual(got, want) {
		t.Fatalf("unexpected values after delete func: got %v want %v", got, want)
	}

	deleted = list.DeleteFunc(func(value user) bool {
		return value.ID == 99
	})
	if deleted {
		t.Fatal("expected delete func to return false for missing item")
	}
}
