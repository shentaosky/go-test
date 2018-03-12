package main

import (
    "net"
    "fmt"
)

func main() {
    ips, err := net.LookupIP("shentao-ThinkPad-T450")
    if err != nil {
         fmt.Println(err)
    }
    fmt.Println(ips)
}
