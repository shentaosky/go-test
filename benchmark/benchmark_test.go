package main

import "testing"

func BenchmarkAdd1(b *testing.T) {
	for i := 0; i < 2; i++ {
		Add(1, 2)
	}
}

func Add(a, b int) int {
	return a + b
}
