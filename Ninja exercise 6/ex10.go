package main

import (
	"fmt"
)

var x int

func main() {
	x = 0
	fmt.Println("in main the value of x is", x)
	foo()
	fmt.Println(x)
}
func foo() {

	x = 25
	fmt.Println("in main the value of x is", x)

}
