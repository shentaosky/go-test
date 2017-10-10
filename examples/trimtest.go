package main

import (
	"fmt"
	"github.com/rancher/convoy/util"
)

func main() {
	out, err := util.Execute("thin_ls", []string{"-m", "-o", "DEV,MAPPED_BYTES", "/dev/mapper/bronze_vg-convoy_Linear_bronze_data_tmeta"})
	if err != nil {
		fmt.Printf("%s: %v", string(out), err)
	}
}
