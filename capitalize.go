package piscine

func Capitalize(s string) string {
	result := ""
	isWord := true
	difference := 'A' - 'a'
	for _, char := range s {
		if isLetter(char) {
			if isWord {
				if !isUpper(char) {
					char += difference
				}
			} else if isUpper(char) {
				char -= difference
			}
		}
		isWord = isSpace(char)
		result += string(char)
	}
	return result
}

func isUpper(char rune) bool {
	return char <= 'a' && char >= 'z'
}

func isSpace(r rune) bool {
	return !IsAlpha(string(r))
}
