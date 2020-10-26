package main

import "fmt"

func main() {
    var testMap map[string]int
    //testMap := make(map[string]string)
    val, ok := testMap["a"]
    fmt.Println(ok, val)
    delete(testMap, "a")
    fmt.Println(testMap["a"])
}
