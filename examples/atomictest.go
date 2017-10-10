// The primary mechanism for managing state in Go is
// communication over channels. We saw this for example
// with [worker pools](worker-pools). There are a few other
// options for managing state though. Here we'll
// look at using the `sync/atomic` package for _atomic
// counters_ accessed by multiple goroutines.

package main

import "fmt"
import (
	"runtime"
	"sync/atomic"
	"time"
)

func testatomic(num <-chan int, ops *int32) {
	atomic.AddInt32(ops, 1)
	a := <-num
	fmt.Println(a)
	// Allow other goroutines to proceed.
	runtime.Gosched() //让出时间片给其他线程
	fmt.Println("after:", a)

}

func main() {

	// We'll use an unsigned integer to represent our
	// (always-positive) counter.
	var ops int32 = 0

	// To simulate concurrent updates, we'll start 50
	// goroutines that each increment the counter about
	// once a millisecond.
	num := make(chan int, 5)
	for i := 0; i < 5; i++ {
		num <- i
		go testatomic(num, &ops)
	}
	time.Sleep(time.Second * 2)

	// Wait a second to allow some ops to accumulate.

	// In order to safely use the counter while it's still
	// being updated by other goroutines, we extract a
	// copy of the current value into `opsFinal` via
	// `LoadUint64`. As above we need to give this
	// function the memory address `&ops` from which to
	// fetch the value.
	opsFinal := atomic.LoadInt32(&ops)
	fmt.Println("ops:", opsFinal)
}
