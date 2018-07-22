package main

import (
	"Pyutils/pyutils/PyList"
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
	// blockchain.PrettyPrint(balance)
}
