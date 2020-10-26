package main

import (
    "fmt"
    "github.com/ceph/go-ceph/rados"
)

func main() {
    conn, err := rados.NewConnWithUser("lvs01cinder")
    if err != nil {
        fmt.Printf("error new conn: %v", err)
        return
    }
    conn.ReadDefaultConfigFile()
    args := []string{ "--mon-host", "10.66.166.39:6789"}
    if err := conn.ParseCmdLineArgs(args); err != nil {
        fmt.Printf("error to parse cmd line: %v", err)
        return
    }
    if err := conn.Connect(); err != nil {
        fmt.Printf("error to connect rados: %v", err)
        return
    }
    pools, err := conn.ListPools()
    if err != nil {
        fmt.Printf("error list pools: %v", err)
        return
    }
    fmt.Println(pools)
}
