package main

import "fmt"

type person struct {
	name string
}

func (p *person) speak() {
	fmt.Println("NAME OF PERSON IS", p.name)
}

type human interface {
	speak()
}

func saysomething(h human) {
	fmt.Println("CALLING FUNC SAYSOMETHING")
	h.speak()
}
func main() {

	p := person{
		name: "bharathi",
	}
	saysomething(&p)
	fmt.Println("CALLING FUNC SPEAK WITHOUT POINTER")
	p.speak()

}
