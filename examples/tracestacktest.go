package main

import (
	"os"
	"os/signal"
	"runtime/debug"
	"runtime/pprof"
	"syscall"
	"time"
)

func main() {
	go a()
	m1()
}
func m1() {
	m2()
}
func m2() {
	m3()
}
func m3() {
	pprof.Lookup("goroutine").WriteTo(os.Stdout, 1)
	debug.PrintStack()
	go setupSigusr1Trap()
	time.Sleep(time.Hour)
}

func setupSigusr1Trap() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGUSR1)
	go func() {
		for range c {
			debug.PrintStack()
		}
	}()
}

func a() {
	time.Sleep(time.Hour)
}
