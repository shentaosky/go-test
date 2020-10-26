package main

import (
    "fmt"
    "time"
)

const timest = time.Hour*72

func main() {
    fmt.Println(time.Now())
    fmt.Println(time.Now().Add((-1)*timest))
}
