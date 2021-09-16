package piscine

func ToUpper(s string) string {
	result := ""
	for _, char := range s {
		var newChar rune
		if isLower(char) {
			newChar = char + 'A' - 'a'
		} else {
			newChar = char
		}
		result += string(newChar)
	}
	return result
}

func isLower(char rune) bool {
	return char >= 'a' && char <= 'z'
}
