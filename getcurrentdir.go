package main

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func main() {
	_, dir, _, _ := runtime.Caller(1)

	fmt.Println(dir)
	fmt.Println(filepath.Join(filepath.Dir(dir), "data.csv"))

}
