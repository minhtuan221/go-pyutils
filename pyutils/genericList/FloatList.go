package genericList

import (
	"fmt"
	"sort"
)

type FloatList struct {
	List []float64
}

func (floatlist *FloatList) Append(x float64) {
	floatlist.List = append(floatlist.List, x)
}

func (floatlist *FloatList) Extend(x []float64) {
	floatlist.List = append(floatlist.List, x...)
}

func (floatlist *FloatList) Insert(i int, x float64) {

	// Make space in the array for a new element. You can assign it any value.
	floatlist.List = append(floatlist.List, 0.0)

	// Copy over elements sourced from index 2, into elements starting at index 3.
	copy(floatlist.List[i+1:], floatlist.List[i:])

	// assign value to index
	floatlist.List[i] = x
}

func (floatlist *FloatList) Remove(x float64) {
	// find value of x
	for i, value := range floatlist.List {
		if value == x {
			// Where a is the slice, and i is the index of the element you want to delete:
			floatlist.List = append(floatlist.List[:i], floatlist.List[i+1:]...)
			break
		}
	}
}

func (floatlist *FloatList) Pop(x ...int) float64 {
	k := len(floatlist.List) - 1
	if len(x) == 0 {
		// make a copy of last item
		res := floatlist.List[k]
		// remove the last item in list
		floatlist.List = append(floatlist.List[:k], floatlist.List[k+1:]...)
		return res
	} else {
		i := x[0]
		res := floatlist.List[i]
		// remove item in index i
		floatlist.List = append(floatlist.List[:i], floatlist.List[i+1:]...)
		return res
	}
}

func (floatlist *FloatList) Popleft() {
	floatlist.Pop(0)
}

func (floatlist *FloatList) Clear() {
	floatlist.List = nil
}

func (floatlist *FloatList) Index(x float64) []int {
	// find value of x
	res := []int{}
	for i, value := range floatlist.List {
		if value == x {
			// Where a is the slice, and i is the index of the element you want to delete:
			res = append(res, i)
		}
	}
	return res
}

func (floatlist *FloatList) Count(x float64) int {
	// find value of x
	res := 0
	for _, value := range floatlist.List {
		if value == x {
			// Where a is the slice, and i is the index of the element you want to delete:
			res++
		}
	}
	return res
}

func (floatlist *FloatList) Sort() {
	sort.Sort(sort.Float64Slice(floatlist.List))
}

func (floatlist *FloatList) Reverse() {
	sort.Sort(sort.Reverse(sort.Float64Slice(floatlist.List)))
}

func (floatlist *FloatList) Copy() FloatList {
	res := FloatList{}
	// process a deepcopy
	res.List = make([]float64, len(floatlist.List))
	copy(res.List, floatlist.List)
	return res
}

func (floatlist FloatList) Len() int {
	return len(floatlist.List)
}

type iterable interface {
	Len() int
}

func Len(list iterable) int {
	return list.Len()
}

func TestFloat() {
	x := FloatList{[]float64{1.2, 2.454, 3.12}}
	// fmt.Println(x)
	for _, value := range []float64{5.4, 6.3, 7.02, 8.0, 9.9} {
		x.Append(value)
	}
	fmt.Println(x)
	x.Extend([]float64{10.10, 11.12})
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

	fmt.Println(x.Index(7.02))
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
