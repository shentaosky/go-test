package main

import (
	"fmt"
)

func main() {
	str := "Hello,123"
	n := len(str)
JLoop:
	for i := 0; i < n; i++ {
		ch := str[i] // 依据下标取字符串中的字符,类型为byte
		fmt.Printf("%d, %v\n", i, ch)
		break JLoop
	}
	fmt.Println("loop")
}
