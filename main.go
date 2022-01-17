package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func testValidity(str string) bool { // validating the string given by the user
	arr := strings.Split(str, "-") // spliting the string on dashes
	if len(arr)%2 != 0 {           // checking the length of the must be even number because of the give scenerio
		return false //return false means not valid if length is Odd
	}
	alphaIt := 1
	result := regexp.MustCompile(`[0-9]`).MatchString // regex for checking is there any digit in string
	for numIt := 0; numIt < len(arr); numIt += 2 {    // loop over the string
		if _, err := strconv.Atoi(arr[numIt]); err != nil { // converting the even indexes of the arr to numbers
			return false // return false if we get any alphabet on even indexes fo array
		}
		if result(arr[alphaIt]) { //  matching the regex of the alphabet that there would be no digit
			return false // return false if found any digit
		}
		alphaIt += 2 // increament of 2 on each iteration.
	}
	return true // return true if everthing was with flow
}
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
	fmt.Println(testValidity("23-ab-48-caba-56-haha"))

}
