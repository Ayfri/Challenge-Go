package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(nbr int) {
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

	str = Reverse(str)

	for _, char := range str {
		z01.PrintRune(char)
	}
}
