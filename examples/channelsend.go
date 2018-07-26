package main

import (
    "fmt"
    "time"
    "os"
    "runtime/pprof"
)

func main() {
    go func() {
        fmt.Println("1")
        go func() {
            fmt.Println("3")
            time.Sleep(time.Second*2)
            fmt.Println("4")
        }()
        pprof.Lookup("goroutine").WriteTo(os.Stdout, 2)
        fmt.Println("2")
    }()
    time.Sleep(time.Second)
    time.Sleep(time.Second*5)
}
