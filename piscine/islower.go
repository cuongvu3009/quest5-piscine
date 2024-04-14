// Write a function that returns true if the string passed as the parameter contains only lowercase characters, otherwise, returns false.
package piscine

func IsLower(s string) bool {
	for _, char := range s {
		if char < 'a' || char > 'z' {
			return false
		}
	}
	return true
}
