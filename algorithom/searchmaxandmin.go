package main

import (
	"fmt"
	"math/rand"
)

const Length = 11

func main() {
	data := []int{}
	for i := 0; i < Length; i++ {
		data = append(data, rand.Intn(10000))
	}
	fmt.Println(data)
	max, min := search2(data, 0, Length-1)
	fmt.Println("max: ", max, ", min: ", min)
}

func search(data []int, begin, end int) (int, int) {
	if begin >= end {
		return data[begin], data[begin]
	}
	maxLeft, minLeft := search(data, begin, begin+(end-begin)/2)
	maxRight, minRight := search(data, begin+(end-begin)/2+1, end)
	max := maxRight
	min := minRight
	if maxLeft > maxRight {
		max = maxLeft
	}
	if minLeft < minRight {
		min = minLeft
	}
	return max, min
}

func search2(data []int, begin, end int) (int, int) {
	max := data[begin]
	min := max
	for _, d := range data {
		if d > max {
			max = d
		}
		if d < min {
			min = d
		}
	}
	return max, min
}