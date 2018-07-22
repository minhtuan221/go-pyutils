package pyList

import (
	"fmt"
)

type List struct {
	List []interface{}
}

func (list *List) Append(x interface{}) {
	list.List = append(list.List, x)
}

func (list *List) Extend(x []interface{}) {
	list.List = append(list.List, x...)
}

func (list *List) Insert(i int, x interface{}) {

	// Make space in the array for a new element. You can assign it any value.
	list.List = append(list.List, 0)

	// Copy over elements sourced from index 2, into elements starting at index 3.
	copy(list.List[i+1:], list.List[i:])

	// assign value to index
	list.List[i] = x
}

func (list *List) Remove(x interface{}) {
	// find value of x
	for i, value := range list.List {
		if value == x {
			// Where a is the slice, and i is the index of the element you want to delete:
			list.List = append(list.List[:i], list.List[i+1:]...)
			break
		}
	}
}

func (list *List) Pop(x ...int) interface{} {
	k := len(list.List) - 1
	if len(x) == 0 {
		// make a copy of last item
		res := list.List[k]
		// remove the last item in list
		list.List = append(list.List[:k], list.List[k+1:]...)
		return res
	} else {
		i := x[0]
		res := list.List[i]
		// remove item in index i
		list.List = append(list.List[:i], list.List[i+1:]...)
		return res
	}
}

func (list *List) Clear() {
	list.List = nil
}

func (list *List) Index(x interface{}) []interface{} {
	// find value of x
	var res []interface{}
	for i, value := range list.List {
		if value == x {
			// Where a is the slice, and i is the index of the element you want to delete:
			res = append(res, i)
		}
	}
	return res
}

func (list *List) Count(x interface{}) interface{} {
	// find value of x
	res := 0
	for _, value := range list.List {
		if value == x {
			// Where a is the slice, and i is the index of the element you want to delete:
			res++
		}
	}
	return res
}

// func (list *List) Sort() {
// 	switch expression {
// 	case condition:

// 	}
// 	sort.Sort(sort.IntSlice(list.List))
// }

// func (List *List) Reverse() {
// 	sort.Sort(sort.Reverse(sort.IntSlice(List.List)))
// }

func (list *List) Copy() List {
	res := List{}
	// process a deepcopy
	res.List = make([]interface{}, len(list.List))
	copy(res.List, list.List)
	return res
}

type test struct {
	value int
}

func TestList() {
	// var dataSlice []int = foo()
	var interfaceSlice []interface{} = make([]interface{}, 3)
	for i, d := range []int{1, 2, 3} {
		interfaceSlice[i] = d
	}
	var x List = List{}
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
