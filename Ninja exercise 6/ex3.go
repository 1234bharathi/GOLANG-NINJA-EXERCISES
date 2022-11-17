package main

import (
	"fmt"
)

func main() {
	defer func1()

	func2()

}
func func1() {
	fmt.Println("I am in function func1 ,and called using defer keyword")
}
func func2() {
	fmt.Println("I am in function func2")
}
