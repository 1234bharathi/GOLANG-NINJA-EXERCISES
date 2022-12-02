package main

import (
	"fmt"
)

func main() {
	m := map[string][]string{
		"pers1": []string{`bond_james`, `Shaken, not stirred`, `Martinis`, `Women`},
		"pers2": []string{`moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`},
		"pers3": []string{`no_dr`, `Being evil`, `Ice cream`, `Sunsets`},
	}
	fmt.Println(m)

	m["pres4"] = []string{`choclate`, `pen`, `laptop`, `bottle`}
	for i, v := range m {
		fmt.Println("record", i)
		fmt.Println(i, v)
		for j, v2 := range v {
			fmt.Printf("\nthe value in %d is %v", j, v2)
		}

	}
	delete(m, `pres2`)

	v, ok := m[`pres2`]
	fmt.Println("\nvalue for person 2 is not found")
	fmt.Println(v, ok)

}
