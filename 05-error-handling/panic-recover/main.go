package main

import "fmt"

func mayPanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	fmt.Println("About to panic")
	panic("Something bad happened")
	fmt.Println("This will not print")
}

func main() {
	mayPanic()
	fmt.Println("Program continues...")
}
