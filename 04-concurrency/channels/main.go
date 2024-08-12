package main

import "fmt"

func main() {
	messages := make(chan string)

	go func() {
		messages <- "Hello, Channel!"
	}()

	msg := <-messages
	fmt.Println(msg)
}
