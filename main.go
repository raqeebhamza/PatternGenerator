package main

import (
	"fmt"
	"os"
)

func ArgumentsValidation(args []string) bool { // argument validation function
	if len(args) != 2 { // must give an argument while running the exe
		fmt.Println("Invalid Number of Arguments!")
		return false // return false if does not contain two arguments
	}
	return true // return true if number of arguments are valid
}
func main() {
	args := os.Args
	if !ArgumentsValidation(args) { // calling validation to check the number of arguments
		os.Exit(0) // return from main if number of arguments are not valid
	}

}
