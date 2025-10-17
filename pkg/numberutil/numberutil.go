package numberutil

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
