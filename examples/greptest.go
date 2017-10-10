package main

import (
	"fmt"
	"os/exec"
)

func main() {
	output, err := exec.Command("cat", "/proc/mounts", "|", "grep", "sysfs").Output()
	if err != nil {
		fmt.Println(string(output))
	} else {
		fmt.Println(err)
	}

}
