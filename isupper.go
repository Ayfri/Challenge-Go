package piscine

func IsUpper(s string) bool {
	containsOtherChars := false
	for _, char := range s {
		if !isLetter(char) {
			containsOtherChars = false
			continue
		}
		if char <= 'A' || char >= 'Z' {
			return false
		}
	}
	if !containsOtherChars {
		return false
	}
	return true
}

func isLetter(char rune) bool {
	return char >= 'A' && char <= 'Z' || char >= 'a' && char <= 'z'
}
