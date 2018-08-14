package main

import (
	"fmt"
	"reflect"
)

type Block struct {
	Cb      func(...interface{}) []interface{}
	Params  []interface{}
	Catcher func(Exception)
	Finally func()
	Error   interface{}
}

func general_func(oldfunc interface{}) func(...interface{}) []interface{} {
	if reflect.TypeOf(oldfunc).Kind() != reflect.Func {
		panic("protected item is not a function")
	}
	return func(args ...interface{}) []interface{} {
		fmt.Println("Protected")
		vargs := make([]reflect.Value, len(args))
		for n, v := range args {
			vargs[n] = reflect.ValueOf(v)
		}
		ret_vals := reflect.ValueOf(oldfunc).Call(vargs)
		to_return := make([]interface{}, len(ret_vals))
		for n, v := range ret_vals {
			to_return[n] = v.Interface()
		}
		return to_return
	}
}

func Callback(params ...interface{}) *Block {
	return &Block{Params: params}
}

func (b *Block) Then(cb func(...interface{}) []interface{}) *Block {
	// exec last cb
	b.Cb = cb
	b.Do()
	// assign
	return b
}

// Catch blocks are only run in the case of an thrown exception. Regular panics
// are ignored and will behave as normal.
//
func (b *Block) Catch(exc func(Exception)) *Block {
	b.Catcher = exc
	return b
}

// Add a finally block. These will be run whether or not an exception was
// thrown. However, if a regular panic occurs this function will not be run and
// the panic will behave as normal.
func (b *Block) Final(finally func()) *Block {
	b.Finally = finally
	return b
}

type Exception interface{}

func Throw(up Exception) {
	panic(up)
}

func (blockcode Block) Do() (Block, interface{}) {
	var is_ok interface{}
	if blockcode.Finally != nil {
		defer blockcode.Finally()
	}
	if blockcode.Catcher != nil {
		defer func() {
			if r := recover(); r != nil {
				blockcode.Catcher(r)
				is_ok = r
			}
		}()
	}
	blockcode.Params = blockcode.Cb(blockcode.Params)
	blockcode.Error = is_ok
	return blockcode, is_ok
}

// func Keyword(kw []interface{}, vtype struct(), index int) {
// 	converted := vtype{}
// }
//  Keyword(struct{ name string, age int})

func main() {
	fmt.Println("Test call back")
	Callback("Hello").Then(func(s ...interface{}) []interface{} {
		fmt.Println(s)
		fmt.Println("Hi")
		return s
	}).Catch(func(e Exception) {
		fmt.Println("Got error: =>", e)
	}).Final(func() {
		fmt.Println("This is finally")
	}).Then(func(s ...interface{}) []interface{} {
		fmt.Println(s)
		fmt.Println("\n Nice to meet you")
		return s
	}).Then(func(s ...interface{}) []interface{} {
		fmt.Println("Nice to meet you too")
		Throw("Error Occur.... !!")
		return s
	})
}
