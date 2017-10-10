package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	//output, err := exec.Command("mount", "|", "grep", "sh").Output()
	//fmt.Println(err)
	//fmt.Println(string(output))
	//
	//output, err= exec.Command("sh", "-c", "dmsetup status").Output()
	//fmt.Println(err)
	//fmt.Println(string(output))

	res, err := exec.Command("sh", "-c", "cat /proc/mounts | grep volume1 ").Output()
	fmt.Println(err)
	r := string(res)
	mounts := strings.Split(r, " ")
	for _, mount := range mounts {
		fmt.Println(mount)
	}

	//err = exec.Command("mount", "| grep", "sh").Run()

}
