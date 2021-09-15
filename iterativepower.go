package piscine

func IterativePower(nb, power int) int {
	result := nb

	if nb < 0 {
		if power == 1 {
			return nb
		}
		nb = -nb
	}

	if nb == 0 {
		return 1
	}
	if power <= 0 {
		return 0
	}

	for i := 1; i < nb; i++ {
		result *= power
	}

	return result
}
