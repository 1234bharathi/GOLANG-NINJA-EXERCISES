package main

import (
	"fmt"
)

func main() {

	slicee := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i, v := range slicee {
		fmt.Printf("%d\t %d\n", i, v)
	}
	fmt.Printf("%T\ttype of slice", slicee)
}
