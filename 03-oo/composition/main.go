package main

import "fmt"

type Engine struct {
	Horsepower int
}

type Car struct {
	Brand string
	Engine
}

func main() {
	car := Car{Brand: "Toyota", Engine: Engine{Horsepower: 150}}
	fmt.Printf("Car Brand: %s, Horsepower: %d\n", car.Brand, car.Horsepower)
}
