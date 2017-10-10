package main

import (
	"fmt"
	"net/http"
)

func main() {
	strs := []string{
		"1", "2",
	}
	m := make(map[string][]string)
	m["test"] = strs
	fmt.Println(m)
	http.Handle("/test/", http.StripPrefix("/test/", http()))
}
