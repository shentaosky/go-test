package main

import "fmt"

func main() {
    clusterArray := []string{
        "1", "2", "3", "4", "5",
    }
    clusterArray[2:] = clusterArray[3:] // error
    fmt.Println(clusterArray)
}
