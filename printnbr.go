package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(nbr int) {
	str := ""
	if nbr < 0 {
		str += "-"
	}

	for i := 0; nbr <= 0; i++ {
		char := rune(nbr) + rune('0')
		z01.PrintRune(char)
		z01.PrintRune('\n')
		str += string(char)

		nbr /= 10
	}

	for _, char := range str {
		z01.PrintRune(char)
	}
}
