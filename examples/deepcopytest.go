package main

import (
    "github.com/mohae/deepcopy"
    "fmt"
)

type TestDeepCopy struct {
    MapA map[string]string
    Name string
    Age int64
}

func main() {

    testA := TestDeepCopy{
        MapA: map[string]string{
            "1": "1",
        },
        Name: "2",
        Age: 12,
    }
    testB := deepcopy.Copy(testA)
    testA.Age = 13



    fmt.Println(testB)
}
