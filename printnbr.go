package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(nbr int) {
	str := ""
	if nbr < 0 {
		str += "-"
	}

	for i := 0; i <= 18; i++ {
		str += string(rune(nbr) + rune('0'))
		nbr /= 10
	}

	for _, char := range str {
		z01.PrintRune(char)
	}
}
