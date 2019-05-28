package tryexcept

import (
	"fmt"

	"github.com/pkg/errors"
)

// Try basic/root struct of try/except
type Try struct {
	Try     func()
	Except  func(interface{})
	Finally func()
	Err     interface{}
}

// TryThrowError try a function and throw error to out side. Example,
/* how to use it:
	```your above code
	userP := &pb.User{}
	tryexcept.TryThrowError(func() {
		userToProto(user, userP)
	})
	return userP```
You can print full traceback with fmt.Printf("%+v\n", err). TryThrowError will print error to stderr by default
*/
func TryThrowError(try func()) error {
	var err error
	except := func(e interface{}) {
		err = Traceback(e)
		fmt.Printf("%+v\n", err)
	}
	just := &Try{try, except, nil, nil}
	just.Do()
	return err
}

// Throw simillar to panic
func Throw(up interface{}) {
	panic(up)
}

// Do execute function try on Try
func (tcf Try) Do() (bool, interface{}) {
	isOk := true
	if tcf.Finally != nil {

		defer tcf.Finally()
	}
	if tcf.Except != nil {
		defer func() {
			if r := recover(); r != nil {
				tcf.Except(r)
				isOk = false
				tcf.Err = r
			}
		}()
	}
	tcf.Try()
	return isOk, tcf.Err
}

// Traceback return err with full stack trace
func Traceback(e interface{}) error {
	err, exp := e.(error)
	if !exp {
		panic("Uncatchable Error: Error while converting interface into error object in traceback function")
	}
	traceback := errors.WithStack(err)
	return traceback
}

// GetTraceBack return full traceback as a string
func GetTraceBack(e interface{}) string {
	err := Traceback(e)
	return fmt.Sprintf("%+v\n", err)
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
