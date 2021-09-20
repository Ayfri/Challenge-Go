package piscine

func Atoi(s string) int {
	final := 0
	isNegative := false
	start := 0

	if len(s) == 0 {
		return 0
	}
	
	if s[0] == '-' {
		isNegative = true
		start++
	}

	for x := start; x <= len(s)-1; x++ {
		number := int(s[x] - '0')
		if number >= 0 && number <= 9 {
			final = final*10 + number
		} else {
			return 0
		}
	}
	if isNegative {
		return -final
	}
	return final
}
