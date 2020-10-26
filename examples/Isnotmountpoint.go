package main

import (
    "flag"
    "fmt"
    "k8s.io/kubernetes/pkg/util/mount"
)

var addr *string = flag.String("addr", "infile", "File contains values for sorting")

func main() {
    flag.Parse()
    mounter := &mount.SafeFormatAndMount{Interface: mount.New(""), Exec: mount.NewOsExec()}
    ismp, err := mounter.IsNotMountPoint(*addr)
    fmt.Printf("ismp: %v, err: %v, addr: %s", ismp, err, *addr)
    str, err := mounter.EvalHostSymlinks(*addr)
    fmt.Printf("str: %v, err: %v, addr: %s", str, err, *addr)
}