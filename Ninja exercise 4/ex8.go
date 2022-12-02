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

	for i, v := range m {
		fmt.Println("record", i)
		fmt.Println(i, v)
		for j, v2 := range v {
			fmt.Printf("\nthe value in %d is %v", j, v2)
		}

	}
}
