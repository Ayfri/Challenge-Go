package piscine

func TrimAtoi(s string) int {
	result := 0
	str := ""
	isNegative := false
	for index, char := range s {
		println("char[", index, "] == ", char)
		if index > 0 && s[index-1] == '-' && len(str) == 0 {
			isNegative = true
		}

		if isNumber(char) {
			str += string(char)
		}
	}

	str = Reverse(str)

	for index, char := range str {
		nb := int(char - '0')
		result += IterativePower(10, index) * nb
	}

	if isNegative {
		result = -result
	}

	return result
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
