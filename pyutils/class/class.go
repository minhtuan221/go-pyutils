package main

import (
	"fmt"
)

func decorator(f func(s string)) func(s string) {

	return func(s string) {
		fmt.Println("Started")
		f(s)
		fmt.Println("Done")
	}
}

func doSomething(s string) {
	fmt.Println(s)
}

type class interface {
	Init()
}

func New(cls *class) {
	cls.Init()
}

type object struct {
	Name string
}

func (o *object) Init() {
	fmt.Println("Object Name is: ", o.Name)
	o.Name = o.Name + " say Hi"
}

func Class(o class) *class {
	return &o
}

func Method(f func()) {
	fmt.Println("this is a method")
	// func (object) f()
}

func abc() {
	fmt.Println("Do some thing in the method")
}

type Console struct {
	X int
	Y int
}

func NewConsole() *Console {
	return &Console{X: 5}
}

var console Console = *NewConsole()

func main() {
	x := Class(object{"new"})
	fmt.Println(*x)
	// x := object{"a class"}
	// x.Method(abc)

}
