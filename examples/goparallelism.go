package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func testParalellism(wg *sync.WaitGroup) {

	x := 0
	for i := 0; i < 10000000000; i++ {
		x += i
	}

	wg.Done()
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(128)
	fmt.Println(runtime.GOMAXPROCS(runtime.NumCPU()))
	fmt.Println(time.Now())
	for i := 0; i < 128; i++ {
		go testParalellism(wg)
	}

	wg.Wait()
	fmt.Println(time.Now())
}

/*
cpu=4, goroutine=4
2017-07-20 12:42:51.837951831 +0800 CST
2017-07-20 12:42:59.429465348 +0800 CST
main ok
7.6

cpu=1, goroutine=4
2017-07-20 12:41:59.010950512 +0800 CST
2017-07-20 12:42:13.219673073 +0800 CST
main ok
14.2s

cpu=2, goroutine=4
2017-07-20 12:43:34.131443008 +0800 CST
2017-07-20 12:43:43.528566805 +0800 CST
main ok
9.4s
*/

/*
cpu : 4 logiccpu: 4 goroutine= 8
2017-07-28 16:30:46.469992425 +0800 CST
2017-07-28 16:31:02.272081716 +0800 CST
15.8s
*/

/*
cpu : 4 logiccpu: 8 goroutine= 8
2017-07-28 16:29:15.018030858 +0800 CST
2017-07-28 16:29:30.237331274 +0800 CST
15.22s
*/

/*
cpu : 4 logiccpu: 4 goroutine= 16
2017-07-28 16:33:50.285149731 +0800 CST
2017-07-28 16:34:21.59633986 +0800 CST
31.311s
*/

/*
cpu : 4 logiccpu: 8 goroutine= 16
2017-07-28 16:35:06.614230251 +0800 CST
2017-07-28 16:35:37.737472576 +0800 CST
31.123s
*/

/* cpu 做切换
cpu : 4 logiccpu: 128 goroutine= 128
2017-07-28 16:37:44.750351603 +0800 CST
2017-07-28 16:41:54.259077836 +0800 CST
4min9.5s
*/

/* go runtime 做切换
cpu : 4 logiccpu: 4 goroutine= 128
2017-07-28 16:42:49.732924102 +0800 CST
2017-07-28 16:47:10.86925634 +0800 CST
4min21.1s
*/
