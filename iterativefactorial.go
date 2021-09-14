package piscine

func IterativeFactorial(nb int) int {
	result := 0
	if nb > 1000000002 || nb < 1 {
		return 0
	}

	for i := nb; i > 0; i-- {
		result += i * nb
	}

	return result
}
