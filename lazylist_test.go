package lazylist

import (
	"fmt"
	"math/rand"
	"testing"
)

// Test assures that a duplicate element cannot be added.
func TestAdd(t *testing.T) {
	fmt.Println("Testing  `Add`...")
	list := NewLazyList()
	// Try to add the same element twice
	list.Add(1)
	if list.Add(1) {
		t.Error("Duplicate entries.")
	}
}

// Tests linear inserts and Contains method.
func TestContains(t *testing.T) {
	fmt.Println("Testing `Contains`...")
	list := NewLazyList()
	// Insert some elements
	list.Add(5)
	list.Add(1)
	list.Add(12)
	// Make sure all elements are in there
	if !(list.Contains(5) && list.Contains(1) && list.Contains(12)) {
		t.Error("Some entries not present.")
	}
}

// Tests linear removal of elements.
func TestRemove(t *testing.T) {
	fmt.Println("Testing `Remove`...")
	list := NewLazyList()
	list.Add(666)
	list.Add(1)
	list.Add(6)
	if !list.Remove(6) {
		t.Error("Removal failed.")
	}
}

// Tests 1000 'concurrent' inserts at a time.
func BenchmarkAdd(b *testing.B) {
	list := NewLazyList()

	for n := 0; n < b.N; n++ {
		for i := 0; i < 1001; i++ {
			go list.Add(uint64(rand.Int()))
		}
	}

}
