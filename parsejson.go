package main

import (
	"fmt"
	"os/exec"
	"strings"
)

const blockSize = int64(66)

const SECTOR_SIZE = 512

func main() {
	out, err := exec.Command("sh", "-c", "vgs docker -o vg_free_count | grep -v Free").Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(strings.Fields(string(out))[0])
	//fmt.Println(string(out))
}
