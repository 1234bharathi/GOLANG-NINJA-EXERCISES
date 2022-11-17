package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}
type human interface {
	speak()
}

func main() {
	p1 := person{
		first: "bharathi",
		last:  "selvakumar",
		age:   22,
	}
	s := p1.speak()
	fmt.Println(s)
}

func (p1 person) speak() string {
	fmt.Println("in speak person")
	return (fmt.Sprintf("My name is %s %s and my age is %d", p1.first, p1.last, p1.age))

}
