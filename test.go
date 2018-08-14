package main

import (
	pylist "Pyutils/pyutils/Pylist"
	"Pyutils/pyutils/exc"
	genericlist "Pyutils/pyutils/genericList"
	"fmt"
	"log"

	ini "gopkg.in/ini.v1"
)

func main() {
	defer fmt.Println("Your test is here: ")
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
	// cfg, err := ini.Load("config.ini")
	// if err != nil {
	// 	fmt.Printf("Fail to read file: %v", err)
	// 	os.Exit(1)
	// }
	// // Classic read of values, default section can be represented as empty string
	// fmt.Println("App Mode:", cfg.Section("env").Key("host").String())
	// z := cfg.Section("env").Key("host").String()
	exc.Try(func() {
		// strings.(z)
		cfg, _ := ini.Load("confie.ini")
		fmt.Println("App Mode:", cfg.Section("env").Key("host").String())
	}).Catch(&exc.Exception{}, func(t exc.Throwable) {
		log.Println(t)
		exc.Rethrow(t, exc.Errorf("rethrow after logging"))
		log.Println("Error have been reported")
	}).Error()
	// tryexcept.Try{
	// 	Try: func() {
	// 		cfg, _ := ini.Load("confie.ini")
	// 		fmt.Println("App Mode:", cfg.Section("env").Key("host").String())
	// 	},
	// 	Except: func(e tryexcept.Exception) {
	// 		fmt.Println(e)
	// 		fmt.Println("Error have been captured")
	// 	},
	// 	Finally: func() {
	// 		fmt.Println("This is finally")
	// 	},
	// }.Do()
	// decorator
	// blockchain.PrettyPrint(balance)
}
