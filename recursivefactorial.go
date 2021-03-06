package piscine

func RecursiveFactorial(nb int) int {
	if nb == 0 {
		return 1
	}

	if nb > 1_000_000_000 || nb < 0 {
		return 0
	}

	return nb * RecursiveFactorial(nb-1)
}
