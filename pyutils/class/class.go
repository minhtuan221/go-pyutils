package main

import (
	"fmt"
	"log"
)

type Console struct {
	X int
	Y int
}

func NewConsole() *Console {
	return &Console{X: 5}
}

var console Console = *NewConsole()

type MyClass struct {
	Name string
}

func (m *MyClass) Amethod(s string) {
	m.Name = s
}

type OtherClass struct {
	Name string
}

func (m *OtherClass) Amethod(s string) {
	m.Name = s
}

type Inheritance struct {
	*MyClass
	OtherClass
	Other string
}

func cast(class interface{}) {
	if c, ok := class.(MyClass); ok { // type assert on it
		fmt.Println(c.Name)
	}
}

func test() {
	x := MyClass{"Old Name"}
	x.Amethod("Name change by method")
	log.Println(x)
	y := Inheritance{&x, OtherClass{"Other Name"}, ""}
	// y.Name = "New Name"
	y.Other = "Other attribute"
	log.Println(y)
	// y.Amethod("Method inheritance")
	log.Println(y)
}

func main() {
	test()
}
