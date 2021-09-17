package piscine

func isAlpha(a rune) bool {
	return (a >= 'A' && a <= 'Z') || (a >= 'a' && a <= 'z') || (a >= '0' && a <= '9')
}

func toUpper(s rune) rune {
	if s-32 < 65 {
		return s
	}
	return rune(s - 32)
}

func toLower(s rune) rune {
	if s+32 > 122 {
		return s
	}
	return rune(s + 32)
}

func Capitalize(s string) string {
	runes := []rune(s)
	length := 0
	ok := true
	for i := range s {
		length = i
	}

	for i := 0; i < length; i++ {
		if isAlpha(runes[i]) == true && ok {
			if runes[i] >= 'a' && runes[i] <= 'z' {
				runes[i] = toUpper(runes[i])
			} else if runes[i] >= '0' && runes[i] <= '9' {
				runes[i+1] = toLower(runes[i+1])
			}
			ok = false
		} else if runes[i] >= 'A' && runes[i] <= 'Z' {
			runes[i] = toLower(runes[i])
		} else if isAlpha(runes[i]) == false {
			ok = true
		}
	}
	return string(runes)
}
