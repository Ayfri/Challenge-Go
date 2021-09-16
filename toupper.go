package piscine

func ToUpper(s string) string {
	result := ""
	for _, char := range s {
		if isLower(char) {
			result += string(char + 'A' - 'z')
		}
	}
	return s
}

func isLower(char rune) bool {
	return char >= 'a' && char <= 'z'
}
