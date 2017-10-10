package main

import "fmt"

type type1 struct {
	var1 int
	var2 string
}

type type2 struct {
	POINT *type1
}

func (type1) test() {
	fmt.Print("test")
}

func main() {
	st := type1{var1: 12, var2: "me"}
	ST := type2{POINT: &st}
	ST.POINT.test()
	//test()   //ERROR: Undefine
	st.test()
	fmt.Printf("%d, %s \n", st.var1, st.var2)
	fmt.Printf("%d, %s \n", (&st).var1, (&st).var2)
	st2 := new(type1)
	st2.var1 = 11
	st2.var2 = "you"
	fmt.Println((*st2).var1, (*st2).var2)
}
