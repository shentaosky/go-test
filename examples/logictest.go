package main

import "fmt"

func main() {
    a := 2
    b := 2
    c := 4
    if a == 1 && b == 3 || c == 4 {
        //report 1
        fmt.Println(1)
    }
    if a == 1 && (b == 3 || c == 4) {
        // not report
        fmt.Println(2)
    }
}
