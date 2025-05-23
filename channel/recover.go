package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("example.txt")
	if err != nil {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered from panic:", r)
			}
		}()

		panic(fmt.Sprintf("Failed to open file: %v", err))
	}
	defer file.Close()

	// Read data from file
	data := make([]byte, 100)
	_, err = file.Read(data)
	if err != nil {
		panic(fmt.Sprintf("Failed to read file: %v", err))
	}

	// Do something with the data
	fmt.Println("File contents:", string(data))
}
