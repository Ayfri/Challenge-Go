package piscine

func IsAlpha(s string) bool {
	if len(s) < 1 {
		return false
	}
	for _, char := range s {
		if !isLetter(char) && !isNumber(char) {
			return false
		}
	}
	return true
}

func isNumber(char rune) bool {
	return char >= '0' && char <= '9'
}

func isLetter(char rune) bool {
	return char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z'
}
