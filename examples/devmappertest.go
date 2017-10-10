package main

import (
	"fmt"
	"github.com/docker/docker/pkg/devicemapper"
	//"os/exec"
	"github.com/rancher/convoy/util"
)

func main() {

	if err := devicemapper.BlockDeviceDiscard("convoy_test_devmapper_pool_4e2"); err != nil {
		fmt.Println(err)
	}
	//
	//err := devicemapper.RemoveDevice("testdeviceautoremount_3e2")
	//if err != nil {
	//    fmt.Println("RemoveDevice:", err)
	//}
	//
	//if err := umountDevice("/home/jenkins/workspace/src/github.com/rancher/convoy/test_dir"); err != nil {
	//    fmt.Println("umountDevice: ", err)
	//}
	//
	//err = devicemapper.ActivateDevice("/dev/mapper/convoy_test_devmapper_pool_62f", "testdeviceautoremount_3e2", 1, 33554432)
	//if err != nil {
	//    fmt.Println("ActivateDevice1: ", err)
	//}

	//err = exec.Command("dmsetup", "remove", "volume1").Run()
	//if err != nil {
	//    fmt.Println("dmsetup remove: ", err)
	//}
	//
	//err = devicemapper.ActivateDevice("/dev/mapper/bronze_vg-convoy_Linear_bronze_data", "volume1", 2, 1073741824)
	//if err != nil {
	//    fmt.Println("ActivateDevice2: ", err)
	//}
}

func umountDevice(mountpoint string) error {
	cmdName := "umount"
	cmdArgs := []string{mountpoint}
	if _, err := util.Execute(cmdName, cmdArgs); err != nil {
		return err
	}
	return nil
}
