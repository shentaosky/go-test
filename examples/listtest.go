package main

import (
	"container/list"
	"fmt"
	"unsafe"
)

type test struct {
	valume int64
	number int32
}

func main() {
	l := list.New()
	l.PushBack(101)
	l.PushBack(102)
	l.PushBack(103)
	value := new(test)
	fmt.Println(unsafe.Offsetof(value.number))
	l.MoveBefore(l.Front().Next(), l.Front())
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	poollist := map[string]string{"1": "1", "2": "2"}
	poollist = map[string]string{}
	if _, exsited := poollist["1"]; exsited == false {
		fmt.Println(poollist)
	}
	poollist["1"] = "2"
}
