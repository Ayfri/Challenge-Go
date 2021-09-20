package piscine

func BasicAtoi2(s string) int {
	arr := []rune(s)
	y := 0
	final := 0
	for range s {
		y++
	}
	for x := 0; x <= y-1; x++ {
		char := int(arr[x] - '0')
		if char >= 0 && char < 10 {
			final *= 10 + char
		} else {
			return 0
		}
	}
	return final
}
