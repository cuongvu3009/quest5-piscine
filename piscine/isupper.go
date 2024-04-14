// Write a function that returns true if the string passed as parameter contains only uppercase characters, otherwise, returns false.
package piscine

func IsUpper(s string) bool {
	// Iterate over each character in the string
	for _, char := range s {
		// Check if the character is not an uppercase letter
		if char < 'A' || char > 'Z' {
			return false
		}
	}
	return true
}
