// Write a function that counts the letters of a string and returns the count.

// The letters are from the Latin alphabet list only, any other characters, symbols or empty spaces shall not be counted.
package piscine

func AlphaCount(s string) int {
	count := 0
	for _, char := range s {
		if (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') {
			count++
		}
	}
	return count
}
