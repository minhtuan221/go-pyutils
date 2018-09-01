package tryexcept

import (
	"fmt"

	"github.com/pkg/errors"
)

type Try struct {
	Try     func()
	Except  func(Exception)
	Finally func()
}

type Exception interface{}

func Throw(up Exception) {
	panic(up)
}

func (tcf Try) Do() bool {
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

func Traceback(e Exception) error {
	err, exp := e.(error)
	if !exp {
		panic("Uncatchable Error: Error while converting interface into error object in traceback function")
	}
	traceback := errors.WithStack(err)
	return traceback
}

type Keywords struct {
	Values []interface{}
}

func (kw *Keywords) Get(i uint, defaultValue interface{}) interface{} {
	x := int(i)
	if x < len(kw.Values) {
		return kw.Values[x]
	}
	return defaultValue
}

// OptionArgs function create keyword arguments
func OptionArgs(Val ...interface{}) *Keywords {
	return &Keywords{Val}
}

func main() {
	fmt.Println("We started")
	x := Try{
		Try: func() {
			fmt.Println("I tried")
			// float64("abc")
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
	// var testUint uint
	// testUint = 0
	// fmt.Println(testUint)
}
