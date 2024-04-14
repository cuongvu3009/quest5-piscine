package piscine

func ToLower(s string) string {
	runes := []rune(s)
	for index, char := range runes {
		if char >= 'A' && char <= 'Z' {
			runes[index] = char - 'A' + 'a'
		}
	}
	return string(runes)
}
