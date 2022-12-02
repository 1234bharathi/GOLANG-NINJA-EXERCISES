
package main

import (
	"fmt"
)

func main() {

	var arr [5]int
	arr[0] = 1
	arr[1] = 1
	arr[2] = 1
	arr[3] = 1
	arr[4] = 1

	for i, v := range arr {
		fmt.Printf("%T\t %d\t %d\n", arr, i, v)
	}

}
