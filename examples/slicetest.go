package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a = [6]int{1, 2, 3, 4, 5, 6}
	b := []int{}
	for _, j := range a {
		b = append(b, j)
		p := unsafe.Pointer(&b)
		fmt.Println(p, cap(b), len(b))
	}
}
