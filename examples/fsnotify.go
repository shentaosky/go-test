package main

import (
    "fmt"
    "github.com/fsnotify/fsnotify"
)

func main() {
    fw, _ := fsnotify.NewWatcher()
    err := fw.Add("/Users/tashen/test/test1")
    if err != nil {
        fmt.Println(err)
    }
    for {
        select {
        case  e := <- fw.Events:
            fmt.Println(e.String())
            fmt.Println(e.Name)
            fmt.Println(e.Op.String())
            return
        default:

        }
    }
}
