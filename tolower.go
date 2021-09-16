package piscine

func ToLower(s string) string {
	result := ""
	for _, char := range s {
		var newChar rune
		if isLetter(char) && !isLower(char) {
			newChar = char - ('A' - 'a')
		} else {
			newChar = char
		}
		result += string(newChar)
	}
	return result
}
