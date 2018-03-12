package main

import (
	"fmt"
	//"sync"
	"runtime"
	"runtime/debug"
	"sync"
)

const maxVolumeNum = 50

func main() {
	fmt.Println(runtime.GOMAXPROCS(4))
	fmt.Println(debug.SetMaxThreads(5))
	wg := sync.WaitGroup{}
	for i := 0; i < maxVolumeNum; i++ {
		volName := fmt.Sprintf("volume_%d", i)
		wg.Add(1)
		go func() {
			fmt.Println(volName)
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("finish")
}
