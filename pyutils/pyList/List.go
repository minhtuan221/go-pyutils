package pylist

import (
	"encoding/json"
	"reflect"

	"github.com/minhtuan221/go-pyutils/pyutils/tryexcept"
)

type iterable interface {
	Len() int
	Contain(key interface{}) bool
}

// Len = simillar to Len but from python
func Len(list iterable) int {
	return list.Len()
}

// IfKeyIn = simillar to contain but from python
func IfKeyIn(key interface{}, list iterable) bool {
	return list.Contain(key)
}

// NewList =  return a new List
func NewList(x ...interface{}) *List {
	return &List{x}
}

// List = Data structure resemble list from python
type List struct {
	Values []interface{}
}

// Append one item to list
func (list *List) Append(x interface{}) {
	list.Values = append(list.Values, x)
}

// Extend list by another list
func (list *List) Extend(x []interface{}) {
	list.Values = append(list.Values, x...)
}

// Insert an item in to a list with specified index position
func (list *List) Insert(i int, x interface{}) {
	if i > list.Len() {
		panic("Insert to out of range position in list")
	} else if i < 0 {
		i = list.Len() + i
	}

	// Make space in the array for a new element. You can assign it any value.
	list.Values = append(list.Values, 0)

	// Copy over elements sourced from index 2, into elements starting at index 3.
	copy(list.Values[i+1:], list.Values[i:])

	// assign value to index
	list.Values[i] = x
}

// Remove an specified item in list
func (list *List) Remove(x interface{}) {
	// find value of x
	for i, value := range list.Values {
		if value == x {
			// Where a is the slice, and i is the index of the element you want to delete:
			list.Values = append(list.Values[:i], list.Values[i+1:]...)
			break
		}
	}
}

// Pop an item in list with specified index
func (list *List) Pop(x ...int) interface{} {
	k := len(list.Values) - 1
	if len(x) == 0 {
		// make a copy of last item
		res := list.Values[k]
		// remove the last item in list
		list.Values = append(list.Values[:k], list.Values[k+1:]...)
		return res
	}
	i := x[0]
	res := list.Values[i]
	// remove item in index i
	list.Values = append(list.Values[:i], list.Values[i+1:]...)
	return res
}

// Popleft the first item add to list
func (list *List) Popleft() {
	list.Pop(0)
}

// Clear = assign nil to list.Value
func (list *List) Clear() {
	list.Values = nil
}

// Index = Delete the item with the specified item
func (list *List) Index(x interface{}) []int {
	// find value of x
	var res []int
	for i, value := range list.Values {
		if value == x {
			// Where a is the slice, and i is the index of the element you want to delete:
			res = append(res, i)
		}
	}
	return res
}

// Contain find item by item. This method only check basic value type => will not work with pointer
func (list List) Contain(x interface{}) bool {
	// find value of x
	for _, value := range list.Values {
		if value == x {
			return true
		}
	}
	return false
}

// ContainDeep find item by item using DeepEqual. This method will return the index of the first value if exist, otherwise, it will return -1
func (list List) ContainDeep(x interface{}) int {
	// find value of x
	for index, value := range list.Values {
		if reflect.DeepEqual(value, x) {
			return index
		}
	}
	return -1
}

// Count how many item in a list. This method only check basic value type => will not work with pointer
func (list *List) Count(x interface{}) int {
	// find value of x
	res := 0
	for _, value := range list.Values {
		if value == x {
			res++
		}
	}
	return res
}

// Copy =  return a copy with different pointer
func (list *List) Copy() List {
	res := List{}
	// process a deepcopy
	res.Values = make([]interface{}, len(list.Values))
	copy(res.Values, list.Values)
	return res
}

// CountDeep how many item in a list. This method will return the number of the values if exist, otherwise, it will return 0
func (list *List) CountDeep(x interface{}) int {
	// find value of x
	res := 0
	for _, value := range list.Values {
		if reflect.DeepEqual(value, x) {
			res++
		}
	}
	return res
}

// FindOneBy => find the first item in a list and return its index. return -1 if not found
func (list *List) FindOneBy(oneItem interface{}, f func(one interface{}, item interface{}) bool) int {
	// find value of x
	for i, value := range list.Values {
		if f(oneItem, value) == true {
			// Where a is the slice, and i is the index of the element you want to delete:
			return i
		}
	}
	return -1
}

// ToJSON = try to convert all to string Json
func (list *List) ToJSON() string {
	data, err := json.Marshal(list.Values)
	if err != nil {
		tryexcept.Throw("Error when converting to json")
		return "{}"
	}
	return string(data)
}

// Len = return the lenght of Value
func (list List) Len() int {
	return len(list.Values)
}

type test struct {
	value int
}
