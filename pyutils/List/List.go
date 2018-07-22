package main

import (
	"fmt"
	// "sort"
)

type List struct {
	List []interface{}
}

func (List *List) Append(x interface{}) {
	List.List = append(List.List, x)
}

func (List *List) Extend(x []interface{}) {
	List.List = append(List.List, x...)
}

func (List *List) Insert(i int, x interface{}) {

	// Make space in the array for a new element. You can assign it any value.
	List.List = append(List.List, 0)

	// Copy over elements sourced from index 2, into elements starting at index 3.
	copy(List.List[i+1:], List.List[i:])

	// assign value to index
	List.List[i] = x
}

func (List *List) Remove(x interface{}) {
	// find value of x
	for i, value := range List.List {
		if value == x {
			// Where a is the slice, and i is the index of the element you want to delete:
			List.List = append(List.List[:i], List.List[i+1:]...)
			break
		}
	}
}

func (List *List) Pop(x ...int) interface{} {
	k := len(List.List) - 1
	if len(x) == 0 {
		// make a copy of last item
		res := List.List[k]
		// remove the last item in list
		List.List = append(List.List[:k], List.List[k+1:]...)
		return res
	} else {
		i := x[0]
		res := List.List[i]
		// remove item in index i
		List.List = append(List.List[:i], List.List[i+1:]...)
		return res
	}
}

func (List *List) Clear() {
	List.List = nil
}

func (mList *List) Index(x interface{}) []interface{} {
	// find value of x
	var res []interface{}
	for i, value := range mList.List {
		if value == x {
			// Where a is the slice, and i is the index of the element you want to delete:
			res = append(res, i)
		}
	}
	return res
}

func (List *List) Count(x interface{}) interface{} {
	// find value of x
	res := 0
	for _, value := range List.List {
		if value == x {
			// Where a is the slice, and i is the index of the element you want to delete:
			res++
		}
	}
	return res
}

// func (List *List) Sort() {
// 	sort.Sort(sort.IntSlice(List.List))
// }

// func (List *List) Reverse() {
// 	sort.Sort(sort.Reverse(sort.IntSlice(List.List)))
// }

func (mList *List) Copy() List {
	res := List{}
	// process a deepcopy
	res.List = make([]interface{}, len(mList.List))
	copy(res.List, mList.List)
	return res
}

type test struct {
	value int
}

func main() {
	// var dataSlice []int = foo()
	var interfaceSlice []interface{} = make([]interface{}, 3)
	for i, d := range []int{1, 2, 3} {
		interfaceSlice[i] = d
	}
	var x List = List{}
	// fmt.Println(x)
	for _, value := range []int{5, 6, 7, 8, 9} {
		x.Append(value)
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
