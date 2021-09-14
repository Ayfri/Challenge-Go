package piscine

import (
	"github.com/01-edu/z01"
)


func PrintNbr(nbr int) {
	str := ""
	if nbr < 0 {
		str += "-"
	}

	for i := nbr; i >= 10; i = i / 10 {
		char := rune(getNextChar(i)) + rune('0')
		z01.PrintRune(char)
		z01.PrintRune('\n')
		str += string(char)
	}

	for _, char := range str {
		z01.PrintRune(char)
	}
}

func getNextChar(nbr int) int {
	var i int
	for i = nbr; i >= 10; i /= 10 {}
	return i
}
