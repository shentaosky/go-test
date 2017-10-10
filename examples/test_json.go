package main

import (
	"io/ioutil"

	"fmt"

	//"github.com/ghodss/yaml"

	"encoding/json"
)

type JsonStruct struct {
}

//func NewJsonStruct () *JsonStruct {
//
//    return &JsonStruct{}
//
//}

//
//func (self *JsonStruct) Load (filename string, v interface{}) {
//
//    data, err := io.ReadFile(filename)
//
//    if err != nil{
//
//        return
//
//    }
//
//    datajson := []byte(data)
//
//
//    err = json.Unmarshal(datajson, v)
//
//    if err != nil{
//
//        return
//
//    }
//
//}

type CephPoolConfig struct {
	Driver string
}

type AWSPoolConfig struct {
	Driver string
}

type ThinPoolConfig struct {
	Driver             string                 `json:"driver,omitempty"`
	PoolName           string                 `json:"poolName"`
	StorageType        string                 `json:"storageType,omitempty"`
	ThinPoolMetaDevice string                 `json:"thinPoolMetaDevice,omitempty"`
	ThinPoolDataDevice string                 `json:"thinPoolDataDevice,omitempty"`
	VolumeList         map[string]*VolumeTest `json:"volumeList,omitempty"`
	BlockSize          string                 `json:"blockSize,omitempty"`
	VolumeSize         string                 `json:"volumeSize,omitempty"`
	FileSystem         string                 `json:"fileSystem,omitempty"`
}

type VolumeTest struct {
	name string
}

type PoolConfig struct {
	StorageType string          `json:"storagetype,omitempty"`
	Driver      string          `json:"driver,omitempty"`
	ThinPool    *ThinPoolConfig `json:"thinpool,omitempty"`
	Ceph        *CephPoolConfig `json:"ceph,omitempty"`
	AWS         *AWSPoolConfig  `json:"aws,omitempty"`
}

type TestPool struct {
	StorageType string `json:"storagetype"`
	Driver      string `json:"driver,omitempty"`
}

func jsonParse(filename string, v interface{}) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	err = json.Unmarshal([]byte(data), v)
	if err != nil {
		return err
	}
	return nil
}

func jsonSave(filename string, v interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	err = ioutil.WriteFile(filename, data, 0666)
	if err != nil {
		return err
	}
	return nil
}

type Pools struct {
	Pools []PoolConfig
}

func main() {
	p := &PoolConfig{
		ThinPool: &ThinPoolConfig{
			PoolName: "123",
		},
	}

	volumes := make(map[string]*VolumeTest)

	volumes["volume1"] = &VolumeTest{
		name: "volume1",
	}
	p.ThinPool.VolumeList = volumes

	jsonSave("testjson.json", p)

	q := &PoolConfig{}
	jsonParse("testjson.json", q)
	fmt.Println(q)
}

func transmitref(v interface{}) ThinPoolConfig {
	value, ok := v.(ThinPoolConfig)
	if ok {
		return value
	}
	return ThinPoolConfig{}
}
