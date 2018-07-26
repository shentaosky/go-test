package main

import (
    "fmt"
    "time"
)

func f(left, right chan int) {
    left <- 1+ <- right
    fmt.Println("run f")
}

func main() {
    const n = 10000
    leftmost := make(chan int)
    left := leftmost
    right := leftmost
    for i := 0; i < 0; i++ {
        right = make(chan int)
        go f(left, right)
        left = right
    }
    go func(c chan int) {c <-12} (right)
    fmt.Println(<-leftmost)
}
//progress:

// {leftmost, left, right}

// {leftmost, left}  <- 1+ <- right

// leftmost <- 1+ <- {left, right}

// leftmost <- 1+ <- left <- 1+ <- right

// ...

// leftmost <- 1+ <- left <- 1+ <- right ...n... <- right < c <- 12
