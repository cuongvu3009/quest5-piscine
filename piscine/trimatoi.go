// Write a function that transforms numbers within a string, into an int.
// If the - sign is encountered before any number it should determine the sign of the returned int.
// This function should only return an int. In the case of an invalid input, the function should return 0.
// Note: There will never be more than one sign in a string in the tests.

package piscine

func TrimAtoi(s string) int {
	// Initialize variables to store result and sign
	var result int
	sign := 1
	started := false
	// Iterate through the string
	for _, char := range s {
		// Check if the character is a digit
		if char >= '0' && char <= '9' {
			started = true
			// Convert the character to digit
			digit := int(char - '0')
			// Update the result
			result = result*10 + digit
		} else if char == '-' && !started {
			// If '-' is encountered before any number, determine the sign
			sign = -1
		} else if started {
			// If non-digit character encountered after number(s), continue loop
			// Remove this condition to continue extracting digits until the end
		}
	}
	// Apply the sign
	result *= sign
	return result
}
