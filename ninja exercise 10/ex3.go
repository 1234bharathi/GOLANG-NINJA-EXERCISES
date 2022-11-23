package main

import (
	"fmt"
)

func main() {
	c := gen()
	receive(c)

	fmt.Println("about to exit")
}

func receive(c <-chan int) {
	for v := range c {
		fmt.Println("val", v)
	}

}

func gen() <-chan int {
	c := make(chan int)
	// use go function or else will reach deadlock and close the chanel as another go routine main is waiting for the values to be read from the channel
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}

		close(c)
	}()

	return c
}
