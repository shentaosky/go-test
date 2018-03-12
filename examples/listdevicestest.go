package main

import (
    "github.com/rancher/convoy/clientapi"
    "fmt"
)

func main() {
    convoyclient, _ := clientapi.NewConvoyClient("0.0.0.0:8808")
    resList, err := convoyclient.ListDevices("bronze")
    if err != nil {
        fmt.Printf("%v", err)
    }
    if resList == nil {
        fmt.Println("resList is nil")
    }
    if resList.Items == nil {
        fmt.Println("items is nil")
    }
    if resList.Items != nil && len(resList.Items) == 0 {
        fmt.Println("items is 0")
    }
    fmt.Println("done")
}


