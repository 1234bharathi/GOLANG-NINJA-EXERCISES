package main

import (
	"fmt"
)

func bar() {
	fmt.Println("In function bar ")

	func() {
		fmt.Println("now in anonymous function ")
	}()
}
func main() {
	bar()
}
