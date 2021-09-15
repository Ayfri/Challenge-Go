package piscine

func AlphaCount(s string) int {
	result := 0
	for _, char := range s {
		if char <= 'a' ||char >= 'z' {
			result++
		}
	}
	return result
}
