package main

import (
	"fmt"
	"strconv"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		fmt.Println("c<-x = " + strconv.Itoa(x) + " cap=" + strconv.Itoa(cap(c)))
		x, y = y, x+y
	}
	close(c)
}

func main1() {
	c := make(chan int, 2)
	go fibonacci(cap(c)+10, c)
	for i := range c {
		fmt.Println(i)
	}
}
