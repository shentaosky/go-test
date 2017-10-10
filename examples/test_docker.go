package main

import (
	"fmt"
	"os/exec"
)

func main() {
	_, err := exec.Command("/usr/bin/docker", "logs", "e412a04c0cf1", ">", "/test_docker.log", "2>&1").Output()
	fmt.Errorf("%v", err)
	fmt.Println()
}
