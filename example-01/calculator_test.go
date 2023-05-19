package main

import "testing"

func TestAdd(t *testing.T) {
	// Arrange เตรียม value
	num1 := 1
	num2 := 2
	expectedResult := 3
	// Act
	result := Add(num1, num2)
	// Assert
	if result != expectedResult {
		t.Errorf("Expected result to be %v, but got %v", expectedResult, result)
	}

}
