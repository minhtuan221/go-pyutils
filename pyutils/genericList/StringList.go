package genericList

import (
	"fmt"
	"sort"
)

type StringList struct {
	List []string
}

func (stringlist *StringList) Append(x string) {
	stringlist.List = append(stringlist.List, x)
}

func (stringlist *StringList) Extend(x []string) {
	stringlist.List = append(stringlist.List, x...)
}

func (stringlist *StringList) Insert(i int, x string) {

	// Make space in the array for a new element. You can assign it any value.
	stringlist.List = append(stringlist.List, "")

	// Copy over elements sourced from index 2, into elements starting at index 3.
	copy(stringlist.List[i+1:], stringlist.List[i:])

	// assign value to index
	stringlist.List[i] = x
}

func (stringlist *StringList) Remove(x string) {
	// find value of x
	for i, value := range stringlist.List {
		if value == x {
			// Where a is the slice, and i is the index of the element you want to delete:
			stringlist.List = append(stringlist.List[:i], stringlist.List[i+1:]...)
			break
		}
	}
}

func (stringlist *StringList) Pop(x ...int) string {
	k := len(stringlist.List) - 1
	if len(x) == 0 {
		// make a copy of last item
		res := stringlist.List[k]
		// remove the last item in list
		stringlist.List = append(stringlist.List[:k], stringlist.List[k+1:]...)
		return res
	} else {
		i := x[0]
		res := stringlist.List[i]
		// remove item in index i
		stringlist.List = append(stringlist.List[:i], stringlist.List[i+1:]...)
		return res
	}
}

func (stringlist *StringList) Popleft() {
	stringlist.Pop(0)
}

func (stringlist *StringList) Clear() {
	stringlist.List = nil
}

func (stringlist *StringList) Index(x string) []int {
	// find value of x
	res := []int{}
	for i, value := range stringlist.List {
		if value == x {
			// Where a is the slice, and i is the index of the element you want to delete:
			res = append(res, i)
		}
	}
	return res
}

func (stringlist *StringList) Count(x string) int {
	// find value of x
	res := 0
	for _, value := range stringlist.List {
		if value == x {
			// Where a is the slice, and i is the index of the element you want to delete:
			res++
		}
	}
	return res
}

func (stringlist *StringList) Sort() {
	sort.Sort(sort.StringSlice(stringlist.List))
}

func (stringlist *StringList) Reverse() {
	sort.Sort(sort.Reverse(sort.StringSlice(stringlist.List)))
}

func (stringlist *StringList) Copy() StringList {
	res := StringList{}
	// process a deepcopy
	res.List = make([]string, len(stringlist.List))
	copy(res.List, stringlist.List)
	return res
}

func (stringlist StringList) Len() int {
	return len(stringlist.List)
}

func TestString() {
	x := StringList{[]string{"Minh", "Tuan", "nguyen"}}
	// fmt.Println(x)
	for _, value := range []string{"nguyen", "thông minh"} {
		x.Append(value)
	}
	fmt.Println(x)
	x.Extend([]string{"10", "Up"})
	fmt.Println(x)
	x.Insert(3, " hay lam")
	fmt.Println(x)

	x.Remove("10")
	fmt.Println(x)

	c := x.Copy()

	y := x.Pop()
	fmt.Println(y)
	z := x.Pop(1)
	fmt.Println(z)

	fmt.Println(x.Index("Up"))
	fmt.Println(x.Count("nguyen"))

	x.Sort()
	fmt.Println(x)
	x.Reverse()
	fmt.Println(x)

	// check copy list
	c.Append("- some thing else")
	fmt.Println(c)
	fmt.Println(Len(c))

	x.Clear()
	fmt.Println(x)

}
