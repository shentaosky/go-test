package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	out, err := ReadAll("/proc/meminfo")
	if err != nil {
		fmt.Printf("read meninfo error: %v", err)
	}

	menInfo := string(out)
	println(menInfo)
	menInfoLines := strings.Split(menInfo, "\n")
	//var totalVM int64
	for _, line := range menInfoLines {
		if len(line) > 1 {
			words := strings.Split(line, " ")
			fmt.Println(words[len(words)-2])
		}
	}
}

func ReadAll(filePth string) ([]byte, error) {
	f, err := os.Open(filePth)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ioutil.ReadAll(f)
}
