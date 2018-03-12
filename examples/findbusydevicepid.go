package main

import (
    "fmt"
    "os"
    "io/ioutil"
)

func main() {
    //name := "pvc-34a44385-df34-11e7-b4c5-0cc47ab1f826"
    //out, err := exec.Command("sh", "-c", "grep pvc-34a44385-df34-11e7-b4c5-0cc47ab1f826 /proc/*/mountstats").CombinedOutput()
    //if err != nil {
    //    fmt.Printf("%s: %v", string(out), err)
    //}
    //info := strings.TrimSpace(string(out))
    //lines := strings.Split(info, "\n")
    //deviceToPidMap := make(map[string][]string)
    //deviceToPidMap["pvc-34a44385-df34-11e7-b4c5-0cc47ab1f826"] = []string{}
    //for _, line := range lines {
    //    str := strings.Split(line, ":")
    //    deviceToPidMap[] = append(deviceToPidMap["pvc-34a44385-df34-11e7-b4c5-0cc47ab1f826"], str[0])
    //}
    //fmt.Println(deviceToPidMap)
    printBusyId("/proc/15")
}

func printBusyId(pidDir string) error {
    procDir, err := os.Open(pidDir)
    if err != nil {
        return fmt.Errorf("open %s dir failed: %v", pidDir, err)
    }
    defer procDir.Close()
    procFiles, err := procDir.Readdirnames(0)
    if err != nil {
        return fmt.Errorf("read proc files failed: %v", err)
    }
    for _, file := range procFiles {
        if file == "cmdline" {
            name, _ := ioutil.ReadFile("/proc/2623/"+ file)
            fmt.Printf("%s", string(name))
        }
    }
    return nil
}
// /proc/123/mountstats
