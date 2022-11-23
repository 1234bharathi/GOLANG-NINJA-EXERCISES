package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		c <- 35
	}()
	// value in channel we get true
	v, ok := <-c
	fmt.Println(v, ok)

	close(c)
	//  no value in channel we get 0 and false
	v, ok = <-c
	fmt.Println(v, ok)
}
