package piscine

func IterativeFactorial(nb int) int {
	result := 1
	if nb > 1_000_000_000 || nb < 0 {
		return 0
	}

	for i := nb; i > 0; i-- {
		result *= i
	}

	return result
}
