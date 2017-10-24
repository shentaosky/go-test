package main

import (
    "os/exec"
    "time"
    "fmt"
    "os"
)


const percent = 1000000000

func main() {
    mountTime := []int64{}
    umountTime := []int64{}

    for i:=0; i < 5; i++ {
        startTime := time.Now().UnixNano()
        if out, err := exec.Command("ionice", "-c", "1", "-n", "0", "mount", "/dev/mapper/volume2", "/var/lib/rancher/convoy/bronze/mounts/volume1").Output(); err != nil {
            fmt.Printf("%s: %v", string(out), err)
        }

        mountTime = append(mountTime, time.Now().UnixNano() - startTime)

        startTime = time.Now().UnixNano()
        if out, err := exec.Command("ionice", "-c", "1", "-n", "0", "umount", "/dev/mapper/volume2").Output(); err != nil {
            fmt.Printf("%s: %v", string(out), err)
        }
        umountTime = append(umountTime, time.Now().UnixNano()- startTime)
    }

    logFile, err := os.OpenFile("/root/mount.record_2", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
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
