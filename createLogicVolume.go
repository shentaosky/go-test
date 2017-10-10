package main

import (
	"fmt"
	"github.com/rancher/convoy/localdisk"
	"os/exec"
	"strconv"
)

const (
	VG_POSTFIX           = "_vg"
	DEFAULT_SEGMENT_TYPE = "Linear"
)

var (
	SEGMENT_TYPES = []string{"SoftRaid0", "SoftRaid1", "SoftRaid5", "Linear"}
)

type ThinPool struct {
	PoolDevices    []string
	PoolName       string
	Type           string
	DataDevice     string `json:"dataDevice,omitempty"`
	MetadataDevice string `json:"metadataDevice,omitempty"`
}

func main() {

	thinPool := ThinPool{
		PoolName: "test",
		PoolDevices: []string{
			"/dev/loop1", "/dev/loop2", "/dev/loop3",
		},
		Type: "SoftRaid1",
		//MetadataDevice: "123",
		//DataDevice: "123",
	}
	if (thinPool.DataDevice == "" || thinPool.MetadataDevice == "") && len(thinPool.PoolDevices) == 0 {
		fmt.Println("data device and metadata device or device list unspecified")
	}

	if len(thinPool.PoolDevices) > 0 && (thinPool.DataDevice != "" || thinPool.MetadataDevice != "") {
		fmt.Println("have specified data and metadata device, don't need specify devices")
	}
	//err := thinPool.CreatePoolFromDeviceList()
	//if err != nil {
	//    fmt.Println(err)
	//}
	err := thinPool.RemovePoolFromDeviceList()
	if err != nil {
		fmt.Println(err)
	}

}

func (d *ThinPool) CreatePoolFromDeviceList() error {
	var err error
	if d.Type == "" {
		d.Type = DEFAULT_SEGMENT_TYPE
	} else {
		for i, segment_type := range SEGMENT_TYPES {
			if segment_type == d.Type {
				break
			} else if i == len(SEGMENT_TYPES)-1 {
				return fmt.Errorf("Error segment type, only support types: %v", SEGMENT_TYPES)
			}
		}
	}

	if d.Type == "SoftRaid5" && len(d.PoolDevices) < 3 {
		return fmt.Errorf("SoftRaid5 type need 3 disks or more")
	}

	if (d.Type == "SoftRaid0" || d.Type == "SoftRaid1") && len(d.PoolDevices) < 2 {
		return fmt.Errorf("SoftRaid0 and SoftRaid1 type need 2 disks or more")
	}

	for _, device := range d.PoolDevices {
		if err = localdisk.CreatePV(device); err != nil {
			return fmt.Errorf("create pv %s fail: %v", device, err)
		}
	}
	vgName := fmt.Sprintf("%s%s", d.PoolName, VG_POSTFIX)
	if err = localdisk.CreateVG(d.PoolDevices, vgName); err != nil {
		return fmt.Errorf("create vg %s fail: %v", vgName, err)
	}

	if poolMetaDevice, poolDataDevice, err := d.createThinpoolMetaAndData(vgName); err != nil {
		return err
	} else {
		d.MetadataDevice = poolMetaDevice
		d.DataDevice = poolDataDevice
	}

	return nil
}

func (d *ThinPool) RemovePoolFromDeviceList() error {
	metaName := fmt.Sprintf("convoy_%s_%s_meta", d.Type, d.PoolName)
	dataName := fmt.Sprintf("convoy_%s_%s_data", d.Type, d.PoolName)
	var err error
	vgName := fmt.Sprintf("%s%s", d.PoolName, VG_POSTFIX)
	if err = localdisk.RemoveLV(vgName, metaName); err != nil {
		return fmt.Errorf("delete lv %s/%s fail: %v", vgName, metaName, err)
	}
	if err = localdisk.RemoveLV(vgName, dataName); err != nil {
		return fmt.Errorf("delete lv %s/%s fail: %v", vgName, dataName, err)
	}
	if err = localdisk.RemoveVG(vgName); err != nil {
		return fmt.Errorf("delete vg %s fail: %v", vgName, err)
	}
	for _, device := range d.PoolDevices {
		if err = localdisk.RemovePV(device); err != nil {
			return fmt.Errorf("delete pv %s fail: %v", device, err)
		}
	}
	return nil
}

func (d *ThinPool) createThinpoolMetaAndData(vgname string) (string, string, error) {
	var meta_args []string
	var data_args []string
	metaName := fmt.Sprintf("convoy_%s_%s_meta", d.Type, d.PoolName)
	dataName := fmt.Sprintf("convoy_%s_%s_data", d.Type, d.PoolName)
	poolMetaDevice := "/dev/mapper/" + metaName
	poolDataDevice := "/dev/mapper/" + dataName
	switch d.Type {
	case "SoftRaid0":
		meta_args = []string{"--type", "raid0", "-i", strconv.Itoa(len(d.PoolDevices)),
			"-Zy", "-l", "5%VG", "--wipesignatures", "y", "-n", metaName, vgname,
		}
		data_args = []string{"--type", "raid0", "-i", strconv.Itoa(len(d.PoolDevices)),
			"-Zy", "-l", "95%VG", "--wipesignatures", "y", "-n", dataName, vgname,
		}
	case "SoftRaid1":
		meta_args = []string{"--type", "raid1",
			"-Zy", "-l", "5%VG", "--wipesignatures", "y", "-n", metaName, vgname,
		}
		data_args = []string{"--type", "raid1",
			"-Zy", "-l", "95%VG", "--wipesignatures", "y", "-n", dataName, vgname,
		}
	case "SoftRaid5":
		meta_args = []string{"--type", "raid5", "-i", strconv.Itoa(len(d.PoolDevices) - 1),
			"-Zy", "-l", "5%VG", "--wipesignatures", "y", "-n", metaName, vgname,
		}
		data_args = []string{"--type", "raid5", "-i", strconv.Itoa(len(d.PoolDevices) - 1),
			"-Zy", "-l", "95%VG", "--wipesignatures", "y", "-n", dataName, vgname,
		}
	default:
		meta_args = []string{
			"-Zy", "-l", "5%VG", "--wipesignatures", "y", "-n", metaName, vgname,
		}
		data_args = []string{
			"-Zy", "-l", "95%VG", "--wipesignatures", "y", "-n", dataName, vgname,
		}
	}
	if output, err := exec.Command("lvcreate", meta_args...).CombinedOutput(); err != nil {
		return "", "", fmt.Errorf("lvcreate %s error, %s: %v", metaName, string(output), err)

	}
	if output, err := exec.Command("lvcreate", data_args...).CombinedOutput(); err != nil {
		if err := localdisk.RemoveLV(vgname, metaName); err != nil {
			fmt.Printf("Clean up created lv %s fail: %v", poolMetaDevice, err)
		}
		return "", "", fmt.Errorf("lvcreate %s error, %s: %v", dataName, string(output), err)
	}
	return poolMetaDevice, poolDataDevice, nil
}
