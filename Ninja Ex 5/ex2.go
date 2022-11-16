package main

import "fmt"

// struct
type person struct {
	first string
	last  string
	// slice
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
	// map
	m := map[string]person{
		p1.first: p1,
		p2.first: p2,
	}

	for k, v := range m {
		fmt.Println(k)
		fmt.Println()
		for j, v2 := range v.icecream {
			fmt.Println(j, v2)
		}
		fmt.Println()
	}
}
