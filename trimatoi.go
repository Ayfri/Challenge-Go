package piscine

func TrimAtoi(s string) int {
	result := 0
	str := ""
	isNegative := false
	for index, char := range s {
		if index > 0 && s[index-1] == '-' && len(str) == 0 {
			isNegative = true
		}

		if isNumber(char) {
			str += string(char)
		}
	}

	for index, char := range str {
		result += IterativePower(int(char-'0'), index-1)
	}
	if isNegative {
		result = -result
	}
	return result
}
