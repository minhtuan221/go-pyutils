package pylist

// type Python struct {
// 	Key interface{}
// }

// // if key in x and ...
// //		do something
// func IfKey(key interface{}) *Python {
// 	return &Python{Key: key}
// }

// func (py *Python) In(list iterable) bool {
// 	return list.Contain(py.Key)
// }

// for key in x and ...
//		do something

// Keywords is the map contain all optional arguments
type Keywords struct {
	Values map[string]interface{}
}

// Get = get the value of optinal argument by their name
func (kw *Keywords) Get(key string, defaultValue interface{}) interface{} {
	if value, exist := kw.Values[key]; exist {
		//return value if key exist
		return value
	}
	return defaultValue
}
