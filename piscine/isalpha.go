// Write a function that returns true if the string passed as the parameter only contains alphanumerical characters or is empty, otherwise, and returns false.
package piscine

func IsAlpha(s string) bool {
	// If the string is empty, return true
	if len(s) == 0 {
		return true
	}
	// Iterate over each character in the string
	for _, char := range s {
		// Check if the character is not an alphanumeric ASCII character
		if (char < 'a' || char > 'z') && (char < 'A' || char > 'Z') && (char < '0' || char > '9') {
			return false
		}
	}
	// All characters are alphanumeric, return true
	return true
}
