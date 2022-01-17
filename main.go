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

func avergeNumber(str string) int { //function to get the average of the numbers
	if testValidity(str) { //checking the validaty of the string
		arr := strings.Split(str, "-") // splting the string to an array
		sum := 0
		for numIt := 0; numIt < len(arr); numIt += 2 { // loop over the numbers
			num, _ := strconv.Atoi(arr[numIt]) // converting string to number
			sum += num                         // calculating sum
		}
		return sum / (len(arr) / 2) // returning the average
	}
	return -1 // return -1 if string is not validate
}

func wholeStory(str string) string {
	if testValidity(str) {
		arr := strings.Split(str, "-")
		words := ""
		for alphaIt := 1; alphaIt < len(arr); alphaIt += 2 {
			words += arr[alphaIt]
			words += " "
		}
		return words[0 : len(words)-1]
	}
	return "-1"
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
	fmt.Println(testValidity(args[1]))
	fmt.Println(avergeNumber(args[1]))
	fmt.Println(wholeStory(args[1]))

}
