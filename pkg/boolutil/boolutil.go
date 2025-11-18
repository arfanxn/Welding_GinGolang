// Package boolutil provides boolean utility functions.
package boolutil

// Ternary is a generic function that returns one of two values based on a condition.
// It's similar to the ternary operator (condition ? a : b) found in other languages.
//
// Parameters:
//   - condition: The boolean condition to evaluate
//   - trueValue: The value to return if condition is true
//   - falseValue: The value to return if condition is false
//
// Returns:
//   - The trueValue if condition is true, falseValue otherwise
//
// Example:
//
//	result := Ternary(someValue > 10, "greater than 10", "10 or less")
func Ternary[T any](condition bool, trueValue, falseValue T) T {
	if condition {
		return trueValue
	}
	return falseValue
}
