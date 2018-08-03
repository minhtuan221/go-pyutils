package pylist

import (
	"fmt"
)

type iterable interface {
	Len() int
	Contain(key interface{}) bool
}

func Len(list iterable) int {
	return list.Len()
}

func IfKeyIn(key interface{}, list iterable) bool {
	return list.Contain(key)
}

type List struct {
	Values []interface{}
}

func (list *List) Append(x interface{}) {
	list.Values = append(list.Values, x)
}

func (list *List) Extend(x []interface{}) {
	list.Values = append(list.Values, x...)
}

func (list *List) Insert(i int, x interface{}) {

	// Make space in the array for a new element. You can assign it any value.
	list.Values = append(list.Values, 0)

	// Copy over elements sourced from index 2, into elements starting at index 3.
	copy(list.Values[i+1:], list.Values[i:])

	// assign value to index
	list.Values[i] = x
}

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

func (list *List) Popleft() {
	list.Pop(0)
}

func (list *List) Clear() {
	list.Values = nil
}

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

func (list List) Contain(x interface{}) bool {
	// find value of x
	for _, value := range list.Values {
		if value == x {
			// Where a is the slice, and i is the index of the element you want to delete:
			return true
		}
	}
	return false
}

func (list *List) Count(x interface{}) int {
	// find value of x
	res := 0
	for _, value := range list.Values {
		if value == x {
			// Where a is the slice, and i is the index of the element you want to delete:
			res++
		}
	}
	return res
}

func (list *List) Copy() List {
	res := List{}
	// process a deepcopy
	res.Values = make([]interface{}, len(list.Values))
	copy(res.Values, list.Values)
	return res
}

func (list List) Len() int {
	return len(list.Values)
}

type test struct {
	value int
}

func TestList() {
	// var dataSlice []int = foo()
	interfaceSlice := make([]interface{}, 3)
	for i, d := range []int{1, 2, 3} {
		interfaceSlice[i] = d
	}
	x := List{}
	// fmt.Println(x)
	for _, value := range []int{5, 6, 7, 8, 9} {
		x.Append(test{value})
	}
	fmt.Println(x)
	for i, d := range []int{10, 11, 13} {
		interfaceSlice[i] = d
	}
	x.Extend(interfaceSlice)
	fmt.Println(x)
	x.Insert(3, 9)
	fmt.Println(x)

	x.Remove(11)
	fmt.Println(x)

	c := x.Copy()

	y := x.Pop()
	fmt.Println(y)
	z := x.Pop(2)
	fmt.Println(z)

	fmt.Println(x.Index(9))
	fmt.Println(x.Count(9))

	fmt.Println(x.Count(test{9}))
	fmt.Println(x.Len())
	fmt.Println("If key 9 in:", IfKeyIn(9, x))

	// x.Sort()
	// fmt.Println(x)
	// x.Reverse()
	// fmt.Println(x)

	// check copy list
	c.Append(-1)
	fmt.Println(c)

	x.Clear()
	fmt.Println(x)

}
