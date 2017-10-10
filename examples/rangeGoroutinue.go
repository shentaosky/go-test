package main

import (
	"fmt"
)

func main() {
	c := newCache()
	//v := Volume{
	//    name: "v1",
	//}
	var v1 Volume
	var ok bool
	if v1, ok = c.get().(Volume); !ok {
		fmt.Println("1")
		return
	}

	fmt.Println(v1.name)
}

type Volume struct {
	name string
}

type cache interface {
	set(interface{})
	get() interface{}
}

type thinpoolcache struct {
	volume interface{}
}

func newCache() cache {
	return &thinpoolcache{}
}

func (c *thinpoolcache) set(v interface{}) {
	c.volume = v
}

func (c *thinpoolcache) get() interface{} {
	return c.volume
}
