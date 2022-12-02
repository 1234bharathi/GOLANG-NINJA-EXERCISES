package main

import (
	"fmt"
)

const (
	// untyped
	a = 2
	// typed
	b int = 2
)

func main() {
	fmt.Println(a, b)
}
