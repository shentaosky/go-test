package main

import (
	"fmt"
	"log"
)

type Rect struct {
	x, y          float64
	width, height float64
	map1          map[string]int
	*log.Logger
}

func (rct *Rect) Area() float64 {
	rct.width = 1000 ///也可以(*rct).width=1000一样,无论指针还是对象，都可以访问内部元素
	fmt.Print(rct)
	fmt.Print((*rct))
	return rct.width * rct.height
}

func main() {
	rct := new(Rect)
	rct.width = 10.0
	rct.height = 10.0
	area := rct.Area()
	fmt.Println(area)
	fmt.Println(rct.width)

	id := fmt.Sprintf("%d", 1)

	fmt.Println(fmt.Sprintf("'create_thin %s'", id))
}
