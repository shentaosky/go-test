package main

import "fmt"

type Personer interface {
	Walk()
	ChangeName(string)
}

type GoodPerson struct {
	name string
	age  int
}

func NewGoodPerson(name string) *GoodPerson {
	return &GoodPerson{name, 12}
}

func (this *GoodPerson) Walk() {
	fmt.Println(this.name, " walking")
	fmt.Println(this.age, " age")
}

//GoodPerson类并没有显示继承Person,但它定义了两个方法Walk和ChangeName,因此可以直接赋值给Personer接口对象

//相当于java中的子类对象赋值给基类对象，然后基类对象可以直接使用子类对象实现的方法,但不能实现基类对象没有的方法

func (this *GoodPerson) ChangeName(name string) {
	this.name = name
}

func (this *GoodPerson) Fly() {
	fmt.Println("man can't fly")
}

func main() {
	var Person2 Personer = &GoodPerson{"STR", 11} //注意go和c一样,struct不是指针,而对象是指针，所以需要加上地址符号
	var Person3 Personer = new(GoodPerson)        //new只能开辟一块内存,但并没有赋值功能，new生成的是对象（指针）
	Person3.ChangeName("123")
	var Person Personer = NewGoodPerson("hxc")
	//var Person2 interface{} = Person

	Person.Walk() // "hxc walking"
	Person.ChangeName("wzm")
	Person.Walk() // "wzm walking"
	Person2.Walk()
	Person3.Walk()
	// Person.Fly() // error Personer have no field or mothed name Fly
	// Person2.Walk() // error
	//Person = 1 //error
}
