package main

import (
    "fmt"
    "strings"
    "strconv"
    "os/exec"
)

func main() {
    mountpoint := "/var/lib/kubelet/pods/4866d3af-adc6-11e71-9190-001e674ff5da/volumes/transwarp.io~tosdisk/transwarp"
    usedCmdArgs := fmt.Sprintf("df -h %s --output=used --block-size=kB |grep -v Used", mountpoint)
    out, err := exec.Command("sh", "-c", usedCmdArgs).CombinedOutput()
    if err != nil {
        fmt.Println("error: ", err)
    }
    end := strings.LastIndex(string(out), "kB")
    start := strings.LastIndex(string(out), " ") + 1
    usedSize, err := strconv.ParseInt(string(out)[start:end], 10, 64)
    if err != nil {
        fmt.Println("error: ", err)
    }
    fmt.Println(usedSize)
}
