package main

import (
    "fmt"
    "k8s.io/kubernetes/pkg/kubectl/cmd/set"
    "reflect"
)

func main() {
    var str []string
    str = append(str, "1")
    str = append(str, "2")
    fmt.Println(len(str))
    fmt.Println(cap(str))
    str = append(str, "3")
    str = append(str, "4")
    str = append(str, "5")
    str = str[1:]
    fmt.Println(len(str))
    fmt.Println(cap(str))
    if str[4] == "" {
        fmt.Println("true")
    }
    fmt.Println(str)
    typed := reflect.TypeOf(str[0])
    typed.NumField()
    var i interface{}
    if typed == i.Fiel {

    }
    set.Queue

}
