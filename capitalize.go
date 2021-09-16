package piscine

func Capitalize(s string) string {
	result := ""
	isWord := false
	for _, char := range s {
		if isWord {
			char += 'A' - 'a'
		}
		isWord = isSpace(char)
		result += string(char)
	}
	return result
}

func isSpace(r rune) bool {
	return r == ' '
}
