package main

import (
	"fmt"
	"reflect"
)

func main() {
	var word []byte
	word = []byte("abc")

	fmt.Println(reflect.TypeOf(word[1]).String())
	c := 'a'
	word[1] = uint8(c)
	fmt.Println(string(word))
}
