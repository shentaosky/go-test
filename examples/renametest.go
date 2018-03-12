package main

import (
    "os"
    "fmt"
)

func main() {
    err := os.Rename("/mnt/test", "/mnt/test_new")
    if err != nil {
        fmt.Println("failed: %v", err)
    }
}
