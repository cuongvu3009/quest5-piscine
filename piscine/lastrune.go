// Write a function that returns the last rune of a string.

package piscine

func LastRune(s string) rune {
	runes := []rune(s)

	return runes[len(runes)-1]
}
