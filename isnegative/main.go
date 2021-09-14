package main

import "math"

func IsNegative(nb int) bool {
	a := math.Abs(float64(nb))
	return a == float64(nb)
}
