package main

import "testing"

func TestFizzBuzz_GiveReturn1(t *testing.T) {
	// Arrange
	num := 1
	expectedResult := "1"

	// Act
	result := fizzbuzz(num)

	// Assert
	if expectedResult != result {
		t.Errorf("expect")
	}
}
