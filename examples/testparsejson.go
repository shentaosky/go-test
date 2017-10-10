package main

import (
	io "io/ioutil"

	"encoding/json"

	"fmt"
)

type JsonStruct struct {
}

func NewJsonStruct() *JsonStruct {

	return &JsonStruct{}

}

func (self *JsonStruct) Load(filename string, v interface{}) {

	data, err := io.ReadFile(filename)

	if err != nil {

		return

	}

	fmt.Printf("read content %s \n", string(data))

	//datajson := []byte(data)

	err = json.Unmarshal(data, v)

	if err != nil {
		fmt.Printf("err：%v", err)
		return

	}

}

type ValueTestAtmp struct {
	StringValue string `json:"StringValue,omitempty"`

	NumericalValue int `json:"NumericalValue,omitempty"`

	BoolValue string `json:"BoolValue,omitempty"`
}

//type testdata struct {
//    config ValueTestAtmp
//
//}

type Pools struct {
	PoolsConfig string `json:"config"`
	Pools       []Pool `json:"pools"`
}

type Pool struct {
	StorageType string          `json:"StorageType,omitempty"`
	ThinPool    *ThinPoolConfig `json:"ThinPool,omitempty"`
}

type ThinPoolConfig struct {
	Driver   string `json:"Driver,omitempty"`
	PoolName string `json:"PoolName,omitempty"`
}

type testdata map[string]ValueTestAtmp

//type pools map[string][]testdata

func main() {

	JsonParse := NewJsonStruct()

	v := Pools{}

	JsonParse.Load("~/gotest/test_json.json", &v)

	//fmt.Println(v.ValueTest.BoolValue)
	fmt.Println(v)
	fmt.Printf("PoolConfig %s\n", v.PoolsConfig)
	fmt.Printf("-- %s\n", v.Pools[0].StorageType)
	fmt.Println(v.Pools[0].ThinPool)
	//fmt.Println(v["ValueTest"][0].config)
	//fmt.Printf("--- %s", v["ValueTest"][0].config.StringValue)

	//fmt.Println(v["ValueTest"][0].storage)

}
