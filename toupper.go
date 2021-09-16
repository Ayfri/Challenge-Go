package piscine

func ToUpper(s string) string {
	for _, char := range s {
		if isLower(char) {
			char += 'A' - 'a'
		}
	}
	return s
}

func isLower(char rune) bool {
	return char >= 'a' && char <= 'z'
}
