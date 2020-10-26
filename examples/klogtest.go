package main

import (
    "fmt"
    "os"
    "time"
)

func main() {
    stop := make(chan struct{})
    go func(stop chan struct{}) {
        for {
            select {
            case <-stop:
                fmt.Println("123")
                return
            default:
                fmt.Println("1")
                time.Sleep(time.Minute*5)
            }
        }
    }(stop)
    time.Sleep(time.Second*5)
    // Run forever
    os.Exit(1)
}
