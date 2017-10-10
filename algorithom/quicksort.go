package main

import (
	"fmt"
	"github.com/docker/docker/pkg/random"
)

const Length = 10

func main() {
	data := []int{}
	for i := 0; i < Length; i++ {
		data = append(data, random.Rand.Intn(100))
	}
	fmt.Println(data)
	quickSort(data, 0, Length-1)
	fmt.Println(data)
}

func quickSort(data []int, begin, end int) {
	if begin > end {
		return
	}
	mid := partition(data, begin, end)
	quickSort(data, 0, mid-1)
	quickSort(data, mid+1, end)
}

func partition(data []int, begin, end int) int {
	mid := begin
	for j := begin; j < end; j++ {
		if data[j] < data[end] {
			swap(data, mid, j)
			mid++
		}
	}
	swap(data, mid, end)
	return mid
}

func swap(data []int, i, j int) {
	if i == j {
		return
	}
	data[i] ^= data[j]
	data[j] ^= data[i]
	data[i] ^= data[j]
}
