package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

const (
	LVM_VERSION = 20161116
	LIB_VERSION = 20161116
)

func main() {
	out, err := exec.Command("sh", "-c", "lvm version | awk {'print $4'}").Output()
	if err != nil {
		fmt.Println(err)
	}
	outs := strings.Split(string(out), "\n")
	lvmVersion := strings.Replace(outs[0][1:len(outs[0])-1], "-", "", -1)
	lvmversion, err := strconv.Atoi(lvmVersion)
	if lvmversion < LVM_VERSION {
		fmt.Println(123)
	}
}
