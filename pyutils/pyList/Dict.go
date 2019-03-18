package pylist

import (
	"encoding/json"
	"fmt"

	"github.com/minhtuan221/go-pyutils/pyutils/tryexcept"
)

// NewDict =  return new instance of Dict
func NewDict() *Dict {
	return &Dict{make(map[string]interface{})}
}

// Dict = Data structure resemble Dict from python
type Dict struct {
	Values map[string]interface{}
}

// Get = return the value or the memory address of pointer
func (dict *Dict) Get(key string) interface{} {
	return dict.Values[key]
}

// Set = set/replace the value or pointer to Dict with a specified key
func (dict *Dict) Set(key string, value interface{}) {
	dict.Values[key] = value
}

// Pop an item in dict with specified key
func (dict *Dict) Pop(key string) interface{} {
	res := dict.Values[key]
	delete(dict.Values, key)
	return res
}

// Update Dict itself by another dict
func (dict *Dict) Update(otherdict map[string]interface{}) {
	for key, value := range otherdict {
		dict.Values[key] = value
	}
}

// Setdefault = set default value for a key. If key already exist, return the value of key in dict
func (dict *Dict) Setdefault(key string, defaultValue interface{}) interface{} {
	if value, exist := dict.Values[key]; exist {
		//do something here
		return value
	}
	dict.Values[key] = defaultValue
	return defaultValue
}

// Clear = assign nil to Values
func (dict *Dict) Clear() {
	dict.Values = nil
}

// Copy =  return a copy with different pointer
func (dict *Dict) Copy() Dict {
	// process a deepcopy
	res := Dict{make(map[string]interface{})}
	// Copy from the original map to the target map
	for key, value := range dict.Values {
		res.Values[key] = value
	}
	return res
}

// ToJSON = try to convert all to string Json
func (dict *Dict) ToJSON() string {
	data, err := json.Marshal(dict.Values)
	if err != nil {
		tryexcept.Throw("Error when converting to json")
		return "{}"
	}
	return string(data)
}

// Len = return the lenght of value
func (dict Dict) Len() int {
	return len(dict.Values)
}

// Contain find item by item. This method only check basic value type => will not work with pointer
func (dict Dict) Contain(x interface{}) bool {
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

// TestDict is the method for test only
func TestDict() {
	// var dataSlice []int = foo()
	mydict := NewDict() //Dict{make(map[string]interface{})}
	mydict.Set("s1", 12)
	mydict.Set("s2", "My Name")
	fmt.Println(mydict.ToJSON())
	fmt.Println(IfKeyIn("s1", mydict))

}
