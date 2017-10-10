package main

import "fmt"

type test1 interface {
	get(string) string
	set(string, string)
}

type test2 interface {
	get(string) string
	set(string, string)
	update(string, string)
}

type Map struct {
	a map[string]string
}

func (mapTest *Map) get(key string) string {
	return mapTest.a[key]
}

func (mapTest *Map) set(key string, value string) {
	mapTest.a[key] = value
}

func (mapTest *Map) update(key string, value string) {
	if _, ok := mapTest.a[key]; ok {
		mapTest.a[key] = value
	}
}

func main() {
	var testMap *Map = &Map{}
	testMap.a = make(map[string]string, 10)
	test := test2(testMap)
	test.set("1", "st")
	fmt.Println(test.get("1"))
}
