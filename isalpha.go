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
