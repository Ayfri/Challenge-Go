package piscine

func Capitalize(s string) string {
	result := ""
	isWord := true
	difference := 'A' - 'a'
	for _, char := range s {
		newChar := char
		if isLetter(char) {
			if isWord {
				if !IsUpperChar(char) {
					newChar += difference
				}
			} else if IsUpperChar(char) {
				newChar -= difference
			}
		}
		isWord = isSpace(char)
		result += string(newChar)
	}
	return result
}

func IsUpperChar(char rune) bool {
	return char >= 'A' && char <= 'Z'
}

func isSpace(r rune) bool {
	return !IsAlpha(string(r))
}
