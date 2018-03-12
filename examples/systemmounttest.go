package main

import "k8s.io/kubernetes/pkg/util/mount"

func main() {
    dir := "/tmp/test"
    m := mount.New("")
    m.Mount("tmpfs", dir, "tmpfs", []string{"size=10240"} /* options */)
}
