package piscine

func Capitalize(s string) string {
	result := ""
	isWord := true
	difference := 'A' - 'a'
	for _, char := range s {
		newChar := char
		if isLetter(char) {
			if isWord {
				if !isUpper(char) {
					newChar += difference
				}
			} else if isUpper(char) {
				newChar -= difference
			}
		}
		isWord = isSpace(char)
		result += string(newChar)
	}
	return result
}

func isUpper(char rune) bool {
	return char <= 'a' && char >= 'z'
}

func isSpace(r rune) bool {
	return !IsAlpha(string(r))
}
