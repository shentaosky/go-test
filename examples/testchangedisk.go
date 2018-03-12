package main

import (
    "fmt"
    "strconv"
    "io/ioutil"
    "strings"
)

func main() {
    data, err := ioutil.ReadFile("/var/lib/rancher/convoy/iron/mounts/changedisk1/test_changedisk.txt")
    if err != nil {
        fmt.Printf("open failed: %v", err)
        return
    }
    datas := strings.Split(string(data), "\n")
    for i, data := range datas {
        index := strings.Index(data, "x")
        if index < 0 {
            fmt.Printf("number %d is error \n", i)
            return
        }
        if data[:index] != strconv.Itoa(i) {
            fmt.Printf("number not consistency at %d \n", i)
        }
    }
}
