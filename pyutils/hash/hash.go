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

func hash32(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func testHash() {
	fmt.Println(hash("HelloWorld"))
	fmt.Println(hash("HelloWorld."))
}
