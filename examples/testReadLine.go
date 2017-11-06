package main

import (
    "bufio"
    "strings"
    "fmt"
)

func main() {
    str := strings.NewReader("123\nsh ent  ao\nhahah a")
    r := bufio.NewReader(str)
    outs, _, err := r.ReadLine()
    if err != nil {
        fmt.Println("err")
    }
    fmt.Println(string(outs))

    outs, _, err = r.ReadLine()
    if err != nil {
        fmt.Println("err")
    }
    fmt.Println(string(outs))
    outs, _, err = r.ReadLine()
    if err != nil {
        fmt.Println("err")
    }
    fmt.Println(string(outs))
    outs, _, err = r.ReadLine()
    if err != nil {
        fmt.Println("err")
    }
    fmt.Println(string(outs))
}
