package errorutil

// Must is a helper function that panics if the error is not nil.
// It's commonly used for variable initialization where the error should be treated as fatal.
//
// Parameters:
//   - value: The value to return if there is no error.
//   - err: The error to check.
//
// Returns:
//   - T: The provided value if err is nil.
//
// Panics:
//   - If err is not nil, it will panic with the provided error.
func Must[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}
