package piscine

func IterativePower(nb, power int) int {
	result := nb

	if nb == 0 || power == 0 {
		return 1
	}

	for i := 1; i < nb; i++ {
		result *= power
	}

	return result
}
