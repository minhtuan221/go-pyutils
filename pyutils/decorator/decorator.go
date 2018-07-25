package main

import (
	"fmt"
)

type class struct {
	X func(string)
}

func Pydecorator(f func(s string)) func(s string) {
	wraps := func(s string) {
		fmt.Println("Start decorator")
		f(s)
		fmt.Println("Done decorator")
	}
	return wraps
}
func test_wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func doSomething(s string) {
	fmt.Println(s)
}

func doSomethingElse(s string) {
	fmt.Println("OTher thing:", s)
}

// def my_decorator(f):
//     @wraps(f)
//     def wrapper(*args, **kwds):
//         print('Calling decorated function')
// 		   return f(*args, **kwds)
//     return wrapper

type Decorator struct {
	Wrapper func(string)
	F       func(string)
}

func (d *Decorator) Wraps(f func(s string)) func(string) {
	d.F = f
	return d.Wrapper
}

func main() {
	// decorator(doSomething)("Do something")
	increment := test_wrapper()
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	// class1 := class{Pydecorator(doSomething)}
	// class1.X("adfasd")

	// fmt.Println("test decorator")
	// x := Pydecorator(doSomething)
	// x("Do inside decorator")

	decor := Decorator{
		F: doSomething,
		Wrapper: func(s string) {
			fmt.Println("Start decorator")
			doSomething(s)
			fmt.Println("Done decorator")
		},
	}
	decor.Wraps(doSomethingElse)("asddsaf")

	decor2 := Decorator{}
	decor2.Wrapper = func(s string) {
		fmt.Println("Start decorator")
		decor2.F(s)
		fmt.Println("Done decorator")
	}
	decor2.Wraps(doSomething)("Test lan 2")
	// x := decorator(
	// 	func(s string) {
	// 		fmt.Println("Do inside decorator: ", s)
	// 	})

}
