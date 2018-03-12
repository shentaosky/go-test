package main

import (
	"fmt"
	"unsafe"
	"math"
)

type node struct {
	left  *node
	right *node
	value int
}

var b = 1

func main() {
	fmt.Println(123)
	a := []int{10, 20, 30, 40, 50}
	fmt.Println(unsafe.Pointer(&a))

	b := a[1:2:3]
	fmt.Println(b)
	fmt.Println(unsafe.Pointer(&b))
// 切片的地址和数组地址不一样
	fmt.Println(unsafe.Pointer(&b[0]))
	fmt.Println(unsafe.Pointer(&a[1]))
	b = append(b, 60)
	fmt.Println(a)
	fmt.Println(b)
	a := 10
	go func() (err error) {
		fmt.Println(err)
	}()

}

type A struct {
	*a
}

type a struct {
	fd int
	name string
}

func traver123Inoder(root *node) []int {
	res := []int{}
	list := []*node{}
	for root != nil || len(list) != 0 {
		for root != nil {
			list = append(list, root)
			root = root.left
		}
		root = list[len(list)-1]
		list = list[:len(list)-1]
		res = append(res, root.value)
		root = root.right
	}
	return res
}
