package tryexcept

import (
	"fmt"
)

type Block struct {
	Try     func()
	Except  func(Exception)
	Finally func()
}

type Exception interface{}

func Throw(up Exception) {
	panic(up)
}

func (tcf Block) Do() bool {
	is_ok := true
	if tcf.Finally != nil {

		defer tcf.Finally()
	}
	if tcf.Except != nil {
		defer func() {
			if r := recover(); r != nil {
				tcf.Except(r)
				is_ok = false
			}
		}()
	}
	tcf.Try()
	return is_ok
}

func TestTryExcept() {
	fmt.Println("We started")
	x := Block{
		Try: func() {
			fmt.Println("I tried")
			Throw("Oh,...sh...")
		},
		Except: func(e Exception) {
			fmt.Printf("Caught %v\n", e)
		},
		Finally: func() {
			fmt.Println("Finally...")
		},
	}.Do()
	fmt.Println("Response", x)
	fmt.Println("We went on")
}
