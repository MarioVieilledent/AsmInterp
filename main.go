package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Create the asm struct
var a asm = asm{}

func main() {
	// Check if the argument has been provided
	if len(os.Args) < 2 {
		fmt.Println("Please provide the path to the file as an argument")
		os.Exit(1)
	}

	// Get the file path from the arguments
	filePath := os.Args[1]

	// Read the file contents
	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", filePath, err)
		os.Exit(1)
	}

	// Interpret the code
	interpLines(string(fileBytes))
}
