package main

import (
    "fmt"
)

func main() {
    a := &volumeTest{name: "123"}
    change(*a)

    fmt.Println(*a)
}

type volumeTest struct {
    name string
}

func change(v volumeTest) {
    a := &v
    a.name = "2"
}
