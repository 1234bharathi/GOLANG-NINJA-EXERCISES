package main

import (
	"fmt"
)

var f func(x ...int) int

func main() {

	f = func(x ...int) int {
		sum := 0
		for _, i := range x {
			sum += i
		}
		return sum

	}

	xi := []int{2, 4, 6, 8, 10}
	foo(xi...)
}
func foo(x ...int) {
	fmt.Println("in foo the sum is", f(x...))
}
