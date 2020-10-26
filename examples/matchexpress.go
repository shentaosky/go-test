package main

import (
    "fmt"
    "regexp"
)

func main() {
    fmt.Println(MatchDisk("sda", "sda"))
    fmt.Println(MatchDisk("fsdfsda", "sda"))
    fmt.Println(MatchDisk("sda123", "sda"))
    fmt.Println(MatchDisk("fdetr(sda", "sda"))
    fmt.Println(MatchDisk("sda asdfsf", "sda"))
    fmt.Println(MatchDisk("!@#$sda@#", "sda"))
}

func MatchDisk(msg, disk string) bool {
    reg := fmt.Sprintf("(^%s$)|(^%s[^a-zA-Z1-9])|([^a-zA-Z1-9]%s[^a-zA-Z1-9])|([^a-zA-Z1-9]%s$)", disk, disk, disk, disk)
    fmt.Println(reg)
    return regexp.MustCompile(reg).MatchString(msg)
}

