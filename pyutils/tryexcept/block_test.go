package tryexcept

import (
	"fmt"
	"testing"
)

func TestTryExcept(t *testing.T) {
	fmt.Println("We started")
	res := "15"
	mapPointer := map[string]string{
		"float":  "12.334",
		"int":    "200",
		"string": "anything",
	}
	x := Try{
		Try: func() {
			fmt.Println("I tried")
			res = mapPointer["abc"]
			fmt.Println(res)
		},
		Except: func(e Exception) {
			fmt.Printf("Caught %v\n", e)
		},
		Finally: func() {
			fmt.Println("Finally...")
			res = "100"
		},
	}.Do()
	fmt.Println("Response", x)
	fmt.Println("We went on", res)
	// var testUint uint
	// testUint = 0
	// fmt.Println(testUint)
}
