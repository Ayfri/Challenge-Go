package piscine

import (
	"github.com/01-edu/z01"
	"strconv"
)

func PrintNbr(nbr int) {
	for _, char := range strconv.Itoa(nbr) {
		z01.PrintRune(char)
	}
}
