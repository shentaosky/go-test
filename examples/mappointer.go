package main

import "fmt"

func main() {
    mapA := make(map[string]string)
    MapTestA := MapTest{
        test: mapA,
    }

    MapTestA.test["1"] = "1"

    fmt.Println(mapA)
}

type MapTest struct {
    test map[string]string
}
