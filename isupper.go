package piscine

func IsUpper(s string) bool {
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			return false
		}
	}
	return true
}
