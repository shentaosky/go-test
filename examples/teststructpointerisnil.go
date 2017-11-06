package main

import "fmt"

type volume struct {
    name string
}

func main() {
    vol := &volume{}
    var v *volume
    if vol == nil {
        fmt.Println("1")
    }
    if v == nil {
        fmt.Println("2")
    }
    fmt.Println("name1:", vol.name)
    fmt.Println("name2:", v.name) //panic
}
