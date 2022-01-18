# PatternGenerator
This pattern Generator will work for a pattern like this one:`23-ab-48-caba-56-haha`. This code will perform patthern validation and
some other calculation in this code.
## Functions
There are four functions I wrote on base on this pattern in go language
1.testValidity
2.averageNumber
3.wholeStory
4.storyStats
There was another function named generator to generate the random valid and invalid pattern
### testValidity(string)
TestValidity function will take an argument the pattern string and I applied the destructuring on this string to perform validation. A 
specific pattern is followed in this string that a number is followed by alphabets. My code inside the testValidation will return false 
if string is not valid pattern and if it is a valid pattern then it will return true.
### averageNumber(string)
This function will be calculating the average of all the numbers that are present in the valid string and return the answer back to the caller.
### wholeStory(string)
WholeStory will return all the words that contain alphabets. The return value would be a string contains words seperated by spaces.
### storyStats(string)
StoryStats function will return the shortest word,longest word and list of words whose length is equal to the round up and down the 
the average length of words present in the pattern.
### generator(bool)
Generator will be used to generate pattern string. It will return two types fo values valid pattern string and invalid pattern string depending upon the given boolean flag.
# How to run 
## Main file
go run main.go 23-ab-48-caba-56-haha
## Test file
go test
