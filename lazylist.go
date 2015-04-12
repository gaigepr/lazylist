package lazylist

import (
	"fmt"
	"sync"
)

const MaxUint = ^uint64(0)
const MinUint = 0

type LazyList struct {
	Head *Entry
	Tail *Entry
}

type Entry struct {
	Key    uint64
	Next   *Entry
	Marked bool
	Lock   sync.Mutex
}

func validate(pred *Entry, curr *Entry) bool {
	return !pred.Marked && !curr.Marked && pred.Next == curr
}

// Function that searches for an Entry by Key.
// Returns true if and only if an Entry with `key` is in the list and not marked.
func (L *LazyList) Contains(key uint64) bool {
	curr := L.Head
	for curr.Key < key {
		curr = curr.Next
	}
	return curr.Key == key && !curr.Marked
}

// An order n size function for the lazylist.
func (L *LazyList) Size() (total uint64) {
	curr := L.Head
	for curr != L.Tail {
		curr = curr.Next
		total += 1
	}
	// We subtract one to account for the head
	return total - 1
}

// Add a new Entry to the list.
// Returns true if and only if the Entry with Key `key` did not exists and does now.
func (L *LazyList) Add(key uint64) bool {
	for {
		pred := L.Head
		curr := L.Head.Next
		for curr.Key < key {
			pred = curr
			curr = curr.Next
		}
		pred.Lock.Lock()
		defer pred.Lock.Unlock()

		curr.Lock.Lock()
		defer curr.Lock.Unlock()

		if validate(pred, curr) {
			if curr.Key == key {
				return false
			} else {
				entry := Entry{key, curr, false, sync.Mutex{}}
				pred.Next = &entry
				return true
			}
		}
	}
}

// Remove a key.
// Returns true if and only if the key existed and was removed.
func (L *LazyList) Remove(key uint64) bool {
	for {
		pred := L.Head
		curr := L.Head.Next
		for curr.Key < key {
			pred = curr
			curr = curr.Next
		}
		pred.Lock.Lock()
		defer pred.Lock.Unlock()

		curr.Lock.Lock()
		defer curr.Lock.Unlock()

		if validate(pred, curr) {
			if curr.Key != key {
				return false
			} else {
				curr.Marked = true
				pred.Next = curr.Next
				return true
			}
		}
	}
}

// Returns a new empty LazyList.
func NewLazyList() LazyList {
	tail := Entry{MaxUint, nil, false, sync.Mutex{}}
	head := Entry{MinUint, &tail, false, sync.Mutex{}}
	list := LazyList{&head, &tail}
	return list
}

// Prints all the values in a lazy list, one line each.
func (L *LazyList) PrintLazyList() {
	curr := L.Head.Next
	for curr.Key < MaxUint {
		fmt.Println(curr.Key)
		curr = curr.Next
	}
}
