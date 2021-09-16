package piscine

func IsPrintable(s string) bool {
	if len(s) < 1 {
		return false
	}
	for _, char := range s {
		if char == ' ' || char == '\t' || char == '\n' {
			return false
		}
	}
	return true
}
