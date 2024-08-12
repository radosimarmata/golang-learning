package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Printf("Number: %d\n", i)
	}
}

func main() {
	go printNumbers() // Jalankan printNumbers sebagai goroutine
	fmt.Println("This message is printed first")
	time.Sleep(6 * time.Second) // Tunggu goroutine selesai
}
