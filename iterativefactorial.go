package piscine

func IterativeFactorial(nb int) int {
	result := 1
	if nb > 1000000000 || nb < 0 {
		return 0
	}

	for i := nb; i > 0; i-- {
		result += i * nb
	}

	return result
}
