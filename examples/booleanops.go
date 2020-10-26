package main

import (
    "fmt"
    "strings"
)

type Colume struct {
    VolumeSource
}

type VolumeSource struct {
    age string
}


func main() {
    c := Colume{
        VolumeSource {
            age: "12",
        },
    }
    fmt.Println(c.age)
    fmt.Println((&c).age)
    teststr := "dev dev "
    teststr = teststr[:len(teststr)-1]
    fmt.Println(teststr)
    ataDiskMap := make(map[string]string)
    res := strings.Contains("1231231", ataDiskMap["sda"])
    fmt.Println(res)
}
