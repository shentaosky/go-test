package main

import "fmt"

func main() {
	var a = [5]int{1, 2, 3, 4, 5}
	var b = a[1:]
	b[0]++
	fmt.Println(a, b)
}
