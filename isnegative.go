package tests

import (
	"github.com/01-edu/z01"
	"math"
)

func IsNegative(nb int) {
	a := math.Abs(float64(nb))
	if a == float64(nb) {
		z01.PrintRune('T')
	} else {
		z01.PrintRune('F')
	}
}
