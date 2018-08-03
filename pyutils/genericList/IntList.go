package genericlist

import (
	"fmt"
	"sort"
)

type IntList struct {
	Values []int
}

func (intlist *IntList) Append(x int) {
	intlist.Values = append(intlist.Values, x)
}

func (intlist *IntList) Extend(x []int) {
	intlist.Values = append(intlist.Values, x...)
}

func (intlist *IntList) Insert(i int, x int) {

	// Make space in the array for a new element. You can assign it any value.
	intlist.Values = append(intlist.Values, 0)

	// Copy over elements sourced from index 2, into elements starting at index 3.
	copy(intlist.Values[i+1:], intlist.Values[i:])

	// assign value to index
	intlist.Values[i] = x
}

func (intlist *IntList) Remove(x int) {
	// find value of x
	for i, value := range intlist.Values {
		if value == x {
			// Where a is the slice, and i is the index of the element you want to delete:
			intlist.Values = append(intlist.Values[:i], intlist.Values[i+1:]...)
			break
		}
	}
}

func (intlist *IntList) Pop(x ...int) int {
	k := len(intlist.Values) - 1
	if len(x) == 0 {
		// make a copy of last item
		res := intlist.Values[k]
		// remove the last item in list
		intlist.Values = append(intlist.Values[:k], intlist.Values[k+1:]...)
		return res
	} else {
		i := x[0]
		res := intlist.Values[i]
		// remove item in index i
		intlist.Values = append(intlist.Values[:i], intlist.Values[i+1:]...)
		return res
	}
}

func (intlist *IntList) Popleft() {
	intlist.Pop(0)
}

func (intlist *IntList) Clear() {
	intlist.Values = nil
}

func (intlist *IntList) Index(x int) []int {
	// find value of x
	res := []int{}
	for i, value := range intlist.Values {
		if value == x {
			// Where a is the slice, and i is the index of the element you want to delete:
			res = append(res, i)
		}
	}
	return res
}

func (intlist *IntList) Count(x int) int {
	// find value of x
	res := 0
	for _, value := range intlist.Values {
		if value == x {
			// Where a is the slice, and i is the index of the element you want to delete:
			res++
		}
	}
	return res
}

func (intlist *IntList) Sort() {
	sort.Sort(sort.IntSlice(intlist.Values))
}

func (intlist *IntList) Reverse() {
	sort.Sort(sort.Reverse(sort.IntSlice(intlist.Values)))
}

func (intlist *IntList) Copy() IntList {
	res := IntList{}
	// process a deepcopy
	res.Values = make([]int, len(intlist.Values))
	copy(res.Values, intlist.Values)
	return res
}

func (list IntList) Contain(x int) bool {
	// find value of x
	for _, value := range list.Values {
		if value == x {
			// Where a is the slice, and i is the index of the element you want to delete:
			return true
		}
	}
	return false
}

func (intlist IntList) Len() int {
	return len(intlist.Values)
}

func TestInt() {
	x := IntList{[]int{1, 2, 3}}
	// fmt.Println(x)
	for _, value := range []int{5, 6, 7, 8, 9} {
		x.Append(value)
	}
	fmt.Println(x)
	x.Extend([]int{10, 11})
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

	x.Sort()
	fmt.Println(x)
	x.Reverse()
	fmt.Println(x)

	// check copy list
	c.Append(-1)
	fmt.Println(c)
	fmt.Println(Len(c))

	x.Clear()
	fmt.Println(x)

}
