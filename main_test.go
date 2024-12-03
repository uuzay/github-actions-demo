package main

import "testing"

func TestFunctionThatReturnsThree(t *testing.T) {
	result := functionThatReturnsThree()
	if result != 3 {
		t.Errorf("functionThatReturnsThree() = %v, expected %v", result, 3)
	}
}
