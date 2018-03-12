package main

import "time"
import (
    "testing"
    "fmt"
)

func TestTimeOut(t *testing.T) {
    for {
        time.Sleep(time.Minute * 30)
        fmt.Println("30m")
    }
}
