// Write a function that returns true if the string passed as a parameter contains only numerical characters, otherwise, returns false.
package piscine

func IsNumeric(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, char := range s {
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}
