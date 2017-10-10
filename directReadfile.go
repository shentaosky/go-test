package main

import (
	"os"
	"os/exec"
	"strings"
)

func main() {
	testContent := "This is a test file"
	file, err := os.OpenFile("testDirect.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend)
	if err != nil {
		println("OpenFile:", err)
	}
	_, err = file.WriteString(testContent)
	if err != nil {
		println("WriteString: ", err)
	}
	file.Sync()
	file.Close()
	res, err := exec.Command("md5sum", "testDirect.txt").Output()
	if err != nil {
		println("md5sum error: ", err)
	}
	println(strings.Split(string(res), " ")[0])

	defer file.Close()
	file, err = os.Open("testDirect.txt")
	if err != nil {
		println("Open err:", err)
	}
	buf := make([]byte, 1024)
	num := 0
	for {
		n, _ := file.Read(buf)
		if 0 == n {
			break
		}
		num += n
	}
	if testContent+testContent == string(buf[0:num]) {
		println(string(buf[0:num]))
	}
	println(string(buf[0:num]))

	vgcreate := []string{"1"}
	vgcreate2 := []string{"2", "3", "4"}
	vgcreate = append(vgcreate, vgcreate2)
	println(vgcreate)
}
