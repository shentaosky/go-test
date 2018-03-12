package main

import (
    "k8s.io/apimachinery/pkg/util/json"
    "fmt"
)

func main() {
    origin := make(map[string]string)
    data := []string{
        "1111", "2222",
    }
    sourceData, err := json.Marshal(&data)
    if err != nil {
        fmt.Printf("err: %v", err)
        return
    }
    origin["data"] = string(sourceData)

    fmt.Println(origin)
    newData := []string{}
    json.Unmarshal([]byte(origin["data"]), &newData)
    fmt.Println(newData)

    a := []int{1, 2, 3, 4, 5}
    // fmt.Println(a[5]) // error index out of range
    fmt.Println(a[5:])
    b := []string{}
    if b == nil {
        fmt.Println("1")
    }
    if len(b) <= 0 {
        fmt.Println("2")
    }
    if len(getNil()) <= 0 {
        fmt.Println("3")
    }

}

func getNil() []string {
    return nil
}
