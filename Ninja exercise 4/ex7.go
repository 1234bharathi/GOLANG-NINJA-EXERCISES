package main

import (
	"fmt"
)

func main() {

	st := []string{"James", "Bond", "Shaken, not stirred"}
	st2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	str := [][]string{st, st2}
	fmt.Println(str)
	for i, v := range str {
		for j, v2 := range v {
			fmt.Println(j, v2)
		}
		fmt.Println(i, v)
	}
}
