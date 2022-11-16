package main

import "fmt"

func main() {
	// anonymous struct
	vehicle := struct {
		doors  string
		colour string
	}{
		`doors:  2`,
		`colour: yellow`,
	}
	fmt.Println(vehicle)
}
