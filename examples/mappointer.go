package main

import "fmt"

func main() {
    mapA := make(map[string]string)
    MapTestA := &MapTest{
        test: mapA,
    }

    if a, ok := mapA["3"]; !ok {
        fmt.Println(a)
    }

    MapTestA.test["1"] = "1"

    for key, value := range mapA {
        mapA[key] = value + "2"
    }

    fmt.Println(mapA)
}

type MapTest struct {
    test map[string]string `json:"-"`
}
