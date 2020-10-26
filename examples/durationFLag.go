package main

import (
    "fmt"
    "strings"
)


func main() {

    fmt.Println(strings.Split("test", "\n")[0])

}
//
//func isVmNode() bool {
//    out, err := exec.Command("sh", "-c", "cat /proc/cpuinfo | grep -q ^flags.*\\ hypervisor\\ && echo 'This is a VM'").CombinedOutput()
//    if err != nil {
//        fmt.Println("This is not a VM")
//        return false
//    }
//    if strings.Split(string(out), "\n")[0] == "This is a VM" {
//        fmt.Println("This is a VM")
//        return true
//    }
//    fmt.Println("Un-exception situation when check is vm node")
//    return false
//}