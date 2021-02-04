package main

import (
	"fmt"
	"os"
)

func main() {

	currDir, err := os.Getwd()

	if err != nil {
		fmt.Println("Error ")
	}

	// fileSize := os

	fmt.Println(currDir)
}

func format() string {
	return "Format"
}
