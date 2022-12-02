package main

import (
	"fmt"
)

func main() {
	type bharathi int
	var x bharathi
	var y int
	fmt.Printf("%T %v", x, x)
	x = 42
	fmt.Printf("%T %v", x, x)
	fmt.Printf("%T %v", y, y)
	y = int(x)
	fmt.Printf("%T %v", y, y)

}
