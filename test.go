package main

import (
	"fmt"

	. "github.com/minhtuan221/go-pyutils/pyutils/Pylist"
	genericlist "github.com/minhtuan221/go-pyutils/pyutils/genericList"
	"github.com/minhtuan221/go-pyutils/pyutils/tryexcept"

	ini "gopkg.in/ini.v1"
)

func main() {
	defer fmt.Println("Your test is here: ")
	var x List
	// = PyList.List{}
	// fmt.Println(x)
	for _, value := range []int{5, 6, 7, 8, 9} {
		x.Append(value)
	}
	fmt.Println(x)
	TestList()

	fmt.Println("End of test list Interface: ")
	fmt.Println("Start of test list Integer: ")
	genericlist.TestInt()
	fmt.Println("Start of test list Float: ")
	genericlist.TestFloat()
	fmt.Println("Start of test list String: ")
	genericlist.TestString()
	fmt.Println("Start of test Dict: ")
	TestDict()
	fmt.Println("Start of test TryCatch: ")
	// cfg, err := ini.Load("config.ini")
	// if err != nil {
	// 	fmt.Printf("Fail to read file: %v", err)
	// 	os.Exit(1)
	// }
	// Method 1: Classic read of values, default section can be represented as empty string
	// fmt.Println("App Mode:", cfg.Section("env").Key("host").String())
	// z := cfg.Section("env").Key("host").String()
	// exc.Try(func() {
	// 	// strings.(z)
	// 	cfg, _ := ini.Load("confie.ini")
	// 	fmt.Println("App Mode:", cfg.Section("env").Key("host").String())
	// }).Catch(&exc.Exception{}, func(t exc.Throwable) {
	// 	log.Println(t)
	// 	exc.Rethrow(t, exc.Errorf("rethrow after logging"))
	// 	log.Println("Error have been reported")
	// }).Error()

	// Method 2:
	// tryexcept.Try{
	// 	Try: func() {
	// 		cfg, _ := ini.Load("confie.ini")
	// 		fmt.Println("App Mode:", cfg.Section("env").Key("host").String())
	// 	},
	// 	Except: func(e tryexcept.Exception) {
	// 		err, _ := e.(error)
	// 		traceback := errors.WithStack(err)
	// 		fmt.Printf("%+v\n", traceback)
	// 		fmt.Println("Error have been captured")
	// 	},
	// 	Finally: func() {
	// 		fmt.Println("This is finally")
	// 	},
	// }.Do()
	// method 3
	tryexcept.Try{
		Try: func() {
			cfg, _ := ini.Load("confie.ini")
			fmt.Println("App Mode:", cfg.Section("env").Key("host").String())
		},
		Except: func(e tryexcept.Exception) {
			fmt.Printf("%+v\n", tryexcept.Traceback(e))
			fmt.Println("Error have been captured")
		},
		Finally: func() {
			fmt.Println("This is finally")
		},
	}.Do()
	// decorator
	// blockchain.PrettyPrint(balance)
}

func TestList() {
	// var dataSlice []int = foo()
	interfaceSlice := make([]interface{}, 3)
	for i, d := range []int{1, 2, 3} {
		interfaceSlice[i] = d
	}
	// x := List{}
	x := NewList()
	// fmt.Println(x)
	for _, value := range []int{5, 6, 7, 8, 9} {
		x.Append(value)
	}
	fmt.Println(x)
	for i, d := range []int{10, 11, 13} {
		interfaceSlice[i] = d
	}
	x.Extend(interfaceSlice)
	fmt.Println(x)
	x.Insert(1, 9)
	fmt.Println(x)

	x.Remove(11)
	fmt.Println(x)

	c := x.Copy()

	y := x.Pop()
	fmt.Println(y)
	z := x.Pop(2)
	fmt.Println(z)

	fmt.Println(x.Index(9))
	fmt.Println(x.Count(9))

	fmt.Println(x.Count(9))
	fmt.Println(x.Len())
	fmt.Println("If key 9 in:", IfKeyIn(9, x))

	// x.Sort()
	// fmt.Println(x)
	// x.Reverse()
	// fmt.Println(x)

	// check copy list
	c.Append(-1)
	fmt.Println(c)

	x.Clear()
	fmt.Println(x)

}
