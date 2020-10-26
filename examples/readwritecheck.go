package main

import (
    "fmt"
    "strings"
    "time"
)

func main() {
    //isFailure, messages := readWriteCheck("/var/vol-pool/nopv/6", "test")
    //out, err := ioutil.ReadFile("/proc/mounts")
    //if err != nil {
    //    fmt.Println(err)
    //
    //    return
    //}
    fmt.Println(strings.ToLower("TEST"))
    t := time.NewTicker(time.Second * 3)

    for {
        select {
        case <- t.C:
            fmt.Println("test1")
            return
        default:
            fmt.Println("test2")
        }

    }
    //for _, value := range strings.Split(string(out), "\n") {
    //    if strings.Contains(value, "/sys/kernel/debug/tracing") {
    //        fmt.Println("yes")
    //        goto haha
    //    }
    //}
    //haha:

}

//s
//// check read write on disk
//func readWriteCheck(mountPath, partition string) (bool, string) {
//    message := ""
//    isFailure := false
//    fmt.Printf("start to read write test %s on %s", mountPath, partition)
//    if err := fileutil.IsDirWriteable(mountPath); err != nil {
//        fmt.Printf("Write file to path %s error: %+v", mountPath, err)
//        // ignore disk full case, as it's not disk error
//        if !strings.Contains(err.Error(), "no space left") {
//            message = fmt.Sprintf("%s write failure;", partition)
//            isFailure = true
//        }
//    }
//    if _, err := fileutil.ReadDir(mountPath); err != nil {
//        fmt.Printf("Read file from path %s error: %+v", mountPath, err)
//        message = message + fmt.Sprintf("%s read failure;", partition)
//        isFailure = true
//    }
//
//    return isFailure, message
//}
