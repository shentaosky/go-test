package main

import (
	"fmt"
	"time"
)

const poolname = "silver"

func main() {
	goRun()
	fmt.Println("stop")
	time.Sleep(time.Second * 3)
}

func testGorun() {
	t := time.NewTicker(time.Second * 10)
	for {
		select {
		case <-t.C:
			return
		default:
			fmt.Println("run")
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func goRun() {
	go testGorun()
	time.Sleep(time.Second * 3)
}
