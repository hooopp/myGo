package main

import (
	"testing"
)

func Add(a, b int) int {
	return a + b
}

func BenchmarkAdd(b *testing.B) {
	for range b.N {
		Add(2, 3)
	}
}
