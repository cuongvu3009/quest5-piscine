// Write a function that returns a concatenated string from the 'strings' passed as arguments.

package piscine

func isAlphaNumeric(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9')
}

func Capitalize(s string) string {
	var result string
	capitalize := true

	for i := 0; i < len(s); i++ {
		if isAlphaNumeric(s[i]) {
			if capitalize {
				if s[i] >= 'a' && s[i] <= 'z' {
					result += string(s[i] - 32)
				} else {
					result += string(s[i])
				}
				capitalize = false
			} else {
				if s[i] >= 'A' && s[i] <= 'Z' {
					result += string(s[i] + 32)
				} else {
					result += string(s[i])
				}
			}
		} else {
			result += string(s[i])
			capitalize = true
		}
	}

	return result
}
