package main

import (
    "sync"
    "fmt"
    "os/exec"
)

func main() {
    wg := &sync.WaitGroup{}
    wg.Add(200)
    i := 0
    for i < 50 {
        go execStart("kubectl", []string{"create", "-f", "zookeeper-instance.yaml"}, wg)
        go execStart("curl", []string{"http://0.0.0.0:8808/metrics"}, wg)
        go execStart("kubectl", []string{"create", "-f", "zookeeper-instance.yaml"}, wg)
        go execStart("convoy", []string{"list"}, wg)
        go execStart("convoy", []string{"pool", "list"}, wg)
        i ++
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