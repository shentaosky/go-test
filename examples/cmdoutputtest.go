package main

import (
    "os/exec"
    "strings"
    "fmt"
)

func main() {
    out, err := exec.Command(KUBECTL, "get", "pod", "testpod2caa-b4b").CombinedOutput()
    if err != nil {
        fmt.Printf("get resouce error: %v", err)
    }
    if strings.Contains(string(out), "NotFound") {
        fmt.Println("true")
    }
}
