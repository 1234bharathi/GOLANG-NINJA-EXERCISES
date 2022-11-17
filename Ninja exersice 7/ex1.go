package main

import (
	"fmt"
)

func main() {
	x := 45
	fmt.Println("value of x", x)
	fmt.Printf(" type of x :%T value: %v\n", x, x)
	fmt.Printf("type of &x :%T value: %v\n", &x, &x)

	y := &x
	fmt.Println("y := &x\nvalue of y", y)
	fmt.Printf("type of y: %T value: %v\n", y, y)
	fmt.Printf("type of *y : %T value: %v\n", *y, *y)

}
