// test function
package main

import (
	"testing"
)

func TestCheckValidity(t *testing.T) {
	actualResult := testValidity("335-cuCZWmIW-cuCZWmIW-69-cuCZWmIW-XXbg-261-XXbg-XXbg-261-XXbg-XXbg-261-XXbg-XXbg-261-XXbg-XXbg-261-XXbg-XXbg-261-XXbg-XXbg-261-XXbg-XXbg-261-XXbg-XXbg-261-XXbg-XXbg-261-XXbg-XXbg-261-XXbg-XXbg-261-XXbg-XXbg-261-XXbg-XXbg-261-XXbg-XXbg-261-uNoEk-uNoEk-262-uNoEk-mOiZBr-152-mOiZBr-mOiZBr-152-mOiZBr-mOiZBr")
	expectedResult := false
	if actualResult != expectedResult {
		t.Errorf("Expected result false is not same as" +
			" actual string true")
	}
}

func TestCheckValidityTrue(t *testing.T) {
	actualResult := testValidity("23-ab-48-caba-56-haha")
	expectedResult := true
	if actualResult != expectedResult {
		t.Errorf("Expected result false is not same as" +
			" actual string true")
	}
}
func TestCheckValiditywithEmpty(t *testing.T) {
	actualResult := testValidity("")
	expectedResult := false
	if actualResult != expectedResult {
		t.Errorf("Expected result false is not same as" +
			" actual string true")
	}
}
func TestAverageNumber(t *testing.T) {
	actualResult := avergeNumber("23-ab-48-caba-56-haha")
	expectedResult := 42
	if actualResult != expectedResult {
		t.Errorf("Expected result false is not same as" +
			" actual string true")
	}
}
func TestWholeStory(t *testing.T) {
	actualResult := wholeStory("23-ab-48-caba-56-haha")
	expectedResult := "ab caba haha"
	if actualResult != expectedResult {
		t.Errorf("Expected result false is not same as" +
			" actual string true")
	}
}
