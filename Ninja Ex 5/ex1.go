package main

import "fmt"

type person struct {
	first    string
	last     string
	icecream []string
}

func main() {
	p1 := person{
		first:    `Bharathi`,
		last:     `Selvakumar`,
		icecream: []string{`chocklate`, `strawberry`},
	}

	p2 := person{
		first:    `Anbarasi`,
		last:     `Selvakumar`,
		icecream: []string{`butterscotch`, `blackcurrent`},
	}

	fmt.Println(p1.first, p1.last)
	for i, v := range p1.icecream {
		fmt.Println(i, v)

	}

	fmt.Println(p2.first, p2.last)
	for i, v := range p2.icecream {
		fmt.Println(i, v)

	}
}
