package main

import (
    "fmt"
    "time"
)

func main() {
    rst := make(chan map[string]string)
    go input(rst)
    for {
        select {
        case i:=<-rst:
            fmt.Printf("rst test: %+v", i)
            for key, value := range i {
                fmt.Println(key, value)
            }
            return
        default:
            time.Sleep(time.Second)
        }

    }

}

func input(rst chan map[string]string) {
    rst1 := make(map[string]string)
    rst1["test1"] = "1"
    rst1 = map[string]string{}
    time.Sleep(time.Second*2)
    rst <- rst1
}
