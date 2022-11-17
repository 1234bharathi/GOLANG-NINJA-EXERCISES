package main

import "fmt"

// 11*noofunits*days
func total(amount func(units int) int) {
	fmt.Println(amount(22) * 30)
}
func main() {
	bill := func(units int) int {
		return units * 11
	}
	total(bill)
}
