package pylist

import (
	"fmt"
)

type Dict struct {
	Values map[interface{}]interface{}
}

func (dict *Dict) Get(key interface{}) interface{} {
	return dict.Values[key]
}

func (dict *Dict) Set(key interface{}, value interface{}) {
	dict.Values[key] = value
}

func (dict *Dict) Pop(key interface{}) interface{} {
	res := dict.Values[key]
	delete(dict.Values, key)
	return res
}

func (dict *Dict) Update(otherdict map[interface{}]interface{}) {
	for key, value := range otherdict {
		dict.Values[key] = value
	}
}

func (dict *Dict) Setdefault(key interface{}, defaultValue interface{}) interface{} {
	if value, exist := dict.Values[key]; exist {
		//do something here
		return value
	}
	dict.Values[key] = defaultValue
	return defaultValue
}

func (dict *Dict) Clear() {
	dict.Values = nil
}
func (dict *Dict) Copy() Dict {
	// process a deepcopy
	res := Dict{make(map[interface{}]interface{})}
	// Copy from the original map to the target map
	for key, value := range dict.Values {
		res.Values[key] = value
	}
	return res
}

func (dict Dict) Len() int {
	return len(dict.Values)
}
func (dict Dict) IfKeyIn(x interface{}) bool {
	// find value of x
	for key := range dict.Values {
		if key == x {
			// Where a is the slice, and i is the index of the element you want to delete:
			return true
		}
	}
	return false
}

type test2 struct {
	value int
}

func TestDict() {
	// var dataSlice []int = foo()
	mydict := Dict{make(map[interface{}]interface{})}
	mydict.Set("s1", 12)
	mydict.Set("s2", "My Name")
	fmt.Println(mydict)
	fmt.Println(IfKeyIn("s1", mydict))

}
