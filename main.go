package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
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

func wholeStory(str string) string { //function to return all the words present the string
	if testValidity(str) { // checking the validity of the string
		arr := strings.Split(str, "-") // split the string by dashes present in the string
		words := ""
		for alphaIt := 1; alphaIt < len(arr); alphaIt += 2 {
			words += arr[alphaIt] // creating a new string and use concatination
			words += " "
		}
		return words[0 : len(words)-1] // returning the words
	}
	return "-1"
}
func storyStats(str string) (string, string, string) {
	if testValidity(str) { // checking validity
		arr := strings.Split(str, "-") // spliting
		sum := 0
		// var strLens []int
		longestlen := 0                                      //storing longest length  string
		longIdx := 0                                         // index of the longest string
		shortlen := len(arr[1])                              // initial string
		shortIdx := 1                                        // initial index
		for alphaIt := 1; alphaIt < len(arr); alphaIt += 2 { // loop over all the words
			sum += len(arr[alphaIt]) //calculating sum
			// strLens = append(strLens, len(arr[alphaIt]))
			if longestlen < len(arr[alphaIt]) { // condition to get the longest string
				longestlen = len(arr[alphaIt]) //
				longIdx = alphaIt              // setting the index
			}
			if shortlen > len(arr[alphaIt]) { // condition for the shortest string
				shortlen = len(arr[alphaIt])
				shortIdx = alphaIt // setting the shortest
			}
		}
		averageLen := float64(sum) / float64(len(arr)/2) // getting average length
		averageStrs := ""
		for alphaIt := 1; alphaIt < len(arr); alphaIt += 2 {
			var averageC int = int(math.Ceil(averageLen))
			var averageF int = int(math.Floor(averageLen))
			if averageC == len(arr[alphaIt]) || averageF == len(arr[alphaIt]) { // matching the round up and down for the average
				averageStrs += arr[alphaIt]
				averageStrs += " "
			}
		}

		return arr[shortIdx], arr[longIdx], averageStrs[0 : len(averageStrs)-1] // return the results
	}
	return "-1", "-1", "-1"
}
func getRandom(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(max-min+1) + min
	return num
}
func getEvenRand(min int, max int) int {
	num := getRandom(min, max)
	for num%2 != 0 {
		num = rand.Intn(max-min+1) + min
	}
	return num
}
func generate(flag bool) string {

	blocks := getEvenRand(1, 20)

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
	fmt.Println(storyStats(args[1]))

}
