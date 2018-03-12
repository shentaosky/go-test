package main

import (
    "fmt"
    "os/exec"
)

func main() {
    if output, err := exec.Command("lvremove", "-y", "test/testlv").CombinedOutput(); err != nil {
        fmt.Printf("RemoveLV error: %v\n", err)
    } else {
        fmt.Printf("succeed: %s", string(output))
    }
}
//
//func RemoveLV(vg, device string) error {
//    fmt.Printf("lvremove: %s", fmt.Sprintf("/dev/mapper/%s-%s\n", vg, device))
//    if output, err := exec.Command("lvremove", "-f", fmt.Sprintf("/dev/mapper/%s-%s", vg, device)).CombinedOutput(); err != nil {
//        fmt.Printf("RemoveLV error: %v\n", err)
//        return fmt.Errorf("%s : %v", string(output), err)
//    }
//    return nil
//}
