package main

import (
    "fmt"
    "github.com/sasha-s/go-deadlock"
    "time"
    "github.com/petermattis/goid"
    "runtime"
)

type MutexTestA struct {
    mutexA *deadlock.RWMutex
    mutexB *deadlock.RWMutex

}

func main() {
    mutexTest := MutexTestA{
        mutexA: &deadlock.RWMutex{},
        mutexB: &deadlock.RWMutex{},
    }
    deadlock.Opts.OnPotentialDeadlock = printINFO
    deadlock.Opts.DeadlockTimeout = time.Second * 20

    go mutexTest.testB()
    mutexTest.testA()
    fmt.Println("succeed")
    time.Sleep(time.Second * 10)
}

func printINFO() {
    fmt.Println("deadlock")
}

func (m *MutexTestA) testA() {
    m.mutexA.RLock()
    fmt.Println("get RlockA and release Rlock")
    time.Sleep(time.Second*2)
    m.mutexB.RLock()
    fmt.Println("get RlockB and release Rlock")
    defer m.mutexA.RUnlock()
    defer m.mutexB.RUnlock()
}

func (m *MutexTestA) testB() {
    m.mutexB.RLock()
    fmt.Println("get RlockB and release Rlock")
    time.Sleep(time.Second*2)
    m.mutexA.RLock()
    fmt.Println("get RlockA and release Rlock")
    defer m.mutexB.RUnlock()
    defer m.mutexA.RUnlock()
}

func (m *MutexTestA) testC() {
    m.mutexB.Lock()
    m.mutexB.Lock()
    gid := goid.Get()
    fmt.Println("123:", gid)
    m.mutexB.Unlock()
    m.mutexB.Unlock()
}

func stacks() []byte {
    buf := make([]byte, 1024*16)
    for {
        n := runtime.Stack(buf, true)
        if n < len(buf) {
            return buf[:n]
        }
        buf = make([]byte, 2*len(buf))
    }
}
