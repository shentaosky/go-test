package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Project struct {
	Name string `json:"name"`
	Url  string `json:"url"`
	Docs string `json:"docs,omitempty"`
	TBool bool `tbool,omitempty`
}

func main() {
	p1 := Project{}

	data, err := ioutil.ReadFile("/Users/tashen/work/devs/src/github.com/go-test/examples/testMarshal.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &p1)
	if err != nil {
		panic(err)
	}

	// p2 则会打印所有
	if p1.TBool{
		fmt.Printf("%v %v\n", p1.TBool, p1)
	}

	var str int
	str = 1 << 1 & 0x01
	d := str
	//i := d | 0x01
	fmt.Println(d)

}
