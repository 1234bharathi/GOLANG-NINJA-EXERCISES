package main

import "fmt"

func send(cs chan<- int) {
	for i := 0; i < 100; i++ {
		cs <- i
	}

	close(cs)
}
func main() {
	cs := make(chan int)
	go send(cs)
	for v := range cs {
		fmt.Println(v)
	}
	fmt.Println("exiting")
}
