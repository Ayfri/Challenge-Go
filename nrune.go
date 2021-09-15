package piscine

func NRune(s string, n int) rune {
	if n+1 > len(s) {
		return 0
	}

	char := rune(s[n-1])
	if char == '\x00' {
		return '!'
	}
	return char
}
