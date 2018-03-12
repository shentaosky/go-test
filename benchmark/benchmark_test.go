package main

import (
	"testing"
	"fmt"
)

func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

//// from fib_test.go
//func BenchmarkFib10(b *testing.B) {
//    // run the Fib function b.N times
//    for n := 0; n < b.N; n++ {
//        Fib(10)
//    }
//}

func TestRunSkip(t *testing.T) {
	fmt.Printf("1")
}

func benchmarkFib(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(i)
	}
}

func BenchmarkFib1(b *testing.B)  { benchmarkFib(1, b) }
func BenchmarkFib2(b *testing.B)  { benchmarkFib(2, b) }
func BenchmarkFib3(b *testing.B)  { benchmarkFib(3, b) }
func BenchmarkFib10(b *testing.B) { benchmarkFib(10, b) }
func BenchmarkFib20(b *testing.B) { benchmarkFib(20, b) }
func BenchmarkFib40(b *testing.B) { benchmarkFib(40, b) }

func BenchmarkFibWrong(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(n)
	}
}

func BenchmarkFibWrong2(b *testing.B) {
	Fib(b.N)
}