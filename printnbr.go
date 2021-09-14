package piscine

import "github.com/01-edu/z01"

func PrintNbr(nbr int) {
	for _, char := range string(nbr) {
		z01.PrintRune(char)
	}
}
