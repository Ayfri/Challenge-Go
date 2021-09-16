package piscine

func IsAlpha(s string) bool {
	if len(s) < 1 {
		return false
	}
	for _, char := range s {
		if !isLetter(char) {
			return false
		}
	}
	return true
}
