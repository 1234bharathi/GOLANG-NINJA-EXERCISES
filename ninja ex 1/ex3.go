package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {

	s := fmt.Sprintf("it is %v that %v is age of %v", z, x, y)
	fmt.Println(s)
}
