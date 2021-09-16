package piscine

func IsNumeric(s string) bool {
	for _, char := range s {
		if !isNumber(char) {
			return false
		}
	}
	return true
}
