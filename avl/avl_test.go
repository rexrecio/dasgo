package avl
package avl

import (
	"math"
	"reflect"
	"testing"
)

func TestNewTreeIsEmpty(t *testing.T) {
	tree := New[int]()
	if !tree.IsEmpty() {
		t.Fatal("expected new tree to be empty")
	}
	if tree.Len() != 0 {
		t.Fatalf("expected length 0, got %d", tree.Len())
	}
	if got := tree.Values(); len(got) != 0 {
		t.Fatalf("expected empty values, got %v", got)
	}
}

func TestInsertAndValuesInOrder(t *testing.T) {
	tree := New[int]()
	inputs := []int{10, 5, 20, 15, 25, 3, 7}
	for _, v := range inputs {
		if inserted := tree.Insert(v); !inserted {
			t.Fatalf("expected insert %d to succeed", v)
		}
	}
	if tree.IsEmpty() {
		t.Fatal("expected non-empty tree")
	}
	if tree.Len() != len(inputs) {
		t.Fatalf("expected length %d, got %d", len(inputs), tree.Len())
	}
	if got, want := tree.Values(), []int{3, 5, 7, 10, 15, 20, 25}; !reflect.DeepEqual(got, want) {
		t.Fatalf("unexpected values: got %v want %v", got, want)
	}
}

func TestInsertDuplicate(t *testing.T) {
	tree := New[int]()
	if inserted := tree.Insert(8); !inserted {
		t.Fatal("expected first insert to succeed")
	}
	if inserted := tree.Insert(8); inserted {
		t.Fatal("expected duplicate insert to fail")
	}
	if tree.Len() != 1 {
		t.Fatalf("expected length 1, got %d", tree.Len())
	}
}

func TestFind(t *testing.T) {
	tree := New[int]()
	tree.Insert(10)
	tree.Insert(4)
	tree.Insert(12)

	v, ok := tree.Find(4)
	if !ok {
		t.Fatal("expected to find value 4")
	}
	if v != 4 {
		t.Fatalf("expected value 4, got %d", v)
	}
	if _, ok := tree.Find(999); ok {
		t.Fatal("expected find of missing value to return false")
	}
}

func TestDeleteFromEmptyTree(t *testing.T) {
	tree := New[int]()
	if deleted := tree.Delete(1); deleted {
		t.Fatal("expected delete on empty tree to return false")
	}
}

func TestDeleteLeafNode(t *testing.T) {
	tree := New[int]()
	for _, v := range []int{10, 5, 20} {
		tree.Insert(v)
	}
	if deleted := tree.Delete(5); !deleted {
		t.Fatal("expected leaf delete to return true")
	}
	if got, want := tree.Values(), []int{10, 20}; !reflect.DeepEqual(got, want) {
		t.Fatalf("unexpected values after leaf delete: got %v want %v", got, want)
	}
	if tree.Len() != 2 {
		t.Fatalf("expected length 2, got %d", tree.Len())
	}
}

func TestDeleteNodeWithOneChild(t *testing.T) {
	tree := New[int]()
	for _, v := range []int{10, 5, 20, 17} {
		tree.Insert(v)
	}
	if deleted := tree.Delete(20); !deleted {
		t.Fatal("expected delete to return true")
	}
	if got, want := tree.Values(), []int{5, 10, 17}; !reflect.DeepEqual(got, want) {
		t.Fatalf("unexpected values after delete: got %v want %v", got, want)
	}
	if tree.Len() != 3 {
		t.Fatalf("expected length 3, got %d", tree.Len())
	}
}

func TestDeleteNodeWithTwoChildren(t *testing.T) {
	tree := New[int]()
	for _, v := range []int{10, 5, 20, 15, 25, 13, 17} {
		tree.Insert(v)
	}
	if deleted := tree.Delete(20); !deleted {
		t.Fatal("expected delete to return true")
	}
	if got, want := tree.Values(), []int{5, 10, 13, 15, 17, 25}; !reflect.DeepEqual(got, want) {
		t.Fatalf("unexpected values after two-child delete: got %v want %v", got, want)
	}
	if tree.Len() != 6 {
		t.Fatalf("expected length 6, got %d", tree.Len())
	}
}

func TestDeleteRootToEmpty(t *testing.T) {
	tree := New[int]()
	tree.Insert(42)
	if deleted := tree.Delete(42); !deleted {
		t.Fatal("expected delete root to return true")
	}
	if !tree.IsEmpty() {
		t.Fatal("expected tree to be empty")
	}
	if tree.Len() != 0 {
		t.Fatalf("expected length 0, got %d", tree.Len())
	}
}

func TestStringValues(t *testing.T) {
	tree := New[string]()
	for _, v := range []string{"pear", "apple", "orange", "banana"} {
		tree.Insert(v)
	}
	if got, want := tree.Values(), []string{"apple", "banana", "orange", "pear"}; !reflect.DeepEqual(got, want) {
		t.Fatalf("unexpected string values: got %v want %v", got, want)
	}
	if _, ok := tree.Find("banana"); !ok {
		t.Fatal("expected to find banana")
	}
	if deleted := tree.Delete("orange"); !deleted {
		t.Fatal("expected delete to return true")
	}
	if got, want := tree.Values(), []string{"apple", "banana", "pear"}; !reflect.DeepEqual(got, want) {
		t.Fatalf("unexpected values after delete: got %v want %v", got, want)
	}
}

// TestBalanceProperty verifies that the tree height stays within the AVL bound
// of 1.44*log2(n+2) even when values are inserted in ascending sorted order,
// which is the worst case for an unbalanced BST.
func TestBalanceProperty(t *testing.T) {
	const n = 1023 // 2^10 - 1
	tree := New[int]()
	for i := 1; i <= n; i++ {
		tree.Insert(i)
	}
	maxHeight := int(math.Ceil(1.44 * math.Log2(float64(n+2))))
	got := nodeHeight(tree.root)
	if got > maxHeight {
		t.Fatalf("tree height %d exceeds AVL bound %d for %d nodes", got, maxHeight, n)
	}
}
