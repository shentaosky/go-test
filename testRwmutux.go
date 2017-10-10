package main

import (
	"sync"
)

type rwmutexTest struct {
	mutex *sync.RWMutex
	name  int
}

func main() {
	u := make(chan int)
	test1 := &rwmutexTest{
		name: 0,
	}
	test1.mutex = &sync.RWMutex{}
	go test1.readTest()
	go test1.writeTest(u)
	<-u
}

func (t *rwmutexTest) readTest() {
	for {
		t.mutex.RLock()
		println("Read name: ", t.name)
		t.mutex.RUnlock()
	}
}

func (t *rwmutexTest) writeTest(u chan int) {
	for {
		println("Write name: ", t.name)
		t.mutex.Lock()
		t.name++
		println("Write done: ", t.name)
		if t.name == 1000 {
			u <- 1
		}
		t.mutex.Unlock()
	}
}
