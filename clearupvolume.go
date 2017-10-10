package main

import (
	"fmt"
	"os/exec"
)

func main() {
	for i := 1; i <= 20; i++ {
		err := exec.Command("sh", "-c", "dmsetup message /dev/mapper/convoy_bronze 0 "+fmt.Sprintf("'delete %d'", i)).Run()
		if err != nil {
			fmt.Println(err)
		}
		err = exec.Command("sh", "-c", "dmsetup message /dev/mapper/convoy_silver 0 "+fmt.Sprintf("'delete %d'", i)).Run()
		if err != nil {
			fmt.Println(err)
		}
		err = exec.Command("sh", "-c", "dmsetup message /dev/mapper/convoy_gold 0 "+fmt.Sprintf("'delete %d'", i)).Run()
		if err != nil {
			fmt.Println(err)
		}
	}
}
