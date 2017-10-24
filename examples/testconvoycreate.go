package main

import (
    "os/exec"
    "time"
    "fmt"
    "github.com/rancher/convoy/util"
    "os"
)


const percent = 1000000000

func main() {
    createTime := []int64{}
    deleteTime := []int64{}
    for i:=0; i < 20; i++ {
        volumeName := "volume" + util.UUID(10)
        startTime := time.Now().UnixNano()
        if out, err := exec.Command("convoy", "create", volumeName).Output(); err != nil {
            fmt.Printf("%s: %v", string(out), err)
        }

        createTime = append(createTime, time.Now().UnixNano() - startTime)

        startTime = time.Now().UnixNano()
        if out, err := exec.Command("convoy", "delete", volumeName).Output(); err != nil {
            fmt.Printf("%s: %v", string(out), err)
        }
        deleteTime = append(deleteTime, time.Now().UnixNano()- startTime)
    }

    logFile, err := os.OpenFile("/root/create.record", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
    if err != nil {
        fmt.Println(err)
        return
    }
    logFile.WriteString("----------------------------------\n")
    defer logFile.Close()
    logFile.WriteString("create time(s): \n")
    maxCreateTime, minCreateTime, avgCreateTime := caculate(logFile, createTime)
    logFile.WriteString("delete time(s): \n")
    maxDeleteTime, minDeleteTime, avgDeleteTime := caculate(logFile, deleteTime)
    logFile.WriteString(fmt.Sprintf("max create time: %5f, min create time: %5f, average create time: %5f\n", float32(maxCreateTime)/percent, float32(minCreateTime)/percent, avgCreateTime/percent))
    logFile.WriteString(fmt.Sprintf("max delete time: %5f, min delete time: %5f, average delete time: %5f\n", float32(maxDeleteTime)/percent, float32(minDeleteTime)/percent, avgDeleteTime/percent))
    logFile.WriteString("----------------------------------\n")
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

