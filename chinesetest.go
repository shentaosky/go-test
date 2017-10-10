package main

import (
	"fmt"
	"unsafe"
)

func main() {
	str := "123沈涛abc"
	// 汉字三个字符，数字字母一个字符
	fmt.Println(len([]byte(str)))
	fmt.Println(str[0:10])
	// rune 就是int32, 只支持4个字节的中文字符，如utf-8
	var r = []rune(str)
	var R = []int32(str)
	length := len(r)
	fmt.Printf(" rune length: %v\n", length)
	fmt.Println(r, R)
	fmt.Println(subString(3, 5, r))

	//分配3个[]int, 预留10个[]int
	newValues := make([][]int, 3, 10)
	line1 := []int{1, 2, 3}
	line2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	line3 := line2[0:4]
	p2 := (*struct {
		array uintptr
		len   int
		cap   int
	})(unsafe.Pointer(&line2))
	p3 := (*struct {
		array uintptr
		len   int
		cap   int
	})(unsafe.Pointer(&line3))
	fmt.Println(p2)
	fmt.Println(p3)
	// 可以改变line2的数据，指向的是同一个地址
	// 而append之后的slice指向的是两个地址，不能改变原值
	line3[0] = 2
	newValues[0] = line1
	newValues = append(newValues, line2)
	fmt.Println(newValues)
}

func subString(start, end int, str []rune) string {
	//实际上取[start, end)左闭右开区间
	return string(str[start:end])
}
