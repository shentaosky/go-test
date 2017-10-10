package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	ops := generageOperations()
	fmt.Println(ops.GetName())
	//classA := ops.(*ClassA)
	//fmt.Println(classA.Name)
}

func Save(filename string, v interface{}) error {
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

type ClassA struct {
	Name string
}

func (a *ClassA) GetName() string {
	return a.Name
}

type ClassB struct {
	*ClassA // 指针作为成员能够继承实现ClassA的所有成员函数, 所以符合Operations接口
	// ClassA // 不加* 也能调用ClassA的函数, 但没有实现GetName, 因为GetName是*ClassA的函数
	Age string
}

type Operations interface {
	GetName() string
}

func generageOperations() Operations {
	return ClassB{
		Age: "16",
		ClassA: &ClassA{
			Name: "st",
		},
	}
}
