package main

import (
	"fmt"
)

func foo(x ...int) int {
	s := 0
	for _, v := range x {
		s = s + v
	}
	return s
}
func main() {
	xi := []int{10, 2, 3, 4, 5, 6, 7, 8, 9}
	sum := foo(xi...)
	fmt.Println("the sum is", sum)

}
