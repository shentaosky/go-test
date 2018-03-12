package main

import (
    "os/exec"
    "fmt"
    "path/filepath"
)

func main() {
    out, err := exec.Command("sh", "-c", "mount |grep "+ "/tmp/stress-acceptance/9d0b/kubelet/pods/47d3f564-249c-11e8-8f65-b083feb83778/volumes/transwarp.io~tosdisk/pvc-47cc5501-249c-11e8-8f65-b083feb83778").CombinedOutput()
    fmt.Println(filepath.Dir("/tmp/stress-acceptance/9d0b/kubelet/pods/47d3f564-249c-11e8-8f65-b083feb83778/volumes/transwarp.io~tosdisk/pvc-47cc5501-249c-11e8-8f65-b083feb83778"))
    fmt.Printf("%s: %v", string(out), err)
}
