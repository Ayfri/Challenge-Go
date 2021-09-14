package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(nbr int) {
	str := ""
	for i := 0; i <= 9; i++ {
		str += string(rune(nbr))
		nbr /= 2
	}

	for _, char := range str {
		z01.PrintRune(char)
	}
}
