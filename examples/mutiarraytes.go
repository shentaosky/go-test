package main

import "fmt"

func main() {
    bitmap := make([][]int, 2)
    for i := 0; i < 2; i ++ {
        bitmap[i] = make([]int, 3)
    }
    bitmap[0][1] = 1
    //bitmap[0][1] = 2
    bitmap[1][0] = 3

    fmt.Println(bitmap)
}
