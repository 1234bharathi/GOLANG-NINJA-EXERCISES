package main

import (
	"fmt"
)

func square(a int) func() int {
	sq := a * a
	return func() int {
		return sq
	}
}
func main() {

	x := square(5)
	fmt.Println(x())
	fmt.Println(square(5)())
}
