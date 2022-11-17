package main

import (
	"fmt"
)

func foo() int {
	fmt.Println("in foo returns int")
	return 26
}
func bar() (int, string) {
	x := "in bar returns an int and string"
	y := 28
	return y, x
}
func main() {
	fmt.Println(foo())
	fmt.Println(bar())
}
