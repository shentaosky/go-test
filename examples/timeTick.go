package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Println(os.Getpid())
    //for {
    //    select {
    //
    //    case <-time.Tick(time.Second):
    //        fmt.Println("1")
    //    }
    //}

    fmt.Println(string([]byte{114,101,97, 108, 116, 105, 109, 101, 58, 32, 112, 114, 105, 111, 32, 48, 10}))
}
