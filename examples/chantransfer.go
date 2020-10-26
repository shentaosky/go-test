package main

import "fmt"

type TestStr struct {
    Data []string
}

func main() {
    var datas []string
    testStr := &TestStr{
        Data: datas,
    }

    datas = append(datas, "123")
    fmt.Println(testStr.Data)
}