package pylist

import (
	"fmt"
	"hash/fnv"
)

func hash(s string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(s))
	return h.Sum64()
}
func testHash() {
	fmt.Println(hash("HelloWorld"))
	fmt.Println(hash("HelloWorld."))
}
