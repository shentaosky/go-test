package main

import (
	"fmt"
	"time"
)

func neverStop(stopCh chan struct{}) {
	fmt.Println("get in")
	ch := make(chan int, 1)
	for {
		select {
		case <-stopCh:
			fmt.Println("case1")
			//*wait <- 1
			return
		case stopCh <- struct{}{}:
			fmt.Println("case2")
			return
		default:
			select {
			case ch <- 0:
			case ch <- 1:
			}
			i := <-ch
			fmt.Println("Value received:", i)
			time.Sleep(time.Millisecond * 100)
		}
	}
	fmt.Println("get out")
}

func testTimeOut(ch chan int) {
	var timeout = make(chan bool)
	go func() {
		time.Sleep(time.Second * 1) // wait 1 sec
		timeout <- true
		close(ch)
	}()
	for {
		select {
		case i := <-ch:
			fmt.Println("Value received:", i)
			return
		case <-timeout:
			fmt.Println("Time out")
			return
		default:
		}
	}
}

func main() {
	//var NeverStop = make(chan struct{})
	var ch = make(chan int)
	//messages := make(chan string)
	//wait := make(chan int)
	//go neverStop(NeverStop)
	//time.Sleep(time.Second*2)
	////NeverStop <- struct {}{}   // case1
	////<-wait
	//<- NeverStop     //case2
	go testTimeOut(ch)
	time.Sleep(time.Second * 3)
	if _, ok := <-ch; ok {
		ch <- 1
	}
	//go func() {
	//    time.Sleep(time.Second*10)
	//    messages <- "string"
	//}()
	//fmt.Println(<-messages)
}
