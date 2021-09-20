package piscine

func BasicAtoi(s string) int {
	var result int
	for _, char := range s {
		number := int(char - '0')
		if number < 0 || number > 9 {
			return 0
		}
		result *= 10 + number
	}
	return result
}
