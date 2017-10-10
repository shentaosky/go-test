package main

import (
	"fmt"
	"github.com/rancher/convoy/util"
)

func main() {
	out, err := util.Execute("umount", []string{"/tmp/test"})
	if err != nil {
		fmt.Println(err)
	}
	println(out)
}
