package main

import (
	"sync"
	//"time"
	"sync/atomic"
)

var a string
var once sync.Once

func setup() {
	a = "hello, world"
	print(a)
}
func doprint() {
	once.Do(setup)
}
func main() {
	//go doprint()
	//go doprint()
	//time.Sleep(1e9)
	a := 0
	b := 1
	var d int64
	atomic.CompareAndSwapInt64(&d, int64(a), int64(b))
	// if d==a : d<-b; return true; else: return false;
	print(d)
	print(a)
	print(b)

}
