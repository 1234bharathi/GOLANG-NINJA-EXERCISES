package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}
type square struct {
	length float64
	width  float64
}
type shape interface {
	area() float64
}

func (s square) area() float64 {
	return s.length * s.width
}
func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}
func info(s shape) {
	fmt.Println(s.area())

}
func main() {
	s1 := square{
		length: 10,
		width:  10,
	}
	c1 := circle{
		radius: 7,
	}

	info(s1)
	info(c1)
}
