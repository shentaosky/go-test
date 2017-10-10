package main

import (
	"fmt"
)

type node struct {
	left  *node
	right *node
	value int
}

var b = 1

func main() {
	fmt.Println(123)
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
