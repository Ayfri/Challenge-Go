package piscine

func IterativePower(nb, power int) int {
	result := 0

	if nb == 0 || power <= 0 {
		return 1
	}

	for i := 0; i < nb; i++ {
		result *= power
	}

	return result
}
