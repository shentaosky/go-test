package main

import (
    "fmt"
    "os/exec"
    "strings"
    "time"
)

const (
    defaultKdumpExpireTime = 12 * time.Hour
    KdumpTimestampFormat = "2006/01/02 15:04:05"
)

func main() {
    out, err := exec.Command("ls", "-l", "-t", "--time-style=+%Y-%m-%d %H:%M:%S", "/tmp").CombinedOutput()
    if err != nil {
        fmt.Printf("ls -l error: %s, %+v", string(out), err)
    }

    lines := strings.Split(string(out), "\n")

    for i, line := range lines {
        if i == 0 || len(line) <= 1 {
            continue
        }

        nameStart := strings.LastIndex(line, " ")
        timeStart := strings.LastIndex(line[:nameStart], " ")
        dataStart := strings.LastIndex(line[:timeStart], " ")
        filename := line[nameStart+1:]
        timestamp := line[dataStart+1:nameStart]
        t, err := time.ParseInLocation(KdumpTimestampFormat, timestamp, time.Local)
        if err != nil {
            fmt.Printf("failed to parse timestamp %s in ipmi:  %+v", timestamp, err)

        }
        fmt.Println(t.String())

    }
}
