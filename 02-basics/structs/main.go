package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{Name: "Golang", Age: 10}
	fmt.Printf("Person: %+v\n", p)
}
