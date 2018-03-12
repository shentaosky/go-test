package main

import (
    "fmt"
    "os/exec"
    "sync"
    "runtime"
)

func main() {
    i := 0
    wg := &sync.WaitGroup{}
    wg.Add(100)
    fmt.Println(runtime.GOMAXPROCS(runtime.NumCPU()))
    for i < 100 {
        go execStart("kubectl", []string{"create", "-f", "/root/zookeeper-instance.yaml"}, wg)
        i++
    }
    wg.Wait()
}

func execStart(commmand string, args []string, wg *sync.WaitGroup)  {
    defer wg.Done()
    //time1 := time.Now().UnixNano()/1000000
    if err := exec.Command(commmand, args...).Start(); err != nil {
        fmt.Println("%v", err)
    }
    //time2 := time.Now().UnixNano()/1000000
    //fmt.Println(args[0], "time: ", time2-time1)
}
