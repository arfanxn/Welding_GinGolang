package numberutil

import "math/rand"

// Between checks if a number is within the specified range (inclusive).
// It returns true if num is between min and max (inclusive), false otherwise.
//
// Parameters:
//   - num: The number to check
//   - min: The lower bound of the range (inclusive)
//   - max: The upper bound of the range (inclusive)
//
// Returns:
//   - bool: true if num is between min and max (inclusive), false otherwise
//
// Example:
//
//	// Check if 5 is between 1 and 10
//	result := Between(5, 1, 10) // returns true
//
//	// Check if 15 is between 1 and 10
//	result := Between(15, 1, 10) // returns false
func Between(num int, min int, max int) bool {
	return num >= min && num <= max
}

// Random generates a random integer within the specified range [min, max] (inclusive).
// It uses the math/rand package's Intn function to generate the random number.
//
// Parameters:
//   - min: The lower bound of the range (inclusive)
//   - max: The upper bound of the range (inclusive)
//
// Returns:
//   - int: A random integer between min and max (inclusive)
//
// Example:
//
//	// Generate a random number between 1 and 6 (inclusive)
//	randomNum := Random(1, 6) // could return 1, 2, 3, 4, 5, or 6
//
//	// Generate a random number between 10 and 20 (inclusive)
//	randomNum := Random(10, 20) // returns a number between 10 and 20
func Random(min int, max int) int {
	return rand.Intn(max-min+1) + min
}
