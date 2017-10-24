package main

import (
    "os/exec"
    "time"
    "fmt"
    "os"
    "strconv"
    "sync"
)

const percent = 1000000000

func main() {
    wg := &sync.WaitGroup{}
    wg.Add(10)
    for i := 20; i <= 29; i++ {
        str := strconv.Itoa(i)
        stopCh := make(chan bool)
        go recordMountUmount(str, stopCh)
        <-stopCh
        wg.Done()
    }
    wg.Wait()
}

func caculate(logFile *os.File, nums []int64) (int64, int64, float64){
    avg := int64(0)
    maxNum := int64(0)
    minNum := int64(999999999999)
    for i, num := range nums {
        logFile.WriteString(fmt.Sprintf("%5f ", float32(num)/percent))
        if num > maxNum {
            maxNum = num
        }
        if num < minNum {
            minNum = num
        }
        avg += num
        if i % 10 == 9 {
            logFile.Write([]byte("\n"))
        }
    }
    logFile.Write([]byte("\n"))
    return maxNum, minNum, float64(avg)/float64(len(nums))
}
//
//
//func recordAwaitTime() {
//    logFile, err := os.OpenFile("/root/mount_await.record", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
//    if err != nil {
//        fmt.Println(err)
//    }
//    t := time.NewTicker(time.Second*1)
//    for {
//
//        out, _ := exec.Command("iostat", "-x", "-k", "/dev/dm-18", "/dev/sdc").CombinedOutput()
//
//        logFile.Write(out)
//        time.Sleep()
//    }
//
//    defer logFile.Close()
//}


func recordMountUmount(str string, stopCh chan bool) {
    mountTime := []int64{}
    umountTime := []int64{}

    for i:=0; i < 5; i++ {
        startTime := time.Now().UnixNano()
        if out, err := exec.Command("convoy", "mount", "volume"+str).Output(); err != nil {
            fmt.Printf("%s: %v", string(out), err)
        }

        mountTime = append(mountTime, time.Now().UnixNano() - startTime)

        startTime = time.Now().UnixNano()
        if out, err := exec.Command("convoy", "umount", "volume"+str).Output(); err != nil {
            fmt.Printf("%s: %v", string(out), err)
        }
        umountTime = append(umountTime, time.Now().UnixNano()- startTime)
    }

    logFile, err := os.OpenFile("/root/mount.record"+str, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
    if err != nil {
        fmt.Println(err)
        return
    }
    logFile.WriteString("----------------------------------\n")
    defer logFile.Close()
    logFile.WriteString("mount time(s): \n")
    maxMountTime, minMountTime, avgMountTime := caculate(logFile, mountTime)
    logFile.WriteString("umount time(s): \n")
    maxUmountTime, minUmountTime, avgUmountTime := caculate(logFile, umountTime)
    logFile.WriteString(fmt.Sprintf("max mount time: %5f, min mount time: %5f, average mount time: %5f\n", float32(maxMountTime)/percent, float32(minMountTime)/percent, avgMountTime/percent))
    logFile.WriteString(fmt.Sprintf("max umount time: %5f, min umount time: %5f, average umount time: %5f\n", float32(maxUmountTime)/percent, float32(minUmountTime)/percent, avgUmountTime/percent))
    stopCh <- true
}
