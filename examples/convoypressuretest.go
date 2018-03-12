package main

import (
    "os/exec"
    "fmt"
    "sync"
    "time"
    "strconv"
    "os"

    "github.com/rancher/convoy/util"
)

const (
    bin = "convoy"
    poolname = "bronze"
    volumesize = "1G"
)

type PrintError struct {
    mutex *sync.Mutex
    file  *os.File
}

func main() {
    t := time.NewTimer(time.Hour * 9)
    fd, err := os.Create("/root/convoypressuretest_error.log")
    if err != nil {
        return
    }
    printError := &PrintError{
        mutex : &sync.Mutex{},
        file: fd,
    }
    defer printError.file.Close()
    defer printError.file.Sync()

    for {
        select {
        case <-t.C:
            return
        default:
            wg := &sync.WaitGroup{}
            wg.Add(200)
            for i := 1; i <= 200; i++ {
                go printError.volumeLifeCycleRun(wg, i)
            }
            wg.Wait()
            time.Sleep(time.Second)
        }
    }
}

func (p *PrintError) printErr(err error) {
    p.mutex.Lock()
    defer p.mutex.Unlock()

    p.file.WriteString(err.Error() + "\n")
}

func (p *PrintError) volumeLifeCycleRun(wait *sync.WaitGroup, i int) {
    defer wait.Done()
    volumeName := "pressure5_" + strconv.Itoa(i) + "_" + util.UUID(32)
    if err := execRun(bin, []string{"create", volumeName, "--poolname", poolname, "--size", volumesize}); err != nil {
        p.printErr(fmt.Errorf("create volumeName error: %v", err))
        return
    }
    if err := execRun(bin, []string{"mount", volumeName}); err != nil {
        p.printErr(fmt.Errorf("mount volumeName error: %v", err))
        return
    }
    if err := execRun(bin, []string{"umount", volumeName}); err != nil {
        p.printErr(fmt.Errorf("umount volumeName error: %v", err))
        return
    }
    if err := execRun(bin, []string{"delete", volumeName}); err != nil {
        p.printErr(fmt.Errorf("delete volumeName error: %v", err))
        return
    }
}

func execRun(commmand string, args []string) error {
    //time1 := time.Now().UnixNano()/1000000
    if out, err := exec.Command(commmand, args...).CombinedOutput(); err != nil {
        return fmt.Errorf("%s: %v", string(out), err)
    }
    //time2 := time.Now().UnixNano()/1000000
    //fmt.Println(args[0], "time: ", time2-time1)
    return nil
}

