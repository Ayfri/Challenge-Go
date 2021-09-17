package piscine

func SplitWhiteSpaces(s string) []string {
	var result []string
	var cache string
	for _, char := range s {
		if char == '\n' {
			result = append(result, cache)
		} else {
			cache += string(char)
		}
	}

	return result
}
