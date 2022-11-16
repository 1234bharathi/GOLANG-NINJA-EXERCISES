package main

import "fmt"

// struct
type vehicle struct {
	doors  string
	colour string
}
type truck struct {
	vehicle
	fourwheel bool
}
type sedan struct {
	vehicle
	luxury bool
}

func main() {
	truck1 := truck{
		vehicle: vehicle{
			`doors:  2`,
			`colour: yellow`,
		},
		fourwheel: true,
	}
	sedan1 := sedan{
		vehicle: vehicle{
			`doors:  4`,
			`colour: silver`,
		},
		luxury: true,
	}

	fmt.Println(truck1)
	fmt.Println(sedan1)
}
