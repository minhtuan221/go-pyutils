package pylist

import (
	"fmt"
)

type Dict struct {
	D map[interface{}]interface{}
}

func (dict *Dict) Get(key interface{}) interface{} {
	return dict.D[key]
}

func (dict *Dict) Set(key interface{}, value interface{}) {
	dict.D[key] = value
}

func (dict *Dict) Pop(key interface{}) interface{} {
	res := dict.D[key]
	delete(dict.D, key)
	return res
}

func (dict *Dict) Update(otherdict map[interface{}]interface{}) {
	for key, value := range otherdict {
		dict.D[key] = value
	}
}

func (dict *Dict) Setdefault(key interface{}, defaultValue interface{}) interface{} {
	if value, exist := dict.D[key]; exist {
		//do something here
		return value
	}
	dict.D[key] = defaultValue
	return defaultValue
}

func (dict *Dict) Values() map[interface{}]interface{} {
	return dict.D
}

func (dict *Dict) Clear() {
	dict.D = nil
}
func (dict *Dict) Copy() Dict {
	// process a deepcopy
	res := Dict{make(map[interface{}]interface{})}
	// Copy from the original map to the target map
	for key, value := range dict.D {
		res.D[key] = value
	}
	return res
}

func (dict Dict) Len() int {
	return len(dict.D)
}
func (dict Dict) IfKeyIn(x interface{}) bool {
	// find value of x
	for key := range dict.D {
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
	var mydict Dict = Dict{make(map[interface{}]interface{})}
	mydict.Set("s1", 12)
	mydict.Set("s2", "My Name")
	fmt.Println(mydict)
	fmt.Println(IfKeyIn("s1", mydict))

}
