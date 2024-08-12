package main

import (
	"errors"
	"fmt"
)

func readFile() error {
	return errors.New("failed to open file")
}

func openFile() error {
	err := readFile()
	if err != nil {
		return fmt.Errorf("openFile: %w", err)
	}
	return nil
}

func main() {
	err := openFile()
	if err != nil {
		fmt.Println("Error occurred:", err)
	}
}
