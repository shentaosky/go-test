package main

import (
    "fmt"
)

type DeepCopy struct {
    Name string
    Age int
    InnerStruct
}

type InnerStruct struct {
    InnerName string
}

type Cache struct {
    Items map[string]interface{}
}

func main() {
    cache := &Cache{
        Items: make(map[string]interface{}),
    }
    item1 := DeepCopy{
        Name: "item1",
        Age:  16,
        InnerStruct: InnerStruct{
            InnerName: "newItem",
        },
    }
    cache.Items[item1.Name] = item1
    item1.Age = 18
    item1.InnerName = "newItem2"
    item2 := cache.Items[item1.Name].(DeepCopy)
    fmt.Println(item2.Age)
    fmt.Println(item2.InnerName)
    item2.Age = 20
    item1.InnerName = "newItem3"
    item2 = cache.Items[item1.Name].(DeepCopy)
    fmt.Println(item2.Age)
    fmt.Println(item2.InnerName)
}




