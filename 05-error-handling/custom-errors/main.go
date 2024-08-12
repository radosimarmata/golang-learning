package main

import "fmt"

type MyError struct {
	Message string
	Code    int
}

func (e *MyError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func riskyOperation() error {
	return &MyError{Message: "something went wrong", Code: 123}
}

func main() {
	err := riskyOperation()
	if err != nil {
		fmt.Println("Custom Error:", err)
	}
}
