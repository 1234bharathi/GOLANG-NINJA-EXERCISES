package main

import (
	"fmt"
)

// struct type person
type person struct {
	first string
	last  string
	age   int
}

func changeme(p *person) {
	(*p).first = "Harry"
	(*p).last = "Poter"
	(*p).age = 35
}
func print(p person) {
	fmt.Println("The first name is", p.first)
	fmt.Println("The last name is", p.last)
	fmt.Println("The age is", p.age)
}
func main() {
	p1 := person{
		first: "Tom",
		last:  "Clarry",
		age:   27,
	}
	println("Before change me")
	print(p1)
	changeme(&p1)
	println("After change me")
	print(p1)

}
