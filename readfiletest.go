package main

import (
	"fmt"
	"github.com/docker/docker/pkg/devicemapper"
	"strings"
)

func main() {
	err := devicemapper.DeleteDevice("/dev/mapper/convoy_test_pool1", 6)
	if strings.Contains(err.Error(), devicemapper.ErrTaskRun.Error()) {
		fmt.Println(1)
	}
}
