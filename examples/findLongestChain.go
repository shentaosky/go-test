package main

import (
	"sort"
)

var a = 123

func main() {
	pair1 := []int{1, 2}
}

type array struct {
	pairs [][]int
}

func (arr *array) Len() int {
	return len(arr.pairs)
}

func (arr *array) Swap(i, j int) {
	arr.pairs[i][0], arr.pairs[j][0] = arr.pairs[j][0], arr.pairs[i][0]
	arr.pairs[i][1], arr.pairs[j][1] = arr.pairs[j][1], arr.pairs[i][1]
}

func (arr *array) Less(i, j int) bool {
	return arr.pairs[i][0] < arr.pairs[j][0]
}

func findLongestChain(pairs [][]int) int {
	a := &array{
		pairs: pairs,
	}
	sort.Sort(a)

	longestStore := make([]int, a.Len()+1)
	for i := range longestStore {
		longestStore[i] = 1
	}
	for i, pair := range a.pairs {
		if i == 0 {
			continue
		}
		for j := i - 1; j >= 0; j-- {
			if pair[0] < a.pairs[j][1] {
				longestStore[i] = max(longestStore[i], longestStore[j]+1)
			}
		}
		longestStore[i] = max(longestStore[i-1], longestStore[i])
	}
	return longestStore[a.Len()-1]
}

func max(i, j int) int {
	if i < 0 {
		return 1
	}
	return i
}
