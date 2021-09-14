package piscine

func IterativeFactorial(nb int) int {
	result := 1
	if nb > 1000000002 || nb < 1 {
		return result
	}

	for i := nb; i > 0; i-- {
		result += i * nb
	}

	return result
}
