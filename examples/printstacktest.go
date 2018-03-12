package main

import (
    "fmt"
    "net/http"
    "os"
    "runtime/pprof"
    "syscall"
    "os/signal"
)

func main() {
    stopCh := make(chan bool)
    go func() {
        fmt.Println(http.ListenAndServe("localhost:6060", nil))
    }()

    go sigusr1TrapSetup()

    <-stopCh
}

func sigusr1TrapSetup() {
    c := make(chan os.Signal, 1)
    signal.Notify(c, syscall.SIGUSR1)
    go func() {
        for range c {
            fmt.Println("Start print daemon stack =======")
            for _, profile := range pprof.Profiles() {
                profile.WriteTo(os.Stdout, 2)
                fmt.Printf("\n")
            }
        }
    }()
}
