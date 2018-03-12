package main

import "fmt"

// iota +1 once,
const (
    mutexLocked = 1 << iota // mutex is locked
    mutexWoken
    mutexWoken2
    mutexWaiterShift = iota // iota=3
    mutexWaiterShift2 = 1 + iota // iota=4
    mutexWaiterShift3 = iota / 2 // iota=5
)

func main() {
    fmt.Println(mutexLocked, mutexWoken, mutexWoken2, mutexWaiterShift, mutexWaiterShift2, mutexWaiterShift3)
    fmt.Println(2 &^ 6)
}
