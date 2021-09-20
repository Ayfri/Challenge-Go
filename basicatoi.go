package piscine

func BasicAtoi(s string) int {
	var result int
	for index, char := range s {
		number := int(char - '0')
		if number < 0 || number > 9 {
			return 0
		}
		result += Pow10(index) * number
	}
	return result
}

func Pow10(nb int) int {
	res := 1
	for i := 1; i <= nb; i++ {
		res *= 10
	}
	return res
}
