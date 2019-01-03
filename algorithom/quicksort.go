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

// 从某一个数开始作为基准点, 基准点左边的数小于它, 右边的数大于它, 将数组划分成了两个子数组,
// 每个子数组再继续做这个操作.
func quickSort(data []int, begin, end int) {
	if begin > end {
		return
	}
	mid := partitiontest2(data, begin, end)
	quickSort(data, 0, mid-1)
	quickSort(data, mid+1, end)
}

// 这里选最后一个数做基准点
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

// 找任意一个基准点都可以, 这金额里选第一个数做基准点
func partitiontest(data []int, begin, end int) int {
	j := begin
	for i := begin + 1; i <= end; i++ {
		if data[i] < data[j] {
			swap(data, j, i)
			// 可以少交换几次, 但是这样写我更能理解一点
			swap(data, j+1, i)
			j++
		}
	}
	return j
}

func partitiontest2(data []int, begin, end int) int {
	benchmark := begin
	j := begin
	for i := begin+1; i <= end; i++ {
		if data[i] < data[benchmark] {
			j++
			swap(data, i ,j)
		}
	}
	swap(data, benchmark, j)
	return j
}