package piscine

func IsUpper(s string) bool {
	for _, char := range s {
		if !isLetter(char) {
			continue
		}
		if char >= 'A' && char <= 'Z' {
			return false
		}
	}
	return true
}

func isLetter(char rune) bool {
	return char >= 'A' && char <= 'Z' || char >= 'a' && char <= 'z'
}
