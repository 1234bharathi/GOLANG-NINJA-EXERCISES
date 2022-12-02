package main

import (
	"fmt"
)

func main() {

	slicee := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Println(append(slicee, 52))

	fmt.Println(append(slicee, 53, 54, 55))
	y := []int{56, 57, 58, 59, 60}

	fmt.Println(append(slicee, y...))
}
