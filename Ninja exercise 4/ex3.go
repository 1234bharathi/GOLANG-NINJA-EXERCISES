package main

import (
	"fmt"
)

func main() {

	slicee := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52}

	// [42 43 44 45 46]
	fmt.Println(slicee[:5])
	// ● [47 48 49 50 51]
	fmt.Println(slicee[5:10])
	//  [44 45 46 47 48]
	fmt.Println(slicee[2:7])
	// [43 44 45 46 47]
	fmt.Println(slicee[1:6])
}
