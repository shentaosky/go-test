package main

import (
	"fmt"
)

type testMap struct {
	name string
}

func main() {
	test2 := map[int]testMap{}
	test2[1] = testMap{"1"}

	fmt.Println(test2[2])

	//test[3]="456"
	//fmt.Println(test)
	//get(test, 2)
}
