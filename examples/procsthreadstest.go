package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    t1 := time.Now()
    wg := &sync.WaitGroup{}
    wg.Add(10)
    for i := 0; i <10; i++ {
        go count(wg)
    }
    wg.Wait()
    fmt.Println(time.Now().Sub(t1).String())
}

func count(wg *sync.WaitGroup) {
    c := 100
    for i := 0; i < 100000000; i++ {
        c += i
    }
    wg.Done()
}
