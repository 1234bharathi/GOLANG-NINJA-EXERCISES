package main

import (
	"fmt"
)

func main() {

	slicee := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Println(append(slicee[:3], slicee[6:]...))
}
