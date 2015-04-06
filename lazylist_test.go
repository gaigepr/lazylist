package main

import (
	"math/rand"
	"testing"
)

func BenchmarkAdd(b *testing.B) {
	list := NewLazyList()
	for n := 0; n < b.N; n++ {
		for i := 0; i < 1001; i++ {
			go list.Add(uint64(rand.Int()))
		}
	}
}
