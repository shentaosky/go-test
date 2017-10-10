package main

import (
	"fmt"
	"sync"
)

const maxVolumeNum = 10

func main() {
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
