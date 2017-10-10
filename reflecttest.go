package main

import (
	//"reflect"
	"fmt"
)

type User struct {
	Id   int
	Name string
	Age  int
	msg  string
}

func (u User) Hello(name string) {
	fmt.Println("hello", name, ",name is ", u.Name)
}

func transmit(v interface{}) User {
	value, ok := v.(User)
	if ok {
		return value
	}
	return User{}
}

func main() {
	u := User{1, "Brain", 21}
	fmt.Println(transmit(u))
	//v := reflect.ValueOf(u)
	//mv := v.MethodByName("Hello")
	//args := []reflect.Value{reflect.ValueOf("wu")}
	//mv.Call(args)

}
