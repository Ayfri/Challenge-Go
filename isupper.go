package piscine

func IsUpper(s string) bool {
	for _, char := range s {
		if char <= 'A' || char >= 'Z' || !isLetter(char) {
			return false
		}
	}
	return true
}

func isLetter(char rune) bool {
	return char >= 'A' && char <= 'Z' || char >= 'a' && char <= 'z'
}
