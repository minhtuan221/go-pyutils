package main

import (
	"Pyutils/pyutils/PyList"
	"Pyutils/pyutils/genericList"
	"fmt"
)

func main() {
	fmt.Println("Your test is here: ")
	var x pyList.List
	// = PyList.List{}
	// fmt.Println(x)
	for _, value := range []int{5, 6, 7, 8, 9} {
		x.Append(value)
	}
	fmt.Println(x)
	pyList.TestList()
	fmt.Println("End of test list Interface: ")
	fmt.Println("Start of test list Integer: ")
	genericList.TestInt()
	fmt.Println("Start of test list Float: ")
	genericList.TestFloat()
	fmt.Println("Start of test list String: ")
	genericList.TestString()
	// blockchain.PrettyPrint(balance)
}
