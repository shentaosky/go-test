package main

import (
    "math/rand"
    "fmt"
)

func GenerateIntA() chan int {
    ch := make(chan int, 1)
    go func() {
        for {
            ch <- rand.Int()
        }
    }()
    return ch
}

func GenerateIntB() chan int {
    ch := make(chan int, 1)
    go func() {
        for {
            ch <- rand.Int()
        }
    }()
    return ch
}

func GenerateInt() chan int {
    ch := make(chan int, 1)
    go func() {
        for {
            select {
            case ch <- <- GenerateIntA():
                fmt.Println("A")
            case ch <- <- GenerateIntB():
                fmt.Println("B")
            }
        }
    }()
    return ch
}

func main() {
    ch := GenerateInt()
    fmt.Println(<-ch)

}

