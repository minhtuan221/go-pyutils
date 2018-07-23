package main

import (
	pylist "Pyutils/pyutils/Pylist"
	genericlist "Pyutils/pyutils/genericList"
	"Pyutils/pyutils/try_except"
	"fmt"
)

func main() {
	fmt.Println("Your test is here: ")
	var x pylist.List
	// = PyList.List{}
	// fmt.Println(x)
	for _, value := range []int{5, 6, 7, 8, 9} {
		x.Append(value)
	}
	fmt.Println(x)
	pylist.TestList()
	fmt.Println("End of test list Interface: ")
	fmt.Println("Start of test list Integer: ")
	genericlist.TestInt()
	fmt.Println("Start of test list Float: ")
	genericlist.TestFloat()
	fmt.Println("Start of test list String: ")
	genericlist.TestString()
	fmt.Println("Start of test Dict: ")
	pylist.TestDict()
	fmt.Println("Start of test TryCatch: ")
	tryexcept.TestTryCatch()
	// blockchain.PrettyPrint(balance)
}
