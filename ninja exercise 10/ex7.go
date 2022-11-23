package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int)
	var wg sync.WaitGroup

	go func() {
		for i := 1; i <= 10; i++ {
			wg.Add(1)
			go func(k int) {

				for j := 1; j <= 10; j++ {
					c <- (j * k)
				}
				wg.Done()
			}(i)

		}

		wg.Wait()
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}
	fmt.Println("exiting")
}
