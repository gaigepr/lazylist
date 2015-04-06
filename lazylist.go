package main

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

func (L *LazyList) Contains(key uint64) bool {
	curr := L.Head
	for curr.Key < key {
		curr = curr.Next
	}
	return curr.Key == key && !curr.Marked
}

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

func NewLazyList() LazyList {
	tail := Entry{MaxUint, nil, false, sync.Mutex{}}
	head := Entry{MinUint, &tail, false, sync.Mutex{}}
	list := LazyList{&head, &tail}
	return list
}

func (L *LazyList) PrintLazyList() {
	curr := L.Head.Next
	for curr.Key < MaxUint {
		fmt.Println(curr.Key)
		curr = curr.Next
	}
}
