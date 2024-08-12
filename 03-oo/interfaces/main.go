package main

import "fmt"

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func printArea(s Shape) {
	fmt.Printf("Area: %f\n", s.Area())
}

func main() {
	c := Circle{Radius: 5}
	printArea(c)
}
