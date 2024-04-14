// Write a function which prints the digits of an int passed in parameter in ascending order. All possible values of type int have to go through, excluding negative numbers. Conversion to int64 is not allowed.
package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	// Handle special case for 0
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	// Initialize a slice to store digits
	digits := make([]int, 0)
	// Extract digits by repeatedly dividing the number by 10
	for n > 0 {
		digit := n % 10
		digits = append(digits, digit)
		n /= 10
	}
	// Sort the digits in ascending order
	for i := 0; i < len(digits)-1; i++ {
		for j := i + 1; j < len(digits); j++ {
			if digits[i] > digits[j] {
				digits[i], digits[j] = digits[j], digits[i]
			}
		}
	}
	// Print the sorted digits
	for _, digit := range digits {
		z01.PrintRune(rune(digit) + '0')
	}
}
