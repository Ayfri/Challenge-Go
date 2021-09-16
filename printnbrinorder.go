package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	result := Reverse(ConvertNbr(n))
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

	if isNegative {
		str += "-"
	} else if nbr == 0 {
		str = "0"
	}

	return Reverse(str)
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
