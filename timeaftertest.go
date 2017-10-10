package main

import "fmt"

type thinPool struct {
	devices []string
}

func main() {
	var thinpool thinPool = thinPool{
		devices: []string{},
	}
	if thinpool.devices == nil {
		fmt.Println("2")
	} else {
		fmt.Println("3")
	}
}
