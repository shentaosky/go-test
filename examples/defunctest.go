package main

import (
	"os/exec"
	"strconv"
	"time"
)

func main() {
	cmd, err := exec1Cmd("etcd", []string{})
	if err != nil {
		return
	}
	exec.Command("kill", strconv.Itoa(cmd.Process.Pid)).Run()
	cmd.Wait()
	time.Sleep(time.Second * 10)
}

func exec1Cmd(commmand string, args []string) (*exec.Cmd, error) {
	cmd := exec.Command(commmand, args...)
	if err := cmd.Start(); err != nil {
		return nil, err
	}
	return cmd, nil
}
