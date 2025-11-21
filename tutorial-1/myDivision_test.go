package main

import (
	"testing"
)

// TestMyDivision_Success_NoRemainder tests a perfect division (e.g., 10 / 2 = 5, remainder 0).
func TestMyDivision_Success_NoRemainder(t *testing.T) {
	// Arrange: Define the test inputs and the expected results.
	const numerator = 10
	const denominator = 2
	const expectedResult = 5
	const expectedRemainder = 0

	// Act: Call the function under test.
	result, remainder, err := myDivision(numerator, denominator)
	// Assert: Check the results.
	// 1. Check for unexpected error
	if err != nil {
		t.Fatalf("Test failed: Expected no error, but got: %v", err)
	}

	// 2. Check the calculated result
	if result != expectedResult {
		t.Errorf("Test failed: Expected result %d, but got %d", expectedResult, result)
	}

	// 3. Check the remainder
	if remainder != expectedRemainder {
		t.Errorf("Test failed: Expected remainder %d, but got %d", expectedRemainder, remainder)
	}
}

// TestMyDivision_Success_WithRemainder tests a division with a fractional part (e.g., 10 / 3 = 3, remainder 1).
func TestMyDivision_Success_WithRemainder(t *testing.T) {
	// Arrange: Define the test inputs and the expected results.
	const numerator = 10
	const denominator = 3
	const expectedResult = 3
	const expectedRemainder = 1

	// Act: Call the function under test.
	result, remainder, err := myDivision(numerator, denominator)
	// Assert: Check the results.
	// 1. Check for unexpected error
	if err != nil {
		t.Fatalf("Test failed: Expected no error, but got: %v", err)
	}

	// 2. Check the calculated result
	if result != expectedResult {
		t.Errorf("Test failed: Expected result %d, but got %d", expectedResult, result)
	}

	// 3. Check the remainder
	if remainder != expectedRemainder {
		t.Errorf("Test failed: Expected remainder %d, but got %d", expectedRemainder, remainder)
	}
}

// TestMyDivision_Error_DivideByZero tests the critical error case.
func TestMyDivision_Error_DivideByZero(t *testing.T) {
	// Arrange: Set up the inputs that will cause an error.
	const numerator = 5
	const denominator = 0
	// We expect the specific error defined in the function.
	const expectedErrorMsg = "cannot divide by zero"

	// Act: Call the function under test.
	result, remainder, err := myDivision(numerator, denominator)

	// Assert: Check the error and ensure returned values are defaults.

	// 1. Check for the expected error
	if err == nil {
		t.Fatal("Test failed: Expected a 'divide by zero' error, but got nil (no error)")
	}

	// 2. Check the error message
	if err.Error() != expectedErrorMsg {
		t.Errorf("Test failed: Expected error '%s', but got '%v'", expectedErrorMsg, err)
	}

	// 3. Check that the returned values are the zero-values (0, 0)
	if result != 0 || remainder != 0 {
		t.Errorf("Test failed: When error occurs, expected (0, 0), but got (%d, %d)", result, remainder)
	}
}
