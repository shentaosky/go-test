package main

import (
	"fmt"
	"github.com/ghodss/yaml"
)

type Person struct {
	Name        string `json:"username"`
	Age         int
	Gender      bool `json:",omitempty"`
	Profile     string
	OmitContent string `json:"-"`
	Count       int    `json:",string"`
}

func main() {

	var p *Person = &Person{
		Name:        "brainwu",
		Age:         21,
		Gender:      true,
		Profile:     "I am ghj1976",
		OmitContent: "OmitConent",
	}
	if bs, err := yaml.Marshal(&p); err != nil {
		panic(err)
	} else {
		//result --> {"username":"brainwu","Age":21,"Gender":true,"Profile":"I am ghj1976","Count":"0"}
		fmt.Println(string(bs))
	}

	var p2 *Person = &Person{
		Name:        "brainwu",
		Age:         21,
		OmitContent: "OmitConent",
	}
	if bs, err := yaml.Marshal(&p2); err != nil {
		panic(err)
	} else {
		//result --> {"username":"brainwu","Age":21,"Profile":"I am ghj1976","Count":"0"}
		fmt.Println(string(bs))
	}

	// slice 序列化为json
	var aStr []string = []string{"Go", "Java", "Python", "Android"}
	if bs, err := yaml.Marshal(aStr); err != nil {
		panic(err)
	} else {
		//result --> ["Go","Java","Python","Android"]
		fmt.Println(string(bs))
	}

	//map 序列化为json
	var m map[string]string = make(map[string]string)
	m["Go"] = "No.1"
	m["Java"] = "No.2"
	m["C"] = "No.3"
	if bs, err := yaml.Marshal(m); err != nil {
		panic(err)
	} else {
		//result --> {"C":"No.3","Go":"No.1","Java":"No.2"}
		fmt.Println(string(bs))
	}
}
