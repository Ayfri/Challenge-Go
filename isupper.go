package piscine

func IsUpper(s string) bool {
	containsLetter := false
	for _, char := range s {
		if !isLetter(char) {
			containsLetter = true
			continue
		}
		if char <= 'A' || char >= 'Z' {
			return false
		}
	}
	if !containsLetter {
		return false
	}
	return true
}

func isLetter(char rune) bool {
	return char >= 'A' && char <= 'Z' || char >= 'a' && char <= 'z'
}
