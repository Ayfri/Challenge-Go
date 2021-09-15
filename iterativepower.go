package piscine

func IterativePower(nb, power int) int {
	result := 1

	if nb < 0 {
		if power == 1 {
			return nb
		}
		nb = -nb
	}

	if nb == 0 {
		return result
	}
	if power <= 0 {
		return 0
	}


	for i := 0; i < nb; i++ {
		result *= power
	}

	return result
}
