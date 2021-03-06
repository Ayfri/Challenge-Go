package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	result := Sort(ConvertNbr(n))
	for _, char := range result {
		z01.PrintRune(char)
	}
}

func ConvertNbr(nbr int) string {
	str := ""
	isNegative := nbr < 0

	for i, remainder := nbr, 0; i != 0; i /= 10 {
		remainder = i % 10
		if isNegative {
			remainder = -remainder
		}
		char := rune(remainder) + rune('0')
		str += string(char)
	}

	if nbr == 0 {
		str = "0"
	}

	return str
}

func Sort(s string) string {
	var result string
	for i := '0'; i <= '9'; i++ {
		for _, char := range s {
			if char == i {
				result += string(char)
			}
		}
	}
	return result
}
