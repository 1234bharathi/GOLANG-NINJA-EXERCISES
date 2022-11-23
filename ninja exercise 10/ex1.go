package main

import (
	"fmt"
)

func main() {
	// using anonymous func
	c := make(chan int)
	// using buffered
	c2 := make(chan int, 1)
	go func() {
		c2 <- 43
	}()
	fmt.Println(<-c2)
	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}
