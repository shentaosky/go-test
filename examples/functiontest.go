package main

import (
	"fmt"
	"io"
	"syscall"
)

func main() {
	func1(true, "123", "231")
}

func func1(a bool, args ...string) {
	stat := syscall.Stat_t{}
	err := syscall.Stat("/dev/sda", &stat)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Major:", uint64(stat.Rdev/256), "Minor:", uint64(stat.Rdev%256))
	fmt.Println(cap([]byte("asdbasd")))
	io.Closer()
}
