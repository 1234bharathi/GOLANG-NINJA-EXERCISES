package main

import (
	"fmt"
)

func main() {
	type bharathi int
	var x bharathi
	fmt.Printf("%T %v", x, x)
	x = 42
	fmt.Printf("%T %v", x, x)
}
