package main

import "fmt"

func main() {
	// Array
	var arr [3]int = [3]int{1, 2, 3}
	fmt.Println("array: ", arr)

	// Slice
	slice := []int{4, 5, 6}
	fmt.Println("slice: ", slice)
}
